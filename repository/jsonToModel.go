package repository

import (
	"encoding/json"
	"lockdown/models"
	"log"
)

func toTraderDetailsDb(str string) models.TraderDetailsDb {
	var details models.TraderDetailsDb
	err := json.Unmarshal([]byte(str), &details)
	if err != nil {
		log.Println("error in unmarshalling Safely returning nil", err)
	}
	return details
}
