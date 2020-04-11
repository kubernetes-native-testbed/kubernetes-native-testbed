package main

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	pb "github.com/__GITHUB_REPO_NAME__/kubernetes-native-testbed/microservices/deliveryStatus/protobuf"
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

	orderQueueUsername string
	orderQueuePassword string
	orderQueueHost     string
	orderQueuePort     int
	orderQueueSubject  string
)

const (
	defaultBindAddr = ":8080"

	componentName     = "deliveryStatus"
	defaultDBUsername = "cassandra"
	defaultDBPassword = "cassandra"
	defaultDBName     = "deliverystatus"
	defaultDBHost     = "delivery-status-db.delivery-status.svc.cluster.local"
	defaultDBPort     = 4000

	defaultOrderQueueUsername = componentName
	defaultOrderQueuePassword = componentName
	defaultOrderQueueHost     = "delivery-status-queue.delivery-status.svc.cluster.local"
	defaultOrderQueuePort     = 4222
	defaultOrderQueueSubject  = "order"
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
	if orderQueueUsername = os.Getenv("ORDER_QUEUE_USERNAME"); orderQueueUsername == "" {
		orderQueueUsername = defaultOrderQueueUsername
	}
	if orderQueuePassword = os.Getenv("ORDER_QUEUE_PASSWORD"); orderQueuePassword == "" {
		orderQueuePassword = defaultOrderQueuePassword
	}
	if orderQueueHost = os.Getenv("ORDER_QUEUE_HOST"); orderQueueHost == "" {
		orderQueueHost = defaultOrderQueueHost
	}
	if orderQueuePort, err = strconv.Atoi(os.Getenv("ORDER_QUEUE_PORT")); err != nil {
		orderQueuePort = defaultOrderQueuePort
		log.Printf("orderQueuePort parse error: %v", err)
	}
	if orderQueueSubject = os.Getenv("ORDER_QUEUE_SUBJECT"); orderQueueSubject == "" {
		orderQueueSubject = defaultOrderQueueSubject
	}
}

type deliveryStatusAPIServer struct {
	deliveryStatusRepository deliveryStatusRepository
	orderQueue               orderQueue
}

func (s *deliveryStatusAPIServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	uuid := req.GetOrderUUID()
	ds, err := s.deliveryStatusRepository.findByUUID(uuid)
	if err != nil {
		return &pb.GetResponse{}, err
	}
	log.Printf("get %s", ds)

	var resp pb.GetResponse
	var cat, uat, dat *timestamp.Timestamp
	if cat, err = ptypes.TimestampProto(ds.CreatedAt); err != nil {
		return &pb.GetResponse{}, err
	}
	if uat, err = ptypes.TimestampProto(ds.UpdatedAt); err != nil {
		return &pb.GetResponse{}, err
	}
	if ds.DeletedAt != nil {
		if dat, err = ptypes.TimestampProto(*ds.DeletedAt); err != nil {
			return &pb.GetResponse{}, err
		}
	}

	resp.DeliveryStatus = &pb.DeliveryStatus{
		OrderUUID:     ds.OrderUUID,
		UserUUID:      ds.UserUUID,
		Status:        int32(ds.Status),
		InquiryNumber: ds.InquiryNumber,
		CreatedAt:     cat,
		UpdatedAt:     uat,
		DeletedAt:     dat,
	}

	return &resp, nil
}

func (s *deliveryStatusAPIServer) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	ds := &DeliveryStatus{
		OrderUUID:     req.GetDeliveryStatus().GetOrderUUID(),
		UserUUID:      req.GetDeliveryStatus().GetUserUUID(),
		Status:        Status(req.GetDeliveryStatus().GetStatus()),
		InquiryNumber: req.GetDeliveryStatus().GetInquiryNumber(),
	}
	log.Printf("set %s", ds)

	uuid, err := s.deliveryStatusRepository.store(ds)
	if err != nil {
		return &pb.SetResponse{}, err
	}

	return &pb.SetResponse{OrderUUID: uuid}, nil
}

func (s *deliveryStatusAPIServer) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	ds := &DeliveryStatus{
		OrderUUID:     req.GetDeliveryStatus().GetOrderUUID(),
		UserUUID:      req.GetDeliveryStatus().GetUserUUID(),
		Status:        Status(req.GetDeliveryStatus().GetStatus()),
		InquiryNumber: req.GetDeliveryStatus().GetInquiryNumber(),
	}
	log.Printf("update %s", ds)

	if err := s.deliveryStatusRepository.update(ds); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *deliveryStatusAPIServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	uuid := req.GetOrderUUID()
	log.Printf("delete {\"uuid\":\"%s\"}", uuid)

	if err := s.deliveryStatusRepository.deleteByUUID(uuid); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *deliveryStatusAPIServer) subscribeOrderQueue() (func() error, error) {
	orderCh, err := s.orderQueue.subscribe()
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			msg := strings.Split(<-orderCh, ":")
			if len(msg) != 2 {
				continue
			}
			orderUUID := msg[0]
			userUUID := msg[1]

			ds := &DeliveryStatus{
				OrderUUID:     orderUUID,
				UserUUID:      userUUID,
				Status:        Waiting,
				InquiryNumber: generateInquiryNumber(),
			}
			log.Printf("[from subscribe] set %s", ds)

			_, err := s.deliveryStatusRepository.store(ds)
			if err != nil {
				log.Printf("[from subscribe] failed to store delivery status (%#v): %v", ds, err)
			}
		}
	}()

	return s.orderQueue.unsubscribe, nil

}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	drConfig := deliveryStatusRepositoryCassandraConfig{
		host:     dbHost,
		port:     dbPort,
		username: dbUsername,
		password: dbPassword,
		keyspace: dbName,
	}
	dr, closeDr, err := drConfig.connect()
	if err != nil {
		log.Fatalf("failed to connect delivery status repository: %v (config=%#v)", err, drConfig)
	}
	defer closeDr()
	log.Printf("successed to connect to delivery status repository")

	oqConfig := orderQueueNATSConfig{
		host:     orderQueueHost,
		port:     orderQueuePort,
		username: orderQueueUsername,
		password: orderQueuePassword,
		subject:  orderQueueSubject,
		retry:    10,
	}
	oq, closeOq, err := oqConfig.connect()
	if err != nil {
		log.Fatalf("failed to connect to order queue: %v (config=%#v)", err, oqConfig)
	}
	defer closeOq()
	log.Printf("successed to connect to order queue")

	s := grpc.NewServer()
	api := &deliveryStatusAPIServer{
		deliveryStatusRepository: dr,
		orderQueue:               oq,
	}
	pb.RegisterDeliveryStatusAPIServer(s, api)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("setup database")
	if err := api.deliveryStatusRepository.initDB(); err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	unsubscribe, err := api.subscribeOrderQueue()
	if err != nil {
		log.Fatalf("failed to subscribe order queue: %v", err)
	}
	defer unsubscribe()

	log.Printf("start deliveryStatus API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
