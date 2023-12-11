package accounts

import (
	"time"

	"github.com/m-posluszny/go-ynab/src/auth"
)

type CheckboxToggle string

func (c CheckboxToggle) Bool() bool {
	return string(c) == "on"
}

type Account struct {
	Uid          string
	UserUid      string `db:"user_uid"`
	Name         string
	Balance      float64
	Offbudget    bool
	CreationDate time.Time `db:"creation_date"`
}

type AccountForm struct {
	Name      string         `form:"name" binding:"required" url:"name"`
	Initial   float64        `form:"initial"  url:"initial"`
	Offbudget CheckboxToggle `form:"budget-type"  url:"budget-type" json:"budget-type"`
}

func (form AccountForm) DbView(creds *auth.Credentials) Account {
	return Account{Name: form.Name, Offbudget: form.Offbudget.Bool(), Balance: form.Initial, UserUid: creds.Uid}
}

type AccountsQuery struct {
	userUid   string
	offBudget bool
}

func DefaultQuery(uid string) AccountsQuery {
	return AccountsQuery{uid, false}
}
