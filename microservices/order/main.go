package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	pb "github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/order/protobuf"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var ()

const (
	defaultBindAddr = ":8080"

	componentName = "order"
)

func init() {
}

type orderAPIServer struct {
	orderRepository orderRepository
}

func (s *orderAPIServer) Get(ctx context.Context, req *pb.OrderGetRequest) (*pb.OrderGetResponse, error) {
	return nil, nil
}

func (s *orderAPIServer) Set(ctx context.Context, req *pb.OrderSetRequest) (*pb.OrderSetResponse, error) {
	return nil, nil
}

func (s *orderAPIServer) Update(ctx context.Context, req *pb.OrderUpdateRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *orderAPIServer) Delete(ctx context.Context, req *pb.OrderDeleteRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	s := grpc.NewServer()
	api := &orderAPIServer{
		orderRepository: &orderRepositoryImpl{},
	}
	pb.RegisterOrderAPIServer(s, api)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("start order API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
