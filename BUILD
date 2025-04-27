load("@gazelle//:def.bzl", "gazelle")
load("@rules_go//go:def.bzl", "go_library", "nogo")
load("@rules_pkg//pkg:tar.bzl", "pkg_tar")

gazelle(name = "gazelle")

#nogo(
#    name = "nogo_vet",
#    config = ":nogo_config.json",
#    vet = True,
#    visibility = ["//visibility:public"],
#)

pkg_tar(
    name = "release",
    srcs = [
        "//app/cli",
        "//app/cli-license",
    ],
)

go_library(
    name = "sqlseed",
    srcs = ["sqlseed.go"],
    importpath = "github.com/knusperleicht/sqlseed",
    visibility = ["//visibility:public"],
    deps = ["//lib/sqlite"],
)
