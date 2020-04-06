package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/buntdb"
	"lockdown/models"
	"testing"
)

func TestBuntDbRepo_AddTraderRegistrationDetails(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := buntdb.Open(":memory:")
	defer db.Close()
	repo := BuntDbRepo{DB: db}
	traderDetails := models.TraderDetailsDb{
		Tehsil:           "Vijaywada",
		DealerType:       "Retail",
		DeliveryLocation: "Jaipur",
		Mobile:           "89289211",
		RegistrationDate: "2019-09-09",
	}
	err := repo.AddTraderRegistrationDetails(traderDetails, "12")
	err = repo.AddTraderRegistrationDetails(traderDetails, "13")
	assert.NoError(t, err)

	//assert in db
	var dbUtil = loadDbUtil(db)
	dbUtil.loadDbData()
	registeredDetails := dbUtil.values["12"]
	assert.NotNil(t, registeredDetails)
	assert.Equal(t, `{"Tehsil":"Vijaywada","DealerType":"Retail","DeliveryLocation":"Jaipur","Mobile":"89289211","RegistrationDate":"2019-09-09","Id":0}`, registeredDetails)

	registeredDetails = dbUtil.values["13"]
	assert.NotNil(t, registeredDetails)
	assert.Equal(t, `{"Tehsil":"Vijaywada","DealerType":"Retail","DeliveryLocation":"Jaipur","Mobile":"89289211","RegistrationDate":"2019-09-09","Id":0}`, registeredDetails)
}

func TestBuntDbRepo_GenerateUniqueId_InEmptyDB(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := buntdb.Open(":memory:")
	defer db.Close()
	repo := BuntDbRepo{DB: db}
	id := repo.GenerateUniqueId()
	assert.Equal(t, 1, id)
}

func TestBuntDbRepo_GenerateUniqueId(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := buntdb.Open(":memory:")
	db.CreateIndex("jsonIndex", "*", buntdb.IndexJSON("id"))
	defer db.Close()
	db.Update(func(tx *buntdb.Tx) error {
		tx.Set("101", `{"id":101}`, nil)
		tx.Set("105", `{"id":105}`, nil)
		tx.Set("103", `{"id":103}`, nil)
		tx.Set("107", `{"id":107}`, nil)

		tx.Set("104", `{"id":104}`, nil)
		return nil
	})
	db.View(func(tx *buntdb.Tx) error {
		tx.Descend("jsonIndex", func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			return false
		})
		return nil
	})
	//db.Update(func(tx *buntdb.Tx) error {
	//	tx.Set("101", "raja", nil)
	//	return nil
	//})
	repo := BuntDbRepo{DB: db}
	id := repo.GenerateUniqueId()
	assert.Equal(t, 108, id)
}

func TestGetAllTraderRegistrationDetails(t *testing.T) {

	//got, err := GetAllTraderRegistrationDetails()

}
