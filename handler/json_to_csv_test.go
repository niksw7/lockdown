package handler

import (
	"github.com/stretchr/testify/assert"
	"lockdown/models"
	"testing"
)



func Test_jsonToCsv(t *testing.T) {
	model1 := models.TraderDetailsDb{
		Tehsil:           "rajasthan",
		DealerType:       "sa",
		DeliveryLocation: "jabalpur",
		Mobile:           "sasa",
	}
	model2 := models.TraderDetailsDb{
		Tehsil:           "kerala",
		DealerType:       "retailer",
		DeliveryLocation: "munar",
		Mobile:           "90898989",
	}
	var models []models.TraderDetailsDb
	models = append(models, model1, model2)
	actualArray := jsonToCsv(models)
	cols := []string{"Tehsil", "DealerType", "DeliveryLocation", "Mobile","RegistrationDate", "Id"}
	assert.Equal(t, cols, actualArray[0])
	assert.Equal(t, "rajasthan", actualArray[1][0])
	assert.Equal(t, "kerala", actualArray[2][0])
}
