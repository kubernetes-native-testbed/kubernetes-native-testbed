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

	s3Host            string
	s3Port            int
	s3AccessKeyID     string
	s3SecretAccessKey string
	s3BucketName      string
	s3PublicHost      string
	s3PublicPort      int
)

const (
	defaultBindAddr = ":8080"

	componentName     = "product"
	defaultDBUsername = componentName
	defaultDBPassword = componentName
	defaultDBName     = componentName
	defaultDBHost     = componentName
	defaultDBPort     = 3306

	defaultS3Host            = "minio-hl-svc.infra.svc.cluster.local"
	defaultS3Port            = 9000
	defaultS3AccessKeyID     = componentName
	defaultS3SecretAccessKey = componentName
	defaultS3BucketName      = componentName + "-image"
	defaultS3PublicHost      = "minio.34.84.105.184.nip.io"
	defaultS3PublicPort      = 443
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

	if s3Host = os.Getenv("S3_HOST"); s3Host == "" {
		s3Host = defaultS3Host
	}
	if s3Port, err = strconv.Atoi(os.Getenv("S3_PORT")); err != nil {
		s3Port = defaultS3Port
		log.Printf("s3Port parse error: %v", err)
	}
	if s3AccessKeyID = os.Getenv("S3_ACCESS_KEY_ID"); s3AccessKeyID == "" {
		s3AccessKeyID = defaultS3AccessKeyID
	}
	if s3SecretAccessKey = os.Getenv("S3_SECRET_ACCESS_KEY"); s3SecretAccessKey == "" {
		s3SecretAccessKey = defaultS3SecretAccessKey
	}
	if s3BucketName = os.Getenv("S3_BUCKET_NAME"); s3BucketName == "" {
		s3BucketName = defaultS3BucketName
	}
	if s3PublicHost = os.Getenv("S3_PUBLIC_HOST"); s3PublicHost == "" {
		s3PublicHost = defaultS3PublicHost
	}
	if s3PublicPort, err = strconv.Atoi(os.Getenv("S3_PUBLIC_PORT")); err != nil {
		s3PublicPort = defaultS3PublicPort
		log.Printf("s3PublicPort parse error: %v", err)
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

	productAPI := &productAPIServer{
		productRepository: mysql,
	}

	minioConfig := product.ImageRepositoryMinIOConfig{
		Host:            s3Host,
		Port:            s3Port,
		AccessKeyID:     s3AccessKeyID,
		SecretAccessKey: s3SecretAccessKey,
		BucketName:      s3BucketName,
		UseSSL:          false,
		PublicHost:      s3PublicHost,
		PublicPort:      s3PublicPort,
	}
	minio, err := minioConfig.Connect()
	if err != nil {
		log.Fatalf("failed to connect object storage: %v (config=%#v)", err, minioConfig)
	}
	log.Printf("succeed to connect object storage")

	imageAPI := &imageAPIServer{
		imageRepository: minio,
	}

	s := grpc.NewServer()
	pb.RegisterProductAPIServer(s, productAPI)
	pb.RegisterImageAPIServer(s, imageAPI)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("setup database")
	if err := productAPI.productRepository.InitDB(); err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	log.Printf("setup object storage")
	if err := imageAPI.imageRepository.InitRepository(); err != nil {
		log.Fatalf("failed to init object storage: %v", err)
	}

	log.Printf("start product API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
