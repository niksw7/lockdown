package controllers

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"github.com/tidwall/buntdb"
	"strconv"
	"time"

	//"encoding/json"
	"github.com/gin-gonic/gin"
	"lockdown/models"
	"log"
	"net/http"
	//"strconv"
	//"time"
)

func Health(c *gin.Context) {
	mood := "high"
	c.JSON(200, gin.H{
		"mood": mood,
	})
}

func RegisterUserDetails(context *gin.Context) {
	var traderDetails models.TraderDetails
	var db = context.MustGet("db").(*buntdb.DB)
	err := context.BindJSON(&traderDetails)

	if err != nil {
		log.Fatal("error in json binding user-data", err)
		return
	}
	traderDetailsAsBytes, jsonError := json.Marshal(traderDetails)
	if jsonError != nil {
		log.Fatal("jsonError formatting", err)
		return
	}

	db.Update(func(tx *buntdb.Tx) error {
		tx.Set(strconv.FormatInt(time.Now().UnixNano(), 10), string(traderDetailsAsBytes), nil)
		return nil
	})
	context.JSON(http.StatusOK, traderDetails)

}

func ReadUserRegisteredDetails(context *gin.Context) {
	var db = context.MustGet("db").(*buntdb.DB)
	var arr []string
	db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			arr = append(arr, value)
			log.Println(key, value)
			return true
		})
		return err

	})
	context.JSON(http.StatusOK, gin.H{"len": arr})
}

func DownloadCsv(context *gin.Context) {
	var db = context.MustGet("db").(*buntdb.DB)
	var traderDetailArray []models.CsvModel
	db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			traderDetailArray = append(traderDetailArray, stringToCsvModel(value, key))
			return true
		})
		return err
	})
	arrayOfStrings := jsonToCsv(traderDetailArray)
	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)
	wr.WriteAll(arrayOfStrings)
	wr.Flush()
	context.Header("Content-Description", "city_requests")
	context.Header("Content-Disposition", "attachment; filename=city_requests.csv")
	context.Data(http.StatusOK, "text/csv", b.Bytes())
}
