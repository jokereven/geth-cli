package main

import (
	"log"

	"github.com/jokereven/geth-cli/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute() failed err: %v", err)
	}
}
