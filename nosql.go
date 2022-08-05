package main

import (
	"bufio"
	"fmt"
	"log"
	"nosql/appinit"
	"os"
)

func main() {

	log.Println("Server is starting")
	// launch web site and serve requests
	// err := http.ListenAndServe(":3000", nil)
	// if err != nil {
	// 	//Throw error
	// }

	// initialize app
	appinit.Init()

	// Test datasource

	i := 0
	fmt.Println("Enter Key: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		i++
		writekey(scanner.Text())
	}
	log.Println("Shutdown complete")
}

func writekey(s string) {

}
