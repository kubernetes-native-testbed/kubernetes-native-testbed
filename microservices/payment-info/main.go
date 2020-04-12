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
	_ "github.com/jinzhu/gorm/dialects/postgres"
	pb "github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/payment-info/protobuf"
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
	dbSSL      string
)

const (
	defaultBindAddr = ":8080"

	componentName     = "paymentInfo"
	defaultDBUser     = componentName
	defaultDBPassword = componentName
	defaultDBName     = componentName
	defaultDBHost     = componentName
	defaultDBHPort    = componentName
	defaultDBSSL      = "require"
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
	if dbSSL = os.Getenv("DB_SSL"); dbSSL == "" {
		dbSSL = defaultDBSSL
	}
}

type paymentInfoAPIServer struct {
	paymentInfoRepository paymentInfoRepository
}

func (s *paymentInfoAPIServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	uuid := req.GetUUID()
	log.Printf("{\"operation\":\"get\", \"uuid\":\"%s\"}", uuid)
	pi, err := s.paymentInfoRepository.findByUUID(uuid)
	if err != nil {
		return &pb.GetResponse{}, err
	}

	var resp pb.GetResponse
	var cat, uat, dat *timestamp.Timestamp
	if cat, err = ptypes.TimestampProto(pi.CreatedAt); err != nil {
		return &pb.GetResponse{}, err
	}
	if uat, err = ptypes.TimestampProto(pi.UpdatedAt); err != nil {
		return &pb.GetResponse{}, err
	}
	if pi.DeletedAt != nil {
		if dat, err = ptypes.TimestampProto(*pi.DeletedAt); err != nil {
			return &pb.GetResponse{}, err
		}
	}

	resp.PaymentInfo = &pb.PaymentInfo{
		UUID:           pi.UUID,
		UserUUID:       pi.UserUUID,
		Name:           pi.Name,
		CardNumber:     pi.CardNumber,
		SecurityCode:   pi.SecurityCode,
		ExpirationDate: pi.ExpirationDate,
		CreatedAt:      cat,
		UpdatedAt:      uat,
		DeletedAt:      dat,
	}

	return &resp, nil
}

func (s *paymentInfoAPIServer) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	pi := &PaymentInfo{
		UserUUID:       req.GetPaymentInfo().GetUserUUID(),
		Name:           req.GetPaymentInfo().GetName(),
		CardNumber:     req.GetPaymentInfo().GetCardNumber(),
		SecurityCode:   req.GetPaymentInfo().GetSecurityCode(),
		ExpirationDate: req.GetPaymentInfo().GetExpirationDate(),
	}
	log.Printf("{\"operation\":\"set\", \"uuid\":\"%s\", \"user_uuid\":\"%s\",\"name\":\"%s\", \"card_number\":\"%s\", \"security_code\":\"%s\", \"expiration_date\":\"%s\"}", pi.Name, pi.UserUUID, pi.CardNumber, pi.SecurityCode, pi.ExpirationDate)

	uuid, err := s.paymentInfoRepository.store(pi)
	if err != nil {
		return &pb.SetResponse{}, err
	}

	return &pb.SetResponse{UUID: uuid}, nil
}

func (s *paymentInfoAPIServer) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	pi := &PaymentInfo{
		UUID:           req.GetPaymentInfo().GetUUID(),
		Name:           req.GetPaymentInfo().GetName(),
		UserUUID:       req.GetPaymentInfo().GetUserUUID(),
		CardNumber:     req.GetPaymentInfo().GetCardNumber(),
		SecurityCode:   req.GetPaymentInfo().GetSecurityCode(),
		ExpirationDate: req.GetPaymentInfo().GetExpirationDate(),
	}
	log.Printf("{\"operation\":\"update\", \"uuid\":\"%s\", \"user_uuid\":\"%s\",\"name\":\"%s\", \"card_number\":\"%s\", \"security_code\":\"%s\", \"expiration_date\":\"%s\"}", pi.Name, pi.UserUUID, pi.CardNumber, pi.SecurityCode, pi.ExpirationDate)

	if err := s.paymentInfoRepository.update(pi); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *paymentInfoAPIServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	uuid := req.GetUUID()
	log.Printf("{\"operation\":\"delete\", \"uuid\":\"%s\"}", uuid)

	if err := s.paymentInfoRepository.deleteByUUID(uuid); err != nil {
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
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			dbHost,
			dbPort,
			dbUser,
			dbPassword,
			dbName,
			dbSSL,
		),
	)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()
	log.Printf("success for connection to %s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)

	s := grpc.NewServer()
	api := &paymentInfoAPIServer{
		paymentInfoRepository: &paymentInfoRepositoryImpl{
			db: db,
		},
	}
	pb.RegisterPaymentInfoAPIServer(s, api)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("setup database")
	if err := api.paymentInfoRepository.initDB(); err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	log.Printf("start paymentInfo API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
