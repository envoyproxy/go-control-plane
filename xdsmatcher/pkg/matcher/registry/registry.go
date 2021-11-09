package registry

import (
	"errors"

	pbcore "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/lyft/xdsmatcher/pkg/matcher/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var InputExtensions map[string]func(proto.Message) (types.DataInput, error) = map[string]func(proto.Message) (types.DataInput, error){}

func ResolveInputExtension(tec *pbcore.TypedExtensionConfig) (types.DataInput, error) {
	factory, ok := InputExtensions[tec.TypedConfig.TypeUrl]
	if !ok {
		return nil, errors.New("extension not implemented: extension " + tec.TypedConfig.TypeUrl)
	}

	m, err := anypb.UnmarshalNew(tec.TypedConfig, proto.UnmarshalOptions{})
	if err != nil {
		return nil, err
	}

	return factory(m)
}
