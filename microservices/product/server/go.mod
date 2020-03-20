module github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/product/server

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.5
	github.com/jinzhu/gorm v1.9.12
	github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/product v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200319234117-63522dbf7eec // indirect
	golang.org/x/sys v0.0.0-20200317113312-5766fd39f98d // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200319113533-08878b785e9c // indirect
	google.golang.org/grpc v1.28.0
)

replace github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/product => ../
