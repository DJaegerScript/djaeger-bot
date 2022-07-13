package todo

import (
	"djaeger-bot/models"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, activity models.Activity) (embedMsg discordgo.MessageEmbed, err error) {
	if result := db.Create(&activity); result.Error != nil {
		err = fmt.Errorf("an error occurred when creating activity")
		return
	}

	embedMsg = discordgo.MessageEmbed{Title: fmt.Sprintf("✅ Task added successfully by %s!", activity.Reporter), Description: "To see the listed task, use `$todo list`"}
	return
}
