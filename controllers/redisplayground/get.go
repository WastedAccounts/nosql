package redisplayground

import (
	"bufio"
	"fmt"
	"nosql/datasources"
	"os"

	"github.com/gomodule/redigo/redis"
)

func getRead() {

	fmt.Println("Nosql: What do you want to read... nerd")
	fmt.Println("Nosql: 1 - all, 2 - one, back - to go back")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "1" {
			getAll()
			getRead()
		}
		if scanner.Text() == "2" {
			fmt.Println("Nosql: Which key?")
			for scanner.Scan() {
				if scanner.Text() == "back" {
					getRead()
				}
				getOne(scanner.Text())
			}
			getRead()
		}
		if scanner.Text() == "back" {
			RedisPlayground()
		}
	}

}

func getAll() {
	// open connection to redis and defer close
	conn := datasources.RedisConnPool.Get()
	defer conn.Close()

	getAll, err := redis.Strings(conn.Do("Keys", "*"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Nosql: Get ALL Keys:\n", getAll)
}

func getOne(key string) {
	// open connection to redis and defer close
	conn := datasources.RedisConnPool.Get()
	defer conn.Close()

	getOne, err := redis.String(conn.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Nosql: Get Key \n", getOne)
}
