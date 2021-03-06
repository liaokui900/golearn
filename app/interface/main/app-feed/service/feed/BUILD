package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_test",
    "go_library",
)

go_test(
    name = "go_default_test",
    srcs = [
        "audit_test.go",
        "banner_test.go",
        "black_test.go",
        "cache_test.go",
        "index2_test.go",
        "index_test.go",
        "infoc_test.go",
        "operate_test.go",
        "player_test.go",
        "service_test.go",
        "tab_test.go",
        "upper_test.go",
    ],
    embed = [":go_default_library"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-card/model/card/ai:go_default_library",
        "//app/interface/main/app-card/model/card/live:go_default_library",
        "//app/interface/main/app-card/model/card/operate:go_default_library",
        "//app/interface/main/app-feed/conf:go_default_library",
        "//app/service/main/archive/model/archive:go_default_library",
        "//vendor/github.com/smartystreets/goconvey/convey:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    srcs = [
        "audit.go",
        "banner.go",
        "black.go",
        "cache.go",
        "index.go",
        "index2.go",
        "infoc.go",
        "operate.go",
        "player.go",
        "rank.go",
        "search.go",
        "service.go",
        "tab.go",
        "tab2.go",
        "upper.go",
    ],
    importpath = "go-common/app/interface/main/app-feed/service/feed",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-card/model:go_default_library",
        "//app/interface/main/app-card/model/bplus:go_default_library",
        "//app/interface/main/app-card/model/card:go_default_library",
        "//app/interface/main/app-card/model/card/ai:go_default_library",
        "//app/interface/main/app-card/model/card/audio:go_default_library",
        "//app/interface/main/app-card/model/card/bangumi:go_default_library",
        "//app/interface/main/app-card/model/card/banner:go_default_library",
        "//app/interface/main/app-card/model/card/cm:go_default_library",
        "//app/interface/main/app-card/model/card/live:go_default_library",
        "//app/interface/main/app-card/model/card/operate:go_default_library",
        "//app/interface/main/app-card/model/card/rank:go_default_library",
        "//app/interface/main/app-card/model/card/show:go_default_library",
        "//app/interface/main/app-feed/conf:go_default_library",
        "//app/interface/main/app-feed/dao/account:go_default_library",
        "//app/interface/main/app-feed/dao/ad:go_default_library",
        "//app/interface/main/app-feed/dao/archive:go_default_library",
        "//app/interface/main/app-feed/dao/article:go_default_library",
        "//app/interface/main/app-feed/dao/audio:go_default_library",
        "//app/interface/main/app-feed/dao/audit:go_default_library",
        "//app/interface/main/app-feed/dao/bangumi:go_default_library",
        "//app/interface/main/app-feed/dao/black:go_default_library",
        "//app/interface/main/app-feed/dao/bplus:go_default_library",
        "//app/interface/main/app-feed/dao/card:go_default_library",
        "//app/interface/main/app-feed/dao/converge:go_default_library",
        "//app/interface/main/app-feed/dao/game:go_default_library",
        "//app/interface/main/app-feed/dao/live:go_default_library",
        "//app/interface/main/app-feed/dao/location:go_default_library",
        "//app/interface/main/app-feed/dao/rank:go_default_library",
        "//app/interface/main/app-feed/dao/recommend:go_default_library",
        "//app/interface/main/app-feed/dao/relation:go_default_library",
        "//app/interface/main/app-feed/dao/resource:go_default_library",
        "//app/interface/main/app-feed/dao/search:go_default_library",
        "//app/interface/main/app-feed/dao/show:go_default_library",
        "//app/interface/main/app-feed/dao/special:go_default_library",
        "//app/interface/main/app-feed/dao/tab:go_default_library",
        "//app/interface/main/app-feed/dao/tag:go_default_library",
        "//app/interface/main/app-feed/dao/upper:go_default_library",
        "//app/interface/main/app-feed/model:go_default_library",
        "//app/interface/main/app-feed/model/feed:go_default_library",
        "//app/interface/main/app-feed/model/live:go_default_library",
        "//app/interface/main/app-feed/model/tag:go_default_library",
        "//app/interface/main/tag/model:go_default_library",
        "//app/interface/openplatform/article/model:go_default_library",
        "//app/service/main/account/model:go_default_library",
        "//app/service/main/archive/api:go_default_library",
        "//app/service/main/archive/model/archive:go_default_library",
        "//app/service/main/feed/model:go_default_library",
        "//app/service/main/location/model:go_default_library",
        "//app/service/main/relation/model:go_default_library",
        "//app/service/main/resource/model:go_default_library",
        "//app/service/openplatform/pgc-season/api/grpc/episode/v1:go_default_library",
        "//library/ecode:go_default_library",
        "//library/log:go_default_library",
        "//library/log/infoc:go_default_library",
        "//library/net/metadata:go_default_library",
        "//library/stat/prom:go_default_library",
        "//library/sync/errgroup:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
