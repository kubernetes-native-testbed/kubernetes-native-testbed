package main

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/product"
	pb "github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/product/protobuf"
)

type productAPIServer struct {
	productRepository product.ProductRepository
}

func (s *productAPIServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	uuid := req.GetUUID()
	p, nferr, err := s.productRepository.FindByUUID(uuid)
	if err != nil {
		return &pb.GetResponse{}, err
	}
	if nferr != nil {
		return &pb.GetResponse{}, nferr
	}
	log.Printf("get %s", p)

	var resp pb.GetResponse
	var cat, uat, dat *timestamp.Timestamp
	if cat, err = ptypes.TimestampProto(p.CreatedAt); err != nil {
		return &pb.GetResponse{}, err
	}
	if uat, err = ptypes.TimestampProto(p.UpdatedAt); err != nil {
		return &pb.GetResponse{}, err
	}
	if p.DeletedAt != nil {
		if dat, err = ptypes.TimestampProto(*p.DeletedAt); err != nil {
			return &pb.GetResponse{}, err
		}
	}

	urls := make([]string, len(p.ImageURLs))
	for i := range p.ImageURLs {
		urls[i] = p.ImageURLs[i].URL
	}

	resp.Product = &pb.Product{
		UUID:      p.UUID,
		Name:      p.Name,
		Price:     p.Price,
		ImageURLs: urls,
		CreatedAt: cat,
		UpdatedAt: uat,
		DeletedAt: dat,
	}

	return &resp, nil
}

func (s *productAPIServer) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	urls := make([]product.ProductImage, len(req.GetProduct().GetImageURLs()))
	for i, url := range req.GetProduct().GetImageURLs() {
		urls[i] = product.ProductImage{URL: url}
	}

	p := &product.Product{
		Name:      req.GetProduct().GetName(),
		Price:     req.GetProduct().GetPrice(),
		ImageURLs: urls,
	}
	log.Printf("set %s", p)

	uuid, err := s.productRepository.Store(p)
	if err != nil {
		return &pb.SetResponse{}, err
	}

	return &pb.SetResponse{UUID: uuid}, nil
}

func (s *productAPIServer) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	urls := make([]product.ProductImage, len(req.GetProduct().GetImageURLs()))
	for i, url := range req.GetProduct().GetImageURLs() {
		urls[i] = product.ProductImage{ProductUUID: req.GetProduct().GetUUID(), URL: url}
	}
	p := &product.Product{
		UUID:      req.GetProduct().GetUUID(),
		Name:      req.GetProduct().GetName(),
		Price:     req.GetProduct().GetPrice(),
		ImageURLs: urls,
	}
	log.Printf("update %s", p)

	if err := s.productRepository.Update(p); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *productAPIServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	uuid := req.GetUUID()
	log.Printf("delete %s", uuid)

	if err := s.productRepository.DeleteByUUID(uuid); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *productAPIServer) IsExists(ctx context.Context, req *pb.IsExistsRequest) (*pb.IsExistsResponse, error) {
	uuid := req.GetUUID()
	log.Printf("isExists: %s", uuid)
	_, nferr, err := s.productRepository.FindByUUID(uuid)
	if err != nil {
		return nil, err
	}
	return &pb.IsExistsResponse{IsExists: nferr == nil}, nil
}
