package embeds

import (
	"time"
	
	"github.com/bwmarrin/discordgo"
)

//HelpEmbed displays the embed on command
func HelpEmbed() *discordgo.MessageEmbed  {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Flankerbot",
			URL:     "https://github.com/rodzy",
			IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/16ea3509f9a41b475966b85565a930fa.webp?size=128",
		},
		Color:       0x66ccff,
		Description: "Git is the open source distributed version control system that facilitates GitHub activities on your laptop or desktop.\n To start just write: ``$<⌘>``",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Install",
				Value:  "```html\n⌘ -> $install```",
				Inline: false,
			},
			{
				Name:   "Create repositories",
				Value:  "```html\n⌘ -> $create```",
				Inline: false,
			},
			{
				Name:   "Configure tooling",
				Value:  "```html\n⌘ -> $config```",
				Inline: false,
			},
			{
				Name:   "The .gitignore file",
				Value:  "```html\n⌘ -> $ignore```",
				Inline: false,
			},
			{
				Name:   "Branches",
				Value:  "```html\n⌘ -> $branch```",
				Inline: false,
			},
			{
				Name:   "Make changes",
				Value:  "```html\n⌘ -> $changes```",
				Inline: false,
			},
			{
				Name:   "Synchronize changes",
				Value:  "```html\n⌘ -> $sync```",
				Inline: false,
			},
			{
				Name:   "Redo commits",
				Value:  "```html\n⌘ -> $redo```",
				Inline: false,
			},
			{
				Name:   "Git Information",
				Value:  "```html\n⌘ -> $info```",
				Inline: false,
			},
			{
				Name:   "Tutorial",
				Value:  "```html\n⌘ -> $tuto```",
				Inline: false,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://gitforwindows.org/img/gwindows_logo.png",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Want to learn more about using GitHub and Git? Email the Training Team or visit the web site for learning event schedules and private class availability.\nservices@github.com\nhttps://services.github.com/",
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "Flankerbot - Git Commands",
	}
	
	return embed
}

