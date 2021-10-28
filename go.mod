module github.com/service-register

go 1.16

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/outreach-golang/etcd v1.1.3
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20201229170055-e5319fda7802 // indirect
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/genproto v0.0.0-20211027162914-98a5263abeca // indirect
	google.golang.org/grpc v1.41.0 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
