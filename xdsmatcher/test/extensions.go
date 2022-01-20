package test

import (
	"errors"

	"google.golang.org/protobuf/proto"

	"github.com/envoyproxy/go-control-plane/xdsmatcher/pkg/matcher/registry"
	"github.com/envoyproxy/go-control-plane/xdsmatcher/pkg/matcher/types"
	pbtest "github.com/envoyproxy/go-control-plane/xdsmatcher/test/proto"
)

// This file contains a number of input and action extensions registered for the purpose of testing the framework.

func init() {
	registry.InputExtensions["type.googleapis.com/xdsmatcher.test.proto.FooInput"] = fooInputFactory
	registry.InputExtensions["type.googleapis.com/xdsmatcher.test.proto.BarInput"] = barInputFactory
	registry.InputExtensions["type.googleapis.com/xdsmatcher.test.proto.MoreDataAvailableInput"] = moreDataAvailableInputFactory
	registry.InputExtensions["type.googleapis.com/xdsmatcher.test.proto.ErrorInput"] = errorInputFactory
	registry.InputExtensions["type.googleapis.com/xdsmatcher.test.proto.NoDataAvailableInput"] = noDataInputFactory

	registry.ActionExtensions["type.googleapis.com/xdsmatcher.test.proto.MatchAction"] = matchActionFactory
}

type fooInput struct {
}

func fooInputFactory(m proto.Message) (interface{}, error) {
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

func barInputFactory(m proto.Message) (interface{}, error) {
	p, ok := m.(*pbtest.BarInput)
	if !ok {
		return nil, errors.New("unexpected proto")
	}

	return barInput{uint64(p.Index)}, nil
}

func (t barInput) Input(d types.MatchingData) (types.DataInputResult, error) {
	metricData, ok := d.(*pbtest.TestData)
	if !ok {
		return types.DataInputResult{}, errors.New("invalid matching data type")
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

type moreDataAvailableInput struct{}

func (moreDataAvailableInput) Input(d types.MatchingData) (types.DataInputResult, error) {
	return types.DataInputResult{
		Availability: types.MoreDataMightBeAvailable,
		Data:         nil,
	}, nil
}

func moreDataAvailableInputFactory(m proto.Message) (interface{}, error) {
	return moreDataAvailableInput{}, nil
}

type errorInput struct{}

func (errorInput) Input(d types.MatchingData) (types.DataInputResult, error) {
	return types.DataInputResult{}, errors.New("test error")
}

func errorInputFactory(m proto.Message) (interface{}, error) {
	return errorInput{}, nil
}

type noDataInput struct{}

func (noDataInput) Input(d types.MatchingData) (types.DataInputResult, error) {
	return types.DataInputResult{
		Availability: types.NotAvailable,
	}, nil
}

func noDataInputFactory(m proto.Message) (interface{}, error) {
	return noDataInput{}, nil
}

type matchAction struct{}

func matchActionFactory(m proto.Message) (interface{}, error) {
	_, ok := m.(*pbtest.MatchAction)
	if !ok {
		return nil, errors.New("unexpected proto")
	}

	return matchAction{}, nil
}
