package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rodzy/Flanker/structs"
)

//MessageHandler controls the way messages are displayed
func MessageHandler(session *discordgo.Session, message *discordgo.Message)  {
	if strings.HasPrefix(message.Content, "$") {
		if message.Author.ID == structs.ID {
			return
		}
	}
}