package embeds

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

//InstallEmbed displays all the information for the git installations
func InstallEmbed() *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Flankerbot",
			URL:     "https://github.com/rodzy/F",
			IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/e70c1562f842ede26dbb80404b1206a5.webp?size=128",
		},
		Color:       0x66ccff,
		Description: "Get git on your computer",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "GitHub for Windows",
				Value:  "https://windows.github.com/",
				Inline: false,
			},
			{
				Name:   "GitHub for Mac",
				Value:  "https://mac.github.com/",
				Inline: false,
			},
			{
				Name:   "Git for All Platforms",
				Value:  "https://git-scm.com/",
				Inline: false,
			},
			{
				Name:   "*",
				Value:  "Git distributions for Linux and POSIX systems are available on the official Git SCM web site.",
				Inline: false,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://gitforwindows.org/img/gwindows_logo.png",
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "Git Cheat Sheet - Install",
	}

	return embed
}
