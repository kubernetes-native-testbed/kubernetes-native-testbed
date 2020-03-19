module github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/order/server

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.5
	github.com/jinzhu/gorm v1.9.12
	github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/order v0.0.0-00010101000000-000000000000
	github.com/nats-io/nats.go v1.9.1
	golang.org/x/crypto v0.0.0-20200317142112-1b76d66859c6 // indirect
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a // indirect
	golang.org/x/sys v0.0.0-20200317113312-5766fd39f98d // indirect
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/genproto v0.0.0-20200319113533-08878b785e9c // indirect
	google.golang.org/grpc v1.28.0
)

replace github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/order => ../
