package handlers

import (
	"github.com/bwmarrin/discordgo"
)

//StateHandler controls the status of the bot
func StateHandler(session *discordgo.Session, event *discordgo.Ready)  {
	session.UpdateListeningStatus("$help")
}