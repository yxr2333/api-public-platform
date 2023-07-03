package utils

import (
	"api-public-platform/config"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// TODO: 对封装的Redis工具类进行测试
type RedisUtil struct {
	client *redis.Client
}

func NewRedisUtil() (*RedisUtil, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.ServerCfg.Redis.Address,
		Password: config.ServerCfg.Redis.Password,
		DB:       config.ServerCfg.Redis.DB,
	})
	// 检查是否连接成功
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("redis connect failed: %v", err)
	}
	return &RedisUtil{client: client}, nil
}

func NewRedisUtilCustom(addr, password string, db int) (*RedisUtil, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	// 检查是否连接成功
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("redis connect failed: %v", err)
	}
	return &RedisUtil{client: client}, nil
}

func (ru *RedisUtil) Get(key string) (string, error) {
	ctx := context.Background()
	val, err := ru.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("key %s does not exist", key)
		} else {
			return "", fmt.Errorf("redis get failed: %v", err)
		}
	}
	return val, nil
}

func (ru *RedisUtil) Set(key string, value interface{}) error {
	ctx := context.Background()
	err := ru.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return fmt.Errorf("redis set failed: %v", err)
	}
	return nil
}

func (ru *RedisUtil) Expire(key string, expiration time.Duration) error {
	ctx := context.Background()
	err := ru.client.Expire(ctx, key, expiration).Err()
	if err != nil {
		return fmt.Errorf("redis expire failed: %v", err)
	}
	return nil
}

func (ru *RedisUtil) Del(key string) error {
	ctx := context.Background()
	err := ru.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("redis del failed: %v", err)
	}
	return nil
}

func (ru *RedisUtil) IsExpired(key string) (bool, error) {
	ctx := context.Background()
	ttl, err := ru.client.TTL(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("failed to get expiration for key: %v", err)
	}
	if ttl.Seconds() < 0 {
		return true, nil
	}
	return false, nil
}

func (ru *RedisUtil) IsExist(key string) (bool, error) {
	ctx := context.Background()
	val, err := ru.client.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check if key exists: %v", err)
	}
	if val == 0 {
		return false, nil
	}
	return true, nil
}

func (ru *RedisUtil) Close() error {
	err := ru.client.Close()
	if err != nil {
		return fmt.Errorf("redis close failed: %v", err)
	}
	return nil
}
