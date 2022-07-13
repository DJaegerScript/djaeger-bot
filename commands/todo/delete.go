package todo

import (
	"djaeger-bot/models"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func Delete(db *gorm.DB, tag string, activity models.Activity) (embedMsg discordgo.MessageEmbed, err error) {
	if result := db.Where("Tag = ?", tag).Delete(&activity); result.Error != nil {
		fmt.Println(err.Error())
		err = fmt.Errorf("an error occurred when deleting activity")
		return
	}

	embedMsg = discordgo.MessageEmbed{Title: fmt.Sprintf("✅ Task [%s] deleted successfully!", tag), Description: "To see the listed task, use `$todo list`"}
	return
}
