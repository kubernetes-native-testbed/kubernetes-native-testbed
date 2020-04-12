package main

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/comment"
	pb "github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/comment/protobuf"
	"github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/user"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var (
	dbUsername string
	dbPassword string
	dbName     string
	dbHost     string
	dbPort     int

	authPublicKey string
)

const (
	defaultBindAddr = ":8080"

	componentName     = "comment"
	defaultDBUsername = componentName
	defaultDBPassword = componentName
	defaultDBName     = componentName
	defaultDBHost     = componentName
	defaultDBPort     = 3306
)

func init() {
	var err error
	if dbUsername = os.Getenv("DB_USERNAME"); dbUsername == "" {
		dbUsername = defaultDBUsername
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
	if authPublicKey = os.Getenv("AUTH_PUBLIC_KEY"); authPublicKey == "" {
		log.Fatal("AUTH_PUBLIC_KEY is required")
	}
}

type commentAPIServer struct {
	commentRepository comment.CommentRepository
}

func (s *commentAPIServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	uuid := req.GetUUID()
	c, err := s.commentRepository.FindByUUID(uuid)
	if err != nil {
		return &pb.GetResponse{}, err
	}
	log.Printf("get %s", c)

	var resp pb.GetResponse
	var cat, uat, dat *timestamp.Timestamp
	if cat, err = ptypes.TimestampProto(c.CreatedAt); err != nil {
		return &pb.GetResponse{}, err
	}
	if uat, err = ptypes.TimestampProto(c.UpdatedAt); err != nil {
		return &pb.GetResponse{}, err
	}
	if c.DeletedAt != nil {
		if dat, err = ptypes.TimestampProto(*c.DeletedAt); err != nil {
			return &pb.GetResponse{}, err
		}
	}

	resp.Comment = &pb.Comment{
		UUID:              c.UUID,
		UserUUID:          c.UserUUID,
		ParentCommentUUID: c.ParentCommentUUID,
		Message:           c.Message,
		CreatedAt:         cat,
		UpdatedAt:         uat,
		DeletedAt:         dat,
	}

	return &resp, nil
}

func (s *commentAPIServer) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	c := &comment.Comment{
		UserUUID:          req.GetComment().GetUserUUID(),
		ParentCommentUUID: req.GetComment().GetParentCommentUUID(),
		Message:           req.GetComment().GetMessage(),
	}
	log.Printf("set %s", c)

	uuid, err := s.commentRepository.Store(c)
	if err != nil {
		return &pb.SetResponse{}, err
	}

	return &pb.SetResponse{UUID: uuid}, nil
}

func (s *commentAPIServer) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	if err := user.VerifyToken(ctx, req.GetComment().GetUserUUID(), authPublicKey); err != nil {
		return nil, err
	}

	c := &comment.Comment{
		UUID:              req.GetComment().GetUUID(),
		UserUUID:          req.GetComment().GetUserUUID(),
		ParentCommentUUID: req.GetComment().GetParentCommentUUID(),
		Message:           req.GetComment().GetMessage(),
	}
	log.Printf("update %s", c)

	if err := s.commentRepository.Update(c); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *commentAPIServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	uuid := req.GetUUID()
	log.Printf("delete %s", uuid)

	if err := s.commentRepository.DeleteByUUID(uuid); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *commentAPIServer) IsExists(ctx context.Context, req *pb.IsExistsRequest) (*pb.IsExistsResponse, error) {
	uuid := req.GetUUID()
	log.Printf("isExists: %s", uuid)
	if _, err := s.commentRepository.FindByUUID(uuid); err != nil {
		return &pb.IsExistsResponse{IsExists: false}, err
	}
	return &pb.IsExistsResponse{IsExists: true}, nil
}

func (s *commentAPIServer) ChildComments(ctx context.Context, req *pb.ChildCommentsRequest) (*pb.ChildCommentsResponse, error) {
	uuid := req.GetParentUUID()
	// bad code
	comments, err := s.commentRepository.List(map[string]interface{}{"parentcommentuuid": uuid})
	if err != nil {
		return &pb.ChildCommentsResponse{}, err
	}
	log.Printf("childComments %s", comments)

	var resp pb.ChildCommentsResponse
	for _, c := range comments {
		var cat, uat, dat *timestamp.Timestamp
		if cat, err = ptypes.TimestampProto(c.CreatedAt); err != nil {
			return &pb.ChildCommentsResponse{}, err
		}
		if uat, err = ptypes.TimestampProto(c.UpdatedAt); err != nil {
			return &pb.ChildCommentsResponse{}, err
		}
		if c.DeletedAt != nil {
			if dat, err = ptypes.TimestampProto(*c.DeletedAt); err != nil {
				return &pb.ChildCommentsResponse{}, err
			}
		}

		pbComment := &pb.Comment{
			UUID:              c.UUID,
			UserUUID:          c.UserUUID,
			ParentCommentUUID: c.ParentCommentUUID,
			Message:           c.Message,
			CreatedAt:         cat,
			UpdatedAt:         uat,
			DeletedAt:         dat,
		}

		resp.ChildComments = append(resp.ChildComments, pbComment)
	}

	return &resp, nil
}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	mongoConfig := comment.CommentRepositoryMongoConfig{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUsername,
		Password: dbPassword,
		DBName:   dbName,
	}
	mongo, err := mongoConfig.Connect()
	if err != nil {
		log.Fatalf("failed to open database: %v (config=%#v)", err, mongoConfig)
	}
	log.Printf("succeed to open database")

	commentAPI := &commentAPIServer{
		commentRepository: mongo,
	}

	s := grpc.NewServer()
	pb.RegisterCommentAPIServer(s, commentAPI)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("start comment API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
