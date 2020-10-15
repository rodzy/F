package embeds

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

//ChangesEmbed displays the changes tab
func ChangesEmbed() *discordgo.MessageEmbed  {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Flankerbot",
			URL:     "https://github.com/rodzy",
			IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/16ea3509f9a41b475966b85565a930fa.webp?size=128",
		},
		Color:       0x66ccff,
		Description: "Browse and inspect the evolution of project files",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Lists version history for the current branch",
				Value:  "```shell\n$ git log```",
				Inline: false,
			},
			{
				Name:   "Lists version history for a file, including renames",
				Value:  "```shell\n$ git log --follow [file]```",
				Inline: false,
			},
			{
				Name:   "Shows content differences between two branches",
				Value:  "```shell\n$ git diff [first-branch]...[second-branch]```",
				Inline: false,
			},
			{
				Name:   "Outputs metadata and content changes of the specified commit",
				Value:  "```shell\n$ git show [commit]```",
				Inline: false,
			},
			{
				Name:   "Snapshots the file in preparation for versioning",
				Value:  "```shell\n$ git add [file]```",
				Inline: false,
			},
			{
				Name:   "Records file snapshots permanently in version history",
				Value:  "```shell\n$ git commit -m [descriptive message]```",
				Inline: false,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://gitforwindows.org/img/gwindows_logo.png",
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "Git Cheat Sheet - Make changes",
	}
	return embed
}