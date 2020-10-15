package embeds

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

//SyncEmbed displays the Sync tab
func SyncEmbed() *discordgo.MessageEmbed  {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Flankerbot",
			URL:     "https://github.com/rodzy/F",
			IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/e70c1562f842ede26dbb80404b1206a5.webp?size=128",
		},
		Color:       0x66ccff,
		Description: "Synchronize your local repository with the remote repository on GitHub.com",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Downloads all history from the remote tracking branches",
				Value:  "```shell\n$ git fetch```",
				Inline: false,
			},
			{
				Name:   "Combines remote tracking branch into current local branch",
				Value:  "```shell\n$ git merge```",
				Inline: false,
			},
			{
				Name:   "Uploads all local branch commits to GitHub",
				Value:  "```shell\n$ git push```",
				Inline: false,
			},
			{
				Name:   "Updates your current local working branch with all new commits from the corresponding remote branch on GitHub. ``git pull`` is a combination of ``git fetch`` and ``git merge``",
				Value:  "```shell\n$ git pull```",
				Inline: false,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "Git Cheat Sheet - Synchronize changes",
	}
	return embed
}