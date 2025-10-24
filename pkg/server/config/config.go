package config

import "github.com/envoyproxy/go-control-plane/pkg/log"

// Opts for individual xDS implementations that can be
// utilized through the functional opts pattern.
type Opts struct {
	// If true respond to ADS requests with a guaranteed resource ordering
	Ordered bool

	Logger log.Logger

	// If true, deactivate legacy wildcard mode for all resource types
	legacyWildcardDeactivated bool

	// Deactivate legacy wildcard mode for specific resource types
	legacyWildcardDeactivatedTypes map[string]struct{}
}

func NewOpts() Opts {
	return Opts{
		Ordered: false,
		Logger:  log.NewDefaultLogger(),
	}
}

// LegacyWildcardDeactivated returns whether legacy wildcard mode is deactivated for all resource types
func (o Opts) LegacyWildcardDeactivated() bool {
	return o.legacyWildcardDeactivated
}

// LegacyWildcardDeactivatedTypes returns the set of resource types for which legacy wildcard mode is deactivated
func (o Opts) LegacyWildcardDeactivatedTypes() map[string]struct{} {
	return o.legacyWildcardDeactivatedTypes
}

// Each xDS implementation should implement their own functional opts.
// It is recommended that config values be added in this package specifically,
// but the individual opts functions should be in their respective
// implementation package so the import looks like the following:
//
// `sotw.WithOrderedADS()`
// `delta.WithOrderedADS()`
//
// this allows for easy inference as to which opt applies to what implementation.
type XDSOption func(*Opts)

// DeactivateLegacyWildcard deactivates legacy wildcard mode for all resource types.
// In legacy wildcard mode, empty requests to a stream, are treated as wildcard requests as long
// as there is no request made with resources or explicit wildcard requests on the same stream.
// When deactivated, empty requests are treated as a request with no subscriptions to any resource.
// This is recommended for when you are using the go-control-plane to serve grpc-xds clients.
// These clients never want to treat an empty request as a wildcard subscription.
func DeactivateLegacyWildcard() XDSOption {
	return func(o *Opts) {
		o.legacyWildcardDeactivated = true
	}
}

// DeactivateLegacyWildcardForTypes deactivates legacy wildcard mode for specific resource types.
// In legacy wildcard mode, empty requests to a stream, are treated as wildcard requests as long
// as there is no request made with resources or explicit wildcard requests on the same stream.
// When deactivated, empty requests are treated as a request with no subscriptions to any resource.
func DeactivateLegacyWildcardForTypes(types []string) XDSOption {
	return func(o *Opts) {
		typeMap := make(map[string]struct{}, len(types))
		for _, t := range types {
			typeMap[t] = struct{}{}
		}
		o.legacyWildcardDeactivatedTypes = typeMap
	}
}
