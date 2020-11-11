package Config

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

var redisCount = 0

func RedisConnection() redis.Conn {
	c, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Println("Not ready, Retry connecting...")
		time.Sleep(time.Second)
		redisCount++
		log.Println(redisCount)
		if redisCount > 30 {
			panic(err)
		}
		return RedisConnection()
	}
	return c
}