module github.com/rabobank/scheduler-plugin

go 1.25

replace (
	github.com/envoyproxy/protoc-gen-validate => github.com/envoyproxy/protoc-gen-validate v1.2.1
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
	github.com/golang/glog => github.com/golang/glog v1.2.5
	github.com/nats-io/nats-server/v2 => github.com/nats-io/nats-server/v2 v2.10.27
	github.com/pkg/sftp => github.com/pkg/sftp v1.13.9
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.17.0
	github.com/yuin/goldmark => github.com/yuin/goldmark v1.7.13
	golang.org/x/crypto => golang.org/x/crypto v0.43.0
	golang.org/x/net => golang.org/x/net v0.46.0
	golang.org/x/text => golang.org/x/text v0.30.0
	golang.org/x/tools => golang.org/x/tools v0.38.0
	google.golang.org/grpc => google.golang.org/grpc v1.76.0
	google.golang.org/protobuf => google.golang.org/protobuf v1.36.10
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.4.0
)

require (
	code.cloudfoundry.org/cli v7.1.0+incompatible
	github.com/cloudfoundry/go-cfclient/v3 v3.0.0-alpha.15
)

require (
	code.cloudfoundry.org/bytefmt v0.54.0 // indirect
	code.cloudfoundry.org/cli-plugin-repo v0.0.0-20251009190434-d2babf059e69 // indirect
	code.cloudfoundry.org/go-log-cache v1.0.0 // indirect
	code.cloudfoundry.org/go-loggregator v7.4.0+incompatible // indirect
	code.cloudfoundry.org/gofileutils v0.0.0-20170111115228-4d0c80011a0f // indirect
	code.cloudfoundry.org/jsonry v1.1.4 // indirect
	code.cloudfoundry.org/rfc5424 v0.0.0-20201103192249-000122071b78 // indirect
	code.cloudfoundry.org/tlsconfig v0.35.0 // indirect
	code.cloudfoundry.org/ykk v0.0.0-20170424192843-e4df4ce2fd4d // indirect
	github.com/SermoDigital/jose v0.9.1 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/bmatcuk/doublestar v1.3.4 // indirect
	github.com/bmizerany/pat v0.0.0-20210406213842-e4b6760bdd6f // indirect
	github.com/charlievieth/fs v0.0.3 // indirect
	github.com/clipperhouse/uax29/v2 v2.2.0 // indirect
	github.com/cloudfoundry/bosh-cli v6.4.1+incompatible // indirect
	github.com/cloudfoundry/bosh-utils v0.0.555 // indirect
	github.com/codegangsta/inject v0.0.0-20150114235600-33e0aa1cb7c0 // indirect
	github.com/cppforlife/go-patch v0.2.0 // indirect
	github.com/creack/pty v1.1.24 // indirect
	github.com/cyphar/filepath-securejoin v0.5.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-martini/martini v0.0.0-20170121215854-22fa46961aab // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/jessevdk/go-flags v1.6.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/pty v1.1.8 // indirect
	github.com/lunixbochs/vtclean v1.0.0 // indirect
	github.com/martini-contrib/render v0.0.0-20150707142108-ec18f8345a11 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.19 // indirect
	github.com/moby/moby v20.10.12+incompatible // indirect
	github.com/moby/term v0.5.2 // indirect
	github.com/nxadm/tail v1.4.11 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/sabhiram/go-gitignore v0.0.0-20210923224102-525f6e181f06 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	github.com/tedsuo/rata v1.0.0 // indirect
	github.com/vito/go-interact v1.0.0 // indirect
	golang.org/x/crypto v0.43.0 // indirect
	golang.org/x/oauth2 v0.32.0 // indirect
	golang.org/x/sys v0.37.0 // indirect
	golang.org/x/term v0.36.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20251007200510-49b9836ed3ff // indirect
	google.golang.org/grpc v1.76.0 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/cheggaaa/pb.v1 v1.0.28 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
