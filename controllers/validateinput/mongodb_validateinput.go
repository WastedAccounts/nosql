package validateinput

import (
	"bufio"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

func scanMongoInput() string {
	var value string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			fmt.Println("Nosql: Later nerd")
			os.Exit(1)
		} else {
			value = scanner.Text()
			break
		}
	}
	return value
}

var Bdoc interface{}

func MarshalBson(value string) {
	// `{"description":"matt"}`
	err := bson.UnmarshalExtJSON([]byte(value), true, &Bdoc)
	if err != nil {
		fmt.Println("Nosql: ERROR: badly formatted json data")
	}
}

// func PromptForInputRedis(prompt string, valid map[string]bool) string {
// 	var validated string
// 	// make choice of which DB you want to play in
// 	fmt.Println(prompt)
// 	//i := 0
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for scanner.Scan() {
// 		if scanner.Text() == "exit" {
// 			fmt.Println("Nosql: Later nerd")
// 			os.Exit(1)
// 		}
// 		validated = validateInput(prompt, scanner.Text(), valid)
// 		//fmt.Println("validated:", validated)
// 		return validated
// 	}
// 	return validated
// }

// func validateInputRedis(prompt string, s string, valid map[string]bool) string {
// 	if !valid[s] {
// 		fmt.Println("Nosql: fool, don't be dumb")
// 		fmt.Println("Nosql:", s, "is not a valid choice")
// 		fmt.Println("Nosql: try again")
// 		PromptForInput(prompt, valid)
// 	} else {
// 		//fmt.Println(valid)
// 		fmt.Println("Nosql:", s, "is an excellent choice")
// 		fmt.Println(s)
// 		return s
// 	}
// 	return s
// }
