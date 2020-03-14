package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	pb "github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/comment/protobuf"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var ()

const (
	defaultBindAddr = ":8080"

	componentName = "comment"
)

func init() {
}

type commentAPIServer struct {
	commentRepository commentRepository
}

func (s *commentAPIServer) Get(ctx context.Context, req *pb.CommentGetRequest) (*pb.CommentGetResponse, error) {
	return nil, nil
}

func (s *commentAPIServer) Set(ctx context.Context, req *pb.CommentSetRequest) (*pb.CommentSetResponse, error) {
	return nil, nil
}

func (s *commentAPIServer) Update(ctx context.Context, req *pb.CommentUpdateRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *commentAPIServer) Delete(ctx context.Context, req *pb.CommentDeleteRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	s := grpc.NewServer()
	api := &commentAPIServer{
		commentRepository: &commentRepositoryImpl{},
	}
	pb.RegisterCommentAPIServer(s, api)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("start comment API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
