package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/ize-302/dayrate-cli/internal/config"
	"github.com/ize-302/dayrate-cli/internal/models"
)

func FormatRating(rating int) string {
	return fmt.Sprintf("%d/5 %s", rating, config.Emojis[rating-1])
}

func IsDayRated(ratings []models.Rating) bool {
	currentTime := time.Now().Format(time.RFC3339) // Get current date/time in ISO string format
	parsedCurrentTime, _ := time.Parse(time.RFC3339, currentTime)
	var result bool

	for i := range ratings {
		parsedDate, err := time.Parse(time.RFC3339, ratings[i].Timestamp)
		if err != nil {
			log.Fatalf("Error parsing time: %v\n", err)
		}

		if parsedDate.Truncate(time.Hour*24) == parsedCurrentTime.Truncate(time.Hour*24) {
			result = true
			break
		}
		result = false
	}
	return result
}

func GetHelp() {
	var message string
	message += fmt.Sprintln("Dayrate CLI - Rate your day")
	message += fmt.Sprintln("usage: dayrate [command]")
	message += fmt.Sprintln("commands:")
	message += fmt.Sprintln("   help    dayrate help")
	message += fmt.Sprintln("   add     Rate current day")
	message += fmt.Sprintln("   list    List ratings")
	fmt.Println(message)
}
