package embeds

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

//RedoEmbed displays the redo tab
func RedoEmbed() *discordgo.MessageEmbed  {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Flankerbot",
			URL:     "https://github.com/rodzy/F",
			IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/e70c1562f842ede26dbb80404b1206a5.webp?size=128",
		},
		Color:       0x66ccff,
		Description: "Erase mistakes and craft replacement history",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Undoes all commits after [commit], preserving changes locally",
				Value:  "```shell\n$ git reset [commit]```",
				Inline: false,
			},
			{
				Name:   "Discards all history and changes back to the specified commit",
				Value:  "```shell\n$ git reset --hard [commit]```",
				Inline: false,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://gitforwindows.org/img/gwindows_logo.png",
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "Git Cheat Sheet - Redo commits",
		Footer: &discordgo.MessageEmbedFooter{
			Text: "CAUTION! Changing history can have nasty side effects. If you need to change commits that exist on GitHub (the remote), proceed with caution. If you need help, reach out at github.community or contact support.",
		},
	}
	return embed
}