package repository

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"lockdown/models"
	"testing"
)

func Test_toTraderDetailsDb(t *testing.T) {
	traderDetailsDb := models.TraderDetailsDb{
		DeliveryLocation: models.DeliveryLocation{
			Area: "Khana Galli",
			City: "Jabalpur"},
	}
	bytes, _ := json.Marshal(traderDetailsDb)
	model := toTraderDetailsDb(string(bytes))
	assert.Equal(t, traderDetailsDb, model)
}
