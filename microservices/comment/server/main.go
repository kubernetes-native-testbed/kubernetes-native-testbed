package main

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/__GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/comment"
	pb "github.com/__GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/comment/protobuf"
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

	componentName     = "comment"
	defaultDBUsername = componentName
	defaultDBPassword = componentName
	defaultDBName     = componentName
	defaultDBHost     = componentName
	defaultDBPort     = 3306
)

func init() {
	var err error
	if dbUsername = os.Getenv("DB_USERNAME"); dbUsername == "" {
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

type commentAPIServer struct {
	commentRepository comment.CommentRepository
}

func (s *commentAPIServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return nil, nil
}

func (s *commentAPIServer) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	return nil, nil
}

func (s *commentAPIServer) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *commentAPIServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	mongoConfig := comment.CommentRepositoryMongoConfig{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUsername,
		Password: dbPassword,
		DBName:   dbName,
	}
	mongo, closeMongo, err := mongoConfig.Connect()
	if err != nil {
		log.Fatalf("failed to open database: %v (config=%#v)", err, mongoConfig)
	}
	defer closeMongo()
	log.Printf("succeed to open database")

	commentAPI := &commentAPIServer{
		commentRepository: mongo,
	}

	s := grpc.NewServer()
	pb.RegisterCommentAPIServer(s, commentAPI)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("start comment API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
