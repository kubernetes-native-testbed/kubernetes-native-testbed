package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	pb "github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/deliveryStatus/protobuf"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var ()

const (
	defaultBindAddr = ":8080"

	componentName = "deliveryStatus"
)

func init() {
}

type deliveryStatusAPIServer struct {
	deliveryStatusRepository deliveryStatusRepository
}

func (s *deliveryStatusAPIServer) Get(ctx context.Context, req *pb.DeliveryStatusGetRequest) (*pb.DeliveryStatusGetResponse, error) {
	return nil, nil
}

func (s *deliveryStatusAPIServer) Set(ctx context.Context, req *pb.DeliveryStatusSetRequest) (*pb.DeliveryStatusSetResponse, error) {
	return nil, nil
}

func (s *deliveryStatusAPIServer) Update(ctx context.Context, req *pb.DeliveryStatusUpdateRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *deliveryStatusAPIServer) Delete(ctx context.Context, req *pb.DeliveryStatusDeleteRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	s := grpc.NewServer()
	api := &deliveryStatusAPIServer{
		deliveryStatusRepository: &deliveryStatusRepositoryImpl{},
	}
	pb.RegisterDeliveryStatusAPIServer(s, api)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("start deliveryStatus API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
