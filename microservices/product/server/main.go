package main

import (
	"log"
	"net"
	"os"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/product"
	pb "github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/product/protobuf"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var (
	dbUsername string
	dbPassword string
	dbName     string
	dbHost     string
	dbPort     int
)

const (
	defaultBindAddr = ":8080"

	componentName     = "product"
	defaultDBUsername = componentName
	defaultDBPassword = componentName
	defaultDBName     = componentName
	defaultDBHost     = componentName
	defaultDBPort     = 3306
)

func init() {
	var err error
	if dbUsername = os.Getenv("DB_USER"); dbUsername == "" {
		dbUsername = defaultDBUsername
	}
	if dbPassword = os.Getenv("DB_PASSWORD"); dbPassword == "" {
		dbPassword = defaultDBPassword
	}
	if dbName = os.Getenv("DB_NAME"); dbName == "" {
		dbName = defaultDBName
	}
	if dbHost = os.Getenv("DB_HOST"); dbHost == "" {
		dbHost = defaultDBHost
	}
	if dbPort, err = strconv.Atoi(os.Getenv("DB_PORT")); err != nil {
		dbPort = defaultDBPort
		log.Printf("dbPort parse error: %v", err)
	}
}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	mysqlConfig := product.ProductRepositoryMySQLConfig{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUsername,
		Password: dbPassword,
		DBName:   dbName,
	}
	mysql, closeMySQL, err := mysqlConfig.Connect()
	if err != nil {
		log.Fatalf("failed to open database: %v (config=%#v)", err, mysqlConfig)
	}
	defer closeMySQL()
	log.Printf("succeed to open database")

	s := grpc.NewServer()
	api := &productAPIServer{
		productRepository: mysql,
	}
	pb.RegisterProductAPIServer(s, api)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("setup database")
	if err := api.productRepository.InitDB(); err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	log.Printf("start product API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
