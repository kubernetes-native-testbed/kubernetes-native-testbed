package main

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/__GITHUB_REPO_NAME__/kubernetes-native-testbed/microservices/rate"
	pb "github.com/__GITHUB_REPO_NAME__/kubernetes-native-testbed/microservices/rate/protobuf"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var (
	sentinelHost string
	sentinelPort int
	kvsPassword  string
	masterName   string
)

const (
	defaultBindAddr = ":8080"

	componentName       = "rate"
	defaultSentinelHost = componentName
	defaultSentinelPort = 26379
	defaultKvsPassword  = componentName
	defaultMasterName   = "mymaster"
)

func init() {
	var err error
	if sentinelHost = os.Getenv("SENTINEL_HOST"); sentinelHost == "" {
		sentinelHost = defaultSentinelHost
	}
	if sentinelPort, err = strconv.Atoi(os.Getenv("SENTINEL_PORT")); err != nil {
		sentinelPort = defaultSentinelPort
		log.Printf("sentinelPort parse error: %v", err)
	}
	if kvsPassword = os.Getenv("KVS_PASSWORD"); kvsPassword == "" {
		kvsPassword = defaultKvsPassword
	}
	if masterName = os.Getenv("REDIS_MASTER_NAME"); masterName == "" {
		masterName = defaultMasterName
	}
}

type rateAPIServer struct {
	rateRepository rate.RateRepository
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

	redisConfig := rate.RateRepositoryRedisConfig{
		SentinelHost: sentinelHost,
		SentinelPort: sentinelPort,
		Password:     kvsPassword,
		MasterName:   masterName,
	}
	redis := redisConfig.Connect()
	log.Printf("succeed to open kvs")

	rateAPI := &rateAPIServer{
		rateRepository: redis,
	}

	s := grpc.NewServer()
	pb.RegisterRateAPIServer(s, rateAPI)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("start rate API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
