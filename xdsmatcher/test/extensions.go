package test

import (
	"errors"

	"github.com/lyft/xdsmatcher/pkg/matcher/registry"
	"github.com/lyft/xdsmatcher/pkg/matcher/types"
	pbtest "github.com/lyft/xdsmatcher/test/proto"
	"google.golang.org/protobuf/proto"
)

func init() {
	registry.InputExtensions["type.googleapis.com/xdsmatcher.test.proto.FooInput"] = fooInputFactory
	registry.InputExtensions["type.googleapis.com/xdsmatcher.test.proto.BarInput"] = barInputFactory
}

type fooInput struct {
}

func fooInputFactory(m proto.Message) (types.DataInput, error) {
	_, ok := m.(*pbtest.FooInput)
	if !ok {
		return nil, errors.New("unexpected proto")
	}

	return fooInput{}, nil
}

func (fooInput) Input(d types.MatchingData) (types.DataInputResult, error) {
	testData, ok := d.(*pbtest.TestData)
	if !ok {
		return types.DataInputResult{}, errors.New("invlaid matching data type")
	}

	return types.DataInputResult{
		Availability: types.AllDataAvailable,
		Data:         &testData.Foo,
	}, nil
}

type barInput struct {
	index uint64
}

func barInputFactory(m proto.Message) (types.DataInput, error) {
	p, ok := m.(*pbtest.BarInput)
	if !ok {
		return nil, errors.New("unexpected proto")
	}

	return barInput{uint64(p.Index)}, nil
}

func (t barInput) Input(d types.MatchingData) (types.DataInputResult, error) {
	metricData, ok := d.(*pbtest.TestData)
	if !ok {
		return types.DataInputResult{}, errors.New("invlaid matching data type")
	}

	if t.index >= uint64(len(metricData.Bar)) {
		return types.DataInputResult{
			Availability: types.AllDataAvailable,
			Data:         nil,
		}, nil

	}
	value := metricData.Bar[t.index]
	return types.DataInputResult{
		Availability: types.AllDataAvailable,
		Data:         &value,
	}, nil

}
