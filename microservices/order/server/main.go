package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/__GITHUB_REPO_NAME__/kubernetes-native-testbed/microservices/order"
	pb "github.com/__GITHUB_REPO_NAME__/kubernetes-native-testbed/microservices/order/protobuf"
	paymentinfopb "github.com/__GITHUB_REPO_NAME__/kubernetes-native-testbed/microservices/payment-info/protobuf"
	userpb "github.com/__GITHUB_REPO_NAME__/kubernetes-native-testbed/microservices/user/protobuf"
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

	deliveryStatusUsername string
	deliveryStatusPassword string
	deliveryStatusHost     string
	deliveryStatusPort     int
	deliveryStatusSubject  string

	pointQueueUsername string
	pointQueuePassword string
	pointQueueHost     string
	pointQueuePort     int
	pointQueueTopic    string

	userHost        string
	userPort        int
	paymentInfoHost string
	paymentInfoPort int
)

const (
	defaultBindAddr = ":8080"

	componentName     = "order"
	defaultDBUsername = componentName + "user"
	defaultDBPassword = componentName
	defaultDBName     = componentName + "db"
	defaultDBHost     = componentName
	defaultDBPort     = 4000

	defaultDeliveryStatusUsername = componentName
	defaultDeliveryStatusPassword = componentName
	defaultDeliveryStatusHost     = "delivery-status-queue.delivery-status.svc.cluster.local"
	defaultDeliveryStatusPort     = 4222
	defaultDeliveryStatusSubject  = componentName

	defaultPointQueueUsername = componentName
	defaultPointQueuePassword = componentName
	defaultPointQueueHost     = componentName
	defaultPointQueuePort     = 4222
	defaultPointQueueTopic    = componentName

	defaultUserHost        = "user.user.svc.cluster.local"
	defaultUserPort        = 8080
	defaultPaymentInfoHost = "payment-info.payment-info.svc.cluster.local"
	defaultPaymentInfoPort = 8080
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

	if deliveryStatusUsername = os.Getenv("DELIVERY_STATUS_USERNAME"); deliveryStatusUsername == "" {
		deliveryStatusUsername = defaultDeliveryStatusUsername
	}
	if deliveryStatusPassword = os.Getenv("DELIVERY_STATUS_PASSWORD"); deliveryStatusPassword == "" {
		deliveryStatusPassword = defaultDeliveryStatusPassword
	}
	if deliveryStatusHost = os.Getenv("DELIVERY_STATUS_HOST"); deliveryStatusHost == "" {
		deliveryStatusHost = defaultDeliveryStatusHost
	}
	if deliveryStatusPort, err = strconv.Atoi(os.Getenv("DELIVERY_STATUS_PORT")); err != nil {
		deliveryStatusPort = defaultDeliveryStatusPort
		log.Printf("deliveryStatusPort parse error: %v", err)
	}
	if deliveryStatusSubject = os.Getenv("DELIVERY_STATUS_SUBJECT"); deliveryStatusSubject == "" {
		deliveryStatusSubject = defaultDeliveryStatusSubject
	}

	if pointQueueUsername = os.Getenv("POINT_QUEUE_USER"); pointQueueUsername == "" {
		pointQueueUsername = defaultPointQueueUsername
	}
	if pointQueuePassword = os.Getenv("POINT_QUEUE_PASSWORD"); pointQueuePassword == "" {
		pointQueuePassword = defaultPointQueuePassword
	}
	if pointQueueHost = os.Getenv("POINT_QUEUE_HOST"); pointQueueHost == "" {
		pointQueueHost = defaultPointQueueHost
	}
	if pointQueuePort, err = strconv.Atoi(os.Getenv("POINT_QUEUE_PORT")); err != nil {
		pointQueuePort = defaultPointQueuePort
		log.Printf("pointQueuePort parse error: %v", err)
	}
	if pointQueueTopic = os.Getenv("POINT_QUEUE_TOPIC"); pointQueueTopic == "" {
		pointQueueTopic = defaultPointQueueTopic
	}

	if userHost = os.Getenv("USER_HOST"); userHost == "" {
		userHost = defaultUserHost
	}
	if userPort, err = strconv.Atoi(os.Getenv("USER_PORT")); err != nil {
		userPort = defaultUserPort
		log.Printf("userPort parse error: %v", err)
	}
	if paymentInfoHost = os.Getenv("PAYMENT_INFO_HOST"); paymentInfoHost == "" {
		paymentInfoHost = defaultPaymentInfoHost
	}
	if paymentInfoPort, err = strconv.Atoi(os.Getenv("PAYMENT_INFO_PORT")); err != nil {
		paymentInfoPort = defaultPaymentInfoPort
		log.Printf("paymentInfoPort parse error: %v", err)
	}
}

type orderAPIServer struct {
	orderRepository              order.OrderRepository
	orderSenderForDeliveryStatus order.OrderSender
	orderSenderForPoint          order.OrderSender

	userClient          userpb.UserAPIClient
	userEndpoint        string
	paymentInfoClient   paymentinfopb.PaymentInfoAPIClient
	paymentInfoEndpoint string
}

func (s *orderAPIServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	uuid := req.GetUUID()
	o, err := s.orderRepository.FindByUUID(uuid)
	if err != nil {
		return &pb.GetResponse{}, err
	}
	log.Printf("get %s", o)

	var resp pb.GetResponse
	var cat, uat, dat *timestamp.Timestamp
	if cat, err = ptypes.TimestampProto(o.CreatedAt); err != nil {
		return &pb.GetResponse{}, err
	}
	if uat, err = ptypes.TimestampProto(o.UpdatedAt); err != nil {
		return &pb.GetResponse{}, err
	}
	if o.DeletedAt != nil {
		if dat, err = ptypes.TimestampProto(*o.DeletedAt); err != nil {
			return &pb.GetResponse{}, err
		}
	}

	orderedProducts := make([]*pb.OrderedProduct, len(o.OrderedProducts))
	for i := range o.OrderedProducts {
		op := o.OrderedProducts[i]

		var opcat, opuat, opdat *timestamp.Timestamp
		if opcat, err = ptypes.TimestampProto(op.CreatedAt); err != nil {
			return &pb.GetResponse{}, err
		}
		if opuat, err = ptypes.TimestampProto(op.UpdatedAt); err != nil {
			return &pb.GetResponse{}, err
		}
		if op.DeletedAt != nil {
			if opdat, err = ptypes.TimestampProto(*op.DeletedAt); err != nil {
				return &pb.GetResponse{}, err
			}
		}

		orderedProducts[i] = &pb.OrderedProduct{
			OrderUUID:   op.OrderUUID,
			ProductUUID: op.ProductUUID,
			Count:       int32(op.Count),
			Price:       int32(op.Price),
			CreatedAt:   opcat,
			UpdatedAt:   opuat,
			DeletedAt:   opdat,
		}
	}

	resp.Order = &pb.Order{
		UUID:            o.UUID,
		OrderedProducts: orderedProducts,
		UserUUID:        o.UserUUID,
		PaymentInfoUUID: o.PaymentInfoUUID,
		AddressUUID:     o.AddressUUID,
		CreatedAt:       cat,
		UpdatedAt:       uat,
		DeletedAt:       dat,
	}

	return &resp, nil
}

func (s *orderAPIServer) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	orderedProducts := make([]order.OrderedProduct, len(req.GetOrder().GetOrderedProducts()))
	for i, op := range req.GetOrder().GetOrderedProducts() {
		orderedProducts[i] = order.OrderedProduct{
			OrderUUID:   op.GetOrderUUID(),
			ProductUUID: op.GetProductUUID(),
			Count:       int(op.GetCount()),
			Price:       int(op.GetPrice()),
		}
	}

	if err := s.orderValidation(ctx, req.GetOrder()); err != nil {
		return nil, err
	}

	o := &order.Order{
		UUID:            req.GetOrder().GetUUID(),
		OrderedProducts: orderedProducts,
		UserUUID:        req.GetOrder().GetUserUUID(),
		PaymentInfoUUID: req.GetOrder().GetPaymentInfoUUID(),
		AddressUUID:     req.GetOrder().GetAddressUUID(),
	}
	log.Printf("set %s", o)

	uuid, err := s.orderRepository.Store(o)
	if err != nil {
		return &pb.SetResponse{}, err
	}

	go func() {
		if err := s.orderSenderForDeliveryStatus.Send(o, order.CreateOperation); err != nil {
			// TODO: save fail information to order table for avoding lost order
			log.Printf("send error: %v", err)
		}
	}()

	go func() {
		if err := s.orderSenderForPoint.Send(o, order.CreateOperation); err != nil {
			// TODO: save fail information to order table for avoding lost order
			log.Printf("send error: %v", err)
		}
	}()

	return &pb.SetResponse{UUID: uuid}, nil
}

func (s *orderAPIServer) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	orderedProducts := make([]order.OrderedProduct, len(req.GetOrder().GetOrderedProducts()))
	for i, op := range req.GetOrder().GetOrderedProducts() {
		orderedProducts[i] = order.OrderedProduct{
			OrderUUID:   op.GetOrderUUID(),
			ProductUUID: op.GetProductUUID(),
			Count:       int(op.GetCount()),
			Price:       int(op.GetPrice()),
		}
	}

	if err := s.orderValidation(ctx, req.GetOrder()); err != nil {
		return nil, err
	}

	o := &order.Order{
		UUID:            req.GetOrder().GetUUID(),
		OrderedProducts: orderedProducts,
		UserUUID:        req.GetOrder().GetUserUUID(),
		PaymentInfoUUID: req.GetOrder().GetPaymentInfoUUID(),
		AddressUUID:     req.GetOrder().GetAddressUUID(),
	}
	log.Printf("update %s", o)

	if err := s.orderRepository.Update(o); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *orderAPIServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	uuid := req.GetUUID()
	log.Printf("delete {\"uuid\":\"%s\"}", uuid)

	if err := s.orderRepository.DeleteByUUID(uuid); err != nil {
		return &empty.Empty{}, err
	}

	go func() {
		if err := s.orderSenderForPoint.Send(&order.Order{UUID: uuid}, order.DeleteOperation); err != nil {
			// TODO: save fail information to order table for avoding lost order
			log.Printf("send error: %v", err)
		}
	}()

	return &empty.Empty{}, nil
}

func (s *orderAPIServer) orderValidation(ctx context.Context, order *pb.Order) error {
	userUUID := order.GetUserUUID()
	targetAddressUUID := order.GetAddressUUID()
	targetPaymentInfoUUID := order.GetPaymentInfoUUID()

	var err error

	retry := 1
	var userResp *userpb.GetResponse
	for i := 0; i < retry; i++ {
		userResp, err = s.userClient.Get(ctx, &userpb.GetRequest{UUID: userUUID})
		if err != nil {
			if err := s.recoverMicroserviceConnection(s.userClient); err != nil {
				return err
			}
			continue
		}
		break
	}
	if err != nil {
		return err
	}
	var addressFound bool
	for _, address := range userResp.GetUser().GetAddresses() {
		if address.GetUUID() == targetAddressUUID {
			addressFound = true
			break
		}
	}
	if !addressFound {
		return fmt.Errorf("valid address uuid is not found: %s (userUUID=%s)", targetAddressUUID, userUUID)
	}

	var paymentInfoResp *paymentinfopb.GetResponse
	for i := 0; i < retry; i++ {
		paymentInfoResp, err = s.paymentInfoClient.Get(ctx, &paymentinfopb.GetRequest{UUID: targetPaymentInfoUUID})
		if err != nil {
			if err := s.recoverMicroserviceConnection(s.paymentInfoClient); err != nil {
				return err
			}
			continue
		}
		break
	}
	if err != nil {
		return err
	}
	if paymentInfoResp.GetPaymentInfo().GetUserUUID() != userUUID {
		return fmt.Errorf("paymentInfo's userUUID is not match: %s (userUUID=%s)", targetPaymentInfoUUID, userUUID)
	}

	return nil
}

func (s *orderAPIServer) recoverMicroserviceConnection(client interface{}) error {
	switch client.(type) {
	case paymentinfopb.PaymentInfoAPIClient:
		conn, err := grpc.Dial(s.paymentInfoEndpoint, grpc.WithInsecure())
		if err != nil {
			return err
		}
		s.paymentInfoClient = paymentinfopb.NewPaymentInfoAPIClient(conn)
	case userpb.UserAPIClient:
		conn, err := grpc.Dial(s.userEndpoint, grpc.WithInsecure())
		if err != nil {
			return err
		}
		s.userClient = userpb.NewUserAPIClient(conn)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	tidbConfig := order.OrderRepositoryTiDBConfig{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUsername,
		Password: dbPassword,
		DBName:   dbName,
	}
	tidb, closeTiDB, err := tidbConfig.Connect()
	if err != nil {
		log.Fatalf("failed to open database: %v (config=%#v)", err, tidbConfig)
	}
	defer closeTiDB()
	log.Printf("succeed to open database")

	natsConfig := order.OrderSenderNATSConfig{
		Host:     deliveryStatusHost,
		Port:     deliveryStatusPort,
		Username: deliveryStatusUsername,
		Password: deliveryStatusPassword,
		Subject:  deliveryStatusSubject,
		Retry:    5,
	}
	nats, closeNats, err := natsConfig.Connect()
	if err != nil {
		log.Fatalf("failed to create connection to delivery status queue: %v (config=%#v)", err, natsConfig)
	}
	defer closeNats()
	log.Printf("succeed to connect to delivery status queue")

	kafkaConfig := order.OrderSenderKafkaConfig{
		Host:     pointQueueHost,
		Port:     pointQueuePort,
		Username: pointQueueUsername,
		Password: pointQueuePassword,
		Topic:    pointQueueTopic,
		Retry:    5,
	}
	kafka, closeKafka, err := kafkaConfig.Connect()
	if err != nil {
		log.Fatalf("failed to create connection to point queue: %v (config=%#v)", err, kafkaConfig)
	}
	defer closeKafka()
	log.Printf("succeed to connect to point queue")

	userEndpoint := fmt.Sprintf("%s:%d", userHost, userPort)
	userConn, err := grpc.Dial(userEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	userClient := userpb.NewUserAPIClient(userConn)
	paymentInfoEndpoint := fmt.Sprintf("%s:%d", paymentInfoHost, paymentInfoPort)
	paymentInfoConn, err := grpc.Dial(paymentInfoEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	paymentInfoClient := paymentinfopb.NewPaymentInfoAPIClient(paymentInfoConn)

	s := grpc.NewServer()
	api := &orderAPIServer{
		orderRepository:              tidb,
		orderSenderForDeliveryStatus: nats,
		orderSenderForPoint:          kafka,
		userClient:                   userClient,
		userEndpoint:                 userEndpoint,
		paymentInfoClient:            paymentInfoClient,
		paymentInfoEndpoint:          paymentInfoEndpoint,
	}
	pb.RegisterOrderAPIServer(s, api)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("setup database")
	if err := api.orderRepository.InitDB(); err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	log.Printf("start order API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
