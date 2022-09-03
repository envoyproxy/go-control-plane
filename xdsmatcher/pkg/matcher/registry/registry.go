package registry

import (
	"errors"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	pbcore "github.com/cncf/xds/go/xds/core/v3"
	"github.com/envoyproxy/go-control-plane/xdsmatcher/pkg/matcher/types"
)

type RegistryType map[string]func(proto.Message) (interface{}, error)

var InputExtensions = make(RegistryType)
var ActionExtensions = make(RegistryType)

func ResolveInputExtension(tec *pbcore.TypedExtensionConfig) (types.DataInput, error) {
	out, err := resolveExtension(tec, InputExtensions)
	if err != nil {
		return nil, err
	}

	return out.(types.DataInput), nil
}

func ResolveActionExtension(tec *pbcore.TypedExtensionConfig) (types.Action, error) {
	return resolveExtension(tec, ActionExtensions)
}

func resolveExtension(tec *pbcore.TypedExtensionConfig, registry RegistryType) (interface{}, error) {
	factory, ok := registry[tec.TypedConfig.TypeUrl]
	if !ok {
		return nil, errors.New("extension not implemented: extension " + tec.TypedConfig.TypeUrl)
	}

	m, err := anypb.UnmarshalNew(tec.TypedConfig, proto.UnmarshalOptions{})
	if err != nil {
		return nil, err
	}

	return factory(m)
}
