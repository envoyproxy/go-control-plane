# TODO(kuat): This can switch back to a point release http_archive at the next
# release (> 3.4.1), we need HEAD proto_library support and
# https://github.com/google/protobuf/pull/3761.
PROTOBUF_SHA = "c4f59dcc5c13debc572154c8f636b8a9361aacde"

# ENVOY_API_SHA = "5b29f002b44f3f6e3921bd3b4535aeb91f89e84a" # Oct 23, 2018
ENVOY_API_SHA = "488be171a42e142d2f8dbaa1b3e176a6ec70bdb4"

def protobuf_dependencies():
    native.http_archive(
        name = "com_google_protobuf",
        strip_prefix = "protobuf-" + PROTOBUF_SHA,
        url = "https://github.com/google/protobuf/archive/" + PROTOBUF_SHA + ".tar.gz",
    )

    # Needed for cc_proto_library, Bazel doesn't support aliases today for repos,
    # see https://groups.google.com/forum/#!topic/bazel-discuss/859ybHQZnuI and
    # https://github.com/bazelbuild/bazel/issues/3219.
    native.http_archive(
        name = "com_google_protobuf_cc",
        strip_prefix = "protobuf-" + PROTOBUF_SHA,
        url = "https://github.com/google/protobuf/archive/" + PROTOBUF_SHA + ".tar.gz",
    )

    native.bind(
        name = "six",
        actual = "@six_archive//:six",
    )

    native.new_http_archive(
        name = "six_archive",
        build_file = "@com_google_protobuf//:six.BUILD",
        sha256 = "105f8d68616f8248e24bf0e9372ef04d3cc10104f1980f54d57b2ce73a5ad56a",
        url = "https://pypi.python.org/packages/source/s/six/six-1.10.0.tar.gz#md5=34eed507548117b2ab523ab14b2f8b55",
    )

def gorules_dependencies():
    native.git_repository(
        name = "io_bazel_rules_go",
        remote = "https://github.com/bazelbuild/rules_go.git",
        commit = "4374be38e9a75ff5957c3922adb155d32086fe14",
    )

def envoy_api_dependencies():
    native.http_archive(
        name = "envoy_api",
        strip_prefix = "data-plane-api-" + ENVOY_API_SHA,
        urls = ["https://github.com/envoyproxy/data-plane-api/archive/" + ENVOY_API_SHA + ".tar.gz"],
    )
