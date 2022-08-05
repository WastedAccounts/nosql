package datasources

import (
	"fmt"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

var (
	RedisConnPool *redis.Pool
	RedisReady    bool
)

func SetupRedis() bool {
	// Get environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Print("No .env file found")
	}
	initRedisPool()
	pingRedisPool(RedisConnPool.Get())
	log.Println("Redis initialization complete")
	return RedisReady
}

func initRedisPool() {
	var redisConnStr string
	REDIS_SVR := os.Getenv("REDIS_IP")
	REDIS_PORT := os.Getenv("REDIS_PORT")
	redisConnStr = REDIS_SVR + ":" + REDIS_PORT
	//fmt.Println("redconnstr", redisConnStr)
	RedisConnPool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", redisConnStr)
			if err != nil {
				log.Printf("ERROR: fail init redis pool: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
	log.Println("Connected to redis database")
}

func pingRedisPool(conn redis.Conn) {
	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		log.Printf("ERROR: fail ping redis conn: %s", err.Error())
		RedisReady = false
		// os.Exit(1)
	}
	log.Printf("Redis database ping succesfull")
	RedisReady = true
	// fmt.Println(pong, err)
}

// func SetDatasource() {
// 	// Get environment variables
// 	if err := godotenv.Load(); err != nil {
// 		log.Print("No .env file found")
// 	}
// 	SQLSERVER := os.Getenv("SQLSERVER")
// 	SQLUSER := os.Getenv("SQLUSER")
// 	SQLPW := os.Getenv("SQLPW")
// 	SQLPORT := os.Getenv("SQLPORT")
// 	SQLDBNAME := os.Getenv("SQLDBNAME")

// 	// Create Datasource string
// 	DataSource = SQLUSER + ":" + SQLPW + "@tcp(" + SQLSERVER + ":" + SQLPORT + ")/" + SQLDBNAME + "?parseTime=true"
// }
