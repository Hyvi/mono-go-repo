load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
def docker_repos():
    http_archive(
        name = "io_bazel_rules_docker",
      sha256 = "27d53c1d646fc9537a70427ad7b034734d08a9c38924cc6357cc973fed300820",
    strip_prefix = "rules_docker-0.24.0",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.24.0/rules_docker-v0.24.0.tar.gz"], 
    )
