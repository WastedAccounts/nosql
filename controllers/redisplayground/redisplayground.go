package redisplayground

import (
	"fmt"
	"nosql/controllers/validateinput"
)

func RedisPlayground() {

	fmt.Println("Nosql: Welcome to Redis land")
	// prompt for input
	prompt := "Nosql: What do you want to do? \nNosql: 1 - Set, 2 - Read, 3 - Delete, back - to go back"
	valid := map[string]bool{"1": true, "2": true, "3": true, "back": true}
	decision := validateinput.PromptForInput(prompt, valid)
	if decision != "back" {
		choosePath(decision)
	}
}

func choosePath(s string) {
	if s == "1" {
		fmt.Println("Nosql: Set it, set it good")
		setToRedis()
		RedisPlayground()
	}
	if s == "2" {
		fmt.Println("Nosql: Reading is for dorks")
		getRead()
		RedisPlayground()
	}
	if s == "3" {
		fmt.Println("Nosql: Let's delete some shit")
		delFromRedis()
		RedisPlayground()
	}
}

// ///// Web stuff if I want to make it a web app

// // import (
// // 	"net/http"
// // 	"regexp"
// // )

// // type redisPlaygroundController struct {
// // 	redisPlaygroundPattern *regexp.Regexp
// // }

// // // entry point from front.go
// // func NewRedisPlaygroundController() *redisPlaygroundController {
// // 	return &redisPlaygroundController{
// // 		redisPlaygroundPattern: regexp.MustCompile(`^/redisplayground.html(\d+)/?`),
// // 	}
// // }

// // // ServeHTTP - Entry point from front.go
// // func (rpgc redisPlaygroundController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// // 	http.ServeFile(w, r, "/redisplayground.html")
// // 	//pageloadRedisPlayground(w, r)
// // }

// // // pageLoadWorkouts - switch to workouts template do work there
// // func pageloadRedisPlayground(w http.ResponseWriter, r *http.Request) {
// // 	http.ServeFile(w, r, "/redisplayground.html")
// // }
