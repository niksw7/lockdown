package handler

import (
	"github.com/stretchr/testify/assert"
	"lockdown/models"
	"testing"
)

func Test_jsonToCsv(t *testing.T) {
	model1 := models.TraderDetailsDb{
		DeliveryLocation: models.DeliveryLocation{
			Area: "Chat Galli",
			City: "Jabalpur"},
	}
	model2 := models.TraderDetailsDb{
		DeliveryLocation: models.DeliveryLocation{
			Area: "Chai Galli",
			City: "Munar"},
	}
	var models []models.TraderDetailsDb
	models = append(models, model1, model2)
	actualArray := jsonToCsv(models)
	cols := []string{"Area", "City", "Name", "Address", "OwnerMobile", "Email", "Type", "HomeDeliveryNumber", "AgentName", "AgentAge", "AgentMobile", "VehicleType", "VehicleNumber", "RegistrationDate", "Id"}
	assert.Equal(t, cols, actualArray[0])
	assert.Equal(t, "Jabalpur", actualArray[1][1])
	assert.Equal(t, "Munar", actualArray[2][1])
}
