package handlers

import (
	"fmt"
	"time"

	"github.com/ize-302/dayrate-cli/internal/utils"
)

func HandleListRatigs() {
	InitialSetup()
	ratings := GetExistingDbData()

	// reverse ratings
	for i, j := 0, len(ratings)-1; i < j; i, j = i+1, j-1 {
		ratings[i], ratings[j] = ratings[j], ratings[i]
	}

	fmt.Printf("Here are your ratings: \n \n")
	fmt.Println("| Date ------------------------  | Rating |")
	if len(ratings) > 0 {
		for _, rating := range ratings {
			parsedTime, _ := time.Parse(time.RFC3339, rating.Timestamp)
			fmt.Printf("| %s  | %s |\n", parsedTime.Format(time.RFC1123), utils.FormatRating(rating.Rating))
		}
	} else {
		fmt.Println("| You have no ratings yet                 |")
	}
	fmt.Println("|=========================================|")
	fmt.Printf("| Total ratings: %d  Average: %d/5          | \n", len(ratings), 0)
	fmt.Println("|-----------------------------------------|")
}
