package main

import (
	"github.com/joho/godotenv"
	"github.com/samhodg1993/toto-todo-cli/cmd"
)

func main() {
	godotenv.Load()
	cmd.Execute()
}
