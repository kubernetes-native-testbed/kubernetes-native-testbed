package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/kubernetes-native-testbed/kubernetes-native-testbed/microservices/cart/protobuf"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var (
	kvsHost string
	kvsPort int
)

const (
	defaultBindAddr = ":8080"

	componentName  = "cart"
	defaultKVSHost = "cart-db-pd.cart.svc.cluster.local"
	defaultKVSPort = 2379
)

func init() {
	var err error
	if kvsHost = os.Getenv("KVS_HOST"); kvsHost == "" {
		kvsHost = defaultKVSHost
	}
	if kvsPort, err = strconv.Atoi(os.Getenv("KVS_PORT")); err != nil {
		kvsPort = defaultKVSPort
		log.Printf("kvsPort parse error: %v", err)
	}
}

type cartAPIServer struct {
	cartRepository cartRepository
}

func (s *cartAPIServer) Show(ctx context.Context, req *pb.ShowRequest) (*pb.ShowResponse, error) {
	userUUID := req.GetUserUUID()
	cart, ok, err := s.cartRepository.findByUUID(userUUID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("cart is not found for %s", userUUID)
	}
	return &pb.ShowResponse{Cart: convertToCartProto(cart)}, nil
}

func (s *cartAPIServer) Add(ctx context.Context, req *pb.AddRequest) (*empty.Empty, error) {
	additionalCart := convertToCart(req.GetCart())
	cart, ok, err := s.cartRepository.findByUUID(additionalCart.UserUUID)
	if err != nil {
		return nil, err
	}
	if !ok {
		if _, err := s.cartRepository.store(additionalCart); err != nil {
			return nil, err
		}
		return &empty.Empty{}, nil
	}

	for productUUID, increaseCount := range additionalCart.CartProducts {
		if _, ok := cart.CartProducts[productUUID]; ok {
			// increase
			cart.CartProducts[productUUID] += increaseCount
		} else {
			// add
			cart.CartProducts[productUUID] = increaseCount
		}
	}

	if err := s.cartRepository.update(cart); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *cartAPIServer) Remove(ctx context.Context, req *pb.RemoveRequest) (*empty.Empty, error) {
	additionalCart := convertToCart(req.GetCart())
	cart, ok, err := s.cartRepository.findByUUID(additionalCart.UserUUID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("cart is not found for %s", additionalCart.UserUUID)
	}

	for productUUID, decreaseCount := range additionalCart.CartProducts {
		if count, ok := cart.CartProducts[productUUID]; ok {
			// decrease and remove
			count -= decreaseCount
			if count <= 0 {
				delete(cart.CartProducts, productUUID)
			} else {
				cart.CartProducts[productUUID] = count
			}
		}
	}

	if err := s.cartRepository.update(cart); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", defaultBindAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on %s", defaultBindAddr)

	crConfig := cartRepositoryTiKVConfig{
		pdAddress: kvsHost,
		pdPort:    kvsPort,
	}
	cr, closeCr, err := crConfig.connect()
	if err != nil {
		log.Fatalf("failed to connect to kvs: %v (config=%#v)", err, crConfig)
	}
	defer closeCr()
	log.Printf("successed to connect to kvs")

	s := grpc.NewServer()
	api := &cartAPIServer{
		cartRepository: cr,
	}
	pb.RegisterCartAPIServer(s, api)

	healthpb.RegisterHealthServer(s, health.NewServer())

	log.Printf("start cart API server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
