# xDS Matching API Evaluation Framework

This library provides a way to evaluate the
[xDS matching API](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/advanced/matching/matching_api)
using the same logic that's implemented within the Envoy proxy. The intended use case right now is 
to support unit testing of generated configuration, but this may be extended to allow for the usage
of this matching API in other contexts.

## Example Usage

To use, first register the input and action extensions that will be used:

```golang
func init() {
	// Register the input type with the register, using the type URL of the describing protobuf.
	registry.InputExtensions["type.googleapis.com/xdsmatcher.test.proto.FooInput"] = fooInputFactory
}

// The implementation of the input.
type fooInput struct {
}

// The factory method of the input, this converts the describing protobuf into the input object.
func fooInputFactory(m proto.Message) (interface{}, error) {
	_, ok := m.(*pbtest.FooInput)
	if !ok {
		return nil, errors.New("unexpected proto")
	}

	return fooInput{}, nil
}

// Implement types.DataInput for the input type.
func (fooInput) Input(d types.MatchingData) (types.DataInputResult, error) {
	// Cast to the expected data type, this should never fail.
	testData, ok := d.(*pbtest.TestData)
	if !ok {
		return types.DataInputResult{}, errors.New("invalid matching data type")
	}

	// Return the value to match on from the data, qualifying it with a data availability.
	return types.DataInputResult{
		Availability: types.AllDataAvailable,
		Data:         &testData.Foo,
	}, nil
}
```

```golang
func init() {
	// Register the action with the register, using the type URL of the describing protobuf.
	registry.Action["type.googleapis.com/xdsmatcher.test.proto.Action"] = actionFactory
}

// The implementation of the action.
type action struct {
}

// The factory method of the action, this converts the describing protobuf into the input object.
func actionFactory(m proto.Message) (interface{}, error) {
	_, ok := m.(*pbtest.Action)
	if !ok {
		return nil, errors.New("unexpected proto")
	}

	return action{}, nil
}
```

Once registered, a match tree can be created and evaluated:

```golang
	yaml := `
matcher_tree:
  input:
    name: foo
    typed_config:
      "@type": "type.googleapis.com/xdsmatcher.test.proto.FooInput"
  exact_match_map:
    map:
      "foo":
        action:
          name: action
          typed_config:
            "@type": type.googleapis.com/xdsmatcher.test.proto.MatchAction
`
	matcher := &pbmatcher.Matcher{}
	_ = iproto.ProtoFromYaml([]byte(yaml), matcher)

	m, _ := matcher.Create(matcher)

	data := &pbtest.TestData{}
	r, _ := m.Match(data)
	if r.MatchResult != nil && r.MatchResult.Action != nil {
		// resulting action
		action, ok := r.MatchResult.Action.(action)
	}
```

See matcher_test.go for more examples of how to use the API.

# Development
To install the required dev tools on Linux:

```
apt-get install -y protobuf-compiler

go install google.golang.org/protobuf/cmd/protoc-gen-go
```
