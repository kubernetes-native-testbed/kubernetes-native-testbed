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
	pb "github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/user/protobuf"
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

	componentName     = "user"
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

type userAPIServer struct {
	userRepository userRepository
}

func (s *userAPIServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	uuid := req.GetUUID()
	log.Printf("{\"operation\":\"get\", \"uuid\":\"%s\"}", uuid)
	u, err := s.userRepository.findByUUID(uuid)
	if err != nil {
		return &pb.GetResponse{}, err
	}

	var resp pb.GetResponse
	var cat, uat, dat *timestamp.Timestamp
	if cat, err = ptypes.TimestampProto(u.CreatedAt); err != nil {
		return &pb.GetResponse{}, err
	}
	if uat, err = ptypes.TimestampProto(u.UpdatedAt); err != nil {
		return &pb.GetResponse{}, err
	}
	if u.DeletedAt != nil {
		if dat, err = ptypes.TimestampProto(*u.DeletedAt); err != nil {
			return &pb.GetResponse{}, err
		}
	}

	addresses := make([]*pb.Address, len(u.Addresses))
	for i := range u.Addresses {
		addresses[i] = &pb.Address{
			UUID:        u.Addresses[i].UUID,
			UserUUID:    u.Addresses[i].UserUUID,
			ZipCode:     u.Addresses[i].ZipCode,
			Country:     u.Addresses[i].Country,
			State:       u.Addresses[i].State,
			City:        u.Addresses[i].City,
			AddressLine: u.Addresses[i].AddressLine,
			Disabled:    u.Addresses[i].Disabled,
		}
	}

	resp.User = &pb.User{
		UUID:                   u.UUID,
		Username:               u.Username,
		FirstName:              u.FirstName,
		LastName:               u.LastName,
		Age:                    u.Age,
		Addresses:              addresses,
		PasswordHash:           u.PasswordHash,
		DefaultPaymentInfoUUID: u.DefaultPaymentInfoUUID,
		CreatedAt:              cat,
		UpdatedAt:              uat,
		DeletedAt:              dat,
	}

	return &resp, nil
}

func (s *userAPIServer) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	addresses := make([]Address, len(req.GetUser().GetAddresses()))
	for i, address := range req.GetUser().GetAddresses() {
		addresses[i] = Address{
			ZipCode:     address.ZipCode,
			Country:     address.Country,
			State:       address.State,
			City:        address.City,
			AddressLine: address.AddressLine,
			Disabled:    address.Disabled,
		}
	}

	u := &User{
		Username:               req.GetUser().GetUsername(),
		FirstName:              req.GetUser().GetFirstName(),
		LastName:               req.GetUser().GetLastName(),
		Age:                    req.GetUser().GetAge(),
		PasswordHash:           req.GetUser().GetPasswordHash(),
		DefaultPaymentInfoUUID: req.GetUser().GetDefaultPaymentInfoUUID(),
		Addresses:              addresses,
	}
	log.Printf("{\"operation\":\"set\", \"uuid\":\"%s\", \"username\":\"%s\", \"first_name\":\"%s\", \"last_name\":\"%s\", \"age\":\"%d\", \"password_hash\":\"%s\", \"default_payment_info_uuid\":\"%s\", \"addresses\":\"%v\"}",
		u.UUID, u.Username, u.FirstName, u.LastName, u.Age, u.PasswordHash, u.DefaultPaymentInfoUUID, u.Addresses)

	uuid, err := s.userRepository.store(u)
	if err != nil {
		return &pb.SetResponse{}, err
	}

	return &pb.SetResponse{UUID: uuid}, nil
}

func (s *userAPIServer) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	addresses := make([]Address, len(req.GetUser().GetAddresses()))
	for i, address := range req.GetUser().GetAddresses() {
		addresses[i] = Address{
			UUID:        address.UUID,
			ZipCode:     address.ZipCode,
			Country:     address.Country,
			State:       address.State,
			City:        address.City,
			AddressLine: address.AddressLine,
			Disabled:    address.Disabled,
		}
	}
	u := &User{
		UUID:                   req.GetUser().GetUUID(),
		Username:               req.GetUser().GetUsername(),
		FirstName:              req.GetUser().GetFirstName(),
		LastName:               req.GetUser().GetLastName(),
		Age:                    req.GetUser().GetAge(),
		PasswordHash:           req.GetUser().GetPasswordHash(),
		DefaultPaymentInfoUUID: req.GetUser().GetDefaultPaymentInfoUUID(),
		Addresses:              addresses,
	}
	log.Printf("{\"operation\":\"update\", \"uuid\":\"%s\", \"username\":\"%s\", \"first_name\":\"%s\", \"last_name\":\"%s\", \"age\":\"%d\", \"password_hash\":\"%s\", \"default_payment_info_uuid\":\"%s\", \"addresses\":\"%v\"}",
		u.UUID, u.Username, u.FirstName, u.LastName, u.Age, u.PasswordHash, u.DefaultPaymentInfoUUID, u.Addresses)

	if err := s.userRepository.update(u); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *userAPIServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	uuid := req.GetUUID()
	log.Printf("{\"operation\":\"delete\", \"uuid\":\"%s\"}", uuid)

	if err := s.userRepository.deleteByUUID(uuid); err != nil {
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
	api := &userAPIServer{
		userRepository: &userRepositoryImpl{
			db: db,
		},
	}
	pb.RegisterUserAPIServer(s, api)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("setup database")
	if err := api.userRepository.initDB(); err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	log.Printf("start user API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
