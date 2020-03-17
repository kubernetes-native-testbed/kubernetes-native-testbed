package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	pb "github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/point/protobuf"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var (
	dbUser     string
	dbPassword string
	dbName     string
	dbHost     string
	dbPort     string
	dbSSL      string
)

const (
	defaultBindAddr = ":8080"

	componentName     = "point"
	defaultDBUser     = componentName
	defaultDBPassword = componentName
	defaultDBName     = componentName
	defaultDBHost     = componentName
	defaultDBHPort    = componentName
	defaultDBSSL      = "enable"

	// defaultKVSUser = componentName
	// defaultKVSPass = componentName
	// defaultKVSHost = componentName
	// defaultKVSPort = componentName

	// defaultQueueHost = componentName
	// defaultQueuePort = componentName
	// defaultQueueTopic = componentName
	// defaultQueueUser = componentName
	// defaultQueuePassword = componentName
)

func init() {
	if dbUser = os.Getenv("DB_USER"); dbUser == "" {
		dbUser = defaultDBUser
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
	if dbPort = os.Getenv("DB_PORT"); dbPort == "" {
		dbPort = defaultDBHPort
	}
	if dbSSL = os.Getenv("DB_SSL"); dbSSL == "" {
		dbSSL = defaultDBHPort
	}
}

type pointAPIServer struct {
	pointRepository pointRepository
}

func (s *pointAPIServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	uuid := req.GetUUID()
	log.Printf("{\"operation\":\"get\", \"uuid\":\"%s\"}", uuid)
	p, err := s.pointRepository.findByUUID(uuid)
	if err != nil {
		return &pb.GetResponse{}, err
	}

	var resp pb.GetResponse
	var cat, uat, dat *timestamp.Timestamp
	if cat, err = ptypes.TimestampProto(p.CreatedAt); err != nil {
		return &pb.GetResponse{}, err
	}
	if uat, err = ptypes.TimestampProto(p.UpdatedAt); err != nil {
		return &pb.GetResponse{}, err
	}
	if p.DeletedAt != nil {
		if dat, err = ptypes.TimestampProto(*p.DeletedAt); err != nil {
			return &pb.GetResponse{}, err
		}
	}

	resp.Point = &pb.Point{
		UUID:        p.UUID,
		UserUUID:    p.UserUUID,
		Balance:     p.Balance,
		Description: p.Description,
		CreatedAt:   cat,
		UpdatedAt:   uat,
		DeletedAt:   dat,
	}

	return &resp, nil
}

func (s *pointAPIServer) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	p := &Point{
		UserUUID:    req.GetPoint().GetUserUUID(),
		Balance:     req.GetPoint().GetBalance(),
		Description: req.GetPoint().GetDescription(),
	}
	log.Printf("{\"operation\":\"set\", \"uuid\":\"%s\", \"user_uuid\":\"%s\",\"balance\":\"%d\", \"description\":\"%s\"}", p.UUID, p.UserUUID, p.Balance, p.Description)

	uuid, err := s.pointRepository.store(p)
	if err != nil {
		return &pb.SetResponse{}, err
	}

	return &pb.SetResponse{UUID: uuid}, nil
}

func (s *pointAPIServer) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	p := &Point{
		UUID:        req.GetPoint().GetUUID(),
		UserUUID:    req.GetPoint().GetUserUUID(),
		Balance:     req.GetPoint().GetBalance(),
		Description: req.GetPoint().GetDescription(),
	}
	log.Printf("{\"operation\":\"update\", \"uuid\":\"%s\", \"user_uuid\":\"%s\",\"balance\":\"%d\", \"description\":\"%s\"}", p.UUID, p.UserUUID, p.Balance, p.Description)

	if err := s.pointRepository.update(p); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *pointAPIServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	uuid := req.GetUUID()
	log.Printf("{\"operation\":\"delete\", \"uuid\":\"%s\"}", uuid)

	if err := s.pointRepository.deleteByUUID(uuid); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			dbHost,
			dbPort,
			dbUser,
			dbPassword,
			dbName,
			dbSSL,
		),
	)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()
	log.Printf("success for connection to %s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)

	s := grpc.NewServer()
	ap := &pointAPIServer{
		pointRepository: &pointRepositoryImpl{},
	}
	pb.RegisterPointAPIServer(s, ap)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("start point API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
