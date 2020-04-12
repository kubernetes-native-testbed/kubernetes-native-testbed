package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/__GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/user"
	pb "github.com/__GITHUB_ORG_NAME__/kubernetes-native-testbed/microservices/user/protobuf"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	passwordHashSalt = "uruamsy"
)

var (
	dbUser     string
	dbPassword string
	dbName     string
	dbHost     string
	dbPort     int

	privateKey string
)

const (
	defaultBindAddr = ":8080"

	componentName     = "user"
	defaultDBUser     = componentName
	defaultDBPassword = componentName
	defaultDBName     = componentName
	defaultDBHost     = componentName
	defaultDBPort     = 3306
)

func init() {
	var err error
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
	if dbPort, err = strconv.Atoi(os.Getenv("DB_PORT")); err != nil {
		dbPort = defaultDBPort
		log.Printf("dbPort parse error: %v", err)
	}
	if privateKey = os.Getenv("PRIVATE_KEY"); privateKey == "" {
		log.Fatal("PRIVATE_KEY is required")
	}
}

type userAPIServer struct {
	userRepository user.UserRepository
}

func (s *userAPIServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	uuid := req.GetUUID()
	log.Printf("{\"operation\":\"get\", \"uuid\":\"%s\"}", uuid)
	u, nferr, err := s.userRepository.FindByUUID(uuid)
	if err != nil {
		return &pb.GetResponse{}, err
	}
	if nferr != nil {
		return &pb.GetResponse{}, nferr
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
	addresses := make([]user.Address, len(req.GetUser().GetAddresses()))
	for i, address := range req.GetUser().GetAddresses() {
		addresses[i] = user.Address{
			ZipCode:     address.ZipCode,
			Country:     address.Country,
			State:       address.State,
			City:        address.City,
			AddressLine: address.AddressLine,
			Disabled:    address.Disabled,
		}
	}

	passwordHash, err := generatePasswordHash(req.GetUser().GetPassword())
	if err != nil {
		return nil, err
	}

	u := &user.User{
		Username:               req.GetUser().GetUsername(),
		FirstName:              req.GetUser().GetFirstName(),
		LastName:               req.GetUser().GetLastName(),
		Age:                    req.GetUser().GetAge(),
		PasswordHash:           passwordHash,
		DefaultPaymentInfoUUID: req.GetUser().GetDefaultPaymentInfoUUID(),
		Addresses:              addresses,
	}
	log.Printf("{\"operation\":\"set\", \"uuid\":\"%s\", \"username\":\"%s\", \"first_name\":\"%s\", \"last_name\":\"%s\", \"age\":\"%d\", \"password_hash\":\"%s\", \"default_payment_info_uuid\":\"%s\", \"addresses\":\"%v\"}",
		u.UUID, u.Username, u.FirstName, u.LastName, u.Age, u.PasswordHash, u.DefaultPaymentInfoUUID, u.Addresses)

	uuid, err := s.userRepository.Store(u)
	if err != nil {
		return &pb.SetResponse{}, err
	}

	return &pb.SetResponse{UUID: uuid}, nil
}

func (s *userAPIServer) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	addresses := make([]user.Address, len(req.GetUser().GetAddresses()))
	for i, address := range req.GetUser().GetAddresses() {
		addresses[i] = user.Address{
			UUID:        address.UUID,
			ZipCode:     address.ZipCode,
			Country:     address.Country,
			State:       address.State,
			City:        address.City,
			AddressLine: address.AddressLine,
			Disabled:    address.Disabled,
		}
	}

	passwordHash, err := generatePasswordHash(req.GetUser().GetPassword())
	if err != nil {
		return nil, err
	}

	u := &user.User{
		UUID:                   req.GetUser().GetUUID(),
		Username:               req.GetUser().GetUsername(),
		FirstName:              req.GetUser().GetFirstName(),
		LastName:               req.GetUser().GetLastName(),
		Age:                    req.GetUser().GetAge(),
		PasswordHash:           passwordHash,
		DefaultPaymentInfoUUID: req.GetUser().GetDefaultPaymentInfoUUID(),
		Addresses:              addresses,
	}
	log.Printf("{\"operation\":\"update\", \"uuid\":\"%s\", \"username\":\"%s\", \"first_name\":\"%s\", \"last_name\":\"%s\", \"age\":\"%d\", \"password_hash\":\"%s\", \"default_payment_info_uuid\":\"%s\", \"addresses\":\"%v\"}",
		u.UUID, u.Username, u.FirstName, u.LastName, u.Age, u.PasswordHash, u.DefaultPaymentInfoUUID, u.Addresses)

	if err := s.userRepository.Update(u); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *userAPIServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	uuid := req.GetUUID()
	log.Printf("{\"operation\":\"delete\", \"uuid\":\"%s\"}", uuid)

	if err := s.userRepository.DeleteByUUID(uuid); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *userAPIServer) IsExists(ctx context.Context, req *pb.IsExistsRequest) (*pb.IsExistsResponse, error) {
	uuid := req.GetUUID()
	log.Printf("isExists: {\"uuid\":\"%s\"}", uuid)
	_, nferr, err := s.userRepository.FindByUUID(uuid)
	if err != nil {
		return nil, err
	}
	return &pb.IsExistsResponse{IsExists: nferr == nil}, nil
}

func (s *userAPIServer) Authentication(ctx context.Context, req *pb.AuthenticationRequest) (*pb.AuthenticationResponse, error) {
	uuid := req.GetUUID()
	log.Printf("authentication: {\"uuid\":\"%s\"}", uuid)
	u, nferr, err := s.userRepository.FindByUUID(uuid)
	if err != nil {
		return nil, err
	}
	if nferr != nil {
		return nil, nferr
	}

	passwordHash, err := generatePasswordHash(req.GetPassword())
	if err != nil {
		return nil, err
	}

	if u.PasswordHash != passwordHash {
		return nil, fmt.Errorf("user_uuid or password is not match (userUUID=%s)", uuid)
	}

	signKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return nil, err
	}

	token := jwt.New(jwt.SigningMethodES512)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_uuid"] = uuid
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return nil, err
	}

	return &pb.AuthenticationResponse{Token: tokenString}, nil
}

func generatePasswordHash(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("empty password is not allowed")
	}
	return fmt.Sprintf("%X", md5.Sum([]byte(passwordHashSalt+password))), nil
}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	mysqlConfig := user.UserRepositoryMySQLConfig{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUser,
		Password: dbPassword,
		DBName:   dbName,
	}
	mysql, closeMySQL, err := mysqlConfig.Connect()
	if err != nil {
		log.Fatalf("failed to open database: %v (config=%#v)", err, mysqlConfig)
	}
	defer closeMySQL()
	log.Printf("succeed to open database")

	s := grpc.NewServer()
	api := &userAPIServer{
		userRepository: mysql,
	}
	pb.RegisterUserAPIServer(s, api)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("setup database")
	if err := api.userRepository.InitDB(); err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	log.Printf("start user API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
