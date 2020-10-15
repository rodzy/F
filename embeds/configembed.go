package embeds

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

//ConfigEmbed displays the embed about the config tab
func ConfigEmbed() *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Flankerbot",
			URL:     "https://github.com/rodzy",
			IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/16ea3509f9a41b475966b85565a930fa.webp?size=128",
		},
		Color:       0x66ccff,
		Description: "Configure user information for all local repositories",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Sets the name you want attached to your commit transactions",
				Value:  "```shell\n$ git config --global user.name [name]```",
				Inline: false,
			},
			{
				Name:   "Sets the email you want attached to your commit transactions",
				Value:  "```shell\n$ git config --global user.email [email address]```",
				Inline: false,
			},
			{
				Name:   "Enables helpful colorization of command line output",
				Value:  "```shell\n$ git config --global color.ui auto```",
				Inline: false,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://gitforwindows.org/img/gwindows_logo.png",
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "Git Cheat Sheet - Configure tooling",
	}
	return embed
}

