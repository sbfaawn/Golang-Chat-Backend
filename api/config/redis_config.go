package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisOption struct {
	Address  string
	Port     string
	Password string
	DbNum    int
}

type redisConnection struct {
	Address  string
	Port     string
	Password string
	DbNum    int
	client   *redis.Client
}

func NewRedisConnection(redisOption RedisOption) *redisConnection {
	return &redisConnection{
		Address:  redisOption.Address,
		Port:     redisOption.Port,
		Password: redisOption.Password,
		DbNum:    redisOption.DbNum,
	}
}

func (rc *redisConnection) ConnectToRedis() error {

	rc.client = redis.NewClient(
		&redis.Options{
			Addr:     rc.Address + ":" + rc.Port,
			Password: rc.Password,
			DB:       rc.DbNum,
		},
	)

	ctx := context.Background()

	if err := rc.client.Conn().ClientSetName(ctx, "myclient").Err(); err != nil {
		panic(err)
	}

	pong, err := rc.client.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("No Connection to Redis\n", err)
		return err
	}

	log.Println("Redis Connection is Established", pong)
	return nil
}

func (rc *redisConnection) GetClient() *redis.Client {
	return rc.client
}
