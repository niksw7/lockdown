package repository

import (
	"encoding/json"
	"github.com/tidwall/buntdb"
	"lockdown/models"
	"log"
	"strconv"
)

type Repo interface {
	AddTraderRegistrationDetails(traderDetails models.TraderDetailsDbRequest, id string) error
	GenerateUniqueId() int
}
type BuntDbRepo struct {
	DB *buntdb.DB
}

func (repo BuntDbRepo) GenerateUniqueId() int {
	var generatedNumber int
	err := repo.DB.Update(func(tx *buntdb.Tx) error {
		err := tx.Descend("jsonIndex", func(key, value string) bool {
			generatedNumber, _ = strconv.Atoi(key)
			generatedNumber++
			return false
		})
		return err
	})
	if err != nil {
		log.Println("::GenerateUniqueId:: This is expected only once!! Generating number 1", err)
		generatedNumber = 1
	}
	return generatedNumber

}
func (repo BuntDbRepo) AddTraderRegistrationDetails(traderDetails models.TraderDetailsDbRequest, id string) error {
	db := repo.DB
	traderDetailsAsBytes, err := json.Marshal(traderDetails)
	if err != nil {
		log.Fatal("::AddTraderRegistrationDetails:: error while marshalling traderDetails", err)
		return err
	}
	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err = tx.Set(id, string(traderDetailsAsBytes), nil)
		return err
	})
	return err
}
