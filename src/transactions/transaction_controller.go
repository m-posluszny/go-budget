package transactions

import (
	"github.com/m-posluszny/go-ynab/src/db"
)

func GetTransactions(userUid string, accUid string) ([]Transaction, error) {
	var transacts []Transaction
	dbx := db.GetDbRead()
	dbx.Select(&transacts, "SELECT * FROM transactions WHERE user_uid = $1 AND account_uid = $2", userUid, accUid)
	return transacts, nil

}

func GetPayees(_userUid string) ([]string, error) {
	return []string{"Auchan", "Kwiaciarnia", "Butcher", "Online"}, nil
}

func GetMemos(_userUid string) ([]string, error) {
	return []string{"Groceries", "Rig", "Bookshelf", "Books"}, nil
}
