package main

import (
	"context"
	"io"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/__TB_GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/product"
	pb "github.com/__TB_GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/product/protobuf"
)

type imageAPIServer struct {
	imageRepository product.ImageRepository
}

func (s *imageAPIServer) Upload(ctx context.Context, req *pb.ImageUploadRequest) (*pb.ImageUploadResponse, error) {
	url, err := s.imageRepository.Store(req.GetImage())
	if err != nil {
		return nil, err
	}
	return &pb.ImageUploadResponse{Url: url}, nil
}

func (s *imageAPIServer) UploadStream(stream pb.ImageAPI_UploadStreamServer) error {
	image := make([]byte, 0, 1_000_000)
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		image = append(image, req.GetImage()...)
	}

	url, err := s.imageRepository.Store(image)
	if err != nil {
		return err
	}

	if err := stream.SendAndClose(&pb.ImageUploadStreamResponse{Url: url}); err != nil {
		return err
	}

	return nil
}

func (s *imageAPIServer) Delete(ctx context.Context, req *pb.ImageDeleteRequest) (*empty.Empty, error) {
	if err := s.imageRepository.Delete(req.GetUrl()); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
