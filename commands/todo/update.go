package todo

import (
	"djaeger-bot/models"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func Update(db *gorm.DB, tag string, status string) (embedMsg discordgo.MessageEmbed, err error) {
	if result := db.Model(&models.Activity{}).Where("Tag = ?", tag).Update("status", status); result.Error != nil {
		fmt.Println(err.Error())
		err = fmt.Errorf("an error occurred when updating activity")
		return
	}

	embedMsg = discordgo.MessageEmbed{Title: fmt.Sprintf("✅ Task [%s] updated successfully!", tag), Description: "To see the listed task, use `$todo list`"}
	return
}
