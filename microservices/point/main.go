package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	pb "github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/point/protobuf"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var ()

const (
	defaultBindAddr = ":8080"

	componentName = "point"
)

func init() {
}

type pointAPIServer struct {
	pointRepository pointRepository
}

func (s *pointAPIServer) Get(ctx context.Context, req *pb.PointGetRequest) (*pb.PointGetResponse, error) {
	return nil, nil
}

func (s *pointAPIServer) Set(ctx context.Context, req *pb.PointSetRequest) (*pb.PointSetResponse, error) {
	return nil, nil
}

func (s *pointAPIServer) Update(ctx context.Context, req *pb.PointUpdateRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *pointAPIServer) Delete(ctx context.Context, req *pb.PointDeleteRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	s := grpc.NewServer()
	api := &pointAPIServer{
		pointRepository: &pointRepositoryImpl{},
	}
	pb.RegisterPointAPIServer(s, api)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("start point API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
