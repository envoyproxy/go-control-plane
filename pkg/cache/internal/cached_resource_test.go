package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
	"time"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const testTypeURL = "type.googleapis.com/google.protobuf.StringValue"

func deterministicMarshal(t *testing.T, msg proto.Message) []byte {
	t.Helper()
	b, err := proto.MarshalOptions{Deterministic: true}.Marshal(msg)
	require.NoError(t, err)
	return b
}

func hashBytes(b []byte) string {
	h := sha256.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func TestGetSotwResource_NoTTL(t *testing.T) {
	res := wrapperspb.String("hello")
	cached := NewCachedResource("my-resource", testTypeURL, res)

	got, err := cached.GetSotwResource(false)
	require.NoError(t, err)

	assert.Equal(t, testTypeURL, got.TypeUrl)
	assert.Equal(t, deterministicMarshal(t, res), got.Value)
}

func TestGetSotwResource_NoTTL_HeartbeatIgnored(t *testing.T) {
	res := wrapperspb.String("hello")
	cached := NewCachedResource("my-resource", testTypeURL, res)

	// Without TTL, heartbeat flag has no effect — resource is always included.
	got, err := cached.GetSotwResource(true)
	require.NoError(t, err)

	assert.Equal(t, testTypeURL, got.TypeUrl)
	assert.Equal(t, deterministicMarshal(t, res), got.Value)
}

func TestGetSotwResource_WithTTL(t *testing.T) {
	res := wrapperspb.String("world")
	ttl := 5 * time.Second
	cached := NewCachedResource("ttl-resource", testTypeURL, res, WithResourceTTL(&ttl))

	got, err := cached.GetSotwResource(false)
	require.NoError(t, err)

	// When TTL is set, the result is wrapped in a discovery.Resource envelope.
	assert.Equal(t, deltaResourceTypeURL, got.TypeUrl)

	// Unwrap and verify the discovery.Resource.
	var wrapper discovery.Resource
	require.NoError(t, proto.Unmarshal(got.Value, &wrapper))

	assert.Equal(t, "ttl-resource", wrapper.Name)
	assert.Equal(t, durationpb.New(ttl), wrapper.Ttl)
	require.NotNil(t, wrapper.Resource)
	assert.Equal(t, testTypeURL, wrapper.Resource.TypeUrl)
	assert.Equal(t, deterministicMarshal(t, res), wrapper.Resource.Value)
}

func TestGetSotwResource_WithTTL_Heartbeat(t *testing.T) {
	res := wrapperspb.String("heartbeat")
	ttl := 3 * time.Second
	cached := NewCachedResource("hb-resource", testTypeURL, res, WithResourceTTL(&ttl))

	got, err := cached.GetSotwResource(true)
	require.NoError(t, err)

	assert.Equal(t, deltaResourceTypeURL, got.TypeUrl)

	var wrapper discovery.Resource
	require.NoError(t, proto.Unmarshal(got.Value, &wrapper))

	assert.Equal(t, "hb-resource", wrapper.Name)
	assert.Equal(t, durationpb.New(ttl), wrapper.Ttl)
	// On heartbeat, the inner resource should be omitted.
	assert.Nil(t, wrapper.Resource)
}

func TestGetDeltaResource(t *testing.T) {
	res := wrapperspb.String("delta-value")
	cached := NewCachedResource("delta-resource", testTypeURL, res)

	got, err := cached.GetDeltaResource()
	require.NoError(t, err)

	assert.Equal(t, "delta-resource", got.Name)

	// Verify the embedded Any.
	require.NotNil(t, got.Resource)
	assert.Equal(t, testTypeURL, got.Resource.TypeUrl)
	marshaled := deterministicMarshal(t, res)
	assert.Equal(t, marshaled, got.Resource.Value)

	// Verify the version is the hash of the deterministic marshaling.
	assert.Equal(t, hashBytes(marshaled), got.Version)
}

func TestGetDeltaResource_WithCustomVersion(t *testing.T) {
	res := wrapperspb.String("custom-version")
	cached := NewCachedResource("cv-resource", testTypeURL, res, WithResourceVersion("v42"))

	got, err := cached.GetDeltaResource()
	require.NoError(t, err)

	assert.Equal(t, "v42", got.Version)
	assert.Equal(t, "cv-resource", got.Name)
}

func TestGetDeltaResource_WithMarshaledResource(t *testing.T) {
	res := wrapperspb.String("pre-marshaled")
	preMarshaled := deterministicMarshal(t, res)
	cached := NewCachedResource("pm-resource", testTypeURL, res, WithMarshaledResource(preMarshaled))

	got, err := cached.GetDeltaResource()
	require.NoError(t, err)

	assert.Equal(t, preMarshaled, got.Resource.Value)
	// Version should be derived from the provided bytes.
	assert.Equal(t, hashBytes(preMarshaled), got.Version)
}

func TestGetSotwAndDeltaConsistency(t *testing.T) {
	// Both methods should produce consistent marshaled bytes for the same resource.
	res := wrapperspb.String("consistent")
	cached := NewCachedResource("cons-resource", testTypeURL, res)

	sotw, err := cached.GetSotwResource(false)
	require.NoError(t, err)

	delta, err := cached.GetDeltaResource()
	require.NoError(t, err)

	// The raw bytes in the Any should match the delta resource's Any.
	assert.Equal(t, sotw.Value, delta.Resource.Value)
	assert.Equal(t, sotw.TypeUrl, delta.Resource.TypeUrl)
}

func TestGetDeltaResource_StableVersion(t *testing.T) {
	// Calling GetDeltaResource multiple times should yield the same version.
	res := wrapperspb.String("stable")
	cached := NewCachedResource("stable-resource", testTypeURL, res)

	r1, err := cached.GetDeltaResource()
	require.NoError(t, err)

	r2, err := cached.GetDeltaResource()
	require.NoError(t, err)

	assert.Equal(t, r1.Version, r2.Version)
}

func TestGetSotwResource_UsesTypeURL(t *testing.T) {
	// Verify the typeURL is correctly propagated in the Any wrapper.
	customTypeURL := "type.googleapis.com/custom.Type"
	res := wrapperspb.String("typed")
	cached := NewCachedResource("typed-resource", customTypeURL, res)

	got, err := cached.GetSotwResource(false)
	require.NoError(t, err)

	assert.Equal(t, customTypeURL, got.TypeUrl)
}

func TestGetDeltaResource_UsesTypeURL(t *testing.T) {
	customTypeURL := "type.googleapis.com/custom.Type"
	res := wrapperspb.String("typed")
	cached := NewCachedResource("typed-resource", customTypeURL, res)

	got, err := cached.GetDeltaResource()
	require.NoError(t, err)

	assert.Equal(t, customTypeURL, got.Resource.TypeUrl)
}

func TestGetDeltaResource_WithMarshaledResource_Empty(t *testing.T) {
	// Passing empty bytes to WithMarshaledResource should be a no-op (fall back to default marshaling).
	res := wrapperspb.String("fallback")
	cached := NewCachedResource("fb-resource", testTypeURL, res, WithMarshaledResource(nil))

	got, err := cached.GetDeltaResource()
	require.NoError(t, err)

	assert.Equal(t, deterministicMarshal(t, res), got.Resource.Value)
}

func TestGetDeltaResource_WithResourceVersion_Empty(t *testing.T) {
	// Passing empty string to WithResourceVersion should be a no-op (fall back to hash-based version).
	res := wrapperspb.String("version-fallback")
	cached := NewCachedResource("vfb-resource", testTypeURL, res, WithResourceVersion(""))

	got, err := cached.GetDeltaResource()
	require.NoError(t, err)

	marshaled := deterministicMarshal(t, res)
	assert.Equal(t, hashBytes(marshaled), got.Version)
}

func TestGetSotwResource_NoTTL_CanBeUnmarshaledBack(t *testing.T) {
	original := wrapperspb.String("roundtrip")
	cached := NewCachedResource("rt-resource", testTypeURL, original)

	got, err := cached.GetSotwResource(false)
	require.NoError(t, err)

	// Unmarshal back from the Any and verify equality.
	var recovered anypb.Any
	recovered.TypeUrl = got.TypeUrl
	recovered.Value = got.Value

	roundtripped := &wrapperspb.StringValue{}
	require.NoError(t, proto.Unmarshal(recovered.Value, roundtripped))
	assert.True(t, proto.Equal(original, roundtripped))
}
