package controllers

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"lockdown/models"
	"testing"
)

func Test_stringToModel(t *testing.T) {
	givenModel := models.TraderDetails{
		Tehsil:           "rajasthan",
		DealerType:       "sa",
		DeliveryLocation: "jabalpur",
		Mobile:           "sasa",
	}
	bytes, _ := json.Marshal(givenModel)
	model := stringToCsvModel(string(bytes), "1357895000000000000")
	expectedModel := models.CsvModel{
		Tehsil:           "rajasthan",
		DealerType:       "sa",
		DeliveryLocation: "jabalpur",
		Mobile:           "sasa",
		ApplicationDate:  "2013-01-11",
	}
	assert.Equal(t, expectedModel, model)
}

func Test_jsonToCsv(t *testing.T) {
	model1 := models.CsvModel{
		Tehsil:           "rajasthan",
		DealerType:       "sa",
		DeliveryLocation: "jabalpur",
		Mobile:           "sasa",
		ApplicationDate:  "2013-01-11",
	}
	model2 := models.CsvModel{
		Tehsil:           "kerala",
		DealerType:       "retailer",
		DeliveryLocation: "munar",
		Mobile:           "90898989",
		ApplicationDate:  "2013-01-11",
	}
	var models []models.CsvModel
	models = append(models, model1, model2)
	actualArray := jsonToCsv(models)
	cols := []string{"Tehsil", "DealerType", "DeliveryLocation", "Mobile", "ApplicationDate"}
	assert.Equal(t, cols, actualArray[0])
	assert.Equal(t, "rajasthan", actualArray[1][0])
	assert.Equal(t, "kerala", actualArray[2][0])
}
