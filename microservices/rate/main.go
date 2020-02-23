package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	pb "github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/rate/protobuf"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var (
	sentinelHost string
	kvsPassword  string
)

const (
	defaultBindAddr = ":8080"

	componentName       = "rate"
	defaultSentinelHost = componentName
	defaultKvsPassword  = componentName
)

func init() {
	if sentinelHost = os.Getenv("SENTINEL_HOST"); sentinelHost == "" {
		sentinelHost = defaultSentinelHost
	}
	if kvsPassword = os.Getenv("KVS_PASSWORD"); kvsPassword == "" {
		kvsPassword = defaultKvsPassword
	}
}

type rateAPIServer struct {
	rateRepository rateRepository
}

func (s *rateAPIServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return nil, nil
}

func (s *rateAPIServer) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	return nil, nil
}

func (s *rateAPIServer) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *rateAPIServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	s := grpc.NewServer()
	api := &rateAPIServer{
		rateRepository: &rateRepositoryImpl{},
	}
	pb.RegisterRateAPIServer(s, api)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("start rate API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
