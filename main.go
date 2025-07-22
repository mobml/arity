package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const red = "\033[31m"
const green = "\033[32m"
const yellow = "\033[33m"
const blue = "\033[34m"
const reset = "\033[0m"

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
	reader := bufio.NewReader(os.Stdin)
	lex := NewLexer("")
	for {
		lex.CleanLexer()
		fmt.Printf("%s[%d]%s ", blue, line, reset)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("An error has ocurred: ", err)
			break
		}

		input = strings.TrimSpace(input)
		line++
		if len(input) == 0 {
			fmt.Printf("%sThe input cannot be empty%s\n", yellow, reset)
			continue
		}

		if input == "exit" {
			break
		}
		lex.SetLine(input)
		lex.ScanTokens()
		lex.Show()
		//fmt.Printf("%s%s%s\n", green, input, reset)
	}
}
