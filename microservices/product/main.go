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
)

var (
	dbUser     string
	dbPassword string
	dbName     string
	dbAddress  string
)

const (
	defaultBindAddr = ":8080"

	componentName     = "product"
	defaultDBUser     = componentName
	defaultDBPassword = componentName
	defaultDBName     = componentName
	defaultDBAddress  = componentName
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
	if dbAddress = os.Getenv("DB_ADDRESS"); dbAddress == "" {
		dbAddress = defaultDBAddress
	}
}

type productAPIServer struct {
	productRepository productRepository
}

func (s *productAPIServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	uuid := req.GetUUID()
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
	if dat, err = ptypes.TimestampProto(*p.DeletedAt); err != nil {
		return &pb.GetResponse{}, err
	}
	resp.Product = &pb.Product{
		UUID:      p.UUID,
		Name:      p.Name,
		Price:     int32(p.Price),
		ImageURLs: p.ImageURLs,
		CreatedAt: cat,
		UpdatedAt: uat,
		DeletedAt: dat,
	}

	return &resp, nil
}

func (s *productAPIServer) Set(ctx context.Context, req *pb.SetRequest) (*empty.Empty, error) {
	p := &product{
		Name:      req.GetProduct().GetName(),
		Price:     int(req.GetProduct().GetPrice()),
		ImageURLs: req.GetProduct().GetImageURLs(),
	}

	_, err := s.productRepository.store(p)
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *productAPIServer) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	p := &product{
		UUID:      req.GetProduct().GetUUID(),
		Name:      req.GetProduct().GetName(),
		Price:     int(req.GetProduct().GetPrice()),
		ImageURLs: req.GetProduct().GetImageURLs(),
	}

	if err := s.productRepository.update(p); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *productAPIServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	uuid := req.GetUUID()

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

	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
			dbUser,
			dbPassword,
			dbAddress,
			dbName,
		),
	)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	log.Printf("start product API server")
	s := grpc.NewServer()
	pb.RegisterProductAPIServer(s, &productAPIServer{
		productRepository: &productRepositoryImpl{
			db: db,
		},
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
