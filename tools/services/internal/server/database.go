package server

import (
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"launcher/internal/utils"
	"log"
	"time"
)

// 全局Gorm客户端
var GormClient *gorm.DB
var mainRedisDB *redis.Client
var mainRedisConnected = false


// 完成服务器的所有DB连接
func InitDBConnection() error {
	err := InitMainGormClient()
	if err != nil { return err }
	// 连接Redis
	go InitMainRedisConnect()
	return nil
}

func InitMainGormClient() error  {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		Config.MySQL.User,
		Config.MySQL.Password,
		Config.MySQL.Host,
		Config.MySQL.Port,
		Config.Server.MySQLDBName,
	)

	var err error
	GormClient, err = utils.NewGormConnection(dsn)
	if err != nil {
		return err
	}
	return nil
}

// 获取数据库
func GetMainRedisConnection() *redis.Client {
	if mainRedisConnected {
		return mainRedisDB
	}
	return nil
}

func InitMainRedisConnect () {
	var err error
	for {
		mainRedisDB, err = utils.NewRedisClient(
			Config.Server.RedisDB,
			fmt.Sprintf("%s:%d", Config.Redis.Host, Config.Redis.Port),
			Config.Redis.Password,
		)
		if err != nil {
			mainRedisConnected = false
			log.Printf("[Redis] %s\n", err.Error())
			time.Sleep(3 * time.Second)  // wait 3 seconds.
			continue
		}
		mainRedisConnected = true
		log.Printf("[redis] DB %d connected.", mainRedisDB)
		for {
			time.Sleep(10 * time.Second)  // ping pre 10 seconds.
			err = mainRedisDB.Ping().Err()
			if err != nil {
				mainRedisConnected = false
				log.Printf("[Redis] %s\n", err.Error())
				time.Sleep(3 * time.Second)  // wait 3 seconds.
				break
			}
		}
	}
}