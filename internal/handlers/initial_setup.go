package handlers

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ize-302/dayrate-cli/internal/config"
)

func InitialSetup() {
	if _, err := os.Stat(config.DatabasePath); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir("../../internal/db/", 0700)
			if err != nil {
				log.Fatalf("Unable to create directory due to %s", err)
			}
			// create file with empty array
			var emptyArray []interface{}

			// Create a file to write to
			file, err := os.Create(config.DatabasePath)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			// Encode the empty array as JSON and write to the file
			encoder := json.NewEncoder(file)
			err = encoder.Encode(emptyArray)
			if err != nil {
				panic(err)
			}
		}
	}
}
