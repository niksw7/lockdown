package controllers

import (
	"encoding/json"
	"github.com/mohae/struct2csv"
	"lockdown/models"
	"log"
)

func jsonToCsv(values []models.TraderDetails) [][]string {
	strings, e := struct2csv.New().Marshal(values)
	if e != nil {
		log.Fatal("struct2csv marshalling error")
		return nil
	}
	return strings
}

func stringToModel(valueAsString string) models.TraderDetails {
	var details models.TraderDetails
	err := json.Unmarshal([]byte(valueAsString), &details)
	if err != nil {
		log.Fatal("error in unmarshalling", err)
	}
	return details
}
