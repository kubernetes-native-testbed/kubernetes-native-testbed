package cart

import (
	"context"
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"

	"github.com/tikv/client-go/config"
	"github.com/tikv/client-go/key"
	"github.com/tikv/client-go/txnkv"
	"github.com/tikv/client-go/txnkv/kv"
)

type CartRepository interface {
	FindByUUID(string) (*Cart, bool, error)
	Store(*Cart) (string, error)
	Update(*Cart) error
	DeleteByUUID(string) error
}

type cartRepositoryTiKV struct {
	client *txnkv.Client
	ctx    context.Context
}

func (cr *cartRepositoryTiKV) FindByUUID(uuid string) (*Cart, bool, error) {
	tx, err := cr.client.Begin(cr.ctx)
	if err != nil {
		return nil, false, err
	}

	v, err := tx.Get(cr.ctx, key.Key(uuid))
	if err != nil {
		if kv.IsErrNotFound(err) {
			return nil, true, nil
		} else {
			return nil, false, err
		}
	}

	hashes := strings.Split(string(v), ",")
	productMap := make(map[string]int, len(hashes))
	for _, hash := range hashes {
		v, err := tx.Get(cr.ctx, key.Key(hash))
		if err != nil {
			if kv.IsErrNotFound(err) {
				continue
			} else {
				return nil, false, err
			}
		}
		cartProductsInfo := strings.Split(string(v), ":")
		if len(cartProductsInfo) != 2 {
			return nil, false, fmt.Errorf("invalid cart products info format: %s", string(v))
		}
		productUUID := cartProductsInfo[0]
		count, err := strconv.Atoi(cartProductsInfo[1])
		if err != nil {
			return nil, false, err
		}
		productMap[productUUID] = count
	}

	if err := tx.Commit(cr.ctx); err != nil {
		return nil, false, err
	}

	return &Cart{UserUUID: uuid, CartProducts: productMap}, false, nil
}

func (cr *cartRepositoryTiKV) Store(cart *Cart) (string, error) {
	tx, err := cr.client.Begin(cr.ctx)
	if err != nil {
		return "", err
	}

	hashes := make([]string, 0, len(cart.CartProducts))
	for productUUID, count := range cart.CartProducts {
		hash := fmt.Sprintf("%X", md5.Sum([]byte(cart.UserUUID+productUUID)))
		hashes = append(hashes, hash)

		v := fmt.Sprintf("%s:%d", productUUID, count)
		if err := tx.Set(key.Key(hash), []byte(v)); err != nil {
			return "", err
		}
	}

	if err := tx.Set(key.Key(cart.UserUUID), []byte(strings.Join(hashes, ","))); err != nil {
		return "", err
	}

	if err := tx.Commit(cr.ctx); err != nil {
		return "", err
	}

	return cart.UserUUID, nil
}

func (cr *cartRepositoryTiKV) Update(cart *Cart) error {
	if _, err := cr.Store(cart); err != nil {
		return err
	}
	return nil
}

func (cr *cartRepositoryTiKV) DeleteByUUID(uuid string) error {
	tx, err := cr.client.Begin(cr.ctx)
	if err != nil {
		return err
	}

	if err := tx.Delete(key.Key(uuid)); err != nil {
		return err
	}

	if err := tx.Commit(cr.ctx); err != nil {
		return err
	}

	return nil
}

type CartRepositoryTiKVConfig struct {
	Ctx       context.Context
	PdAddress string
	PdPort    int
}

func (c *CartRepositoryTiKVConfig) Connect() (CartRepository, func() error, error) {
	client, err := txnkv.NewClient(c.Ctx, []string{fmt.Sprintf("%s:%d", c.PdAddress, c.PdPort)}, config.Default())
	if err != nil {
		return nil, nil, err
	}

	return &cartRepositoryTiKV{client: client, ctx: c.Ctx}, client.Close, nil

}
