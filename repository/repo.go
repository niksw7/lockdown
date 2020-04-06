package repository

import (
	"encoding/json"
	"github.com/tidwall/buntdb"
	"lockdown/models"
	"log"
	"strconv"
)

type Repo interface {
	AddTraderRegistrationDetails(traderDetails models.TraderDetailsDb, id string) error
	GenerateUniqueId() int
	GetAllTraderRegistrationDetails() ([]models.TraderDetailsDb, error)
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

func (repo BuntDbRepo) AddTraderRegistrationDetails(traderDetails models.TraderDetailsDb, id string) error {
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

func (repo BuntDbRepo) GetAllTraderRegistrationDetails() ([]models.TraderDetailsDb, error) {
	var traderDetailArray []models.TraderDetailsDb
	err := repo.DB.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("jsonIndex", func(key, value string) bool {
			traderDetailArray = append(traderDetailArray, toTraderDetailsDb(value))
			return true
		})
		return err
	})
	if err != nil {
		log.Println("::GetAllTraderRegistrationDetails:: error in viewing data in db", err)
		return nil, err
	}
	return traderDetailArray, nil
}
