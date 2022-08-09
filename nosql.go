package main

import (
	"log"
	"nosql/appinit"
	"nosql/controllers/enterapp"
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

	// let's start
	close := false
	for close != true {
		enterapp.LetsStart()
	}

	// Shutdown message
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
