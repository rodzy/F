package structs

import (
	"github.com/bwmarrin/discordgo"
)

var (
	//ID var
	ID string

	//Session var
	Session *discordgo.Session
)


type bot struct{
	ID string
	Session *discordgo.Session
}