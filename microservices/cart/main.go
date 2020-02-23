package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	pb "github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/cart/protobuf"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var ()

const (
	defaultBindAddr = ":8080"

	componentName = "cart"
)

func init() {
}

type cartAPIServer struct {
	cartRepository cartRepository
}

func (s *cartAPIServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return nil, nil
}

func (s *cartAPIServer) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	return nil, nil
}

func (s *cartAPIServer) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *cartAPIServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	s := grpc.NewServer()
	api := &cartAPIServer{
		cartRepository: &cartRepositoryImpl{},
	}
	pb.RegisterCartAPIServer(s, api)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("start cart API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
