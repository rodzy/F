package bot

import (
	"fmt"

	"github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
)

//ID var
var ID string

//Session var
var Session *discordgo.Session

//Start initializates the process of conection
func Start(token string) {
	Session, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := Session.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ID = user.ID

	Session.AddHandler(StateHandler)

}
