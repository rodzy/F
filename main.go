package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	// "github.com/rodzy/Flanker/config"
)

func main()  {
	err:=godotenv.Load();
	if err != nil {
		log.Fatal("Error loading the env file")
	}

	test:=os.Getenv("DISCORD_TOKEN")
	fmt.Print(test)

	// erre:=config.ReadConfig
	// if erre != nil {
	// 	fmt.Print("Connection failed")
	// }
}