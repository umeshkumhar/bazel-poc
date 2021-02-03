
## General rules
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

## rules_go
http_archive(
     name = "io_bazel_rules_go",
     urls = [
          "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/v0.20.1/rules_go-v0.20.1.tar.gz",
  "https://github.com/bazelbuild/rules_go/releases/download/v0.20.1/rules_go-v0.20.1.tar.gz",
],
sha256="842ec0e6b4fbfdd3de6150b61af92901eeb73681fd4d185746644c338f51d4c0",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()


## Gazelle
http_archive(
     name = "bazel_gazelle",
     urls = [
       "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/0.18.2/bazel-gazelle-0.18.2.tar.gz",
"https://github.com/bazelbuild/bazel-gazelle/releases/download/0.18.2/bazel-gazelle-0.18.2.tar.gz",
],
     sha256 = "7fc87f4170011201b1690326e8c16c5d802836e3a0d617d8f75c3af2b23180c4",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()


## rules_docker
http_archive(
    name = "io_bazel_rules_docker",
    strip_prefix = "rules_docker-0.15.0",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.15.0/rules_docker-v0.15.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load("@io_bazel_rules_docker//go:image.bzl", _go_image_repos = "repositories")
_go_image_repos()






load("@io_bazel_rules_docker//container:pull.bzl", "container_pull")

container_pull(
    name = "alpine_linux_amd64",
    registry = "index.docker.io",
    repository = "library/golang",
    tag = "1.15-alpine",
)

