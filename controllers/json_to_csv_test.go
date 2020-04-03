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
	model := stringToModel(string(bytes))
	assert.Equal(t, givenModel, model)
}

func Test_jsonToCsv(t *testing.T) {
	model1 := models.TraderDetails{
		Tehsil:           "rajasthan",
		DealerType:       "sa",
		DeliveryLocation: "jabalpur",
		Mobile:           "sasa",
	}
	model2 := models.TraderDetails{
		Tehsil:           "kerala",
		DealerType:       "retailer",
		DeliveryLocation: "munar",
		Mobile:           "90898989",
	}
	var models []models.TraderDetails
	models = append(models, model1, model2)
	actualArray := jsonToCsv(models)
	cols := []string{"Tehsil", "DealerType", "DeliveryLocation", "Mobile"}
	assert.Equal(t, cols, actualArray[0])
	assert.Equal(t, "rajasthan", actualArray[1][0])
	assert.Equal(t, "kerala", actualArray[2][0])
}
