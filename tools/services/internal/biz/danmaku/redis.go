package danmaku

import (
	"github.com/go-redis/redis"
	"launcher/internal/server"
	"launcher/internal/utils"
	"log"
	"time"
)

var danmakuRedisDB *redis.Client
var danmakuRedisConnected = false

func GetDanmakuRedisConnection() *redis.Client {
	if danmakuRedisConnected {
		return danmakuRedisDB
	}
	return nil
}

func InitRedisConnect () {
	var err error
	for {
		danmakuRedisDB, err = utils.NewRedisClient(server.DanmakuRedisDB)
		if err != nil {
			danmakuRedisConnected = false
			log.Printf("[Redis] %s\n", err.Error())
			time.Sleep(3 * time.Second)  // wait 3 seconds.
			continue
		}
		danmakuRedisConnected = true
		log.Printf("[redis] DB %d connected.", server.DanmakuRedisDB)
		for {
			time.Sleep(10 * time.Second)  // ping pre 10 seconds.
			err = danmakuRedisDB.Ping().Err()
			if err != nil {
				danmakuRedisConnected = false
				log.Printf("[Redis] %s\n", err.Error())
				time.Sleep(3 * time.Second)  // wait 3 seconds.
				break
			}
		}
	}
}

