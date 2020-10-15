package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/rodzy/Flanker/handlers"
	"github.com/rodzy/Flanker/structs"
)


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

	structs.ID = user.ID

	Session.AddHandler(handlers.StateHandler)
	Session.AddHandler(handlers.MessageHandler)

	err = Session.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Flanker running, type $help for commands")
	<-make(chan struct{})
	return
}
