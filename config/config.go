package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//Token and bot command by default
var (
	Token  string
	BotCom string
	config *configstruct
)

type configstruct struct {
	Token  string ` json:"Token"`
	BotCom string ` json:"BotCom"`
}

//ReadConfig reads directly from the config file
func ReadConfig() error {
	file, err := ioutil.ReadFile("./token.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotCom = config.BotCom
	return nil

}
