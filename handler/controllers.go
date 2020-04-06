package handler

import (
	"bytes"
	"encoding/csv"
	"github.com/tidwall/buntdb"
	"lockdown/repository"

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

func HealthChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		mood := "high"
		c.JSON(200, gin.H{
			"mood": mood,
		})
	}
}

func UserDetailsRegistrar(repo repository.Repo) gin.HandlerFunc {
	return func(context *gin.Context) {
		var traderDetails models.TraderDetailsDb
		err := context.BindJSON(&traderDetails)
		if err != nil {
			log.Println("error in json binding user-data", err)
			return
		}
		traderDetails.RegistrationDate = time.Now().Format(time.RFC850)
		traderDetails.Id = repo.GenerateUniqueId()

		err = repo.AddTraderRegistrationDetails(traderDetails, strconv.Itoa(traderDetails.Id))
		if err != nil {
			log.Println("::UserDetailsRegistrar:: db save failed", err)
			context.AbortWithError(500, err)
			return
		}
		context.JSON(http.StatusOK, traderDetails)

	}
}

func UserDetailsReader() gin.HandlerFunc {
	return func(context *gin.Context) {
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
}

func CsvDownloader(repo repository.Repo) gin.HandlerFunc {
	return func(context *gin.Context) {
		b := &bytes.Buffer{}
		wr := csv.NewWriter(b)
		//ideally csb write should be passed to db reader here
		traderDetailArray, err := repo.GetAllTraderRegistrationDetails()
		if err != nil {
			context.AbortWithError(500, err)
			return
		}
		arrayOfStrings := jsonToCsv(traderDetailArray)
		err = wr.WriteAll(arrayOfStrings)
		if err != nil {
			log.Println("::CsvDownloader:: writeall failed", err)
		}
		wr.Flush()
		context.Header("Content-Description", "city_requests")
		context.Header("Content-Disposition", "attachment; filename=city_requests.csv")
		context.Data(http.StatusOK, "text/csv", b.Bytes())
	}
}
