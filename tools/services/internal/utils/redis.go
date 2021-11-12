package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"launcher/internal/server"
	"time"
)

// NewRedisClient 初始化redis连接
func NewRedisClient(db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", server.Config.Redis.Host, server.Config.Redis.Port), // Redis地址
		Password:    server.Config.Redis.Password,                                             // Redis账号
		DB:          db,                                                                       // Redis库
		PoolSize:    16,                                                                       // Redis连接池大小
		MaxRetries:  3,                                                                        // 最大重试次数
		IdleTimeout: 10 * time.Second,                                                         // 空闲链接超时时间
	})
	_, err := client.Ping().Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("[redis] connection failed")
	} else if err != nil {
		return nil, fmt.Errorf("[redis] connection failed:%s", err)
	}
	return client, nil
}
