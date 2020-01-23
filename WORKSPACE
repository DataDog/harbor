workspace(name = "com_github_goharbor_harbor")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "9c3579a83b585d898c1d732950a5d2cfc5f6500b81cd9b351346e532a3084a18",
    strip_prefix = "rules_go-4c28450eae71d4179d946dfc930b9f13dee9a76c",
    urls = [
        "https://github.com/bazelbuild/rules_go/archive/4c28450eae71d4179d946dfc930b9f13dee9a76c.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

git_repository(
    name = "io_bazel_rules_docker",
    commit = "07a636fb8090df80c85f76ad1ca5d5bad403f0cf",
    remote = "https://github.com/bazelbuild/rules_docker",
    shallow_since = "1564591745 -0400",
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

git_repository(
    name = "com_google_protobuf",
    commit = "09745575a923640154bcf307fba8aedff47f240a",
    remote = "https://github.com/protocolbuffers/protobuf",
    shallow_since = "1558721209 -0700",
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

http_archive(
    name = "bazel_gazelle",
    sha256 = "be9296bfd64882e3c08e3283c58fcb461fa6dd3c171764fcc4cf322f60615a9b",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/0.18.1/bazel-gazelle-0.18.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.18.1/bazel-gazelle-0.18.1.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "co_honnef_go_tools",
    commit = "3f1c8253044a",
    importpath = "honnef.co/go/tools",
)

go_repository(
    name = "com_github_agl_ed25519",
    commit = "5312a6153412",
    importpath = "github.com/agl/ed25519",
)

go_repository(
    name = "com_github_alecthomas_template",
    commit = "a0175ee3bccc",
    importpath = "github.com/alecthomas/template",
)

go_repository(
    name = "com_github_alecthomas_units",
    commit = "2efee857e7cf",
    importpath = "github.com/alecthomas/units",
)

go_repository(
    name = "com_github_aliyun_alibaba_cloud_sdk_go",
    commit = "cd293c93fd97",
    importpath = "github.com/aliyun/alibaba-cloud-sdk-go",
)

go_repository(
    name = "com_github_aliyun_aliyun_oss_go_sdk",
    commit = "86c17b95fcd5",
    importpath = "github.com/aliyun/aliyun-oss-go-sdk",
)

go_repository(
    name = "com_github_apache_thrift",
    importpath = "github.com/apache/thrift",
    tag = "v0.12.0",
)

go_repository(
    name = "com_github_armon_consul_api",
    commit = "eb2c6b5be1b6",
    importpath = "github.com/armon/consul-api",
)

go_repository(
    name = "com_github_astaxie_beego",
    importpath = "github.com/astaxie/beego",
    tag = "v1.9.0",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    importpath = "github.com/aws/aws-sdk-go",
    tag = "v1.19.47",
)

go_repository(
    name = "com_github_baiyubin_aliyun_sts_go_sdk",
    commit = "cfa1a18b161f",
    importpath = "github.com/baiyubin/aliyun-sts-go-sdk",
)

go_repository(
    name = "com_github_beego_i18n",
    commit = "e87155e8f0c0",
    importpath = "github.com/beego/i18n",
)

go_repository(
    name = "com_github_beorn7_perks",
    importpath = "github.com/beorn7/perks",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_bitly_go_hostpool",
    commit = "a3a6125de932",
    importpath = "github.com/bitly/go-hostpool",
)

go_repository(
    name = "com_github_bitly_go_simplejson",
    importpath = "github.com/bitly/go-simplejson",
    tag = "v0.5.0",
)

go_repository(
    name = "com_github_bmatcuk_doublestar",
    importpath = "github.com/bmatcuk/doublestar",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_bmizerany_assert",
    commit = "b7ed37b82869",
    importpath = "github.com/bmizerany/assert",
)

go_repository(
    name = "com_github_bugsnag_bugsnag_go",
    importpath = "github.com/bugsnag/bugsnag-go",
    tag = "v1.5.2",
)

go_repository(
    name = "com_github_bugsnag_panicwrap",
    importpath = "github.com/bugsnag/panicwrap",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    tag = "v0.3.1",
)

go_repository(
    name = "com_github_casbin_casbin",
    importpath = "github.com/casbin/casbin",
    tag = "v1.7.0",
)

go_repository(
    name = "com_github_cenkalti_backoff",
    importpath = "github.com/cenkalti/backoff",
    tag = "v2.1.1",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    tag = "v0.3.4",
)

go_repository(
    name = "com_github_cloudflare_cfssl",
    commit = "9c027c93ba9e",
    importpath = "github.com/cloudflare/cfssl",
)

go_repository(
    name = "com_github_coreos_bbolt",
    importpath = "github.com/coreos/bbolt",
    tag = "v1.3.2",
)

go_repository(
    name = "com_github_coreos_etcd",
    importpath = "github.com/coreos/etcd",
    tag = "v3.3.10",
)

go_repository(
    name = "com_github_coreos_go_oidc",
    importpath = "github.com/coreos/go-oidc",
    tag = "v2.0.0",
)

go_repository(
    name = "com_github_coreos_go_semver",
    importpath = "github.com/coreos/go-semver",
    tag = "v0.2.0",
)

go_repository(
    name = "com_github_coreos_go_systemd",
    commit = "95778dfbb74e",
    importpath = "github.com/coreos/go-systemd",
)

go_repository(
    name = "com_github_coreos_pkg",
    commit = "399ea9e2e55f",
    importpath = "github.com/coreos/pkg",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_denisenkom_go_mssqldb",
    commit = "731ef375ac02",
    importpath = "github.com/denisenkom/go-mssqldb",
)

go_repository(
    name = "com_github_dghubble_sling",
    importpath = "github.com/dghubble/sling",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_dgrijalva_jwt_go",
    importpath = "github.com/dgrijalva/jwt-go",
    tag = "v3.2.0",
)

go_repository(
    name = "com_github_docker_distribution",
    importpath = "github.com/docker/distribution",
    tag = "v2.7.1",
)

go_repository(
    name = "com_github_docker_docker",
    importpath = "github.com/docker/docker",
    tag = "v1.13.1",
)

go_repository(
    name = "com_github_docker_go",
    commit = "d30aec9fd63c",
    importpath = "github.com/docker/go",
)

go_repository(
    name = "com_github_docker_go_connections",
    importpath = "github.com/docker/go-connections",
    tag = "v0.4.0",
)

go_repository(
    name = "com_github_docker_go_metrics",
    commit = "b84716841b82",
    importpath = "github.com/docker/go-metrics",
)

go_repository(
    name = "com_github_docker_go_units",
    importpath = "github.com/docker/go-units",
    tag = "v0.4.0",
)

go_repository(
    name = "com_github_docker_libtrust",
    commit = "aabc10ec26b7",
    importpath = "github.com/docker/libtrust",
)

go_repository(
    name = "com_github_eapache_go_resiliency",
    importpath = "github.com/eapache/go-resiliency",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_eapache_go_xerial_snappy",
    commit = "776d5712da21",
    importpath = "github.com/eapache/go-xerial-snappy",
)

go_repository(
    name = "com_github_eapache_queue",
    importpath = "github.com/eapache/queue",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_erikstmartin_go_testdb",
    commit = "8d10e4a1bae5",
    importpath = "github.com/erikstmartin/go-testdb",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    tag = "v1.4.7",
)

go_repository(
    name = "com_github_garyburd_redigo",
    importpath = "github.com/garyburd/redigo",
    tag = "v1.6.0",
)

go_repository(
    name = "com_github_ghodss_yaml",
    importpath = "github.com/ghodss/yaml",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_go_kit_kit",
    importpath = "github.com/go-kit/kit",
    tag = "v0.8.0",
)

go_repository(
    name = "com_github_go_logfmt_logfmt",
    importpath = "github.com/go-logfmt/logfmt",
    tag = "v0.4.0",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    importpath = "github.com/go-sql-driver/mysql",
    tag = "v1.4.1",
)

go_repository(
    name = "com_github_go_stack_stack",
    importpath = "github.com/go-stack/stack",
    tag = "v1.8.0",
)

go_repository(
    name = "com_github_gobwas_glob",
    importpath = "github.com/gobwas/glob",
    tag = "v0.2.3",
)

go_repository(
    name = "com_github_gocraft_work",
    importpath = "github.com/gocraft/work",
    tag = "v0.5.1",
)

go_repository(
    name = "com_github_gofrs_uuid",
    importpath = "github.com/gofrs/uuid",
    tag = "v3.2.0",
)

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    tag = "v1.2.1",
)

go_repository(
    name = "com_github_goji_httpauth",
    commit = "2da839ab0f4d",
    importpath = "github.com/goji/httpauth",
)

go_repository(
    name = "com_github_golang_glog",
    commit = "23def4e6c14b",
    importpath = "github.com/golang/glog",
)

go_repository(
    name = "com_github_golang_groupcache",
    commit = "5b532d6fd5ef",
    importpath = "github.com/golang/groupcache",
)

go_repository(
    name = "com_github_golang_migrate_migrate",
    importpath = "github.com/golang-migrate/migrate",
    tag = "v3.3.0",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    tag = "v1.3.1",
)

go_repository(
    name = "com_github_golang_snappy",
    commit = "2e65f85255db",
    importpath = "github.com/golang/snappy",
)

go_repository(
    name = "com_github_gomodule_redigo",
    importpath = "github.com/gomodule/redigo",
    tag = "v2.0.0",
)

go_repository(
    name = "com_github_google_btree",
    importpath = "github.com/google/btree",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_google_certificate_transparency_go",
    importpath = "github.com/google/certificate-transparency-go",
    tag = "v1.0.21",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    tag = "v0.2.0",
)

go_repository(
    name = "com_github_google_go_querystring",
    commit = "53e6ce116135",
    importpath = "github.com/google/go-querystring",
)

go_repository(
    name = "com_github_google_gofuzz",
    commit = "24818f796faf",
    importpath = "github.com/google/gofuzz",
)

go_repository(
    name = "com_github_google_martian",
    importpath = "github.com/google/martian",
    tag = "v2.1.0",
)

go_repository(
    name = "com_github_google_pprof",
    commit = "3ea8567a2e57",
    importpath = "github.com/google/pprof",
)

go_repository(
    name = "com_github_googleapis_gax_go_v2",
    importpath = "github.com/googleapis/gax-go/v2",
    tag = "v2.0.4",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    commit = "0766667cb4d1",
    importpath = "github.com/gopherjs/gopherjs",
)

go_repository(
    name = "com_github_gorilla_context",
    importpath = "github.com/gorilla/context",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_gorilla_handlers",
    importpath = "github.com/gorilla/handlers",
    tag = "v1.3.0",
)

go_repository(
    name = "com_github_gorilla_mux",
    importpath = "github.com/gorilla/mux",
    tag = "v1.6.2",
)

go_repository(
    name = "com_github_gorilla_websocket",
    importpath = "github.com/gorilla/websocket",
    tag = "v1.4.0",
)

go_repository(
    name = "com_github_graph_gophers_dataloader",
    importpath = "github.com/graph-gophers/dataloader",
    tag = "v5.0.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_middleware",
    importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_prometheus",
    importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    tag = "v1.9.0",
)

go_repository(
    name = "com_github_hailocab_go_hostpool",
    commit = "e80d13ce29ed",
    importpath = "github.com/hailocab/go-hostpool",
)

go_repository(
    name = "com_github_hashicorp_golang_lru",
    importpath = "github.com/hashicorp/golang-lru",
    tag = "v0.5.0",
)

go_repository(
    name = "com_github_hashicorp_hcl",
    importpath = "github.com/hashicorp/hcl",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_hpcloud_tail",
    importpath = "github.com/hpcloud/tail",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_jinzhu_gorm",
    importpath = "github.com/jinzhu/gorm",
    tag = "v1.9.8",
)

go_repository(
    name = "com_github_jinzhu_inflection",
    commit = "04140366298a",
    importpath = "github.com/jinzhu/inflection",
)

go_repository(
    name = "com_github_jinzhu_now",
    importpath = "github.com/jinzhu/now",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    commit = "c2b33e8439af",
    importpath = "github.com/jmespath/go-jmespath",
)

go_repository(
    name = "com_github_jonboulle_clockwork",
    importpath = "github.com/jonboulle/clockwork",
    tag = "v0.1.0",
)

go_repository(
    name = "com_github_json_iterator_go",
    importpath = "github.com/json-iterator/go",
    tag = "v1.1.6",
)

go_repository(
    name = "com_github_jstemmer_go_junit_report",
    commit = "af01ea7f8024",
    importpath = "github.com/jstemmer/go-junit-report",
)

go_repository(
    name = "com_github_jtolds_gls",
    importpath = "github.com/jtolds/gls",
    tag = "v4.20.0",
)

go_repository(
    name = "com_github_julienschmidt_httprouter",
    importpath = "github.com/julienschmidt/httprouter",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_justinas_alice",
    commit = "03f45bd4b7da",
    importpath = "github.com/justinas/alice",
)

go_repository(
    name = "com_github_kardianos_osext",
    commit = "2bc1f35cddc0",
    importpath = "github.com/kardianos/osext",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    importpath = "github.com/kisielk/errcheck",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_knetic_govaluate",
    importpath = "github.com/Knetic/govaluate",
    tag = "v3.0.0",
)

go_repository(
    name = "com_github_konsorten_go_windows_terminal_sequences",
    importpath = "github.com/konsorten/go-windows-terminal-sequences",
    tag = "v1.0.2",
)

go_repository(
    name = "com_github_kr_logfmt",
    commit = "b84e30acd515",
    importpath = "github.com/kr/logfmt",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    tag = "v0.1.0",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    tag = "v0.1.0",
)

go_repository(
    name = "com_github_lib_pq",
    importpath = "github.com/lib/pq",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_magiconair_properties",
    importpath = "github.com/magiconair/properties",
    tag = "v1.8.0",
)

go_repository(
    name = "com_github_masterminds_semver",
    importpath = "github.com/Masterminds/semver",
    tag = "v1.4.2",
)

go_repository(
    name = "com_github_mattn_go_runewidth",
    importpath = "github.com/mattn/go-runewidth",
    tag = "v0.0.4",
)

go_repository(
    name = "com_github_mattn_go_sqlite3",
    importpath = "github.com/mattn/go-sqlite3",
    tag = "v1.10.0",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
    tag = "v1.0.1",
)

go_repository(
    name = "com_github_microsoft_go_winio",
    importpath = "github.com/Microsoft/go-winio",
    tag = "v0.4.12",
)

go_repository(
    name = "com_github_miekg_pkcs11",
    commit = "7283ca79f35e",
    importpath = "github.com/miekg/pkcs11",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    importpath = "github.com/mitchellh/mapstructure",
    tag = "v1.1.2",
)

go_repository(
    name = "com_github_modern_go_concurrent",
    commit = "bacd9c7ef1dd",
    importpath = "github.com/modern-go/concurrent",
)

go_repository(
    name = "com_github_modern_go_reflect2",
    importpath = "github.com/modern-go/reflect2",
    tag = "v1.0.1",
)

go_repository(
    name = "com_github_mwitkow_go_conntrack",
    commit = "cc309e4a2223",
    importpath = "github.com/mwitkow/go-conntrack",
)

go_repository(
    name = "com_github_olekukonko_tablewriter",
    importpath = "github.com/olekukonko/tablewriter",
    tag = "v0.0.1",
)

go_repository(
    name = "com_github_onsi_ginkgo",
    importpath = "github.com/onsi/ginkgo",
    tag = "v1.7.0",
)

go_repository(
    name = "com_github_onsi_gomega",
    importpath = "github.com/onsi/gomega",
    tag = "v1.4.3",
)

go_repository(
    name = "com_github_opencontainers_go_digest",
    importpath = "github.com/opencontainers/go-digest",
    tag = "v1.0.0-rc0",
)

go_repository(
    name = "com_github_opencontainers_image_spec",
    importpath = "github.com/opencontainers/image-spec",
    tag = "v1.0.1",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    importpath = "github.com/opentracing/opentracing-go",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_openzipkin_zipkin_go",
    importpath = "github.com/openzipkin/zipkin-go",
    tag = "v0.1.6",
)

go_repository(
    name = "com_github_pelletier_go_toml",
    importpath = "github.com/pelletier/go-toml",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_pierrec_lz4",
    importpath = "github.com/pierrec/lz4",
    tag = "v2.0.5",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    tag = "v0.8.1",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_pquerna_cachecontrol",
    commit = "1555304b9b35",
    importpath = "github.com/pquerna/cachecontrol",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    importpath = "github.com/prometheus/client_golang",
    tag = "v0.9.4",
)

go_repository(
    name = "com_github_prometheus_client_model",
    commit = "fd36f4220a90",
    importpath = "github.com/prometheus/client_model",
)

go_repository(
    name = "com_github_prometheus_common",
    importpath = "github.com/prometheus/common",
    tag = "v0.4.1",
)

go_repository(
    name = "com_github_prometheus_procfs",
    importpath = "github.com/prometheus/procfs",
    tag = "v0.0.2",
)

go_repository(
    name = "com_github_rcrowley_go_metrics",
    commit = "3113b8401b8a",
    importpath = "github.com/rcrowley/go-metrics",
)

go_repository(
    name = "com_github_robfig_cron",
    importpath = "github.com/robfig/cron",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_rogpeppe_fastuuid",
    commit = "6724a57986af",
    importpath = "github.com/rogpeppe/fastuuid",
)

go_repository(
    name = "com_github_satori_go_uuid",
    importpath = "github.com/satori/go.uuid",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_shopify_logrus_bugsnag",
    commit = "577dee27f20d",
    importpath = "github.com/Shopify/logrus-bugsnag",
)

go_repository(
    name = "com_github_shopify_sarama",
    importpath = "github.com/Shopify/sarama",
    tag = "v1.19.0",
)

go_repository(
    name = "com_github_shopify_toxiproxy",
    importpath = "github.com/Shopify/toxiproxy",
    tag = "v2.1.4",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    tag = "v1.4.1",
)

go_repository(
    name = "com_github_smartystreets_assertions",
    commit = "b2de0cb4f26d",
    importpath = "github.com/smartystreets/assertions",
)

go_repository(
    name = "com_github_smartystreets_goconvey",
    commit = "68dc04aab96a",
    importpath = "github.com/smartystreets/goconvey",
)

go_repository(
    name = "com_github_soheilhy_cmux",
    importpath = "github.com/soheilhy/cmux",
    tag = "v0.1.4",
)

go_repository(
    name = "com_github_spf13_afero",
    importpath = "github.com/spf13/afero",
    tag = "v1.1.2",
)

go_repository(
    name = "com_github_spf13_cast",
    importpath = "github.com/spf13/cast",
    tag = "v1.3.0",
)

go_repository(
    name = "com_github_spf13_jwalterweatherman",
    importpath = "github.com/spf13/jwalterweatherman",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_spf13_pflag",
    importpath = "github.com/spf13/pflag",
    tag = "v1.0.3",
)

go_repository(
    name = "com_github_spf13_viper",
    importpath = "github.com/spf13/viper",
    tag = "v1.4.0",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    tag = "v0.1.1",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    tag = "v1.3.0",
)

go_repository(
    name = "com_github_theupdateframework_notary",
    importpath = "github.com/theupdateframework/notary",
    tag = "v0.6.1",
)

go_repository(
    name = "com_github_tmc_grpc_websocket_proxy",
    commit = "0ad062ec5ee5",
    importpath = "github.com/tmc/grpc-websocket-proxy",
)

go_repository(
    name = "com_github_ugorji_go",
    importpath = "github.com/ugorji/go",
    tag = "v1.1.4",
)

go_repository(
    name = "com_github_unknwon_goconfig",
    commit = "5f601ca6ef4d",
    importpath = "github.com/Unknwon/goconfig",
)

go_repository(
    name = "com_github_xiang90_probing",
    commit = "43a291ad63a2",
    importpath = "github.com/xiang90/probing",
)

go_repository(
    name = "com_github_xordataexchange_crypt",
    commit = "b2862e3d0a77",
    importpath = "github.com/xordataexchange/crypt",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    tag = "v0.37.4",
)

go_repository(
    name = "in_gopkg_alecthomas_kingpin_v2",
    importpath = "gopkg.in/alecthomas/kingpin.v2",
    tag = "v2.2.6",
)

go_repository(
    name = "in_gopkg_asn1_ber_v1",
    commit = "4e86f4367175",
    importpath = "gopkg.in/asn1-ber.v1",
)

go_repository(
    name = "in_gopkg_check_v1",
    commit = "788fd7840127",
    importpath = "gopkg.in/check.v1",
)

go_repository(
    name = "in_gopkg_dancannon_gorethink_v3",
    importpath = "gopkg.in/dancannon/gorethink.v3",
    tag = "v3.0.5",
)

go_repository(
    name = "in_gopkg_fatih_pool_v2",
    importpath = "gopkg.in/fatih/pool.v2",
    tag = "v2.0.0",
)

go_repository(
    name = "in_gopkg_fsnotify_v1",
    importpath = "gopkg.in/fsnotify.v1",
    tag = "v1.4.7",
)

go_repository(
    name = "in_gopkg_gorethink_gorethink_v3",
    importpath = "gopkg.in/gorethink/gorethink.v3",
    tag = "v3.0.5",
)

go_repository(
    name = "in_gopkg_inf_v0",
    importpath = "gopkg.in/inf.v0",
    tag = "v0.9.1",
)

go_repository(
    name = "in_gopkg_ini_v1",
    importpath = "gopkg.in/ini.v1",
    tag = "v1.42.0",
)

go_repository(
    name = "in_gopkg_ldap_v2",
    importpath = "gopkg.in/ldap.v2",
    tag = "v2.5.0",
)

go_repository(
    name = "in_gopkg_resty_v1",
    importpath = "gopkg.in/resty.v1",
    tag = "v1.12.0",
)

go_repository(
    name = "in_gopkg_square_go_jose_v2",
    importpath = "gopkg.in/square/go-jose.v2",
    tag = "v2.3.0",
)

go_repository(
    name = "in_gopkg_tomb_v1",
    commit = "dd632973f1e7",
    importpath = "gopkg.in/tomb.v1",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    tag = "v2.2.2",
)

go_repository(
    name = "io_etcd_go_bbolt",
    importpath = "go.etcd.io/bbolt",
    tag = "v1.3.2",
)

go_repository(
    name = "io_k8s_api",
    commit = "5cb15d344471",
    importpath = "k8s.io/api",
)

go_repository(
    name = "io_k8s_apimachinery",
    commit = "f534d624797b",
    importpath = "k8s.io/apimachinery",
)

go_repository(
    name = "io_k8s_client_go",
    importpath = "k8s.io/client-go",
    tag = "v8.0.0",
)

go_repository(
    name = "io_k8s_helm",
    importpath = "k8s.io/helm",
    tag = "v2.9.1",
)

go_repository(
    name = "io_opencensus_go",
    importpath = "go.opencensus.io",
    tag = "v0.20.1",
)

go_repository(
    name = "org_golang_google_api",
    importpath = "google.golang.org/api",
    tag = "v0.3.1",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    tag = "v1.4.0",
)

go_repository(
    name = "org_golang_google_genproto",
    commit = "64821d5d2107",
    importpath = "google.golang.org/genproto",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    tag = "v1.21.0",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "a5d413f7728c",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_exp",
    commit = "509febef88a4",
    importpath = "golang.org/x/exp",
)

go_repository(
    name = "org_golang_x_lint",
    commit = "d0100b6bd8b3",
    importpath = "golang.org/x/lint",
)

go_repository(
    name = "org_golang_x_net",
    commit = "f3200d17e092",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_oauth2",
    commit = "e64efc72b421",
    importpath = "golang.org/x/oauth2",
)

go_repository(
    name = "org_golang_x_sync",
    commit = "e225da77a7e6",
    importpath = "golang.org/x/sync",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "d0b11bdaac8a",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "org_golang_x_text",
    commit = "17ff2d5776d2",
    importpath = "golang.org/x/text",
)

go_repository(
    name = "org_golang_x_time",
    commit = "9d24e82272b4",
    importpath = "golang.org/x/time",
)

go_repository(
    name = "org_golang_x_tools",
    commit = "ab21143f2384",
    importpath = "golang.org/x/tools",
)

go_repository(
    name = "org_uber_go_atomic",
    importpath = "go.uber.org/atomic",
    tag = "v1.4.0",
)

go_repository(
    name = "org_uber_go_multierr",
    importpath = "go.uber.org/multierr",
    tag = "v1.1.0",
)

go_repository(
    name = "org_uber_go_zap",
    importpath = "go.uber.org/zap",
    tag = "v1.10.0",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    tag = "v1.1.1",
)
