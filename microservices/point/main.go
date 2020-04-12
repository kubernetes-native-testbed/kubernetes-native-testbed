package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/order"
	pb "github.com/kubernetes-native-testbed-user/kubernetes-native-testbed/microservices/point/protobuf"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var (
	dbUser             string
	dbPassword         string
	dbName             string
	dbHost             string
	dbPort             string
	dbSSL              string
	kvsHost            string
	kvsPort            string
	orderQueueUsername string
	orderQueuePassword string
	orderQueueHost     string
	orderQueuePort     int
	orderQueueTopic    string
)

const (
	defaultBindAddr = ":8080"
	pointRatio      = 0.01

	componentName     = "point"
	defaultDBUser     = componentName
	defaultDBPassword = componentName
	defaultDBName     = componentName
	defaultDBHost     = componentName
	defaultDBHPort    = componentName
	defaultDBSSL      = "require"

	// defaultKVSUser = componentName
	// defaultKVSPass = componentName
	defaultKVSHost = componentName
	defaultKVSPort = componentName

	defaultQueueHost     = componentName
	defaultQueuePort     = 0
	defaultQueueTopic    = componentName
	defaultQueueUser     = componentName
	defaultQueuePassword = componentName
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
	if dbPort = os.Getenv("DB_PORT"); dbPort == "" {
		dbPort = defaultDBHPort
	}
	if dbSSL = os.Getenv("DB_SSL"); dbSSL == "" {
		dbSSL = defaultDBSSL
	}

	if kvsHost = os.Getenv("KVS_HOST"); kvsHost == "" {
		kvsHost = defaultKVSHost
	}
	if kvsPort = os.Getenv("KVS_PORT"); kvsPort == "" {
		kvsPort = defaultKVSPort
	}

	if orderQueueUsername = os.Getenv("QUEUE_USER"); orderQueueUsername == "" {
		orderQueueUsername = defaultQueueUser
	}
	if orderQueuePassword = os.Getenv("QUEUE_PASSWORD"); orderQueuePassword == "" {
		orderQueuePassword = defaultQueuePassword
	}
	if orderQueueHost = os.Getenv("QUEUE_HOST"); orderQueueHost == "" {
		orderQueueHost = defaultQueueHost
	}
	if orderQueuePort, err = strconv.Atoi(os.Getenv("QUEUE_PORT")); err != nil {
		orderQueuePort = defaultQueuePort
		log.Printf("orderQueuePort parse error: %v", err)
	}
	if orderQueueTopic = os.Getenv("QUEUE_TOPIC"); orderQueueTopic == "" {
		orderQueueTopic = defaultQueueTopic
	}
}

type pointAPIServer struct {
	pointRepository      pointRepository
	pointCacheRepository pointCacheRepository
	orderQueue           orderQueue
}

func (s *pointAPIServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	uuid := req.GetUUID()
	log.Printf("{\"operation\":\"get\", \"uuid\":\"%s\"}", uuid)
	p, err := s.pointRepository.findByUUID(uuid)
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

	resp.Point = &pb.Point{
		UUID:        p.UUID,
		UserUUID:    p.UserUUID,
		Balance:     p.Balance,
		Description: p.Description,
		CreatedAt:   cat,
		UpdatedAt:   uat,
		DeletedAt:   dat,
	}

	return &resp, nil
}

func (s *pointAPIServer) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	p := &Point{
		UserUUID:    req.GetPoint().GetUserUUID(),
		Balance:     req.GetPoint().GetBalance(),
		Description: req.GetPoint().GetDescription(),
	}
	log.Printf("{\"operation\":\"set\", \"uuid\":\"%s\", \"user_uuid\":\"%s\",\"balance\":\"%d\", \"description\":\"%s\"}", p.UUID, p.UserUUID, p.Balance, p.Description)

	uuid, err := s.pointRepository.store(p)
	if err != nil {
		return &pb.SetResponse{}, err
	}

	if err := s.updateTotalAmountCache(p.UserUUID); err != nil {
		log.Printf("failed to update amout cache for %s: %w", p.UserUUID, err)
	}

	return &pb.SetResponse{UUID: uuid}, nil
}

func (s *pointAPIServer) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	p := &Point{
		UUID:        req.GetPoint().GetUUID(),
		UserUUID:    req.GetPoint().GetUserUUID(),
		Balance:     req.GetPoint().GetBalance(),
		Description: req.GetPoint().GetDescription(),
	}
	log.Printf("{\"operation\":\"update\", \"uuid\":\"%s\", \"user_uuid\":\"%s\",\"balance\":\"%d\", \"description\":\"%s\"}", p.UUID, p.UserUUID, p.Balance, p.Description)

	if err := s.pointRepository.update(p); err != nil {
		return &empty.Empty{}, err
	}

	if err := s.updateTotalAmountCache(p.UserUUID); err != nil {
		log.Printf("failed to update amout cache for %s: %w", p.UserUUID, err)
	}

	return &empty.Empty{}, nil
}

func (s *pointAPIServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	uuid := req.GetUUID()
	log.Printf("{\"operation\":\"delete\", \"uuid\":\"%s\"}", uuid)

	if err := s.pointRepository.deleteByUUID(uuid); err != nil {
		return &empty.Empty{}, err
	}

	p, err := s.pointRepository.findByUUID(uuid)
	if err != nil {
		return &empty.Empty{}, err
	}

	if err := s.updateTotalAmountCache(p.UserUUID); err != nil {
		log.Printf("failed to update amout cache for %s: %w", p.UserUUID, err)
	}

	return &empty.Empty{}, nil
}

func (s *pointAPIServer) GetTotalAmount(ctx context.Context, req *pb.GetTotalAmountRequest) (*pb.GetTotalAmountResponse, error) {
	user_uuid := req.GetUserUUID()
	pc, err := s.pointCacheRepository.findByUUID(user_uuid)
	if err != nil {
		log.Printf("cannot find cache for %s", user_uuid)
		if err := s.updateTotalAmountCache(user_uuid); err != nil {
			log.Printf("failed to update amout cache for %s: %w", user_uuid, err)
			return nil, err
		} else {
			pc, err = s.pointCacheRepository.findByUUID(user_uuid)
			if err != nil {
				log.Printf("cannot store cache for %s", user_uuid)
				return nil, err
			}
		}
	}

	return &pb.GetTotalAmountResponse{UserUUID: pc.UserUUID, TotalAmount: pc.TotalAmount}, nil
}

func (s *pointAPIServer) updateTotalAmountCache(useruuid string) error {
	amount, err := s.pointRepository.getTotalAmount(useruuid)
	if err != nil {
		return err
	}

	pc := &PointCache{
		UserUUID:    useruuid,
		TotalAmount: amount,
	}
	if err := s.pointCacheRepository.store(pc); err != nil {
		return err
	}
	log.Printf("Total amount is updated. UserUUID=%s, TotalAmount=%d", useruuid, amount)
	return nil
}

func (s *pointAPIServer) subscribeOrderQueue() (func() error, error) {
	orderCh, err := s.orderQueue.subscribe()
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			msg := strings.Split(<-orderCh, "::")
			if len(msg) != 2 {
				log.Printf("received data is invalid: %s (size: %d)\n", msg, len(msg))
				continue
			}
			operation := msg[0]
			data := msg[1]

			o := &order.Order{}
			err = json.Unmarshal([]byte(data), o)
			if err != nil {
				log.Printf("received body data is invalid: %s\n", data)
				continue
			}

			switch order.Operation(operation) {
			case order.CreateOperation:
				cost := float64(0)
				for _, op := range o.OrderedProducts {
					cost = cost + float64(op.Price*op.Count)*pointRatio
				}

				p := &Point{
					UserUUID:    o.UserUUID,
					Balance:     int32(cost),
					Description: fmt.Sprintf("Order point [order id = %s, order date = %s]", o.UUID, o.CreatedAt),
				}
				_, err := s.pointRepository.store(p)
				if err != nil {
					log.Printf("[from subscribe] failed to store point (%#v): %v", p, err)
				}

			case order.DeleteOperation:
				cost := 0
				for _, op := range o.OrderedProducts {
					cost = cost + op.Price*op.Count
				}

				p := &Point{
					UserUUID:    o.UserUUID,
					Balance:     -int32(cost),
					Description: fmt.Sprintf("[Cancel] Order point [order id = %s, order date = %s]", o.UUID, o.CreatedAt),
				}
				_, err := s.pointRepository.store(p)
				if err != nil {
					log.Printf("[from subscribe] failed to store point (%#v): %v", p, err)
				}
			case order.UpdateOperation:
				log.Printf("order update is not permitted")
			default:
				log.Printf("Unknown operation [%s]: %s", operation, msg)
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

	cache := memcache.New(
		fmt.Sprintf("%s:%s",
			kvsHost,
			kvsPort,
		),
	)

	oqConfig := orderQueueKafkaConfig{
		host:     orderQueueHost,
		port:     orderQueuePort,
		username: orderQueueUsername,
		password: orderQueuePassword,
		topic:    orderQueueTopic,
		group:    "point-consumer",
		retry:    10,
	}
	oq, closeOq, err := oqConfig.connect()
	if err != nil {
		log.Fatalf("failed to connect to order queue: %v (config=%#v)", err, oqConfig)
	}
	defer closeOq()
	log.Printf("successed to connect to order queue")

	s := grpc.NewServer()
	api := &pointAPIServer{
		pointRepository: &pointRepositoryImpl{
			db: db,
		},
		pointCacheRepository: &pointRepositoryMemcache{
			cache: cache,
		},
		orderQueue: oq,
	}
	pb.RegisterPointAPIServer(s, api)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("setup database")
	if err := api.pointRepository.initDB(); err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	unsubscribe, err := api.subscribeOrderQueue()
	if err != nil {
		log.Fatalf("failed to subscribe order queue: %v", err)
	}
	defer unsubscribe()

	log.Printf("start point API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
