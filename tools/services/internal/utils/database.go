package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// NewRedisClient 初始化redis连接
func NewRedisClient(db int, addr, passwd string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:        addr, // Redis地址
		Password:    passwd,                                             // Redis账号
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


// NewGormConnection 建立Gorm客户端
func NewGormConnection(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{
		//DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DSN: dsn,
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
}
