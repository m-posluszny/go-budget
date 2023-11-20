package accounts

import (
	"github.com/m-posluszny/go-ynab/src/auth"
)

type Account struct {
	Uid        string
	UserUid    string `db:"user_uid"`
	Name       string
	Balance    float64
	BudgetType BudgetType
}

type AccountForm struct {
	Name       string     `form:"name" binding:"required" url:"name"`
	Initial    float64    `form:"initial" binding:"required" url:"initial"`
	BudgetType BudgetType `form:"budget-type"  url:"budget-type" json:"budget-type"`
}

func (form AccountForm) DbView(creds auth.Credentials) Account {
	return Account{Name: form.Name, BudgetType: form.BudgetType, Balance: form.Initial, UserUid: creds.Uid}
}

type AccountsQuery struct {
	userUid   string
	offBudget bool
}

func DefaultQuery(uid string) AccountsQuery {
	return AccountsQuery{uid, false}
}
