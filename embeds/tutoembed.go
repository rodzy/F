package embeds

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

//TutoEmbed displays the tuto embed
func TutoEmbed() *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Flankerbot",
			URL:     "https://github.com/rodzy",
			IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/16ea3509f9a41b475966b85565a930fa.webp?size=128",
		},
		Color:       0x66ccff,
		Description: "Just in case - Quick setup\n **REMEMBER: YOU NEED TO CHANGE EVERYTHING THAT'S ON SQUARE BRACKETS**",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Create a new repository",
				Value:  "```shell\n$ echo # [repo-name] >> README.md\n$ git init git add README.md\n$ git commit -m [first commit]\n$ git remote add origin https://github.com/[user]/[repo-name].git\n$ git push -u origin master```",
				Inline: false,
			},
			{
				Name:   "Pushing an existing repository from the command line",
				Value:  "```shell\n$ git remote add origin https://github.com/[user]/[repo-name].git\n$ git push -u origin master```",
				Inline: false,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://gitforwindows.org/img/gwindows_logo.png",
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "GitHub Repo - Quick setup",
	}

	return embed
}
