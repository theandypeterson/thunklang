package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {

	if len(os.Args) < 2 {
		setupRepl()
	}

	filename := os.Args[1]

	dat, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", filename, err.Error())
	}

	evaluated, errors := repl.Run(string(dat))

	if len(errors) > 0 {
		printParserErrors(errors)
	}

	if evaluated != nil {
		fmt.Println(evaluated.Inspect())
	}
}

func setupRepl() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}

func printParserErrors(errors []string) {
	for _, msg := range errors {
		fmt.Println("\t" + msg + "\n")
	}
}
