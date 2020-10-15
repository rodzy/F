package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/rodzy/Flanker/handlers"
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

	Session.AddHandler(handlers.StateHandler)

	err = Session.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Flanker running, type $help for commands")
	<-make(chan struct{})
	return
}
