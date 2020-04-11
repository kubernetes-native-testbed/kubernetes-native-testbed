module github.com/__GITHUB_REPO_NAME__/kubernetes-native-testbed/microservices/rate/server

go 1.13

require (
	github.com/golang/protobuf v1.3.5
	github.com/jinzhu/gorm v1.9.12
	github.com/__GITHUB_REPO_NAME__/kubernetes-native-testbed/microservices/rate v0.0.0-00010101000000-000000000000
	github.com/__GITHUB_REPO_NAME__/kubernetes-native-testbed/microservices/rate/protobuf v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.26.0
)

replace (
	github.com/__GITHUB_REPO_NAME__/kubernetes-native-testbed/microservices/rate => ../
	github.com/__GITHUB_REPO_NAME__/kubernetes-native-testbed/microservices/rate/protobuf => ../protobuf
)
