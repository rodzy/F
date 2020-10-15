package embeds

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

//HelpEmbed displays the embed on command
func HelpEmbed() *discordgo.MessageEmbed {
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
				Inline: true,
			},
			{
				Name:   "Create repositories",
				Value:  "```html\n⌘ -> $create```",
				Inline: true,
			},
			{
				Name:   "Configure tooling",
				Value:  "```html\n⌘ -> $config```",
				Inline: true,
			},
			{
				Name:   "The .gitignore file",
				Value:  "```html\n⌘ -> $ignore```",
				Inline: true,
			},
			{
				Name:   "Branches",
				Value:  "```html\n⌘ -> $branch```",
				Inline: true,
			},
			{
				Name:   "Make changes",
				Value:  "```html\n⌘ -> $changes```",
				Inline: true,
			},
			{
				Name:   "Synchronize changes",
				Value:  "```html\n⌘ -> $sync```",
				Inline: true,
			},
			{
				Name:   "Redo commits",
				Value:  "```html\n⌘ -> $redo```",
				Inline: true,
			},
			{
				Name:   "Git Information",
				Value:  "```html\n⌘ -> $info```",
				Inline: true,
			},
			{
				Name:   "Tutorial",
				Value:  "```html\n⌘ -> $tuto```",
				Inline: false,
			},
		},
		// Thumbnail: &discordgo.MessageEmbedThumbnail{
		// 	URL: "https://gitforwindows.org/img/gwindows_logo.png",
		// },
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Learn more about using GitHub and Git: \nservices@github.com\nhttps://services.github.com/",
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "Flankerbot - Git Commands",
	}

	return embed
}
