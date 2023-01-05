package redis

import (
	"fmt"
	"github.com/ccloud/phase2/pkg/config"
	"time"

	"log"

	"github.com/go-redis/redis"
)

var Cache *RedisRepo

type RedisRepo struct {
	client                *redis.Client
	defaultExpirationTime int
}

func SetupRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Conf.Redis.Host, config.Conf.Redis.Port),
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.Db,
	})
	// test connection
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	// return pong if server is online
	fmt.Println(pong)
	fmt.Println("redis connected!!!")
	Cache = &RedisRepo{
		client:                client,
		defaultExpirationTime: config.Conf.Redis.ExpireTime,
	}
}

func (r *RedisRepo) SetPrice(name string, price float64) error {
	return r.client.Set(name, price, time.Duration(r.defaultExpirationTime)*time.Second).Err()
}
func (r *RedisRepo) CheckPrice(name string) (float64, error) {
	value, err := r.client.Get(name).Float64()
	if err != nil {
		return -1.0, err
	}
	return value, nil
}
