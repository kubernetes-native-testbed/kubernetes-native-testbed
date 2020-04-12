module github.com/__TB_GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/payment-info

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.5
	github.com/google/uuid v1.1.1
	github.com/jinzhu/gorm v1.9.12
	github.com/__TB_GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/payment-info/protobuf v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2 // indirect
	golang.org/x/sys v0.0.0-20200212091648-12a6c2dcc1e4 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/genproto v0.0.0-20200212174721-66ed5ce911ce // indirect
	google.golang.org/grpc v1.28.0
)

replace github.com/__TB_GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/payment-info/protobuf => ./protobuf
