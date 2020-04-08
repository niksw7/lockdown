package handler

import (
	"github.com/stretchr/testify/assert"
	"lockdown/models"
	"testing"
)

func Test_jsonToCsv(t *testing.T) {
	model1 := models.TraderDetailsDb{
		City:           "rajasthan",
		DealerType:       "sa",
		DeliveryLocation: "jabalpur",
		Mobile:           "sasa",
	}
	model2 := models.TraderDetailsDb{
		City:           "kerala",
		DealerType:       "retailer",
		DeliveryLocation: "munar",
		Mobile:           "90898989",
	}
	var models []models.TraderDetailsDb
	models = append(models, model1, model2)
	actualArray := jsonToCsv(models)
	cols := []string{"City", "DealerType", "DeliveryLocation", "Mobile", "ShopName", "ShopAddress", "PhoneNumber", "Email", "ShopType", "HomeDeliveryNumber", "AgentName", "AgentAge", "AgentMobile", "Type", "Number", "RegistrationDate", "Id"}
	assert.Equal(t, cols, actualArray[0])
	assert.Equal(t, "rajasthan", actualArray[1][0])
	assert.Equal(t, "kerala", actualArray[2][0])
}
