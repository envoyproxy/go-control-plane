workspace(name="com_github_envoyproxy_go_control_plane")

load(":bazel/repositories.bzl",
     "protobuf_dependencies",
     "gorules_dependencies",
     "envoy_api_dependencies")

protobuf_dependencies()
gorules_dependencies()
envoy_api_dependencies()

load("@envoy_api//:bazel/repositories.bzl",
     "api_dependencies")

api_dependencies()

load("@io_bazel_rules_go//go:def.bzl",
     "go_rules_dependencies",
     "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()

load("@io_bazel_rules_go//proto:def.bzl",
     "proto_register_toolchains")
proto_register_toolchains()
