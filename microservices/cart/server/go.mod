module github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/cart/server

go 1.13

require (
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d // indirect
	github.com/coreos/etcd v3.3.19+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/golang/protobuf v1.3.5
	github.com/google/btree v1.0.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/cart v0.0.0-00010101000000-000000000000
	github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/order/protobuf v0.0.0-00010101000000-000000000000
	github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/product/protobuf v0.0.0-00010101000000-000000000000
	github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/user v0.0.0-00010101000000-000000000000
	github.com/pingcap/goleveldb v0.0.0-20191226122134-f82aafb29989 // indirect
	github.com/pingcap/kvproto v0.0.0-20200317112120-78042b285b75 // indirect
	github.com/pingcap/log v0.0.0-20200117041106-d28c14d3b1cd // indirect
	github.com/pingcap/pd v2.1.19+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.5.1 // indirect
	github.com/prometheus/procfs v0.0.11 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20190728182440-6a916e37a237 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	go.etcd.io/etcd v3.3.19+incompatible // indirect
	go.uber.org/zap v1.14.1 // indirect
	golang.org/x/crypto v0.0.0-20200317142112-1b76d66859c6 // indirect
	google.golang.org/grpc v1.26.0
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/cart => ../
	github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/order/protobuf => ../../order/protobuf
	github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/product/protobuf => ../../product/protobuf
	github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/user => ../../user
)
