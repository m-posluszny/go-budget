package transactions

import "time"

type Transaction struct {
	Uid         string
	AccountUid  string `db:"account_uid"`
	Memo        string
	Payee       string
	Category    string
	CategoryUid string
	Date        time.Time
	Value       float64
}
