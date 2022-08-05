package redisplayground

import (
	"fmt"
	"nosql/controllers/validateinput"
)

func RedisPlayground() {
	fmt.Println("Nosql: Welcome to Redis land")
	// prompt for input
	// promptForInput()
	////promptForInput()
	prompt := "Nosql: What do you want to do? \n Nosql: 1 - Create, 2 - Read, 3 - Update, 4 - Delete"
	valid := map[string]bool{"1": true, "2": true, "3": true, "4": true}
	decision := validateinput.PromptForInput(prompt, valid)
	choosePath(decision)
}

func choosePath(s string) {
	valid := map[string]bool{"1": true, "2": true}
	if valid[s] {
		if s == "1" {
			fmt.Println("Nosql: Create it is")
		}
		if s == "2" {
			fmt.Println("Nosql: Reading is for dorks")
		}
		if s == "3" {
			fmt.Println("Nosql: Updating in prod, brave")
		}
		if s == "4" {
			fmt.Println("Nosql: Let's delete some shit")
		}
	}
}

// func promptForInput() {
// 	fmt.Println("Nosql: What do you want to do?")
// 	fmt.Println("Nosql: 1 - Create, 2 - Read, 3 - Update, 4 - Delete")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for scanner.Scan() {
// 		validateInput(scanner.Text())
// 	}
// }

// func validateInput(s string) {
// 	valid := map[string]bool{"1": true, "2": true, "3": true, "4": true, "exit": true}
// 	if valid[s] {
// 		if s == "1" {
// 			fmt.Println("Nosql: Create it is")
// 		}
// 		if s == "2" {
// 			fmt.Println("Nosql: Reading is for dorks")
// 		}
// 		if s == "3" {
// 			fmt.Println("Nosql: Updating in prod, brave")
// 		}
// 		if s == "4" {
// 			fmt.Println("Nosql: Let's delete some shit")
// 		}
// 		if s == "exit" {
// 			fmt.Println("Nosql: Later nerd")
// 			os.Exit(1)
// 		}
// 	} else {
// 		fmt.Println("Nosql: fool, don't be dumb")
// 		fmt.Println("Nosql: try again")
// 		promptForInput()
// 	}
// }

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
