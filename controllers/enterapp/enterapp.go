package enterapp

import (
	"fmt"
	"nosql/controllers/mongodbplayground"
	"nosql/controllers/redisplayground"
	"nosql/controllers/validateinput"
)

func LetsStart() {
	// prompt for input
	// make choice of which DB you want to play in
	fmt.Println("Nosql: You can type 'exit' at any prompt to close the app")
	prompt := "Nosql: Make a choice: Enter 1 for Mongo or 2 for Redis"
	valid := map[string]bool{"1": true, "2": true, "exit": true}
	decision := validateinput.PromptForInput(prompt, valid)
	choosePath(decision)
}

func choosePath(s string) {
	valid := map[string]bool{"1": true, "2": true}
	if valid[s] {
		if s == "1" {
			fmt.Println("Nosql: You've chosen: Mongo")
			fmt.Println("Nosql: Let's begin")
			mongodbplayground.MongoPlayground()
		}
		if s == "2" {
			fmt.Println("Nosql: You've chosen: Redis")
			fmt.Println("Nosql: Let's begin")
			redisplayground.RedisPlayground()
		}
	}
}
