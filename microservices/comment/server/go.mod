module github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/comment/server

go 1.13

require (
	github.com/golang/protobuf v1.3.5
	github.com/jinzhu/gorm v1.9.12
	github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/comment v0.0.0-00010101000000-000000000000
	github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/comment/protobuf v0.0.0-00010101000000-000000000000
	github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/user v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.26.0
)

replace (
	github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/comment => ../
	github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/comment/protobuf => ../protobuf
	github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/user => ../../user
)
