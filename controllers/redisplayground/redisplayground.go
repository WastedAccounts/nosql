package redisplayground

import (
	"fmt"
	"nosql/controllers/validateinput"
	"os"
)

func RedisPlayground() {
	fmt.Println("Nosql: Welcome to Redis land")
	// prompt for input
	prompt := "Nosql: What do you want to do? \nNosql: 1 - Create, 2 - Read, 3 - Update, 4 - Delete, back - to go back"
	valid := map[string]bool{"1": true, "2": true, "3": true, "4": true}
	decision := validateinput.PromptForInput(prompt, valid)
	choosePath(decision)
}

func choosePath(s string) {
	var key, val, ex string

	if s == "1" {
		fmt.Println("Nosql: Create it is")
		key = validateinput.GetKey()
		val = validateinput.GetValue()
		ex = validateinput.GetExpire()
		setToRedis(key, val, ex)
		RedisPlayground()
	}
	if s == "2" {
		fmt.Println("Nosql: Reading is for dorks")
		getRead()
		RedisPlayground()
	}
	if s == "3" {
		fmt.Println("Nosql: Updating in prod, brave")

		RedisPlayground()
	}
	if s == "4" {
		fmt.Println("Nosql: Let's delete some shit")

		RedisPlayground()
	}
	if s == "back" {
		fmt.Println("Nosql: going back")
		os.Exit(2)
	}
}

func create() {

}
func read() {

}
func update() {

}
func delete() {

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
