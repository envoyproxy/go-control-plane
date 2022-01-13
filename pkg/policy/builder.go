package policy

// Builder defines a global API for various policy engines
// that are targeted towards various service mesh core functionality pieces.
// These can be used by xDS/Envoy based control-planes to help
// internally construct config.
// Builder cannot tie itself to specific pieces of config, but
// should return generic bytes (proto config).
type Builder interface {
	// Build returns a generic byte array marshaled from protobuf.
	Build() ([]byte, error)

	// BuildJSON returns JSON bytes rendered from `protojson.Marshal()`.
	BuildJSON() ([]byte, error)
}
