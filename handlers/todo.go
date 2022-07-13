package handlers

import (
	"djaeger-bot/commands/todo"
	"djaeger-bot/models"
	"djaeger-bot/utils"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func (h Handler) Todo(session *discordgo.Session, message *discordgo.MessageCreate) {
	var err error
	var embedMsg discordgo.MessageEmbed

	if session.State.User.ID == message.Author.ID {
		return
	}

	msgArr := strings.SplitN(message.Content, " ", 3)
	command := msgArr[1]

	if msgArr[0] != "$todo" {
		return
	}

	if command == "create" {
		tag := utils.GenerateTag(5)

		activity := models.Activity{
			Name:     msgArr[2],
			Reporter: message.Author.Username,
			Tag:      tag,
			GuildID:  message.GuildID,
			UserID:   message.Author.ID,
		}

		embedMsg, err = todo.Create(h.DB, activity)
	}

	if command == "list" {
		var activities []models.Activity
		embedMsg, err = todo.List(h.DB, activities)
	}

	if command == "update" {
		content := strings.SplitN(msgArr[2], " to ", 2)

		embedMsg, err = todo.Update(h.DB, content[0], content[1])
	}

	if command == "delete" {
		deleteActivity := models.Activity{}
		embedMsg, err = todo.Delete(h.DB, msgArr[2], deleteActivity)
	}

	if err != nil {
		embedCreateErrMsg := discordgo.MessageEmbed{Title: err.Error()}
		_, err := session.ChannelMessageSendEmbed(message.ChannelID, &embedCreateErrMsg)
		if err != nil {
			return
		}
		return
	}

	_, err = session.ChannelMessageSendEmbed(message.ChannelID, &embedMsg)
	if err != nil {
		return
	}
}
