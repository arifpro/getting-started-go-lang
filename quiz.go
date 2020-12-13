package main

import (
	"fmt"
	"strings"
)

func main() {
	var userInput string
	fmt.Print("Please enter text:")
	fmt.Scan(&userInput)
	userInput = strings.ToLower(userInput)
	var startWithI bool = strings.HasPrefix(userInput, "i")
	var endWithN bool = strings.HasSuffix(userInput, "n")
	var containA bool = strings.Contains(userInput, "a")
	if startWithI && endWithN && containA {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}