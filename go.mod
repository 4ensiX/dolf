module example.com/dolf

go 1.17

require (
	dolf/util v0.0.0-00010101000000-000000000000
	github.com/docker/docker v20.10.10+incompatible
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2
)

require (
	github.com/Microsoft/go-winio v0.4.17 // indirect
	github.com/containerd/containerd v1.5.7 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/genproto v0.0.0-20201110150050-8816d57aaa9a // indirect
	google.golang.org/grpc v1.42.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)

replace example.com/util => ./util

replace dolf/util => ./util