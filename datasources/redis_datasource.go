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

// SetupRedis -- exposed call for appinit to reach and start up the connection pool for redis
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

// initRedisPool -- initializes pool and create global var
func initRedisPool() *redis.Pool {
	REDIS_SVR := os.Getenv("REDIS_IP")
	REDIS_PORT := os.Getenv("REDIS_PORT")
	redisConnStr := REDIS_SVR + ":" + REDIS_PORT
	redisConnPool := &redis.Pool{
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

	// Set Global var for redis connection pool and return it
	RedisConnPool = redisConnPool
	return RedisConnPool
}

// pingRedisPool -- Pings pool to confirm it's working and open initiation connection
func pingRedisPool(conn redis.Conn) {
	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		log.Printf("ERROR: fail ping redis conn: %s", err.Error())
		RedisReady = false
	}
	log.Printf("Redis database ping succesfull")
	RedisReady = true
}
