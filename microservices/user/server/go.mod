module github.com/__GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/user/server

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.5
	github.com/jinzhu/gorm v1.9.12
	github.com/__GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/user v0.0.0-00010101000000-000000000000
	github.com/__GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/user/protobuf v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa // indirect
	golang.org/x/sys v0.0.0-20200124204421-9fbb57f87de9 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200122232147-0452cf42e150 // indirect
	google.golang.org/grpc v1.28.0
)

replace (
	github.com/__GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/user => ../
	github.com/__GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/user/protobuf => ../protobuf
)
