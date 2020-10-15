package embeds

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

//InfoEmbed displays the info tab
func InfoEmbed() *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Flankerbot",
			URL:     "https://github.com/rodzy/F",
			IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/e70c1562f842ede26dbb80404b1206a5.webp?size=128",
		},
		Color:       0x66ccff,
		Description: "Some intresting facts",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Sets the name you want attached to your commit transactions",
				Value:  "**git:** an open source, distributed version-control system\n\n **GitHub:** a platform for hosting and collaborating on Git repositories\n\n **commit:** a Git object, a snapshot of your entire repository compressed into a SHA\n\n **branch:** a lightweight movable pointer to a commit\n\n **clone:** a local version of a repository, including all commits and branches\n\n **remote:** a common repository on GitHub that all team member use to exchange their changes\n\n **fork:** a copy of a repository on GitHub owned by a different user\n\n **pull request:** a place to compare and discuss the differences introduced on a branch with reviews, comments, integrated tests, and more\n\n **HEAD:** representing your current working directory, the HEAD pointer can be moved to different branches, tags, or commits when using ``git checkout``",
				Inline: false,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://gitforwindows.org/img/gwindows_logo.png",
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://arccwiki.uwyo.edu/images/1/19/GitHub_Flow_steps.png",
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "Git Cheat Sheet - Information",
	}
	return embed
}
