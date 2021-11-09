package matcher

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	pbmatcher "github.com/envoyproxy/go-control-plane/envoy/config/common/matcher/v3"
	iproto "github.com/envoyproxy/go-control-plane/xdsmatcher/internal/proto"
	_ "github.com/envoyproxy/go-control-plane/xdsmatcher/test"
	pbtest "github.com/envoyproxy/go-control-plane/xdsmatcher/test/proto"
)

func TestSimple(t *testing.T) {
	configuration := func(inputType string) string {
		return fmt.Sprintf(`
matcher_tree:
  input:
    name: foo
    typed_config:
      "@type": %s
  exact_match_map:
    map:
      "foo":
        action:
          name: action
          typed_config:
            "@type": type.googleapis.com/xdsmatcher.test.proto.MatchAction
`, inputType)
	}

	testHelper(t, configuration)
}

func TestConjunction(t *testing.T) {
	conjunctionMatcherConfig := func(inputType string) string {
		return fmt.Sprintf(`
matcher_list:
  matchers:
  - predicate:
      and_matcher:
        predicate:
        - single_predicate:
            input:
              name: bar
              typed_config:
                "@type": type.googleapis.com/xdsmatcher.test.proto.BarInput
                index: 1
            value_match:
              safe_regex:
                regex: two
                google_re2: {}
        - single_predicate:
            input:
              name: foo
              typed_config:
                "@type": %s
            value_match:
              safe_regex:
                regex: foo
                google_re2: {}
    on_match:
      action:
        name: action
        typed_config:
          "@type": type.googleapis.com/xdsmatcher.test.proto.MatchAction
          `, inputType)
	}

	testHelper(t, conjunctionMatcherConfig)
}

func TestDisjunctionMatch(t *testing.T) {
	disjunnctionMatcherConfig := func(inputType string) string {
		return fmt.Sprintf(`
matcher_list:
  matchers:
  - predicate:
      or_matcher:
        predicate:
        - single_predicate:
            input:
              name: bar
              typed_config:
                "@type": type.googleapis.com/xdsmatcher.test.proto.BarInput
                index: 1
            value_match:
              safe_regex:
                regex: wrong # something that won't match
                google_re2: {}
        - single_predicate:
            input:
              name: bar
              typed_config:
                "@type": %s
            value_match:
              safe_regex:
                regex: ^foo$
                google_re2: {}
    on_match:
      action:
        name: action
        typed_config:
          "@type": type.googleapis.com/xdsmatcher.test.proto.MatchAction
          `, inputType)
	}

	testHelper(t, disjunnctionMatcherConfig)
}

func testHelper(t *testing.T, configurationGenerator func(typeName string) string) {
	t.Helper()
	tcs := []struct {
		name          string
		inputType     string
		input         interface{}
		hasMatch      bool
		needsMoreData bool
		matchingError bool
	}{
		{
			name:      "successful matching FooInput against foo",
			inputType: "type.googleapis.com/xdsmatcher.test.proto.FooInput",
			input:     testData1,
			hasMatch:  true,
		},
		{
			name:      "failure to match FooInput against 'not foo'",
			inputType: "type.googleapis.com/xdsmatcher.test.proto.FooInput",
			input:     testData2,
		},
		{
			name:          "failure to match against nil, more data available",
			inputType:     "type.googleapis.com/xdsmatcher.test.proto.MoreDataAvailableInput",
			input:         testData1,
			needsMoreData: true,
		},
		{
			name:          "input returning error",
			inputType:     "type.googleapis.com/xdsmatcher.test.proto.ErrorInput",
			input:         testData1,
			needsMoreData: true,
			matchingError: true,
		},
		{
			name:          "not available input",
			inputType:     "type.googleapis.com/xdsmatcher.test.proto.NoDataAvailableInput",
			input:         testData1,
			needsMoreData: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			configuration := configurationGenerator(tc.inputType)
			m, err := createMatcher(t, configuration)
			assert.NoError(t, err)

			r, err := m.Match(tc.input)

			if tc.matchingError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			if tc.hasMatch {
				assert.NotNil(t, r.MatchResult)
				assert.NotNil(t, r.MatchResult.Action)
			} else {
				assert.Nil(t, r.MatchResult)
			}

			assert.Equal(t, tc.needsMoreData, r.NeedMoreData)
		})
	}
}

func TestNested(t *testing.T) {
	configuration := `
matcher_tree:
  input:
    name: foo
    typed_config:
      "@type": type.googleapis.com/xdsmatcher.test.proto.FooInput
  exact_match_map:
    map:
      "foo":
        matcher:
          matcher_list:
            matchers:
            - predicate:
                single_predicate:
                  input:
                    name: bar
                    typed_config:
                     "@type": type.googleapis.com/xdsmatcher.test.proto.BarInput
                     index: 1
                  value_match:
                    safe_regex:
                      regex: t.o
                      google_re2: {}
              on_match:
                action:
                  name: action
                  typed_config:
                    "@type": type.googleapis.com/xdsmatcher.test.proto.MatchAction
`

	m, err := createMatcher(t, configuration)
	assert.NoError(t, err)
	r, err := m.Match(testData1)
	assert.NoError(t, err)

	log.Printf("%+v", m)
	assert.NotNil(t, r.MatchResult)
	assert.NotNil(t, r.MatchResult.Action)
}

func TestCreateInvalidExtensions(t *testing.T) {
	config := `
matcher_tree:
  input:
    name: foo
    typed_config:
      "@type": type.googleapis.com/xdsmatcher.test.proto.NotRegistered
  exact_match_map:
    map:
      "foo":
        action:
          name: action
          typed_config:
            "@type": type.googleapis.com/xdsmatcher.test.proto.NotRegistered
`

	m, err := createMatcher(t, config)
	assert.Error(t, err)
	assert.Nil(t, m)

	config = `
matcher_tree:
  input:
    name: foo
    typed_config:
      "@type": type.googleapis.com/xdsmatcher.test.proto.FooInput
  exact_match_map:
    map:
      "foo":
        action:
          name: action
          typed_config:
            "@type": type.googleapis.com/xdsmatcher.test.proto.NotRegistered
`

	m, err = createMatcher(t, config)
	assert.Error(t, err)
	assert.Nil(t, m)
}

var testData1 = &pbtest.TestData{
	Foo: "foo",
	Bar: []string{"one", "two"},
}

var testData2 = &pbtest.TestData{
	Foo: "not-foo",
}

func createMatcher(t *testing.T, yaml string) (*MatcherTree, error) {
	matcher := &pbmatcher.Matcher{}
	assert.NoError(t, iproto.ProtoFromYaml([]byte(yaml), matcher))

	return Create(matcher)
}
