package rate

import (
	"errors"
	"fmt"
	"time"

	"github.com/FZambia/sentinel"
	"github.com/gomodule/redigo/redis"
)

type RateRepository interface {
	FindByUUID(string) (*Rate, error)
	Store(*Rate) (string, error)
	Update(*Rate) error
	DeleteByUUID(string) error
}

type rateRepositoryRedis struct {
	config *RateRepositoryRedisConfig
	pool   *redis.Pool
}

func (rr *rateRepositoryRedis) FindByUUID(uuid string) (*Rate, error) {
	return nil, nil
}

func (rr *rateRepositoryRedis) Store(r *Rate) (string, error) {
	return "", nil
}

func (rr *rateRepositoryRedis) Update(r *Rate) error {
	return nil
}

func (rr *rateRepositoryRedis) DeleteByUUID(uuid string) error {
	return nil
}

type RateRepositoryRedisConfig struct {
	SentinelHost string
	SentinelPort int
	Password     string
	MasterName   string
}

func (c *RateRepositoryRedisConfig) Connect() RateRepository {
	sntnl := &sentinel.Sentinel{
		Addrs:      []string{fmt.Sprintf("%s:%d", c.SentinelHost, c.SentinelPort)},
		MasterName: c.MasterName,
		Dial: func(addr string) (redis.Conn, error) {
			timeout := 500 * time.Millisecond
			c, err := redis.DialTimeout("tcp", addr, timeout, timeout, timeout)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	pool := &redis.Pool{
		MaxIdle:     3,
		MaxActive:   64,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			masterAddr, err := sntnl.MasterAddr()
			if err != nil {
				return nil, err
			}
			c, err := redis.Dial("tcp", masterAddr)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if !sentinel.TestRole(c, "master") {
				return errors.New("Role check failed")
			} else {
				return nil
			}
		},
	}

	return &rateRepositoryRedis{pool: pool, config: c}
}
