package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	pb "github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/product/protobuf"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var (
	dbUser     string
	dbPassword string
	dbName     string
	dbHost     string
	dbPort     string
)

const (
	defaultBindAddr = ":8080"

	componentName     = "product"
	defaultDBUser     = componentName
	defaultDBPassword = componentName
	defaultDBName     = componentName
	defaultDBHost     = componentName
	defaultDBHPort    = componentName
)

func init() {
	if dbUser = os.Getenv("DB_USER"); dbUser == "" {
		dbUser = defaultDBUser
	}
	if dbPassword = os.Getenv("DB_PASSWORD"); dbPassword == "" {
		dbPassword = defaultDBPassword
	}
	if dbName = os.Getenv("DB_NAME"); dbName == "" {
		dbName = defaultDBName
	}
	if dbHost = os.Getenv("DB_HOST"); dbHost == "" {
		dbHost = defaultDBHost
	}
	if dbPort = os.Getenv("DB_PORT"); dbPort == "" {
		dbPort = defaultDBHPort
	}
}

type productAPIServer struct {
	productRepository productRepository
}

func (s *productAPIServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	uuid := req.GetUUID()
	log.Printf("{\"operation\":\"get\", \"uuid\":\"%s\"}", uuid)
	p, err := s.productRepository.findByUUID(uuid)
	if err != nil {
		return &pb.GetResponse{}, err
	}

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
	urls := make([]ProductImage, len(req.GetProduct().GetImageURLs()))
	for i, url := range req.GetProduct().GetImageURLs() {
		urls[i] = ProductImage{URL: url}
	}

	p := &Product{
		Name:      req.GetProduct().GetName(),
		Price:     req.GetProduct().GetPrice(),
		ImageURLs: urls,
	}
	log.Printf("{\"operation\":\"set\", \"uuid\":\"%s\", \"name\":\"%s\", \"price\":\"%d\", \"image_urls\":\"%v\"}", p.UUID, p.Name, p.Price, p.ImageURLs)

	uuid, err := s.productRepository.store(p)
	if err != nil {
		return &pb.SetResponse{}, err
	}

	return &pb.SetResponse{UUID: uuid}, nil
}

func (s *productAPIServer) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	urls := make([]ProductImage, len(req.GetProduct().GetImageURLs()))
	for i, url := range req.GetProduct().GetImageURLs() {
		urls[i] = ProductImage{ProductUUID: req.GetProduct().GetUUID(), URL: url}
	}
	p := &Product{
		UUID:      req.GetProduct().GetUUID(),
		Name:      req.GetProduct().GetName(),
		Price:     req.GetProduct().GetPrice(),
		ImageURLs: urls,
	}
	log.Printf("{\"operation\":\"update\", \"uuid\":\"%s\", \"name\":\"%s\", \"price\":\"%d\", \"image_urls\":\"%v\"}", p.UUID, p.Name, p.Price, p.ImageURLs)

	if err := s.productRepository.update(p); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *productAPIServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	uuid := req.GetUUID()
	log.Printf("{\"operation\":\"delete\", \"uuid\":\"%s\"}", uuid)

	if err := s.productRepository.deleteByUUID(uuid); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbUser,
			dbPassword,
			dbHost,
			dbPort,
			dbName,
		),
	)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()
	log.Printf("success for connection to %s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	s := grpc.NewServer()
	api := &productAPIServer{
		productRepository: &productRepositoryImpl{
			db: db,
		},
	}
	pb.RegisterProductAPIServer(s, api)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("setup database")
	if err := api.productRepository.initDB(); err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	log.Printf("start product API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
