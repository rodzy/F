package embeds

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

//BranchEmbed displays all the information about how manipulate branches
func BranchEmbed() *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Flankerbot",
			URL:     "https://github.com/rodzy",
			IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/16ea3509f9a41b475966b85565a930fa.webp?size=128",
		},
		Color:       0x66ccff,
		Description: "Branches are an important part of working with Git. Any commits you make will be made on the branch you're currently “checked out” to. Use ``git status`` to see which branch that is.",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Creates a new branch",
				Value:  "```shell\n$ git branch [branch-name]```",
				Inline: false,
			},
			{
				Name:   "Switches to the specified branch and updates the working directory",
				Value:  "```shell\n$ git checkout [branch-name]```",
				Inline: false,
			},
			{
				Name:   "Combines the specified branch’s history into the current branch. This is usually done in pull requests, but is an important Git operation.",
				Value:  "```shell\n$ git merge [branch]```",
				Inline: false,
			},
			{
				Name:   "Deletes the specified branch",
				Value:  "```shell\n$ git branch -d [branch-name]```",
				Inline: false,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://gitforwindows.org/img/gwindows_logo.png",
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "Git Cheat Sheet - Branches",
	}

	return embed
}
