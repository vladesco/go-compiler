package main

import (
	"compiler/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	currentUser, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is REPL, feel fre to input any command\n", currentUser.Username)
	repl.Start(os.Stdin, os.Stdout)

}
