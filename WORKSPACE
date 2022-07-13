load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Download the Go rules.
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "69de5c704a05ff37862f7e0f5534d4f479418afc21806c887db544a316f3cb6b",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.27.0/rules_go-v0.27.0.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.27.0/rules_go-v0.27.0.tar.gz",
    ],
)

# Download Gazelle.
http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
    ],
)

# Load macros and repository rules.
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

# Currently, it seems this order matters (need to seed a definition of Sirupsen/sirupsen before
# gazelle auto-defines one by reference).  Experiment later with migrating back to sort-order
go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    #importpath = "github.com/Sirupsen/logrus",
    #replace = "github.com/sirupsen/logrus",
    sum = "h1:dJKuHgqk1NNQlqoA6BTlM1Wf9DOH3NBjQyu0h9+AZZE=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:nwc3DEeHmmLAfoZucVR881uASk0Mfjw8xYJ99tb5CcY=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_testcontainers_testcontainers_go",
    importpath = "github.com/testcontainers/testcontainers-go",
    sum = "h1:FiYsB83LSGbiawoV8TpAZGfcCUbtaeeg1SXqEKUxh08=",
    version = "v0.11.1",
)

go_repository(
    name = "in_gopkg_yaml_v3",
    importpath = "gopkg.in/yaml.v3",
    sum = "h1:dUUwHk2QECo/6vqA44rthZ8ie2QXMNeKRTHCNY2nXvo=",
    version = "v3.0.0-20200313102051-9f266ea9e77c",
)

go_repository(
    name = "co_honnef_go_tools",
    importpath = "honnef.co/go/tools",
    sum = "h1:sXmLre5bzIR6ypkjXCDI3jHPssRhc8KD/Ome589sc3U=",
    version = "v0.0.1-2020.1.3",
)

go_repository(
    name = "com_github_alecthomas_template",
    importpath = "github.com/alecthomas/template",
    sum = "h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=",
    version = "v0.0.0-20190718012654-fb15b899a751",
)

go_repository(
    name = "com_github_alecthomas_units",
    importpath = "github.com/alecthomas/units",
    sum = "h1:Hs82Z41s6SdL1CELW+XaDYmOH4hkBN4/N9og/AsOv7E=",
    version = "v0.0.0-20190717042225-c3de453c63f4",
)

go_repository(
    name = "com_github_alexflint_go_filemutex",
    importpath = "github.com/alexflint/go-filemutex",
    sum = "h1:AMzIhMUqU3jMrZiTuW0zkYeKlKDAFD+DG20IoO421/Y=",
    version = "v0.0.0-20171022225611-72bdc8eae2ae",
)

go_repository(
    name = "com_github_asaskevich_govalidator",
    importpath = "github.com/asaskevich/govalidator",
    sum = "h1:idn718Q4B6AGu/h5Sxe66HYVdqdGu2l9Iebqhi/AEoA=",
    version = "v0.0.0-20190424111038-f61b66f89f4a",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    importpath = "github.com/aws/aws-sdk-go",
    sum = "h1:m45+Ru/wA+73cOZXiEGLDH2d9uLN3iHqMc0/z4noDXE=",
    version = "v1.15.11",
)

go_repository(
    name = "com_github_azure_azure_sdk_for_go",
    importpath = "github.com/Azure/azure-sdk-for-go",
    sum = "h1:KnPIugL51v3N3WwvaSmZbxukD1WuWXOiE9fRdu32f2I=",
    version = "v16.2.1+incompatible",
)

go_repository(
    name = "com_github_azure_go_ansiterm",
    importpath = "github.com/Azure/go-ansiterm",
    sum = "h1:w+iIsaOQNcT7OZ575w+acHgRric5iCyQh+xv+KJ4HB8=",
    version = "v0.0.0-20170929234023-d6e3b3328b78",
)

go_repository(
    name = "com_github_azure_go_autorest",
    importpath = "github.com/Azure/go-autorest",
    sum = "h1:V5VMDjClD3GiElqLWO7mz2MxNAK/vTfRHdAubSIPRgs=",
    version = "v14.2.0+incompatible",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest",
    importpath = "github.com/Azure/go-autorest/autorest",
    sum = "h1:eVvIXUKiTgv++6YnWb42DUA1YL7qDugnKP0HljexdnQ=",
    version = "v0.11.1",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_adal",
    importpath = "github.com/Azure/go-autorest/autorest/adal",
    sum = "h1:Y3bBUV4rTuxenJJs41HU3qmqsb+auo+a3Lz+PlJPpL0=",
    version = "v0.9.5",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_date",
    importpath = "github.com/Azure/go-autorest/autorest/date",
    sum = "h1:7gUk1U5M/CQbp9WoqinNzJar+8KY+LPI6wiWrP/myHw=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_mocks",
    importpath = "github.com/Azure/go-autorest/autorest/mocks",
    sum = "h1:K0laFcLE6VLTOwNgSxaGbUcLPuGXlNkbVvq4cW4nIHk=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_azure_go_autorest_logger",
    importpath = "github.com/Azure/go-autorest/logger",
    sum = "h1:e4RVHVZKC5p6UANLJHkM4OfR1UKZPj8Wt8Pcx+3oqrE=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_azure_go_autorest_tracing",
    importpath = "github.com/Azure/go-autorest/tracing",
    sum = "h1:TYi4+3m5t6K48TGI9AUdb+IzbnSxvnvUMfuitfgcfuo=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_beorn7_perks",
    importpath = "github.com/beorn7/perks",
    sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_bgentry_speakeasy",
    importpath = "github.com/bgentry/speakeasy",
    sum = "h1:ByYyxL9InA1OWqxJqqp2A5pYHUrCiAL6K3J+LKSsQkY=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_bitly_go_simplejson",
    importpath = "github.com/bitly/go-simplejson",
    sum = "h1:6IH+V8/tVMab511d5bn4M7EwGXZf9Hj6i2xSwkNEM+Y=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_blang_semver",
    importpath = "github.com/blang/semver",
    sum = "h1:cQNTCjp13qL8KC3Nbxr/y2Bqb63oX6wdnnjpJbkM4JQ=",
    version = "v3.5.1+incompatible",
)

go_repository(
    name = "com_github_bmizerany_assert",
    importpath = "github.com/bmizerany/assert",
    sum = "h1:DDGfHa7BWjL4YnC6+E63dPcxHo2sUxDIu8g3QgEJdRY=",
    version = "v0.0.0-20160611221934-b7ed37b82869",
)

go_repository(
    name = "com_github_bshuster_repo_logrus_logstash_hook",
    importpath = "github.com/bshuster-repo/logrus-logstash-hook",
    sum = "h1:pgAtgj+A31JBVtEHu2uHuEx0n+2ukqUJnS2vVe5pQNA=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_buger_jsonparser",
    importpath = "github.com/buger/jsonparser",
    sum = "h1:y853v6rXx+zefEcjET3JuKAqvhj+FKflQijjeaSv2iA=",
    version = "v0.0.0-20180808090653-f4dd9f5a6b44",
)

go_repository(
    name = "com_github_bugsnag_bugsnag_go",
    importpath = "github.com/bugsnag/bugsnag-go",
    sum = "h1:rFt+Y/IK1aEZkEHchZRSq9OQbsSzIT/OrI8YFFmRIng=",
    version = "v0.0.0-20141110184014-b1d153021fcd",
)

go_repository(
    name = "com_github_bugsnag_osext",
    importpath = "github.com/bugsnag/osext",
    sum = "h1:otBG+dV+YK+Soembjv71DPz3uX/V/6MMlSyD9JBQ6kQ=",
    version = "v0.0.0-20130617224835-0dd3f918b21b",
)

go_repository(
    name = "com_github_bugsnag_panicwrap",
    importpath = "github.com/bugsnag/panicwrap",
    sum = "h1:nvj0OLI3YqYXer/kZD8Ri1aaunCxIEsOst1BVJswV0o=",
    version = "v0.0.0-20151223152923-e2c28503fcd0",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_burntsushi_xgb",
    importpath = "github.com/BurntSushi/xgb",
    sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
    version = "v0.0.0-20160522181843-27f122750802",
)

go_repository(
    name = "com_github_cenkalti_backoff",
    importpath = "github.com/cenkalti/backoff",
    sum = "h1:tNowT99t7UNflLxfYYSlKYsBpXdEet03Pg2g16Swow4=",
    version = "v2.2.1+incompatible",
)

go_repository(
    name = "com_github_census_instrumentation_opencensus_proto",
    importpath = "github.com/census-instrumentation/opencensus-proto",
    sum = "h1:glEXhBS5PSLLv4IXzLA5yPRVX4bilULVyxxbrfOtDAk=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_cespare_xxhash_v2",
    importpath = "github.com/cespare/xxhash/v2",
    sum = "h1:6MnRN8NT7+YBpUIWxHtefFZOKTAPgGjpQSxqLNn0+qY=",
    version = "v2.1.1",
)

go_repository(
    name = "com_github_checkpoint_restore_go_criu_v4",
    importpath = "github.com/checkpoint-restore/go-criu/v4",
    sum = "h1:WW2B2uxx9KWF6bGlHqhm8Okiafwwx7Y2kcpn8lCpjgo=",
    version = "v4.1.0",
)

go_repository(
    name = "com_github_chzyer_logex",
    importpath = "github.com/chzyer/logex",
    sum = "h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
    version = "v1.1.10",
)

go_repository(
    name = "com_github_chzyer_readline",
    importpath = "github.com/chzyer/readline",
    sum = "h1:fY5BOSpyZCqRo5OhCuC+XN+r/bBCmeuuJtjz+bCNIf8=",
    version = "v0.0.0-20180603132655-2972be24d48e",
)

go_repository(
    name = "com_github_chzyer_test",
    importpath = "github.com/chzyer/test",
    sum = "h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
    version = "v0.0.0-20180213035817-a1ea475d72b1",
)

go_repository(
    name = "com_github_cilium_ebpf",
    importpath = "github.com/cilium/ebpf",
    sum = "h1:Fv93L3KKckEcEHR3oApXVzyBTDA8WAm6VXhPE00N3f8=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
    version = "v0.3.4",
)

go_repository(
    name = "com_github_cncf_udpa_go",
    importpath = "github.com/cncf/udpa/go",
    sum = "h1:WBZRG4aNOuI15bLRrCgN8fCq8E5Xuty6jGbmSNEvSsU=",
    version = "v0.0.0-20191209042840-269d4d468f6f",
)

go_repository(
    name = "com_github_cockroachdb_datadriven",
    importpath = "github.com/cockroachdb/datadriven",
    sum = "h1:OaNxuTZr7kxeODyLWsRMC+OD03aFUH+mW6r2d+MWa5Y=",
    version = "v0.0.0-20190809214429-80d97fb3cbaa",
)

go_repository(
    name = "com_github_containerd_aufs",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/containerd/aufs",
    sum = "h1:eDnTE2gn5fNykNuQdUIhr5XHqLJR52FXCGxhWNHG4FM=",
    version = "v0.0.0-20210316121734-20793ff83c97",
)

go_repository(
    name = "com_github_containerd_btrfs",
    importpath = "github.com/containerd/btrfs",
    sum = "h1:jNaRXAhCLuQ2x3k4nZNxaZN27/BCMhLfVkCEokYPI5Q=",
    version = "v0.0.0-20210316141732-918d888fb676",
)

go_repository(
    name = "com_github_containerd_cgroups",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/containerd/cgroups",
    sum = "h1:hkGVFjz+plgr5UfxZUTPFbUFIF/Km6/s+RVRIRHLrrY=",
    version = "v0.0.0-20210114181951-8a68de567b68",
)

go_repository(
    name = "com_github_containerd_console",
    importpath = "github.com/containerd/console",
    sum = "h1:u7SFAJyRqWcG6ogaMAx3KjSTy1e3hT9QxqX7Jco7dRc=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_containerd_containerd",
    importpath = "github.com/containerd/containerd",
    sum = "h1:zjz4MOAOFgdBlwid2nNUlJ3YLpVi/97L36lfMYJex60=",
    version = "v1.5.0-beta.4",
)

go_repository(
    name = "com_github_containerd_continuity",
    importpath = "github.com/containerd/continuity",
    sum = "h1:6JKvHHt396/qabvMhnhUZvWaHZzfVfldxE60TK8YLhg=",
    version = "v0.0.0-20210208174643-50096c924a4e",
)

go_repository(
    name = "com_github_containerd_fifo",
    importpath = "github.com/containerd/fifo",
    sum = "h1:u6sWqdNGAy7+O8qG/r1dqdnZE7IdEjteK3WGuvbfreo=",
    version = "v0.0.0-20210316144830-115abcc95a1d",
)

go_repository(
    name = "com_github_containerd_go_cni",
    importpath = "github.com/containerd/go-cni",
    sum = "h1:VXr2EkOPD0v1gu7CKfof6XzEIDzsE/dI9yj/W7PSWLs=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_containerd_go_runc",
    importpath = "github.com/containerd/go-runc",
    sum = "h1:e+50zk22gvHLJKe8+d+xSMyA88PPQk/XfWuUw1BdnPA=",
    version = "v0.0.0-20201020171139-16b287bc67d0",
)

go_repository(
    name = "com_github_containerd_imgcrypt",
    importpath = "github.com/containerd/imgcrypt",
    sum = "h1:gjN37M4b3GvtvJyN6fnahr5A5fhj6PHLvzoNmdzoA74=",
    version = "v1.1.1-0.20210312161619-7ed62a527887",
)

go_repository(
    name = "com_github_containerd_nri",
    importpath = "github.com/containerd/nri",
    sum = "h1:ap4DriuHs+kzozfFpPH+qccXB219gbChW/F4KW8IGZk=",
    version = "v0.0.0-20210316161719-dbaa18c31c14",
)

go_repository(
    name = "com_github_containerd_ttrpc",
    importpath = "github.com/containerd/ttrpc",
    sum = "h1:2/O3oTZN36q2xRolk0a2WWGgh7/Vf/liElg5hFYLX9U=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_containerd_typeurl",
    importpath = "github.com/containerd/typeurl",
    sum = "h1:PvuK4E3D5S5q6IqsPDCy928FhP0LUIGcmZ/Yhgp5Djw=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_containerd_zfs",
    importpath = "github.com/containerd/zfs",
    sum = "h1:HXgv62Qod4KvnxeNWszVQ4ZR+NrVY9ODFxlaqKzDQMs=",
    version = "v0.0.0-20210315114300-dde8f0fda960",
)

go_repository(
    name = "com_github_containernetworking_cni",
    importpath = "github.com/containernetworking/cni",
    sum = "h1:BT9lpgGoH4jw3lFC7Odz2prU5ruiYKcgAjMCbgybcKI=",
    version = "v0.8.0",
)

go_repository(
    name = "com_github_containernetworking_plugins",
    importpath = "github.com/containernetworking/plugins",
    sum = "h1:npZTLiMa4CRn6m5P9+1Dz4O1j0UeFbm8VYN6dlsw568=",
    version = "v0.8.6",
)

go_repository(
    name = "com_github_containers_ocicrypt",
    importpath = "github.com/containers/ocicrypt",
    sum = "h1:A6UzSUFMla92uxO43O6lm86i7evMGjTY7wTKB2DyGPY=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_coreos_go_iptables",
    importpath = "github.com/coreos/go-iptables",
    sum = "h1:DpHb9vJrZQEFMcVLFKAAGMUVX0XoRC0ptCthinRYm38=",
    version = "v0.4.5",
)

go_repository(
    name = "com_github_coreos_go_oidc",
    importpath = "github.com/coreos/go-oidc",
    sum = "h1:sdJrfw8akMnCuUlaZU3tE/uYXFgfqom8DBE9so9EBsM=",
    version = "v2.1.0+incompatible",
)

go_repository(
    name = "com_github_coreos_go_semver",
    importpath = "github.com/coreos/go-semver",
    sum = "h1:wkHLiw0WNATZnSG7epLsujiMCgPAc9xhjJ4tgnAxmfM=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_coreos_go_systemd",
    importpath = "github.com/coreos/go-systemd",
    sum = "h1:Wf6HqHfScWJN9/ZjdUKyjop4mf3Qdd+1TvvltAvM3m8=",
    version = "v0.0.0-20190321100706-95778dfbb74e",
)

go_repository(
    name = "com_github_coreos_go_systemd_v22",
    importpath = "github.com/coreos/go-systemd/v22",
    sum = "h1:kq/SbG2BCKLkDKkjQf5OWwKWUKj1lgs3lFI4PxnR5lg=",
    version = "v22.1.0",
)

go_repository(
    name = "com_github_coreos_pkg",
    importpath = "github.com/coreos/pkg",
    sum = "h1:lBNOc5arjvs8E5mO2tbpBpLoyyu8B6e44T7hJy6potg=",
    version = "v0.0.0-20180928190104-399ea9e2e55f",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man_v2",
    importpath = "github.com/cpuguy83/go-md2man/v2",
    sum = "h1:EoUDS0afbrsXAZ9YQ9jdu/mZ2sXgT1/2yyNng4PGlyM=",
    version = "v2.0.0",
)

go_repository(
    name = "com_github_creack_pty",
    importpath = "github.com/creack/pty",
    sum = "h1:07n33Z8lZxZ2qwegKbObQohDhXDQxiMMz1NOUGYlesw=",
    version = "v1.1.11",
)

go_repository(
    name = "com_github_cyphar_filepath_securejoin",
    importpath = "github.com/cyphar/filepath-securejoin",
    sum = "h1:jCwT2GTP+PY5nBz3c/YL5PAIbusElVrPujOBSCj8xRg=",
    version = "v0.2.2",
)

go_repository(
    name = "com_github_d2g_dhcp4",
    importpath = "github.com/d2g/dhcp4",
    sum = "h1:Xo2rK1pzOm0jO6abTPIQwbAmqBIOj132otexc1mmzFc=",
    version = "v0.0.0-20170904100407-a1d1b6c41b1c",
)

go_repository(
    name = "com_github_d2g_dhcp4client",
    importpath = "github.com/d2g/dhcp4client",
    sum = "h1:suYBsYZIkSlUMEz4TAYCczKf62IA2UWC+O8+KtdOhCo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_d2g_dhcp4server",
    importpath = "github.com/d2g/dhcp4server",
    sum = "h1:+CpLbZIeUn94m02LdEKPcgErLJ347NUwxPKs5u8ieiY=",
    version = "v0.0.0-20181031114812-7d4a0a7f59a5",
)

go_repository(
    name = "com_github_d2g_hardwareaddr",
    importpath = "github.com/d2g/hardwareaddr",
    sum = "h1:itqmmf1PFpC4n5JW+j4BU7X4MTfVurhYRTjODoPb2Y8=",
    version = "v0.0.0-20190221164911-e7d9fbe030e4",
)

go_repository(
    name = "com_github_denverdino_aliyungo",
    importpath = "github.com/denverdino/aliyungo",
    sum = "h1:p6poVbjHDkKa+wtC8frBMwQtT3BmqGYBjzMwJ63tuR4=",
    version = "v0.0.0-20190125010748-a747050bb1ba",
)

go_repository(
    name = "com_github_dgrijalva_jwt_go",
    importpath = "github.com/dgrijalva/jwt-go",
    sum = "h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
    version = "v3.2.0+incompatible",
)

go_repository(
    name = "com_github_dnaeon_go_vcr",
    importpath = "github.com/dnaeon/go-vcr",
    sum = "h1:r8L/HqC0Hje5AXMu1ooW8oyQyOFv4GxqpL0nRP7SLLY=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_docker_distribution",
    importpath = "github.com/docker/distribution",
    sum = "h1:a5mlkVzth6W5A4fOsS3D2EO5BUmsJpcB+cRlLU7cSug=",
    version = "v2.7.1+incompatible",
)

go_repository(
    name = "com_github_docker_docker",
    importpath = "github.com/docker/docker",
    sum = "h1:RVqD337BgQicVCzYrrlhLDWhq6OAD2PJDUg2LsEUvKM=",
    version = "v20.10.8+incompatible",
)

go_repository(
    name = "com_github_docker_go_connections",
    importpath = "github.com/docker/go-connections",
    sum = "h1:El9xVISelRB7BuFusrZozjnkIM5YnzCViNKohAFqRJQ=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_docker_go_events",
    importpath = "github.com/docker/go-events",
    sum = "h1:+pKlWGMw7gf6bQ+oDZB4KHQFypsfjYlq/C4rfL7D3g8=",
    version = "v0.0.0-20190806004212-e31b211e4f1c",
)

go_repository(
    name = "com_github_docker_go_metrics",
    importpath = "github.com/docker/go-metrics",
    sum = "h1:AgB/0SvBxihN0X8OR4SjsblXkbMvalQ8cjmtKQ2rQV8=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_docker_go_units",
    importpath = "github.com/docker/go-units",
    sum = "h1:3uh0PgVws3nIA0Q+MwDC8yjEPf9zjRfZZWXZYDct3Tw=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_docker_libtrust",
    importpath = "github.com/docker/libtrust",
    sum = "h1:ZClxb8laGDf5arXfYcAtECDFgAgHklGI8CxgjHnXKJ4=",
    version = "v0.0.0-20150114040149-fa567046d9b1",
)

go_repository(
    name = "com_github_docker_spdystream",
    importpath = "github.com/docker/spdystream",
    sum = "h1:cenwrSVm+Z7QLSV/BsnenAOcDXdX4cMv4wP0B/5QbPg=",
    version = "v0.0.0-20160310174837-449fdfce4d96",
)

go_repository(
    name = "com_github_docopt_docopt_go",
    importpath = "github.com/docopt/docopt-go",
    sum = "h1:bWDMxwH3px2JBh6AyO7hdCn/PkvCZXii8TGj7sbtEbQ=",
    version = "v0.0.0-20180111231733-ee0de3bc6815",
)

go_repository(
    name = "com_github_dustin_go_humanize",
    importpath = "github.com/dustin/go-humanize",
    sum = "h1:VSnTsYCnlFHaM2/igO1h6X3HA71jcobQuxemgkq4zYo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_elazarl_goproxy",
    importpath = "github.com/elazarl/goproxy",
    sum = "h1:yUdfgN0XgIJw7foRItutHYUIhlcKzcSf5vDpdhQAKTc=",
    version = "v0.0.0-20180725130230-947c36da3153",
)

go_repository(
    name = "com_github_emicklei_go_restful",
    importpath = "github.com/emicklei/go-restful",
    sum = "h1:spTtZBk5DYEvbxMVutUuTyh1Ao2r4iyvLdACqsl/Ljk=",
    version = "v2.9.5+incompatible",
)

go_repository(
    name = "com_github_envoyproxy_go_control_plane",
    importpath = "github.com/envoyproxy/go-control-plane",
    sum = "h1:rEvIZUSZ3fx39WIi3JkQqQBitGwpELBIYWeBVh6wn+E=",
    version = "v0.9.4",
)

go_repository(
    name = "com_github_envoyproxy_protoc_gen_validate",
    importpath = "github.com/envoyproxy/protoc-gen-validate",
    sum = "h1:EQciDnbrYxy13PgWoY8AqoxGiPrpgBZ1R8UNe3ddc+A=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_evanphx_json_patch",
    importpath = "github.com/evanphx/json-patch",
    sum = "h1:kLcOMZeuLAJvL2BPWLMIj5oaZQobrkAqrL+WFZwQses=",
    version = "v4.9.0+incompatible",
)

go_repository(
    name = "com_github_fatih_color",
    importpath = "github.com/fatih/color",
    sum = "h1:DkWD4oS2D8LGGgTQ6IvwJJXSL5Vp2ffcQg58nFV38Ys=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_form3tech_oss_jwt_go",
    importpath = "github.com/form3tech-oss/jwt-go",
    sum = "h1:TcekIExNqud5crz4xD2pavyTgWiPvpYe4Xau31I0PRk=",
    version = "v3.2.2+incompatible",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    sum = "h1:hsms1Qyu0jgnwNXIxa+/V/PDsU6CfLf6CNO8H7IWoS4=",
    version = "v1.4.9",
)

go_repository(
    name = "com_github_fullsailor_pkcs7",
    importpath = "github.com/fullsailor/pkcs7",
    sum = "h1:RDBNVkRviHZtvDvId8XSGPu3rmpmSe+wKRcEWNgsfWU=",
    version = "v0.0.0-20190404230743-d7302db945fa",
)

go_repository(
    name = "com_github_garyburd_redigo",
    importpath = "github.com/garyburd/redigo",
    sum = "h1:LofdAjjjqCSXMwLGgOgnE+rdPuvX9DxCqaHwKy7i/ko=",
    version = "v0.0.0-20150301180006-535138d7bcd7",
)

go_repository(
    name = "com_github_ghodss_yaml",
    importpath = "github.com/ghodss/yaml",
    sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_gl_glfw",
    importpath = "github.com/go-gl/glfw",
    sum = "h1:QbL/5oDUmRBzO9/Z7Seo6zf912W/a6Sr4Eu0G/3Jho0=",
    version = "v0.0.0-20190409004039-e6da0acd62b1",
)

go_repository(
    name = "com_github_go_gl_glfw_v3_3_glfw",
    importpath = "github.com/go-gl/glfw/v3.3/glfw",
    sum = "h1:WtGNWLvXpe6ZudgnXrq0barxBImvnnJoMEhXAzcbM0I=",
    version = "v0.0.0-20200222043503-6f7a984d4dc4",
)

go_repository(
    name = "com_github_go_ini_ini",
    importpath = "github.com/go-ini/ini",
    sum = "h1:Mujh4R/dH6YL8bxuISne3xX2+qcQ9p0IxKAP6ExWoUo=",
    version = "v1.25.4",
)

go_repository(
    name = "com_github_go_kit_kit",
    importpath = "github.com/go-kit/kit",
    sum = "h1:wDJmvq38kDhkVxi50ni9ykkdUr1PKgqKOoi01fa0Mdk=",
    version = "v0.9.0",
)

go_repository(
    name = "com_github_go_logfmt_logfmt",
    importpath = "github.com/go-logfmt/logfmt",
    sum = "h1:MP4Eh7ZCb31lleYCFuwm0oe4/YGak+5l1vA2NOE80nA=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_go_logr_logr",
    importpath = "github.com/go-logr/logr",
    sum = "h1:QvGt2nLcHH0WK9orKa+ppBPAxREcH364nPUedEpK0TY=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_go_openapi_jsonpointer",
    importpath = "github.com/go-openapi/jsonpointer",
    sum = "h1:gihV7YNZK1iK6Tgwwsxo2rJbD1GTbdm72325Bq8FI3w=",
    version = "v0.19.3",
)

go_repository(
    name = "com_github_go_openapi_jsonreference",
    importpath = "github.com/go-openapi/jsonreference",
    sum = "h1:5cxNfTy0UVC3X8JL5ymxzyoUZmo8iZb+jeTWn7tUa8o=",
    version = "v0.19.3",
)

go_repository(
    name = "com_github_go_openapi_spec",
    importpath = "github.com/go-openapi/spec",
    sum = "h1:0XRyw8kguri6Yw4SxhsQA/atC88yqrk0+G4YhI2wabc=",
    version = "v0.19.3",
)

go_repository(
    name = "com_github_go_openapi_swag",
    importpath = "github.com/go-openapi/swag",
    sum = "h1:lTz6Ys4CmqqCQmZPBlbQENR1/GucA2bzYTE12Pw4tFY=",
    version = "v0.19.5",
)

go_repository(
    name = "com_github_go_redis_redis",
    importpath = "github.com/go-redis/redis",
    sum = "h1:K0pv1D7EQUjfyoMql+r/jZqCLizCGKFlFgcHWWmHQjg=",
    version = "v6.15.9+incompatible",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    importpath = "github.com/go-sql-driver/mysql",
    sum = "h1:BCTh4TKNUYmOmMUcQ3IipzF5prigylS7XXjEkfCHuOE=",
    version = "v1.6.0",
)

go_repository(
    name = "com_github_go_stack_stack",
    importpath = "github.com/go-stack/stack",
    sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_godbus_dbus",
    importpath = "github.com/godbus/dbus",
    sum = "h1:BWhy2j3IXJhjCbC68FptL43tDKIq8FladmaTs3Xs7Z8=",
    version = "v0.0.0-20190422162347-ade71ed3457e",
)

go_repository(
    name = "com_github_godbus_dbus_v5",
    importpath = "github.com/godbus/dbus/v5",
    sum = "h1:ZqHaoEF7TBzh4jzPmqVhE/5A1z9of6orkAe5uHoAeME=",
    version = "v5.0.3",
)

go_repository(
    name = "com_github_gogo_googleapis",
    importpath = "github.com/gogo/googleapis",
    sum = "h1:zgVt4UpGxcqVOw97aRGxT4svlcmdK35fynLNctY32zI=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_gogo_protobuf",
    build_directives = [
        "gazelle:exclude **/testdata/**",
        "gazelle:exclude test/**",
        "gazelle:exclude conformance/**",
    ],
    importpath = "github.com/gogo/protobuf",
    sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_golang_glog",
    importpath = "github.com/golang/glog",
    sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
    version = "v0.0.0-20160126235308-23def4e6c14b",
)

go_repository(
    name = "com_github_golang_groupcache",
    importpath = "github.com/golang/groupcache",
    sum = "h1:1r7pUrabqp18hOBcwBwiTsbnFeTZHV9eER/QT5JVZxY=",
    version = "v0.0.0-20200121045136-8c9f03a8e57e",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    sum = "h1:ocYkMQY5RrXTYgXl7ICpV0IXwlEQGwKIsery4gyXa1U=",
    version = "v1.4.1",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    sum = "h1:JjCZWpVbqXDqFVmTfYWEVTMIYrL/NPdPSCHPJ0T/raM=",
    version = "v1.4.3",
)

go_repository(
    name = "com_github_google_btree",
    importpath = "github.com/google/btree",
    sum = "h1:0udJVsspx3VBr5FwtLhQQtuAsVc79tTq0ocGIPAU6qo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:L8R9j+yAqZuZjsqh/z+F1NCffTKKLShY6zXTItVIZ8M=",
    version = "v0.5.4",
)

go_repository(
    name = "com_github_google_gofuzz",
    importpath = "github.com/google/gofuzz",
    sum = "h1:Hsa8mG0dQ46ij8Sl2AYJDUv1oA9/d6Vk+3LG99Oe02g=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_google_martian",
    importpath = "github.com/google/martian",
    sum = "h1:/CP5g8u/VJHijgedC/Legn3BAbAaWPgecwXBIDzw5no=",
    version = "v2.1.0+incompatible",
)

go_repository(
    name = "com_github_google_pprof",
    importpath = "github.com/google/pprof",
    sum = "h1:SRgJV+IoxM5MKyFdlSUeNy6/ycRUF2yBAKdAQswoHUk=",
    version = "v0.0.0-20200229191704-1ebb73c60ed3",
)

go_repository(
    name = "com_github_google_renameio",
    importpath = "github.com/google/renameio",
    sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    sum = "h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_googleapis_gax_go_v2",
    importpath = "github.com/googleapis/gax-go/v2",
    sum = "h1:sjZBwGj9Jlw33ImPtvFviGYvseOtDM7hkSKB7+Tv3SM=",
    version = "v2.0.5",
)

go_repository(
    name = "com_github_googleapis_gnostic",
    importpath = "github.com/googleapis/gnostic",
    sum = "h1:DLJCy1n/vrD4HPjOvYcT8aYQXpPIzoRZONaYwyycI+I=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    importpath = "github.com/gopherjs/gopherjs",
    sum = "h1:EGx4pi6eqNxGaHF6qqu48+N2wcFQ5qg5FXgOdqsJ5d8=",
    version = "v0.0.0-20181017120253-0766667cb4d1",
)

go_repository(
    name = "com_github_gorilla_handlers",
    importpath = "github.com/gorilla/handlers",
    sum = "h1:893HsJqtxp9z1SF76gg6hY70hRY1wVlTSnC/h1yUDCo=",
    version = "v0.0.0-20150720190736-60c7bfde3e33",
)

go_repository(
    name = "com_github_gorilla_mux",
    importpath = "github.com/gorilla/mux",
    sum = "h1:zoNxOV7WjqXptQOVngLmcSQgXmgk4NMz1HibBchjl/I=",
    version = "v1.7.2",
)

go_repository(
    name = "com_github_gorilla_websocket",
    importpath = "github.com/gorilla/websocket",
    sum = "h1:+/TMaTYc4QFitKJxsQ7Yye35DkWvkdLcvGKqM+x0Ufc=",
    version = "v1.4.2",
)

go_repository(
    name = "com_github_gregjones_httpcache",
    importpath = "github.com/gregjones/httpcache",
    sum = "h1:pdN6V1QBWetyv/0+wjACpqVH+eVULgEjkurDLq3goeM=",
    version = "v0.0.0-20180305231024-9cad4c3443a7",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_middleware",
    importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
    sum = "h1:z53tR0945TRRQO/fLEVPI6SMv7ZflF0TEaTAoU7tOzg=",
    version = "v1.0.1-0.20190118093823-f849b5445de4",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_prometheus",
    importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
    sum = "h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    sum = "h1:UImYN5qQ8tuGpGE16ZmjvcTtTw24zw1QAp/SlnNrZhI=",
    version = "v1.9.5",
)

go_repository(
    name = "com_github_hashicorp_errwrap",
    importpath = "github.com/hashicorp/errwrap",
    sum = "h1:hLrqtEDnRye3+sgx6z4qVLNuviH3MR5aQ0ykNJa/UYA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_multierror",
    importpath = "github.com/hashicorp/go-multierror",
    sum = "h1:iVjPR7a6H0tWELX5NxNe7bYopibicUzc7uPribsnS6o=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_golang_lru",
    importpath = "github.com/hashicorp/golang-lru",
    sum = "h1:0hERBMJE1eitiLkihrMvRVBYAkpHzc/J3QdDN+dAcgU=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_hpcloud_tail",
    importpath = "github.com/hpcloud/tail",
    sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_ianlancetaylor_demangle",
    importpath = "github.com/ianlancetaylor/demangle",
    sum = "h1:UDMh68UUwekSh5iP2OMhRRZJiiBccgV7axzUG8vi56c=",
    version = "v0.0.0-20181102032728-5e5cf60278f6",
)

go_repository(
    name = "com_github_imdario_mergo",
    importpath = "github.com/imdario/mergo",
    sum = "h1:3tnifQM4i+fbajXKBHXWEH+KvNHqojZ778UH75j3bGA=",
    version = "v0.3.11",
)

go_repository(
    name = "com_github_inconshreveable_mousetrap",
    importpath = "github.com/inconshreveable/mousetrap",
    sum = "h1:Z8tu5sraLXCXIcARxBp/8cbvlwVa7Z1NHg9XEKhtSvM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_j_keck_arping",
    importpath = "github.com/j-keck/arping",
    sum = "h1:742eGXur0715JMq73aD95/FU0XpVKXqNuTnEfXsLOYQ=",
    version = "v0.0.0-20160618110441-2cf9dc699c56",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    importpath = "github.com/jmespath/go-jmespath",
    sum = "h1:SMvOWPJCES2GdFracYbBQh93GXac8fq7HeN6JnpduB8=",
    version = "v0.0.0-20160803190731-bd40a432e4c7",
)

go_repository(
    name = "com_github_jonboulle_clockwork",
    importpath = "github.com/jonboulle/clockwork",
    sum = "h1:VKV+ZcuP6l3yW9doeqz6ziZGgcynBVQO+obU0+0hcPo=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_json_iterator_go",
    importpath = "github.com/json-iterator/go",
    sum = "h1:Kz6Cvnvv2wGdaG/V8yMvfkmNiXq9Ya2KUv4rouJJr68=",
    version = "v1.1.10",
)

go_repository(
    name = "com_github_jstemmer_go_junit_report",
    importpath = "github.com/jstemmer/go-junit-report",
    sum = "h1:6QPYqodiu3GuPL+7mfx+NwDdp2eTkp9IfEUpgAwUN0o=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_jtolds_gls",
    importpath = "github.com/jtolds/gls",
    sum = "h1:xdiiI2gbIgH/gLH7ADydsJ1uDOEzR8yvV7C0MuV77Wo=",
    version = "v4.20.0+incompatible",
)

go_repository(
    name = "com_github_julienschmidt_httprouter",
    importpath = "github.com/julienschmidt/httprouter",
    sum = "h1:TDTW5Yz1mjftljbcKqRcrYhd4XeOoI98t+9HbQbYf7g=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    importpath = "github.com/kisielk/errcheck",
    sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_klauspost_compress",
    importpath = "github.com/klauspost/compress",
    sum = "h1:dB4Bn0tN3wdCzQxnS8r06kV74qN/TAfaIS0bVE8h3jc=",
    version = "v1.11.3",
)

go_repository(
    name = "com_github_konsorten_go_windows_terminal_sequences",
    importpath = "github.com/konsorten/go-windows-terminal-sequences",
    sum = "h1:CE8S1cTafDpPvMhIxNJKvHsGVBgn1xWYf1NbHQhywc8=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_kr_logfmt",
    importpath = "github.com/kr/logfmt",
    sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
    version = "v0.0.0-20140226030751-b84e30acd515",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    sum = "h1:s5hAObm+yFO5uHYt5dYjxi2rXrsnmRpJx4OYvIWUaQs=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    sum = "h1:hyz3dwM5QLc1Rfoz4FuWJQG5BN7tc6K1MndAUnGpQr4=",
    version = "v1.1.5",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    sum = "h1:45sCR5RtlFHMR4UwH9sdQ5TC8v0qDQCHnXt+kaKSTVE=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_mailru_easyjson",
    importpath = "github.com/mailru/easyjson",
    sum = "h1:aizVhC/NAAcKWb+5QsU1iNOZb4Yws5UO2I+aIprQITM=",
    version = "v0.7.0",
)

go_repository(
    name = "com_github_marstr_guid",
    importpath = "github.com/marstr/guid",
    sum = "h1:/M4H/1G4avsieL6BbUwCOBzulmoeKVP5ux/3mQNnbyI=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    importpath = "github.com/mattn/go-colorable",
    sum = "h1:UVL0vNpWh04HeJXV0KLcaT7r06gOH2l4OW6ddYRUIY4=",
    version = "v0.0.9",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    importpath = "github.com/mattn/go-isatty",
    sum = "h1:bnP0vzxcAdeI1zdubAl5PjU6zsERjGZb7raWodagDYs=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_mattn_go_runewidth",
    importpath = "github.com/mattn/go-runewidth",
    sum = "h1:UnlwIPBGaTZfPQ6T1IGzPI0EkYAQmT9fAEJ/poFC63o=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_mattn_go_shellwords",
    importpath = "github.com/mattn/go-shellwords",
    sum = "h1:K/VxK7SZ+cvuPgFSLKi5QPI9Vr/ipOf4C1gN+ntueUk=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
    sum = "h1:I0XW9+e1XWDxdcEniV4rQAIOPUGDq67JSCiRCgGCZLI=",
    version = "v1.0.2-0.20181231171920-c182affec369",
)

go_repository(
    name = "com_github_microsoft_go_winio",
    importpath = "github.com/Microsoft/go-winio",
    sum = "h1:mw6pDQqv38/WGF1cO/jF5t/jyAJ2yi7CmtFLLO5tGFI=",
    version = "v0.4.17-0.20210211115548-6eac466e5fa3",
)

go_repository(
    name = "com_github_microsoft_hcsshim",
    importpath = "github.com/Microsoft/hcsshim",
    sum = "h1:8/auA4LFIZFTGrqfKhGBSXwM6/4X1fHa/xniyEHu8ac=",
    version = "v0.8.16",
)

go_repository(
    name = "com_github_microsoft_hcsshim_test",
    importpath = "github.com/Microsoft/hcsshim/test",
    sum = "h1:4FA+QBaydEHlwxg0lMN3rhwoDaQy6LKhVWR4qvq4BuA=",
    version = "v0.0.0-20210227013316-43a75bb4edd3",
)

go_repository(
    name = "com_github_miekg_pkcs11",
    importpath = "github.com/miekg/pkcs11",
    sum = "h1:iMwmD7I5225wv84WxIG/bmxz9AXjWvTWIbM/TYHvWtw=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_mistifyio_go_zfs",
    importpath = "github.com/mistifyio/go-zfs",
    sum = "h1:aKW/4cBs+yK6gpqU3K/oIwk9Q/XICqd3zOX/UFuvqmk=",
    version = "v2.1.2-0.20190413222219-f784269be439+incompatible",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    importpath = "github.com/mitchellh/mapstructure",
    sum = "h1:fmNYVwqnSfB9mZU6OS2O6GsXM+wcskZDuKQzvN1EDeE=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_mitchellh_osext",
    importpath = "github.com/mitchellh/osext",
    sum = "h1:2+myh5ml7lgEU/51gbeLHfKGNfgEQQIWrlbdaOsidbQ=",
    version = "v0.0.0-20151018003038-5e2d6d41470f",
)

go_repository(
    name = "com_github_moby_sys_mount",
    importpath = "github.com/moby/sys/mount",
    sum = "h1:WhCW5B355jtxndN5ovugJlMFJawbUODuW8fSnEH6SSM=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_moby_sys_mountinfo",
    importpath = "github.com/moby/sys/mountinfo",
    sum = "h1:1O+1cHA1aujwEwwVMa2Xm2l+gIpUHyd3+D+d7LZh1kM=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_moby_sys_symlink",
    importpath = "github.com/moby/sys/symlink",
    sum = "h1:MTFZ74KtNI6qQQpuBxU+uKCim4WtOMokr03hCfJcazE=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_moby_term",
    importpath = "github.com/moby/term",
    sum = "h1:rzf0wL0CHVc8CEsgyygG0Mn9CNCCPZqOPaz8RiiHYQk=",
    version = "v0.0.0-20201216013528-df9cb8a40635",
)

go_repository(
    name = "com_github_modern_go_concurrent",
    importpath = "github.com/modern-go/concurrent",
    sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
    version = "v0.0.0-20180306012644-bacd9c7ef1dd",
)

go_repository(
    name = "com_github_modern_go_reflect2",
    importpath = "github.com/modern-go/reflect2",
    sum = "h1:9f412s+6RmYXLWZSEzVVgPGK7C2PphHj5RJrvfx9AWI=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_morikuni_aec",
    importpath = "github.com/morikuni/aec",
    sum = "h1:nXxl5PrvVm2L/wCy8dQu6DMTwH4oIuGN8GJDAlqDdVE=",
    version = "v0.0.0-20170113033406-39771216ff4c",
)

go_repository(
    name = "com_github_mrunalp_fileutils",
    importpath = "github.com/mrunalp/fileutils",
    sum = "h1:NKzVxiH7eSk+OQ4M+ZYW1K6h27RUV3MI6NUTsHhU6Z4=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_munnerz_goautoneg",
    importpath = "github.com/munnerz/goautoneg",
    sum = "h1:C3w9PqII01/Oq1c1nUAm88MOHcQC9l5mIlSMApZMrHA=",
    version = "v0.0.0-20191010083416-a7dc8b61c822",
)

go_repository(
    name = "com_github_mwitkow_go_conntrack",
    importpath = "github.com/mwitkow/go-conntrack",
    sum = "h1:F9x/1yl3T2AeKLr2AMdilSD8+f9bvMnNN8VS5iDtovc=",
    version = "v0.0.0-20161129095857-cc309e4a2223",
)

go_repository(
    name = "com_github_mxk_go_flowrate",
    importpath = "github.com/mxk/go-flowrate",
    sum = "h1:y5//uYreIhSUg3J1GEMiLbxo1LJaP8RfCpH6pymGZus=",
    version = "v0.0.0-20140419014527-cca7078d478f",
)

go_repository(
    name = "com_github_ncw_swift",
    importpath = "github.com/ncw/swift",
    sum = "h1:4DQRPj35Y41WogBxyhOXlrI37nzGlyEcsforeudyYPQ=",
    version = "v1.0.47",
)

go_repository(
    name = "com_github_nytimes_gziphandler",
    importpath = "github.com/NYTimes/gziphandler",
    sum = "h1:lsxEuwrXEAokXB9qhlbKWPpo3KMLZQ5WB5WLQRW1uq0=",
    version = "v0.0.0-20170623195520-56545f4a5d46",
)

go_repository(
    name = "com_github_olekukonko_tablewriter",
    importpath = "github.com/olekukonko/tablewriter",
    sum = "h1:58+kh9C6jJVXYjt8IE48G2eWl6BjwU5Gj0gqY84fy78=",
    version = "v0.0.0-20170122224234-a0225b3f23b5",
)

go_repository(
    name = "com_github_onsi_ginkgo",
    importpath = "github.com/onsi/ginkgo",
    sum = "h1:JAKSXpt1YjtLA7YpPiqO9ss6sNXEsPfSGdwN0UHqzrw=",
    version = "v1.11.0",
)

go_repository(
    name = "com_github_onsi_gomega",
    importpath = "github.com/onsi/gomega",
    sum = "h1:K0jcRCwNQM3vFGh1ppMtDh/+7ApJrjldlX8fA0jDTLQ=",
    version = "v1.7.1",
)

go_repository(
    name = "com_github_opencontainers_go_digest",
    importpath = "github.com/opencontainers/go-digest",
    sum = "h1:apOUWs51W5PlhuyGyz9FCeeBIOUDA/6nW8Oi/yOhh5U=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_opencontainers_image_spec",
    importpath = "github.com/opencontainers/image-spec",
    sum = "h1:JMemWkRwHx4Zj+fVxWoMCFm/8sYGGrUVojFA6h/TRcI=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_opencontainers_runc",
    importpath = "github.com/opencontainers/runc",
    sum = "h1:x2UMpOOVf3kQ8arv/EsDGwim8PTNqzL1/EYDr/+scOM=",
    version = "v1.0.0-rc93",
)

go_repository(
    name = "com_github_opencontainers_runtime_spec",
    importpath = "github.com/opencontainers/runtime-spec",
    sum = "h1:pNa8metDkwZjb9g4T8s+krQ+HRgZAkqnXml+wNir/+s=",
    version = "v1.0.3-0.20200929063507-e6143ca7d51d",
)

go_repository(
    name = "com_github_opencontainers_runtime_tools",
    importpath = "github.com/opencontainers/runtime-tools",
    sum = "h1:H7DMc6FAjgwZZi8BRqjrAAHWoqEr5e5L6pS4V0ezet4=",
    version = "v0.0.0-20181011054405-1d69bd0f9c39",
)

go_repository(
    name = "com_github_opencontainers_selinux",
    importpath = "github.com/opencontainers/selinux",
    sum = "h1:+77ba4ar4jsCbL1GLbFL8fFM57w6suPfSS9PDLDY7KM=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_peterbourgon_diskv",
    importpath = "github.com/peterbourgon/diskv",
    sum = "h1:UBdAOUP5p4RWqPBg048CAvpKN+vxiaj6gdUUzhl4XmI=",
    version = "v2.0.1+incompatible",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_pquerna_cachecontrol",
    importpath = "github.com/pquerna/cachecontrol",
    sum = "h1:0XM1XL/OFFJjXsYXlG30spTkV/E9+gmd5GD1w2HE8xM=",
    version = "v0.0.0-20171018203845-0dec1b30a021",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    importpath = "github.com/prometheus/client_golang",
    sum = "h1:NTGy1Ja9pByO+xAeH/qiWnLrKtr3hJPNjaVUwnjpdpA=",
    version = "v1.7.1",
)

go_repository(
    name = "com_github_prometheus_client_model",
    importpath = "github.com/prometheus/client_model",
    sum = "h1:uq5h0d+GuxiXLJLNABMgp2qUWDPiLvgCzz2dUR+/W/M=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_prometheus_common",
    importpath = "github.com/prometheus/common",
    sum = "h1:RyRA7RzGXQZiW+tGMr7sxa85G1z0yOpM1qq5c8lNawc=",
    version = "v0.10.0",
)

go_repository(
    name = "com_github_prometheus_procfs",
    importpath = "github.com/prometheus/procfs",
    sum = "h1:mxy4L2jP6qMonqmq+aTtOx1ifVWUgG/TAmntgbh3xv4=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_puerkitobio_purell",
    importpath = "github.com/PuerkitoBio/purell",
    sum = "h1:WEQqlqaGbrPkxLJWfBwQmfEAE1Z7ONdDLqrN38tNFfI=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_puerkitobio_urlesc",
    importpath = "github.com/PuerkitoBio/urlesc",
    sum = "h1:d+Bc7a5rLufV/sSk/8dngufqelfh6jnri85riMAaF/M=",
    version = "v0.0.0-20170810143723-de5bf2ad4578",
)

go_repository(
    name = "com_github_rogpeppe_fastuuid",
    importpath = "github.com/rogpeppe/fastuuid",
    sum = "h1:gu+uRPtBe88sKxUCEXRoeCvVG90TJmwhiqRpvdhQFng=",
    version = "v0.0.0-20150106093220-6724a57986af",
)

go_repository(
    name = "com_github_rogpeppe_go_internal",
    importpath = "github.com/rogpeppe/go-internal",
    sum = "h1:RR9dF3JtopPvtkroDZuVD7qquD0bnHlKSqaQhgwt8yk=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_russross_blackfriday_v2",
    importpath = "github.com/russross/blackfriday/v2",
    sum = "h1:lPqVAte+HuHNfhJ/0LC98ESWRz8afy9tM/0RK8m9o+Q=",
    version = "v2.0.1",
)

go_repository(
    name = "com_github_safchain_ethtool",
    importpath = "github.com/safchain/ethtool",
    sum = "h1:2c1EFnZHIPCW8qKWgHMH/fX2PkSabFc5mrVzfUNdg5U=",
    version = "v0.0.0-20190326074333-42ed695e3de8",
)

go_repository(
    name = "com_github_satori_go_uuid",
    importpath = "github.com/satori/go.uuid",
    sum = "h1:0uYX9dsZ2yD7q2RtLRtPSdGDWzjeM3TbMJP9utgA0ww=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_seccomp_libseccomp_golang",
    importpath = "github.com/seccomp/libseccomp-golang",
    sum = "h1:NJjM5DNFOs0s3kYE1WUOr6G8V97sdt46rlXTMfXGWBo=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_shopify_logrus_bugsnag",
    importpath = "github.com/Shopify/logrus-bugsnag",
    sum = "h1:UrqY+r/OJnIp5u0s1SbQ8dVfLCZJsnvazdBP5hS4iRs=",
    version = "v0.0.0-20171204204709-577dee27f20d",
)

go_repository(
    name = "com_github_shurcool_sanitized_anchor_name",
    importpath = "github.com/shurcooL/sanitized_anchor_name",
    sum = "h1:PdmoCO6wvbs+7yrJyMORt4/BmY5IYyJwS/kOiWx8mHo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_smartystreets_assertions",
    importpath = "github.com/smartystreets/assertions",
    sum = "h1:zE9ykElWQ6/NYmHa3jpm/yHnI4xSofP+UP6SpjHcSeM=",
    version = "v0.0.0-20180927180507-b2de0cb4f26d",
)

go_repository(
    name = "com_github_smartystreets_goconvey",
    importpath = "github.com/smartystreets/goconvey",
    sum = "h1:pa8hGb/2YqsZKovtsgrwcDH1RZhVbTKCjLp47XpqCDs=",
    version = "v0.0.0-20190330032615-68dc04aab96a",
)

go_repository(
    name = "com_github_soheilhy_cmux",
    importpath = "github.com/soheilhy/cmux",
    sum = "h1:0HKaf1o97UwFjHH9o5XsHUOF+tqmdA7KEzXLpiyaw0E=",
    version = "v0.1.4",
)

go_repository(
    name = "com_github_spf13_afero",
    importpath = "github.com/spf13/afero",
    sum = "h1:5jhuqJyZCZf2JRofRvN/nIFgIWNzPa3/Vz8mYylgbWc=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_spf13_cobra",
    importpath = "github.com/spf13/cobra",
    sum = "h1:ZlrZ4XsMRm04Fr5pSFxBgfND2EBVa1nLpiy1stUsX/8=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_spf13_pflag",
    importpath = "github.com/spf13/pflag",
    sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
    version = "v1.0.5",
)

go_repository(
    name = "com_github_stefanberger_go_pkcs11uri",
    importpath = "github.com/stefanberger/go-pkcs11uri",
    sum = "h1:lIOOHPEbXzO3vnmx2gok1Tfs31Q8GQqKLc8vVqyQq/I=",
    version = "v0.0.0-20201008174630-78d3cae3a980",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:Hbg2NidpLE8veEBkEZTL3CvlkUIVzuU9jDplZO54c48=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_syndtr_gocapability",
    importpath = "github.com/syndtr/gocapability",
    sum = "h1:kdXcSzyDtseVEc4yCz2qF8ZrQvIDBJLl4S1c3GCXmoI=",
    version = "v0.0.0-20200815063812-42c35b437635",
)

go_repository(
    name = "com_github_tchap_go_patricia",
    importpath = "github.com/tchap/go-patricia",
    sum = "h1:JvoDL7JSoIP2HDE8AbDH3zC8QBPxmzYe32HHy5yQ+Ck=",
    version = "v2.2.6+incompatible",
)

go_repository(
    name = "com_github_tmc_grpc_websocket_proxy",
    importpath = "github.com/tmc/grpc-websocket-proxy",
    sum = "h1:LnC5Kc/wtumK+WB441p7ynQJzVuNRJiqddSIE3IlSEQ=",
    version = "v0.0.0-20190109142713-0ad062ec5ee5",
)

go_repository(
    name = "com_github_urfave_cli",
    importpath = "github.com/urfave/cli",
    sum = "h1:gsqYFH8bb9ekPA12kRo0hfjngWQjkJPlN9R0N78BoUo=",
    version = "v1.22.2",
)

go_repository(
    name = "com_github_vishvananda_netlink",
    importpath = "github.com/vishvananda/netlink",
    sum = "h1:1iyaYNBLmP6L0220aDnYQpo1QEV4t4hJ+xEEhhJH8j0=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_vishvananda_netns",
    importpath = "github.com/vishvananda/netns",
    sum = "h1:OviZH7qLw/7ZovXvuNyL3XQl8UFofeikI1NW1Gypu7k=",
    version = "v0.0.0-20191106174202-0a2b9b5464df",
)

go_repository(
    name = "com_github_willf_bitset",
    importpath = "github.com/willf/bitset",
    sum = "h1:N7Z7E9UvjW+sGsEl7k/SJrvY2reP1A07MrGuCjIOjRE=",
    version = "v1.1.11",
)

go_repository(
    name = "com_github_xeipuuv_gojsonpointer",
    importpath = "github.com/xeipuuv/gojsonpointer",
    sum = "h1:J9EGpcZtP0E/raorCMxlFGSTBrsSlaDGf3jU/qvAE2c=",
    version = "v0.0.0-20180127040702-4e3ac2762d5f",
)

go_repository(
    name = "com_github_xeipuuv_gojsonreference",
    importpath = "github.com/xeipuuv/gojsonreference",
    sum = "h1:EzJWgHovont7NscjpAxXsDA8S8BMYve8Y5+7cuRE7R0=",
    version = "v0.0.0-20180127040603-bd5ef7bd5415",
)

go_repository(
    name = "com_github_xeipuuv_gojsonschema",
    importpath = "github.com/xeipuuv/gojsonschema",
    sum = "h1:mvXjJIHRZyhNuGassLTcXTwjiWq7NmjdavZsUnmFybQ=",
    version = "v0.0.0-20180618132009-1d523034197f",
)

go_repository(
    name = "com_github_xiang90_probing",
    importpath = "github.com/xiang90/probing",
    sum = "h1:eY9dn8+vbi4tKz5Qo6v2eYzo7kUS51QINcR5jNpbZS8=",
    version = "v0.0.0-20190116061207-43a291ad63a2",
)

go_repository(
    name = "com_github_yuin_goldmark",
    importpath = "github.com/yuin/goldmark",
    sum = "h1:ruQGxdhGHe7FWOJPT0mKs5+pD2Xs1Bm/kdGlHO04FmM=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_yvasiyarov_go_metrics",
    importpath = "github.com/yvasiyarov/go-metrics",
    sum = "h1:+lm10QQTNSBd8DVTNGHx7o/IKu9HYDvLMffDhbyLccI=",
    version = "v0.0.0-20140926110328-57bccd1ccd43",
)

go_repository(
    name = "com_github_yvasiyarov_gorelic",
    importpath = "github.com/yvasiyarov/gorelic",
    sum = "h1:hlE8//ciYMztlGpl/VA+Zm1AcTPHYkHJPbHqE6WJUXE=",
    version = "v0.0.0-20141212073537-a9bba5b9ab50",
)

go_repository(
    name = "com_github_yvasiyarov_newrelic_platform_go",
    importpath = "github.com/yvasiyarov/newrelic_platform_go",
    sum = "h1:ERexzlUfuTvpE74urLSbIQW0Z/6hF9t8U4NsJLaioAY=",
    version = "v0.0.0-20140908184405-b21fdbd4370f",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    sum = "h1:3ithwDMr7/3vpAMXiH+ZQnYbuIsh+OPhUPMFC9enmn0=",
    version = "v0.54.0",
)

go_repository(
    name = "com_google_cloud_go_bigquery",
    importpath = "cloud.google.com/go/bigquery",
    sum = "h1:xE3CPsOgttP4ACBePh79zTKALtXwn/Edhcr16R5hMWU=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_datastore",
    importpath = "cloud.google.com/go/datastore",
    sum = "h1:/May9ojXjRkPBNVrq+oWLqmWCkr4OU5uRY29bu0mRyQ=",
    version = "v1.1.0",
)

go_repository(
    name = "com_google_cloud_go_pubsub",
    importpath = "cloud.google.com/go/pubsub",
    sum = "h1:Lpy6hKgdcl7a3WGSfJIFmxmcdjSpP6OmBEfcOv1Y680=",
    version = "v1.2.0",
)

go_repository(
    name = "com_google_cloud_go_storage",
    importpath = "cloud.google.com/go/storage",
    sum = "h1:UDpwYIwla4jHGzZJaEJYx1tOejbgSoNqsAfHAUYe2r8=",
    version = "v1.6.0",
)

go_repository(
    name = "com_shuralyov_dmitri_gpu_mtl",
    importpath = "dmitri.shuralyov.com/gpu/mtl",
    sum = "h1:VpgP7xuJadIUuKccphEpTJnWhS2jkQyMt6Y7pJCD7fY=",
    version = "v0.0.0-20190408044501-666a987793e9",
)

go_repository(
    name = "in_gopkg_airbrake_gobrake_v2",
    importpath = "gopkg.in/airbrake/gobrake.v2",
    sum = "h1:7z2uVWwn7oVeeugY1DtlPAy5H+KYgB1KeKTnqjNatLo=",
    version = "v2.0.9",
)

go_repository(
    name = "in_gopkg_alecthomas_kingpin_v2",
    importpath = "gopkg.in/alecthomas/kingpin.v2",
    sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
    version = "v2.2.6",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:YR8cESwS4TdDjEe65xsg0ogRM/Nc3DYOhEAlW+xobZo=",
    version = "v1.0.0-20190902080502-41f04d3bba15",
)

go_repository(
    name = "in_gopkg_cheggaaa_pb_v1",
    importpath = "gopkg.in/cheggaaa/pb.v1",
    sum = "h1:Ev7yu1/f6+d+b3pi5vPdRPc6nNtP1umSfcWiEfRqv6I=",
    version = "v1.0.25",
)

go_repository(
    name = "in_gopkg_errgo_v2",
    importpath = "gopkg.in/errgo.v2",
    sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
    version = "v2.1.0",
)

go_repository(
    name = "in_gopkg_fsnotify_v1",
    importpath = "gopkg.in/fsnotify.v1",
    sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
    version = "v1.4.7",
)

go_repository(
    name = "in_gopkg_gemnasium_logrus_airbrake_hook_v2",
    importpath = "gopkg.in/gemnasium/logrus-airbrake-hook.v2",
    sum = "h1:OAj3g0cR6Dx/R07QgQe8wkA9RNjB2u4i700xBkIT4e0=",
    version = "v2.1.2",
)

go_repository(
    name = "in_gopkg_inf_v0",
    importpath = "gopkg.in/inf.v0",
    sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
    version = "v0.9.1",
)

go_repository(
    name = "in_gopkg_natefinch_lumberjack_v2",
    importpath = "gopkg.in/natefinch/lumberjack.v2",
    sum = "h1:1Lc07Kr7qY4U2YPouBjpCLxpiyxIVoxqXgkXLknAOE8=",
    version = "v2.0.0",
)

go_repository(
    name = "in_gopkg_resty_v1",
    importpath = "gopkg.in/resty.v1",
    sum = "h1:CuXP0Pjfw9rOuY6EP+UvtNvt5DSqHpIxILZKT/quCZI=",
    version = "v1.12.0",
)

go_repository(
    name = "in_gopkg_square_go_jose_v2",
    importpath = "gopkg.in/square/go-jose.v2",
    sum = "h1:7odma5RETjNHWJnR32wx8t+Io4djHE1PqxCFx3iiZ2w=",
    version = "v2.5.1",
)

go_repository(
    name = "in_gopkg_tomb_v1",
    importpath = "gopkg.in/tomb.v1",
    sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
    version = "v1.0.0-20141024135613-dd632973f1e7",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
    version = "v2.4.0",
)

go_repository(
    name = "io_etcd_go_bbolt",
    importpath = "go.etcd.io/bbolt",
    sum = "h1:XAzx9gjCb0Rxj7EoqcClPD1d5ZBxZJk0jbuoPHenBt0=",
    version = "v1.3.5",
)

go_repository(
    name = "io_etcd_go_etcd",
    importpath = "go.etcd.io/etcd",
    sum = "h1:1JFLBqwIgdyHN1ZtgjTBwO+blA6gVOmZurpiMEsETKo=",
    version = "v0.5.0-alpha.5.0.20200910180754-dd1b699fc489",
)

go_repository(
    name = "io_k8s_api",
    importpath = "k8s.io/api",
    sum = "h1:ud1c3W3YNzGd6ABJlbFfKXBKXO+1KdGfcgGGNgFR03E=",
    version = "v0.20.1",
)

go_repository(
    name = "io_k8s_apimachinery",
    importpath = "k8s.io/apimachinery",
    sum = "h1:LAhz8pKbgR8tUwn7boK+b2HZdt7MiTu2mkYtFMUjTRQ=",
    version = "v0.20.1",
)

go_repository(
    name = "io_k8s_apiserver",
    importpath = "k8s.io/apiserver",
    sum = "h1:yEqdkxlnQbxi/3e74cp0X16h140fpvPrNnNRAJBDuBk=",
    version = "v0.20.1",
)

go_repository(
    name = "io_k8s_client_go",
    importpath = "k8s.io/client-go",
    sum = "h1:Qquik0xNFbK9aUG92pxHYsyfea5/RPO9o9bSywNor+M=",
    version = "v0.20.1",
)

go_repository(
    name = "io_k8s_component_base",
    importpath = "k8s.io/component-base",
    sum = "h1:6OQaHr205NSl24t5wOF2IhdrlxZTWEZwuGlLvBgaeIg=",
    version = "v0.20.1",
)

go_repository(
    name = "io_k8s_cri_api",
    importpath = "k8s.io/cri-api",
    sum = "h1:b4l7SZ9+VPfIrrJnMXzm0HR9wAsHwHh9+QcmK31nQMI=",
    version = "v0.20.1",
)

go_repository(
    name = "io_k8s_gengo",
    importpath = "k8s.io/gengo",
    sum = "h1:sAvhNk5RRuc6FNYGqe7Ygz3PSo/2wGWbulskmzRX8Vs=",
    version = "v0.0.0-20200413195148-3a45101e95ac",
)

go_repository(
    name = "io_k8s_klog_v2",
    importpath = "k8s.io/klog/v2",
    sum = "h1:7+X0fUguPyrKEC4WjH8iGDg3laWgMo5tMnRTIGTTxGQ=",
    version = "v2.4.0",
)

go_repository(
    name = "io_k8s_kube_openapi",
    importpath = "k8s.io/kube-openapi",
    sum = "h1:sOHNzJIkytDF6qadMNKhhDRpc6ODik8lVC6nOur7B2c=",
    version = "v0.0.0-20201113171705-d219536bb9fd",
)

go_repository(
    name = "io_k8s_kubernetes",
    importpath = "k8s.io/kubernetes",
    sum = "h1:qTfB+u5M92k2fCCCVP2iuhgwwSOv1EkAkvQY1tQODD8=",
    version = "v1.13.0",
)

go_repository(
    name = "io_k8s_sigs_apiserver_network_proxy_konnectivity_client",
    importpath = "sigs.k8s.io/apiserver-network-proxy/konnectivity-client",
    sum = "h1:TihvEz9MPj2u0KWds6E2OBUXfwaL4qRJ33c7HGiJpqk=",
    version = "v0.0.14",
)

go_repository(
    name = "io_k8s_sigs_structured_merge_diff_v4",
    importpath = "sigs.k8s.io/structured-merge-diff/v4",
    sum = "h1:YHQV7Dajm86OuqnIR6zAelnDWBRjo+YhYV9PmGrh1s8=",
    version = "v4.0.2",
)

go_repository(
    name = "io_k8s_sigs_yaml",
    importpath = "sigs.k8s.io/yaml",
    sum = "h1:kr/MCeFWJWTwyaHoR9c8EjH9OumOmoF9YGiZd7lFm/Q=",
    version = "v1.2.0",
)

go_repository(
    name = "io_k8s_utils",
    importpath = "k8s.io/utils",
    sum = "h1:CbnUZsM497iRC5QMVkHwyl8s2tB3g7yaSHkYPkpgelw=",
    version = "v0.0.0-20201110183641-67b214c5f920",
)

go_repository(
    name = "io_opencensus_go",
    importpath = "go.opencensus.io",
    sum = "h1:8sGtKOrtQqkN1bp2AtX+misvLIlOmsEsNd+9NIcPEm8=",
    version = "v0.22.3",
)

go_repository(
    name = "io_rsc_binaryregexp",
    importpath = "rsc.io/binaryregexp",
    sum = "h1:HfqmD5MEmC0zvwBuF187nq9mdnXjXsSivRiXN7SmRkE=",
    version = "v0.2.0",
)

go_repository(
    name = "io_rsc_quote_v3",
    importpath = "rsc.io/quote/v3",
    sum = "h1:9JKUTTIUgS6kzR9mK1YuGKv6Nl+DijDNIc0ghT58FaY=",
    version = "v3.1.0",
)

go_repository(
    name = "io_rsc_sampler",
    importpath = "rsc.io/sampler",
    sum = "h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=",
    version = "v1.3.0",
)

go_repository(
    name = "org_bazil_fuse",
    importpath = "bazil.org/fuse",
    sum = "h1:SC+c6A1qTFstO9qmB86mPV2IpYme/2ZoEQ0hrP+wo+Q=",
    version = "v0.0.0-20160811212531-371fbbdaa898",
)

go_repository(
    name = "org_golang_google_api",
    importpath = "google.golang.org/api",
    sum = "h1:jz2KixHX7EcCPiQrySzPdnYT7DbINAypCqKZ1Z7GM40=",
    version = "v0.20.0",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    sum = "h1:tycE03LOZYQNhDpS27tcQdAzLCVMaj7QT2SXxebnpCM=",
    version = "v1.6.5",
)

go_repository(
    name = "org_golang_google_cloud",
    importpath = "google.golang.org/cloud",
    sum = "h1:Cpp2P6TPjujNoC5M2KHY6g7wfyLYfIWRZaSdIKfDasA=",
    version = "v0.0.0-20151119220103-975617b05ea8",
)

go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    sum = "h1:pOwg4OoaRYScjmR4LlLgdtnyoHYTSAVhhqe5uPdpII8=",
    version = "v0.0.0-20201110150050-8816d57aaa9a",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:EQyQC3sa8M+p6Ulc8yy9SWSS2GVwyRc83gAbG8lrl4o=",
    version = "v1.33.2",
)

go_repository(
    name = "org_golang_google_protobuf",
    importpath = "google.golang.org/protobuf",
    sum = "h1:Ejskq+SyPohKW+1uil0JJMtmHCgJPJ/qWTxr8qp+R4c=",
    version = "v1.25.0",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:hb9wdF1z5waM+dSIICn1l0DkLVDT3hqhhQsDNUmHPRE=",
    version = "v0.0.0-20201002170205-7f63de1d35b0",
)

go_repository(
    name = "org_golang_x_exp",
    importpath = "golang.org/x/exp",
    sum = "h1:QE6XYQK6naiK1EPAe1g/ILLxN5RBoH5xkJk3CqlMI/Y=",
    version = "v0.0.0-20200224162631-6cc2880d07d6",
)

go_repository(
    name = "org_golang_x_image",
    importpath = "golang.org/x/image",
    sum = "h1:+qEpEAPhDZ1o0x3tHzZTQDArnOixOzGD9HUJfcg0mb4=",
    version = "v0.0.0-20190802002840-cff245a6509b",
)

go_repository(
    name = "org_golang_x_lint",
    importpath = "golang.org/x/lint",
    sum = "h1:Wh+f8QHJXR411sJR8/vRBTZ7YapZaRvUcLFFJhusH0k=",
    version = "v0.0.0-20200302205851-738671d3881b",
)

go_repository(
    name = "org_golang_x_mobile",
    importpath = "golang.org/x/mobile",
    sum = "h1:4+4C/Iv2U4fMZBiMCc98MG1In4gJY5YRhtpDNeDeHWs=",
    version = "v0.0.0-20190719004257-d2bd2a29d028",
)

go_repository(
    name = "org_golang_x_mod",
    importpath = "golang.org/x/mod",
    sum = "h1:RM4zey1++hCTbCVQfnWeKs9/IEsaBLA8vTkd0WVtmH4=",
    version = "v0.3.0",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:iFwSg7t5GZmB/Q5TjiEAsdoLDrdJRC1RiF2WhuV29Qw=",
    version = "v0.0.0-20201224014010-6772e930b67b",
)

go_repository(
    name = "org_golang_x_oauth2",
    importpath = "golang.org/x/oauth2",
    sum = "h1:TzXSXBo42m9gQenoE3b9BGiEpg5IG2JkU5FkPIawgtw=",
    version = "v0.0.0-20200107190931-bf48bf16ab8d",
)

go_repository(
    name = "org_golang_x_sync",
    importpath = "golang.org/x/sync",
    sum = "h1:DcqTD9SDLc+1P/r1EmRBwnVsrOwW+kk2vWf9n+1sGhs=",
    version = "v0.0.0-20201207232520-09787c993a3a",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:Paq34FxTluEPvVyayQqMPgHm+vTOrIifmcYxFBx9TLg=",
    version = "v0.0.0-20210324051608-47abb6519492",
)

go_repository(
    name = "org_golang_x_term",
    importpath = "golang.org/x/term",
    sum = "h1:v+OssWQX+hTHEmOBgwxdZxK4zHq3yOs8F9J7mk0PY8E=",
    version = "v0.0.0-20201126162022-7de9c90e9dd1",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:0YWbFKbhXG/wIiuHDSKpS0Iy7FSA+u45VtBMfQcFTTc=",
    version = "v0.3.4",
)

go_repository(
    name = "org_golang_x_time",
    importpath = "golang.org/x/time",
    sum = "h1:EHBhcS0mlXEAVwNyO2dLfjToGsyY4j24pTs2ScHnX7s=",
    version = "v0.0.0-20200630173020-3af7569d3a1e",
)

go_repository(
    name = "org_golang_x_xerrors",
    importpath = "golang.org/x/xerrors",
    sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
)

go_repository(
    name = "org_mozilla_go_pkcs7",
    importpath = "go.mozilla.org/pkcs7",
    sum = "h1:A/5uWzF44DlIgdm/PQFwfMkW0JX+cIcQi/SwLAmZP5M=",
    version = "v0.0.0-20200128120323-432b2356ecb1",
)

go_repository(
    name = "org_uber_go_atomic",
    importpath = "go.uber.org/atomic",
    sum = "h1:cxzIVoETapQEqDhQu3QfnvXAV4AlzcvUCxkVUFw3+EU=",
    version = "v1.4.0",
)

go_repository(
    name = "org_uber_go_multierr",
    importpath = "go.uber.org/multierr",
    sum = "h1:HoEmRHQPVSqub6w2z2d2EOVs2fjyFRGyofhKuyDq0QI=",
    version = "v1.1.0",
)

go_repository(
    name = "org_uber_go_zap",
    importpath = "go.uber.org/zap",
    sum = "h1:ORx85nbTijNz8ljznvCMR1ZBIPKFn3jQrag10X2AsuM=",
    version = "v1.10.0",
)

go_repository(
    name = "tools_gotest",
    importpath = "gotest.tools",
    sum = "h1:VsBPFP1AI068pPrMxtb/S8Zkgf9xEmTLJjfM+P5UIEo=",
    version = "v2.2.0+incompatible",
)

go_repository(
    name = "tools_gotest_v3",
    importpath = "gotest.tools/v3",
    sum = "h1:4AuOwCGf4lLR9u3YOe2awrHygurzhO/HeQ6laiA6Sx0=",
    version = "v3.0.3",
)

go_repository(
    name = "com_github_go_redis_redis_v8",
    importpath = "github.com/go-redis/redis/v8",
    sum = "h1:AcZZR7igkdvfVmQTPnu9WE37LRrO/YrBH5zWyjDC0oI=",
    version = "v8.11.5",
)

go_repository(
    name = "com_github_dgryski_go_rendezvous",
    importpath = "github.com/dgryski/go-rendezvous",
    sum = "h1:lO4WD4F/rVNCu3HqELle0jiPLLBs70cWOduZpkS1E78=",
    version = "v0.0.0-20200823014737-9f7001d12a5f",
)

# Declare indirect dependencies and register toolchains.
go_rules_dependencies()

go_register_toolchains(version = "1.17")

gazelle_dependencies()

http_archive(
    name = "com_google_protobuf",
    sha256 = "d0f5f605d0d656007ce6c8b5a82df3037e1d8fe8b121ed42e536f569dec16113",
    strip_prefix = "protobuf-3.14.0",
    urls = [
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
        "https://github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
    ],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

http_archive(
    name = "com_github_atlassian_bazel_tools",
    sha256 = "a1ebb48984f0a8c38c5ab155043f9ed63a3880ef53df9785154ebab946f4821f",
    strip_prefix = "bazel-tools-541e546def26783063c7630495a618d48be2b4cb",
    urls = [
        "https://github.com/atlassian/bazel-tools/archive/541e546def26783063c7630495a618d48be2b4cb.tar.gz",
    ],
)

# vim: ts=8 et sw=4 sts=4 :
