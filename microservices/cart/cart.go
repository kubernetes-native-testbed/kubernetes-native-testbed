package cart

import (
	"encoding/json"

	pb "github.com/__GITHUB_REPO_NAME__/kubernetes-native-testbed/microservices/cart/protobuf"
)

type Cart struct {
	UserUUID     string         `json:"user_uuid"`
	CartProducts map[string]int `json:"cart_products"`
}

func (cart *Cart) String() string {
	b, _ := json.Marshal(cart)
	return string(b)
}

func ConvertToCart(pbCart *pb.Cart) *Cart {
	pbCartProducts := pbCart.GetCartProducts()
	cartProducts := make(map[string]int, len(pbCartProducts))
	for k, v := range pbCartProducts {
		cartProducts[k] = int(v)
	}
	return &Cart{
		UserUUID:     pbCart.GetUserUUID(),
		CartProducts: cartProducts,
	}
}

func ConvertToCartProto(cart *Cart) *pb.Cart {
	pbCartProducts := make(map[string]int32, len(cart.CartProducts))
	for k, v := range cart.CartProducts {
		pbCartProducts[k] = int32(v)
	}
	return &pb.Cart{
		UserUUID:     cart.UserUUID,
		CartProducts: pbCartProducts,
	}
}
