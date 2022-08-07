package validateinput

import (
	"bufio"
	"fmt"
	"os"
)

func PromptForInput(prompt string, valid map[string]bool) string {
	var validated string
	// make choice of which DB you want to play in
	fmt.Println(prompt)
	//i := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			fmt.Println("Nosql: Later nerd")
			os.Exit(1)
		}
		if scanner.Text() == "back" {
			os.Exit(1)
		}
		validated = validateInput(prompt, scanner.Text(), valid)
		break
	}
	return validated
}

func validateInput(prompt string, s string, valid map[string]bool) string {
	if !valid[s] {
		fmt.Println("Nosql: fool, don't be dumb")
		fmt.Println("Nosql:", s, "is not a valid choice")
		fmt.Println("Nosql: try again")
		PromptForInput(prompt, valid)
	} else {
		fmt.Println("Nosql:", s, "is an excellent choice")
		fmt.Println(s)
	}
	return s
}
