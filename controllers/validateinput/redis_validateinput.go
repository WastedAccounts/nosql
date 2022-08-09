package validateinput

import (
	"bufio"
	"fmt"
	"os"
)

func scanInput() string {
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

func GetKey() string {
	var key string
	// var err error
	fmt.Println("Nosql: key")
	key = scanInput()
	return key
}

func GetValue() string {
	var value string
	fmt.Println("Nosql: value")
	value = scanInput()
	return value
}

func GetExpire() string {
	var exStr string
	fmt.Println("Nosql: expire (default 3600")
	exStr = scanInput()
	if exStr == "" {
		exStr = "3600"
	}
	return exStr
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
