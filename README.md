# control-plane
This repository contains a Go-based implementation of an API server that
implements the discovery service APIs defined in
[data-plane-api](https://github.com/envoyproxy/data-plane-api).

## Scope

Due to the variety of platforms out there, there is no single
control plane implementation that can satisfy everyone's needs. Hence this
code base does not attempt to be a full scale control plane for a fleet of
Envoy proxies. Instead, it provides infrastructure that is shared by
multiple different control plane implementations. The components provided
by this library are:

* _API Server:_ A generic gRPC based API server that implements xDS APIs as defined
  in the
  [data-plane-api](https://github.com/envoyproxy/data-plane-api). The API
  server is responsible for pushing configuration updates to
  Envoys. Consumers should be able to import this go library and use the
  API server as is, in production deployments.

* _Configuration Cache:_ The library will cache Envoy configurations in
memory in an attempt to provide fast response to consumer Envoys. It is the
responsibility of the consumer of this library to populate the cache as
well as invalidate it when necessary. The cache will be keyed based on a
pre-defined hash function whose keys are based on the
[Node information](https://github.com/envoyproxy/data-plane-api/blob/d4988844024d0bcff4bcd030552eabe3396203fa/api/base.proto#L26-L36).

At this moment, this repository will not tackle translating platform
specific representation of resources (e.g., services, instances of
services, etc.) into Envoy-style configuration. Based on usage and
feedback, we might decided to revisit this aspect at a later point in time.

## Requirements

1. Go 1.9.x

## Quick start

1. Setup tools and dependencies

```sh
make tools
make depend.install
```

2. Edit the code in your favorite IDE

3. Format, vet and lint the code

```sh
make check
```

3. Build and test

```sh
make build
make test
```
