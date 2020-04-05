package handler

import (
	"encoding/json"
	"github.com/mohae/struct2csv"
	"lockdown/models"
	"log"
	"strconv"
	"time"
)

func jsonToCsv(values []models.CsvModel) [][]string {
	strings, e := struct2csv.New().Marshal(values)
	if e != nil {
		log.Fatal("struct2csv marshalling error")
		return nil
	}
	return strings
}

func stringToCsvModel(valueAsString string, timestampInNanos string) models.CsvModel {
	var details models.CsvModel
	err := json.Unmarshal([]byte(valueAsString), &details)
	applicationTime, timeerror := strconv.ParseInt(timestampInNanos, 10, 64)
	if timeerror != nil {
		log.Fatal("::stringToCsvModel Error in Parsing int", timeerror)
	}
	details.ApplicationDate = time.Unix(0, applicationTime).Format("2006-01-02")
	if err != nil {
		log.Fatal("error in unmarshalling", err)
	}
	return details
}
