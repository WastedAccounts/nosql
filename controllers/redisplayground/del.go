package redisplayground

import (
	"fmt"
	"log"
	"nosql/controllers/validateinput"
	"nosql/datasources"

	"github.com/gomodule/redigo/redis"
)

func delFromRedis() {
	var key string
	key = validateinput.GetKey()
	delOne(key)
}

func delOne(key string) {
	// open connection to redis and defer close
	conn := datasources.RedisConnPool.Get()
	defer conn.Close()

	fmt.Println("Nosql: Deleting key:", key)

	// set value to redis
	del, err := redis.Int64(conn.Do("DEL", key))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Nosql: Deleting key-value result:", del)
}
