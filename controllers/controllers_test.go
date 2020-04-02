package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/buntdb"
	"lockdown/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterUserDetails(t *testing.T) {
	gin.SetMode(gin.TestMode)
	traderDetails := models.TraderDetails{
		Tehsil:           "VijayWada",
		DealerType:       "Retail",
		DeliveryLocation: "Jaipur",
		Mobile:           "89289211",
	}
	responseRecorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(responseRecorder)
	db, _ := buntdb.Open(":memory:")
	defer db.Close()
	context.Set("db", db)
	traderDetailsAsString := contextForRegisterDetailsRequest(traderDetails, context)
	RegisterUserDetails(context)
	//Assert Response
	assert.Equal(t, traderDetailsAsString, responseRecorder.Body.String())
	//Assert Database
	db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			assert.Equal(t, traderDetailsAsString, value)
			return true
		})
		return err
	})
}

func contextForRegisterDetailsRequest(traderDetails models.TraderDetails, context *gin.Context) string {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(traderDetails)
	traderDetailsAsString := string(bytes.Trim(b.Bytes(), "\n"))
	context.Request = httptest.NewRequest(http.MethodPost, "/register-details", b)
	return traderDetailsAsString
}
