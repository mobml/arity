package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) > 0 {
		inlineMode(&args)
	}
	promptMode()
}

func inlineMode(args *[]string) {
	fmt.Println("Running in inline mode, args: ", *args)
	os.Exit(1)
}

func promptMode() {
	line := 1
	for {
		var input string
		fmt.Printf("[%d] ", line)
		fmt.Scan(&input)
		fmt.Printf("%s\n", input)
		line++
		if input == "exit" {
			break
		}
	}
}
