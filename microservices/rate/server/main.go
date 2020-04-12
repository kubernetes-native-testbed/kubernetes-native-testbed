package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	commentpb "github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/comment/protobuf"
	"github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/rate"
	pb "github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/rate/protobuf"
	"github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/user"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var (
	sentinelHost string
	sentinelPort int
	kvsPassword  string
	masterName   string

	authPublicKey string

	commentHost string
	commentPort int
)

const (
	defaultBindAddr = ":8080"

	componentName       = "rate"
	defaultSentinelHost = componentName
	defaultSentinelPort = 26379
	defaultKvsPassword  = componentName
	defaultMasterName   = "mymaster"

	defaultCommentHost = "comment.comment.svc.cluster.local"
	defaultCommentPort = 8080
)

func init() {
	var err error
	if sentinelHost = os.Getenv("SENTINEL_HOST"); sentinelHost == "" {
		sentinelHost = defaultSentinelHost
	}
	if sentinelPort, err = strconv.Atoi(os.Getenv("SENTINEL_PORT")); err != nil {
		sentinelPort = defaultSentinelPort
		log.Printf("sentinelPort parse error: %v", err)
	}
	if kvsPassword = os.Getenv("KVS_PASSWORD"); kvsPassword == "" {
		kvsPassword = defaultKvsPassword
	}
	if masterName = os.Getenv("REDIS_MASTER_NAME"); masterName == "" {
		masterName = defaultMasterName
	}
	if authPublicKey = os.Getenv("AUTH_PUBLIC_KEY"); authPublicKey == "" {
		log.Fatal("AUTH_PUBLIC_KEY is required")
	}
	if commentHost = os.Getenv("COMMENT_HOST"); commentHost == "" {
		commentHost = defaultCommentHost
	}
	if commentPort, err = strconv.Atoi(os.Getenv("COMMENT_PORT")); err != nil {
		commentPort = defaultCommentPort
		log.Printf("commentPort parse error: %v", err)
	}
}

type rateAPIServer struct {
	rateRepository  rate.RateRepository
	commentClient   commentpb.CommentAPIClient
	commentEndpoint string
}

func (s *rateAPIServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	uuid := req.GetUUID()
	r, err := s.rateRepository.FindByUUID(uuid)
	if err != nil {
		if rate.IsNotFound(err) {
			return &pb.GetResponse{}, err
		}
		return &pb.GetResponse{}, err
	}
	log.Printf("get %s", r)

	var resp pb.GetResponse
	var cat, uat, dat *timestamp.Timestamp
	if cat, err = ptypes.TimestampProto(time.Unix(r.CreatedAt, 0)); err != nil {
		return &pb.GetResponse{}, err
	}
	if uat, err = ptypes.TimestampProto(time.Unix(r.UpdatedAt, 0)); err != nil {
		return &pb.GetResponse{}, err
	}
	if dat, err = ptypes.TimestampProto(time.Unix(r.DeletedAt, 0)); err != nil {
		return &pb.GetResponse{}, err
	}

	resp.Rate = &pb.Rate{
		UUID:        r.UUID,
		UserUUID:    r.UserUUID,
		CommentUUID: r.CommentUUID,
		ProductUUID: r.ProductUUID,
		Rating:      r.Rating,
		CreatedAt:   cat,
		UpdatedAt:   uat,
		DeletedAt:   dat,
	}

	return &resp, nil
}

func (s *rateAPIServer) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	r := &rate.Rate{
		UserUUID:    req.GetRate().GetUserUUID(),
		CommentUUID: req.GetRate().GetCommentUUID(),
		ProductUUID: req.GetRate().GetProductUUID(),
		Rating:      req.GetRate().GetRating(),
	}
	log.Printf("set %s", r)

	retry := 5
	var err error
	for i := 0; i < retry; i++ {
		if _, err = s.commentClient.Get(ctx, &commentpb.GetRequest{UUID: r.CommentUUID}); err != nil {
			if err := s.recoverMicroserviceConnection(s.commentClient); err != nil {
				return nil, err
			}
			continue
		}
		break
	}
	if err != nil {
		return nil, fmt.Errorf("comment request error: %w", err)
	}

	uuid, err := s.rateRepository.Store(r)
	if err != nil {
		return &pb.SetResponse{}, err
	}

	return &pb.SetResponse{UUID: uuid}, nil
}

func (s *rateAPIServer) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	if err := user.VerifyToken(ctx, req.GetRate().GetUserUUID(), authPublicKey); err != nil {
		return nil, err
	}

	r := &rate.Rate{
		UUID:        req.GetRate().GetUUID(),
		UserUUID:    req.GetRate().GetUserUUID(),
		CommentUUID: req.GetRate().GetCommentUUID(),
		ProductUUID: req.GetRate().GetProductUUID(),
		Rating:      req.GetRate().GetRating(),
	}
	log.Printf("update %s", r)

	retry := 5
	var err error
	for i := 0; i < retry; i++ {
		if _, err = s.commentClient.Get(ctx, &commentpb.GetRequest{UUID: r.CommentUUID}); err != nil {
			if err := s.recoverMicroserviceConnection(s.commentClient); err != nil {
				return nil, err
			}
			continue
		}
		break
	}
	if err != nil {
		return nil, fmt.Errorf("comment request error: %w", err)
	}

	if err := s.rateRepository.Update(r); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *rateAPIServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	uuid := req.GetUUID()
	log.Printf("delete %s", uuid)

	if err := s.rateRepository.DeleteByUUID(uuid); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *rateAPIServer) recoverMicroserviceConnection(client interface{}) error {
	switch client.(type) {
	case commentpb.CommentAPIClient:
		conn, err := grpc.Dial(s.commentEndpoint, grpc.WithInsecure())
		if err != nil {
			return err
		}
		s.commentClient = commentpb.NewCommentAPIClient(conn)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	redisConfig := rate.RateRepositoryRedisConfig{
		SentinelHost: sentinelHost,
		SentinelPort: sentinelPort,
		Password:     kvsPassword,
		MasterName:   masterName,
	}
	redis := redisConfig.Connect()
	log.Printf("succeed to open kvs")

	commentEndpoint := fmt.Sprintf("%s:%d", commentHost, commentPort)
	commentConn, err := grpc.Dial(commentEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	commentClient := commentpb.NewCommentAPIClient(commentConn)

	rateAPI := &rateAPIServer{
		rateRepository:  redis,
		commentClient:   commentClient,
		commentEndpoint: commentEndpoint,
	}

	s := grpc.NewServer()
	pb.RegisterRateAPIServer(s, rateAPI)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("start rate API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
