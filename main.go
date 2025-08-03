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

	repl.Run(os.Stdin, os.Stdout, string(dat))
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
