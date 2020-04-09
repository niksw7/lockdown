package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"lockdown/models"
	mockrepository "lockdown/repository"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterUserDetails(t *testing.T) {
	gin.SetMode(gin.TestMode)
	traderDetails := sampleTraderDetails()
	marshal, _ := json.Marshal(sampleTraderDetails())
	fmt.Println(string(marshal))
	responseRecorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(responseRecorder)
	contextForRegisterDetailsRequest(traderDetails, context)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mockrepository.NewMockRepo(ctrl)
	m.EXPECT().GenerateUniqueId().Return(1).Times(1)
	m.EXPECT().AddTraderRegistrationDetails(gomock.Any(), "1").Return(nil).Times(1)
	UserDetailsRegistrar(m)(context)

	//Don't care much about response
}

func contextForRegisterDetailsRequest(traderDetails models.TraderDetailsRequest, context *gin.Context) string {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(traderDetails)
	traderDetailsAsString := string(bytes.Trim(b.Bytes(), "\n"))
	context.Request = httptest.NewRequest(http.MethodPost, "/register-details", b)
	return traderDetailsAsString
}

func TestDownloadCsv(t *testing.T) {
	gin.SetMode(gin.TestMode)
	responseRecorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(responseRecorder)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mockrepository.NewMockRepo(ctrl)
	dbs := []models.TraderDetailsDb{{
		DeliveryLocation: models.DeliveryLocation{
			Area: "Khana Galli",
			City: "Jaipur"},
		RegistrationDate: "2019",
		Id:               0,
	}}
	m.EXPECT().GetAllTraderRegistrationDetails().Return(dbs, nil).Times(1)
	CsvDownloader(m)(context)
	content, _ := ioutil.ReadFile("test.csv")
	expectedContent := string(content)
	assert.Equal(t, expectedContent, responseRecorder.Body.String())

}

func buildTraderDetails() string {
	details := models.TraderDetailsRequest{
		DeliveryLocation: models.DeliveryLocation{
			Area: "Peena Galli",
			City: "Murdabad"},
	}
	marshal, _ := json.Marshal(details)
	return string(marshal)
}

func sampleTraderDetails() models.TraderDetailsRequest {
	return models.TraderDetailsRequest{
		DeliveryLocation: models.DeliveryLocation{
			Area: "Khana Galli",
			City: "Jaipur"},
		ShopDetails: models.ShopDetails{
			Name:        "Ramlal mitaiwaala",
			Address:     "RustomJee Area, Kalakand",
			OwnerMobile: "90881910",
			Email:       "jackson@gmail.com",
			Type:        "Retail",
		},
		HomeDeliveryInfo: models.HomeDeliveryInfo{
			HomeDeliveryNumber: "98001010101",
			AgentInfo: models.AgentInfo{
				AgentName:   "Ramchandani",
				AgentAge:    45,
				AgentMobile: "99092029292",
			},
			VehicleInfo: models.VehicleInfo{
				VehicleType:   "Car",
				VehicleNumber: "MH091111",
			},
		},
	}
}
