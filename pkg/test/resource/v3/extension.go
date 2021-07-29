// Copyright 2018 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Package resource creates test xDS resources
package resource

import (
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	common "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/common/fault/v3"
	fault "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/fault/v3"
	Type "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	"github.com/golang/protobuf/ptypes"
	duration "github.com/golang/protobuf/ptypes/duration"
)

// MakeExtensionConfig make HTTP fault as a Extension Config
func MakeExtensionConfig(extension string) *core.TypedExtensionConfig {
	Fault := &fault.HTTPFault{
		Delay: &common.FaultDelay{
			FaultDelaySecifier: &common.FaultDelay_FixedDelay{
				FixedDelay: &duration.Duration{
					Seconds: 3,
				},
			},
			Percentage: &Type.FractionalPercent{
				Numerator:   0,
				Denominator: Type.FractionalPercent_HUNDRED,
			},
		},
		Abort: &fault.FaultAbort{
			ErrorType: &fault.FaultAbort_HttpStatus{
				HttpStatus: 503,
			},
			Percentage: &Type.FractionalPercent{
				Numerator:   0,
				Denominator: Type.FractionalPercent_HUNDRED,
			},
		},
	}
	extensionConfig, err := ptypes.MarshalAny(Fault)
	if err != nil {
		panic(err)
	}
	return &core.TypedExtensionConfig{
		Name:        extension,
		TypedConfig: extensionConfig,
	}
}
