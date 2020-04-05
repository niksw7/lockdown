package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/buntdb"
	"io/ioutil"
	"lockdown/models"
	mockrepository "lockdown/repository"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestRegisterUserDetails(t *testing.T) {
	gin.SetMode(gin.TestMode)
	traderDetails := models.TraderDetailsRequest{
		Tehsil:           "VijayWada",
		DealerType:       "Retail",
		DeliveryLocation: "Jaipur",
		Mobile:           "89289211",
	}
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
	db, _ := buntdb.Open(":memory:")
	defer db.Close()
	context.Set("db", db)
	//insert into database
	db.Update(func(tx *buntdb.Tx) error {
		tx.Set(strconv.FormatInt(time.Now().UnixNano(), 10), buildTraderDetails(), nil)
		tx.Set(strconv.FormatInt(time.Now().UnixNano(), 10), buildTraderDetails(), nil)
		return nil
	})
	context.Request = httptest.NewRequest(http.MethodGet, "/download", nil)
	CsvDownloader()(context)
	content, _ := ioutil.ReadFile("test.csv")
	expectedContent := string(content)
	assert.Equal(t, expectedContent, responseRecorder.Body.String())

}

func buildTraderDetails() string {
	details := models.TraderDetailsRequest{
		Tehsil:           "ramaPura",
		DealerType:       "retail",
		DeliveryLocation: "muradabad",
		Mobile:           "976112233",
	}
	marshal, _ := json.Marshal(details)
	return string(marshal)
}
