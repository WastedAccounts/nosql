package redisplayground

import (
	//"context"

	"fmt"
	"log"
	"nosql/controllers/validateinput"
	"nosql/datasources"

	"github.com/gomodule/redigo/redis"
)

func setToRedis() {
	var key, val, ex string
	key = validateinput.GetKey()
	val = validateinput.GetValue()
	ex = validateinput.GetExpire()
	setOne(key, val, ex)
}

func setOne(key string, val string, ex string) {
	// open connection to redis and defer close
	conn := datasources.RedisConnPool.Get()
	defer conn.Close()

	log.Println(key, val, ex)

	// set value to redis
	set, err := redis.String(conn.Do("SET", key, val, "EX", ex))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Nosql: Save key-value pair result:", set)
}

// func setMany() {

// }
