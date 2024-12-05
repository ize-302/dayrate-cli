package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ize-302/dayrate-cli/internal/config"
	"github.com/ize-302/dayrate-cli/internal/models"
	"github.com/ize-302/dayrate-cli/internal/utils"
	"github.com/manifoldco/promptui"
)

func HandleRatingEntry() {
	InitialSetup()

	file, err := os.ReadFile(config.DatabasePath)
	if err != nil {
		log.Fatalf("Unable to read file due to %s", err)
	}

	currentTime := time.Now().Format(time.RFC3339) // Get current date/time in ISO string format

	var data []models.Rating
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalf("Unable to unmarshal JSON due to %s", err)
	}

	if utils.IsDayRated(data) {
		fmt.Println("Oops! You already logged today")
	} else {
		emojisOptions := []models.RatingOption{}
		for i, emoji := range config.Emojis {
			emojisOptions = append(emojisOptions, models.RatingOption{Name: fmt.Sprintf("%s %d/5", emoji, i+1), Value: i + 1})
		}
		templates := &promptui.SelectTemplates{
			Label:    "{{ .Name }}",
			Active:   "⏵ {{ .Name | yellow }}",
			Inactive: "   {{ .Name }}",
			Selected: "💡 You are rating today: {{ .Value }}/5",
		}
		prompt := promptui.Select{
			Label:     "Select Rating",
			Items:     emojisOptions,
			Templates: templates,
		}
		i, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt exited %v\n", err)
			return
		}

		// Proceed?
		proceedPrompt := promptui.Prompt{
			Label:     "Proceed and save rating?",
			IsConfirm: true,
		}
		result, err := proceedPrompt.Run()

		if err != nil {
			fmt.Printf("Canceled %v\n", err)
			return
		}
		fmt.Printf("You choose %q\n", result)

		todayRating := models.Rating{
			Timestamp: currentTime,
			Rating:    emojisOptions[i].Value,
		}

		ratings := GetExistingDbData()
		ratings = append(ratings, todayRating)

		jsonData, err := json.Marshal(ratings)

		err = os.WriteFile(config.DatabasePath, []byte(jsonData), 0666)
		if err != nil {
			log.Fatalf("Unable to write to file due to %s", err)
		}
		fmt.Println("✅ Rating saved!!")
	}
}
