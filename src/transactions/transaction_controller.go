package transactions

import (
	"math/rand"
	"strconv"
	"time"
)

type Transaction struct {
	Uid         string
	Memo        string
	Payee       string
	Category    string
	CategoryUid string
	Date        time.Time
	Value       float64
}

func GetTransactions(userUid string, accUid string) ([]Transaction, error) {
	var transacts []Transaction
	for i := 0; i < 100; i++ {
		val := -1000 + rand.Float64()*(2000)
		transacts = append(transacts, Transaction{strconv.Itoa(i), "Memo", "Payee", "Groceries", "GroUID", time.Now(), val})
	}
	return transacts, nil

}

func GetPayees(_userUid string) ([]string, error) {
	return []string{"Auchan", "Kwiaciarnia", "Butcher", "Online"}, nil
}

func GetMemos(_userUid string) ([]string, error) {
	return []string{"Groceries", "Rig", "Bookshelf", "Books"}, nil
}
