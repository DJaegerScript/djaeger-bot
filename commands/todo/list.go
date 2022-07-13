package todo

import (
	"djaeger-bot/models"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
	"strings"
)

func List(db *gorm.DB, activities []models.Activity) (embedMsg discordgo.MessageEmbed, err error) {
	if result := db.Find(&activities); result.Error != nil {
		err = fmt.Errorf("an error occurred when retrieving activities")
		return
	}

	var activityList []string

	for _, activity := range activities {
		status := "⚪️️"
		if activity.Status == "in progress" {
			status = "🟡"
		} else if activity.Status == "done" {
			status = "🟢"
		}

		activityItem := fmt.Sprintf("%s **[%s] - %s** (%s)", status, activity.Tag, activity.Name, activity.Reporter)
		activityList = append(activityList, activityItem)
	}

	description := fmt.Sprintf("To update activity status, use `$todo update`\n\n%s", strings.Join(activityList, "\n\n"))

	embedMsg = discordgo.MessageEmbed{Title: "📝 Task", Description: description}
	return
}
