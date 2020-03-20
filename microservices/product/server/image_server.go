package main

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/product"
	pb "github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/product/protobuf"
)

type imageAPIServer struct {
	imageRepository product.ImageRepository
}

func (s *imageAPIServer) Upload(stream pb.ImageAPI_UploadServer) error {
	return nil

}

func (s *imageAPIServer) Delete(ctx context.Context, req *pb.ImageDeleteRequest) (*empty.Empty, error) {
	return nil, nil

}
