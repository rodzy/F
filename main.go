package main

import (
	"fmt"

	"github.com/rodzy/Flanker/config"
)

func main()  {
	err:=config.ReadConfig
	if err != nil {
		fmt.Print("Connection failed")
	}
}