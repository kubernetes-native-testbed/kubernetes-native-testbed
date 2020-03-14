package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	pb "github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/paymentInfo/protobuf"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var ()

const (
	defaultBindAddr = ":8080"

	componentName = "paymentInfo"
)

func init() {
}

type paymentInfoAPIServer struct {
	paymentInfoRepository paymentInfoRepository
}

func (s *paymentInfoAPIServer) Get(ctx context.Context, req *pb.PaymentInfoGetRequest) (*pb.PaymentInfoGetResponse, error) {
	return nil, nil
}

func (s *paymentInfoAPIServer) Set(ctx context.Context, req *pb.PaymentInfoSetRequest) (*pb.PaymentInfoSetResponse, error) {
	return nil, nil
}

func (s *paymentInfoAPIServer) Update(ctx context.Context, req *pb.PaymentInfoUpdateRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *paymentInfoAPIServer) Delete(ctx context.Context, req *pb.PaymentInfoDeleteRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	s := grpc.NewServer()
	api := &paymentInfoAPIServer{
		paymentInfoRepository: &paymentInfoRepositoryImpl{},
	}
	pb.RegisterPaymentInfoAPIServer(s, api)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("start paymentInfo API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
