package accounts

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/db"
	"github.com/m-posluszny/go-ynab/src/misc"
	"github.com/m-posluszny/go-ynab/src/panel"
)

type AccountsView struct {
	panel.PanelView
	Selected    Account
	Budget      []Account
	BudgetTypes []BudgetType
	Offbudget   []Account
}

func GetAccountsView(dbx db.Queryable, panel panel.PanelView) AccountsView {
	q := DefaultQuery(panel.UserUid)

	q.offBudget = false
	onAccs, err := GetAccountsFromUserUid(dbx, q)

	q.offBudget = true
	offAccs, errOff := GetAccountsFromUserUid(dbx, q)

	errs := errors.Join(err, errOff)
	if errs != nil {
		slog.Error("GetAccountsView", "err", errs)
	}
	return AccountsView{panel, Account{}, onAccs, BudgetTypes, offAccs}
}

func RenderPanel(c *gin.Context) {
	dbx := db.GetDbRead()
	creds, err := auth.GetCredsFromSession(dbx, c)
	if err != nil {
		panic(err)
	}
	panel := panel.GetPanelView(creds, misc.Accounts, "")
	c.HTML(http.StatusOK, "accounts.html", GetAccountsView(dbx, panel))
}

func RenderPanelDelete(c *gin.Context) {
	dbx := db.GetDbRead()
	accUid := c.Param("uid")
	creds, err := auth.GetCredsFromSession(dbx, c)
	if err != nil {
		panic(err)
	}
	panel := panel.GetPanelView(creds, misc.Accounts, "")
	accs := GetAccountsView(dbx, panel)
	acc, err := GetAccountFromUid(dbx, accUid)
	if err != nil {
		panic(err)
	}
	accs.Selected = acc
	c.HTML(http.StatusOK, "accounts-delete.html", accs)
}

func PostDeleteAccount(c *gin.Context) {
	creds := panel.MustGetCreds(c)
	err := db.Transact(func(dbx *sqlx.Tx) error {
		accUid := c.Param("uid")
		err := DeleteAccount(dbx, accUid)
		if err != nil {
			c.Redirect(http.StatusFound, fmt.Sprintf("/panel/accounts/%s/delete", accUid))
			return err
		}
		c.Redirect(http.StatusFound, "/panel/accounts")
		return nil
	})
	if err != nil {
		slog.Error("Unhandled", "err", err)
		panel.RenderPanelWithErr(c, "Unknown error occured", creds)
	}
}

func validateForm(c *gin.Context, form *AccountForm) error {
	if err := c.ShouldBind(form); err != nil {
		slog.Error("Bind Error", err)
		return errors.New("invalid form")
	}
	if !misc.ValidateLength(form.Name, 4, 24) {
		return errors.New("account name has to have between 4 and 24 characters")
	}
	return nil
}

func PostCreateAccount(c *gin.Context) {
	creds := panel.MustGetCreds(c)
	err := db.Transact(func(dbx *sqlx.Tx) error {
		var form AccountForm
		var acc Account
		if err := validateForm(c, &form); err != nil {
			slog.Error("Validation Error", "msg", err)
			RenderAccountError(err.Error(), c, creds, dbx)
			return err
		}

		acc = form.DbView(creds)
		uid, err := CreateAccount(dbx, acc)
		if err != nil {
			RenderAccountError("Unknown error occured", c, creds, dbx)
			return err
		}
		c.Redirect(http.StatusFound, fmt.Sprintf("/panel/accounts/%s", uid))
		return nil

	})
	if err != nil {
		slog.Error("Create Account", "err", err)
		panel.RenderPanelWithErr(c, "Unknown error occured", creds)
	}

}

func RenderAccountError(errMsg string, c *gin.Context, creds *auth.Credentials, dbx db.Queryable) {
	panel := panel.GetPanelView(creds, misc.Accounts, errMsg)
	c.HTML(http.StatusBadRequest, "accounts.html", GetAccountsView(dbx, panel))
}
