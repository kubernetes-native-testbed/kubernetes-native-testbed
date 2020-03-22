package main

import (
	"fmt"
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
)

type pointCacheRepository interface {
	findByUUID(string) (*PointCache, error)
	store(*PointCache) error
	deleteByUUID(string) error
}

type pointRepositoryMemcache struct {
	cache *memcache.Client
}

func (pr *pointRepositoryMemcache) findByUUID(uuid string) (*PointCache, error) {
	item, err := pr.cache.Get(uuid)
	if err != nil {
		return nil, fmt.Errorf("findByID error: %w (uuid: %s)", err, uuid)
	}
	amount, err := strconv.ParseInt(string(item.Value), 10, 32)
	if err != nil {
		return nil, fmt.Errorf("findByID error: %w (uuid: %s)", err, uuid)
	}

	pc := &PointCache{UserUUID: uuid, TotalAmount: int32(amount)}
	return pc, nil
}

func (pr *pointRepositoryMemcache) store(pc *PointCache) error {
	if err := pr.cache.Set(&memcache.Item{Key: pc.UserUUID, Value: []byte(string(pc.TotalAmount))}); err != nil {
		return fmt.Errorf("store error: %w (pointCache: %v)", err, pc)
	}
	return nil
}

func (pr *pointRepositoryMemcache) deleteByUUID(uuid string) error {
	if err := pr.cache.Delete(uuid); err != nil {
		return fmt.Errorf("deleteByID error: %w (uuid: %s)", err, uuid)
	}
	return nil
}
