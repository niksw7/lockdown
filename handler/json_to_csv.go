package handler

import (
	"github.com/mohae/struct2csv"
	"lockdown/models"
	"log"
)

func jsonToCsv(values []models.TraderDetailsDb) [][]string {
	strings, e := struct2csv.New().Marshal(values)
	if e != nil {
		log.Println("::jsonToCsv:: struct2csv marshalling error")
		return nil
	}
	return strings
}
