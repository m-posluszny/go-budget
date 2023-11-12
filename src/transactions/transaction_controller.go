package transactions

import "time"

type Transaction struct {
	Uid         string
	Memo        string
	Payee       string
	Category    string
	CategoryUid string
	Date        time.Time
	Value       float64
}
