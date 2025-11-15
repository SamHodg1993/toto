package main

import (
	"github.com/joho/godotenv"
	"github.com/odgy8/toto/cmd"
)

func main() {
	godotenv.Load()
	cmd.Execute()
}
