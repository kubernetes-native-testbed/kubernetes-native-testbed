package rate

import (
	"errors"
	"fmt"
	"time"

	"github.com/FZambia/sentinel"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

type RateRepository interface {
	FindByUUID(string) (*Rate, error)
	Store(*Rate) (string, error)
	Update(*Rate) error
	DeleteByUUID(string) error
}

type RateNotFoundError error

func IsNotFound(err error) bool {
	switch err.(type) {
	case RateNotFoundError:
		return true
	}
	return false
}

type rateRepositoryRedis struct {
	config *RateRepositoryRedisConfig
	pool   *redis.Pool
}

func (rr *rateRepositoryRedis) FindByUUID(uuid string) (*Rate, error) {
	conn := rr.pool.Get()
	defer conn.Close()

	v, err := redis.Values(conn.Do("HGETALL", uuid))
	if err != nil {
		return nil, err
	}

	var r Rate
	if err := redis.ScanStruct(v, &r); err != nil {
		return nil, err
	}

	if r.DeletedAt != 0 {
		return nil, fmt.Errorf("%s is deleted", uuid)
	}

	return &r, nil
}

func (rr *rateRepositoryRedis) Store(r *Rate) (string, error) {
	conn := rr.pool.Get()
	defer conn.Close()

	r.UUID = uuid.New().String()
	r.CreatedAt = time.Now().Unix()

	reply, err := conn.Do("EXISTS", r.UUID)
	if err != nil {
		return "", err
	}
	if reply == 1 { // exists
		return "", fmt.Errorf("the uuid is already exists: %s", r.UUID)
	}

	reply, err = conn.Do("HMSET", redis.Args{}.Add(r.UUID).AddFlat(r)...)
	if err != nil {
		return "", err
	}
	if ret, ok := reply.(string); !ok {
		return "", fmt.Errorf("assert return value failed: %#v", ret)
	} else if ret != "OK" {
		return "", fmt.Errorf("failed to store: %s", ret)
	}

	return r.UUID, nil
}

func (rr *rateRepositoryRedis) Update(r *Rate) error {
	conn := rr.pool.Get()
	defer conn.Close()

	r.UpdatedAt = time.Now().Unix()

	reply, err := conn.Do("EXISTS", r.UUID)
	if err != nil {
		return err
	}
	if reply == 0 { // not exists
		return fmt.Errorf("the uuid is not found: %s", r.UUID)
	}

	reply, err = conn.Do("HMSET", redis.Args{}.Add(r.UUID).AddFlat(r)...)
	if err != nil {
		return err
	}
	if ret, ok := reply.(string); !ok {
		return fmt.Errorf("assert return value failed: %#v", ret)
	} else if ret != "OK" {
		return fmt.Errorf("failed to store: %s", ret)
	}

	return nil
}

func (rr *rateRepositoryRedis) DeleteByUUID(uuid string) error {
	conn := rr.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("EXISTS", uuid)
	if err != nil {
		return err
	}
	if reply == 0 { // not exists
		return fmt.Errorf("the uuid is not found: %s", uuid)
	}

	reply, err = conn.Do("HSET", uuid, "DeletedAt", time.Now().Unix())
	if err != nil {
		return err
	}
	if reply == 0 || reply == 1 { // succeed to store exists key or non-exists key
		return fmt.Errorf("set DeletedAt is failed: %s", uuid)
	}

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
			conn, err := redis.Dial("tcp", masterAddr)
			if err != nil {
				return nil, err
			}
			if _, err := conn.Do("AUTH", c.Password); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if !sentinel.TestRole(conn, "master") {
				return errors.New("Role check failed")
			} else {
				return nil
			}
		},
	}

	return &rateRepositoryRedis{pool: pool, config: c}
}
