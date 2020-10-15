package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rodzy/Flanker/bot"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the env file")
	}
	token := os.Getenv("DISCORD_TOKEN")
	
	if token == "" {
		fmt.Println("Couldn't allocate for Token")
		return
	}
	bot.Start(token)

}
