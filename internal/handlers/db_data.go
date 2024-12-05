package handlers

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ize-302/dayrate-cli/internal/config"
	"github.com/ize-302/dayrate-cli/internal/models"
)

func GetExistingDbData() []models.Rating {
	f, err := os.ReadFile(config.DatabasePath)
	if err != nil {
		log.Fatalf("Unable to read file due to %s\n", err)
	}

	var data []models.Rating

	err = json.Unmarshal(f, &data)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}
	return data
}
