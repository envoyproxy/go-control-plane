package matcher

import (
	"testing"

	pbmatcher "github.com/envoyproxy/go-control-plane/envoy/config/common/matcher/v3"
	iproto "github.com/lyft/xdsmatcher/internal/proto"
	_ "github.com/lyft/xdsmatcher/test"
	pbtest "github.com/lyft/xdsmatcher/test/proto"
	"github.com/stretchr/testify/assert"
)

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

	matcher := &pbmatcher.Matcher{}
	assert.NoError(t, iproto.ProtoFromYaml([]byte(configuration), matcher))

	m, err := Create(matcher)
	assert.NoError(t, err)

	d := &pbtest.TestData{
		Foo: "foo",
		Bar: []string{"one", "two"},
	}
	r, err := m.Match(d)
	assert.NoError(t, err)

	assert.NotNil(t, r.MatchResult)
	assert.NotNil(t, r.MatchResult.Action)
}
