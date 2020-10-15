package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rodzy/Flanker/embeds"
	"github.com/rodzy/Flanker/structs"
)

//MessageHandler controls the way messages are displayed
func MessageHandler(session *discordgo.Session, message *discordgo.Message) {
	if strings.HasPrefix(message.Content, "$") {
		if message.Author.ID == structs.ID {
			return
		}
	}
	if message.Content == "$help" {
		_, _ = session.ChannelMessageSendEmbed(message.ChannelID, embeds.HelpEmbed())
	}
	if message.Content == "$install" {
		_, _ = session.ChannelMessageSendEmbed(message.ChannelID, embeds.InstallEmbed())
	}
	if message.Content == "$branch" {
		_, _ = session.ChannelMessageSendEmbed(message.ChannelID, embeds.BranchEmbed())
	}
	if message.Content == "$create" {
		_, _ = session.ChannelMessageSendEmbed(message.ChannelID, embeds.CreateEmbed())
	}
	if message.Content == "$config" {
		_, _ = session.ChannelMessageSendEmbed(message.ChannelID, embeds.ConfigEmbed())
	}
	if message.Content == "$changes" {
		_, _ = session.ChannelMessageSendEmbed(message.ChannelID, embeds.ChangesEmbed())
	}
	if message.Content == "$redo" {
		_, _ = session.ChannelMessageSendEmbed(message.ChannelID, embeds.RedoEmbed())
	}
	if message.Content == "$info" {
		_, _ = session.ChannelMessageSendEmbed(message.ChannelID, embeds.InfoEmbed())
	}
}
