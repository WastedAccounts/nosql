package main

import (
	"fmt"
	"log"
	"nosql/appinit"
	"nosql/controllers/mongodbplayground"
	"nosql/controllers/redisplayground"
	"nosql/controllers/validateinput"
)

func main() {

	log.Println("Server is starting")
	// //launch web site and serve requests
	// err := http.ListenAndServe(":3000", nil)
	// if err != nil {
	// 	//Throw error
	// }

	// initialize app
	appinit.Init()

	// prompt for input
	// make choice of which DB you want to play in
	fmt.Println("Nosql: You can type 'exit' at any prompt to close the app")
	prompt := "Nosql: Make a choice: Enter 1 for Mongo or 2 for Redis"
	valid := map[string]bool{"1": true, "2": true, "exit": true}
	decision := validateinput.PromptForInput(prompt, valid)
	choosePath(decision)
	log.Println("Shutdown complete")
}

// func promptForInput() {
// 	// make choice of which DB you want to play in
// 	fmt.Println("Nosql: Make a choice: Enter 1 for Mongo or 2 for Redis")
// 	fmt.Println("Nosql: You can type 'exit' to close the app at any time as well")
// 	//i := 0
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for scanner.Scan() {
// 		validateInput(scanner.Text())
// 	}
// }

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

//////// Just some junk I'm hording

// i := 0
// fmt.Println("Enter Key: ")
// scanner := bufio.NewScanner(os.Stdin)
// for scanner.Scan() {
// 	if len(scanner.Text()) == 0 {
// 		break
// 	}
// 	i++
// 	writekey(scanner.Text())
// }
