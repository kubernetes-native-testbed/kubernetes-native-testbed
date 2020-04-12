module github.com/__TB_GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/order/server

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.5
	github.com/jinzhu/gorm v1.9.12
	github.com/__TB_GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/order v0.0.0-00010101000000-000000000000
	github.com/__TB_GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/order/protobuf v0.0.0-00010101000000-000000000000
	github.com/__TB_GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/payment-info/protobuf v0.0.0-00010101000000-000000000000
	github.com/__TB_GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/user/protobuf v0.0.0-00010101000000-000000000000
	github.com/nats-io/nats.go v1.9.1
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/genproto v0.0.0-20200319113533-08878b785e9c // indirect
	google.golang.org/grpc v1.28.0
)

replace (
	github.com/__TB_GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/order => ../
	github.com/__TB_GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/order/protobuf => ../protobuf
	github.com/__TB_GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/payment-info/protobuf => ../../payment-info/protobuf
	github.com/__TB_GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/user/protobuf => ../../user/protobuf
)
