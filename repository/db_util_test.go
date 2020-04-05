package repository

import (
	"github.com/tidwall/buntdb"
)

type DbUtil struct {
	*buntdb.DB
	values map[string]string
}

func loadDbUtil(db *buntdb.DB) DbUtil {
	return DbUtil{db, make(map[string]string)}
}

func (dbutil DbUtil) loadDbData() {
	dbutil.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			dbutil.values[key] = value
			return true
		})
		return err
	})
}
