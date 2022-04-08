package kvengine

import (
	"context"

	"github.com/Raobian/bgofs/pkg/common/log"
	"github.com/go-redis/redis/v8"
)

const (
	Addr   = "127.0.0.1:6379"
	Passwd = ""
)

type KVEngineRedis struct {
	cli *redis.Client
}

func NewRedisKV() *KVEngineRedis {
	op := &redis.Options{
		Addr:     Addr,
		Password: Passwd,
		DB:       0,
	}
	cli := redis.NewClient(op)

	log.DINFO("redis connect success")

	return &KVEngineRedis{
		cli: cli,
	}
}

func (kv *KVEngineRedis) Close() error {
	return kv.cli.Close()
}

func (kv *KVEngineRedis) Set(key string, value []byte) error {
	return kv.cli.Set(context.Background(), key, value, 0).Err()
}

func (kv *KVEngineRedis) Get(key string) ([]byte, error) {
	result, err := kv.cli.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, ENotFound
	}

	return []byte(result), nil
}

func (kv *KVEngineRedis) Delete(key string) error {
	return kv.cli.Del(context.Background(), key).Err()
}
