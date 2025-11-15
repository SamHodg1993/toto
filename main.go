package main

import (
	"github.com/joho/godotenv"
	"github.com/ODGY8/toto/cmd"
)

func main() {
	godotenv.Load()
	cmd.Execute()
}
