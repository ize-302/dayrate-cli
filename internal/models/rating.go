package models

type (
	Rating struct {
		Timestamp string `json:"timestamp"`
		Rating    int    `json:"rating"`
	}
)
