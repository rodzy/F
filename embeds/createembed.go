package embeds

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

//CreateEmbed displays all the messaging for the create tab
func CreateEmbed() *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Flankerbot",
			URL:     "https://github.com/rodzy/F",
			IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/e70c1562f842ede26dbb80404b1206a5.webp?size=128",
		},
		Color:       0x66ccff,
		Description: "When starting out with a new repository, you only need to do it once; either locally, then push to GitHub, or by cloning an existing repository.",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Turn an existing directory into a git repository",
				Value:  "```shell\n$ git init```",
				Inline: false,
			},
			{
				Name:   "Clone (download) a repository that already exists on GitHub, including all of the files, branches, and commits",
				Value:  "```shell\n$ git clone [url]```",
				Inline: false,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://gitforwindows.org/img/gwindows_logo.png",
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "Git Cheat Sheet - Create repositories",
	}
	return embed
}
