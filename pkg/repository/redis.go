package repository

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"os"

	"github.com/go-redis/redis/v8"

	"github.com/dimaunx/go-clean-example/pkg/entity"
)

var redisHost = os.Getenv("REDIS_HOST")

type RedisRepo struct{}

func NewRedisRepository() *RedisRepo {
	return &RedisRepo{}
}

// NewRedis returns a client to the Redis Server
func NewRedis(host string, enableTls bool, passwd string, db int) *redis.Client {
	c := redis.NewClient(&redis.Options{Addr: host, DB: db})
	if enableTls {
		c.Options().TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	if passwd != "" {
		c.Options().Password = passwd
	}
	return c
}

func (RedisRepo) Save(ctx context.Context, d *entity.Device) (string, error) {
	r := NewRedis(redisHost, false, "", 0)
	data, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	err = r.Set(ctx, d.Id, data, 0).Err()
	if err != nil {
		return "", err
	}
	return d.Id, nil
}

func (RedisRepo) FindAll(ctx context.Context) ([]entity.Device, error) {
	var results []entity.Device
	r := NewRedis(redisHost, false, "", 0)
	iter := r.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		var d entity.Device
		item, err := r.Get(ctx, iter.Val()).Bytes()
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(item, &d)
		if err != nil {
			return nil, err
		}
		results = append(results, d)
	}
	if err := iter.Err(); err != nil {
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}

func (RedisRepo) FindById(ctx context.Context, id string) (*entity.Device, error) {
	r := NewRedis(redisHost, false, "", 0)
	device, err := r.Get(ctx, id).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var result entity.Device
	err = json.Unmarshal(device, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
