package accounts

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/db"
	"github.com/m-posluszny/go-ynab/src/misc"
	"github.com/m-posluszny/go-ynab/src/panel"
)

type AccountsView struct {
	panel.PanelView
	Budget      []Account
	BudgetTypes []BudgetType
	Offbudget   []Account
}

func GetAccountsView(dbx *db.DBRead, panel panel.PanelView) AccountsView {
	q := DefaultQuery(panel.UserUid)
	q.offBudget = false
	onAccs, _ := GetAccountsFromUserUid(dbx, q)
	q.offBudget = false
	offAccs, _ := GetAccountsFromUserUid(dbx, q)
	return AccountsView{panel, onAccs, BudgetTypes, offAccs}
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
	var form AccountForm
	dbx := db.GetDbWrite()
	creds, err := auth.GetCredsFromSession(dbx, c)
	if err != nil {
		panic(err)
	}

	if err := validateForm(c, &form); err != nil {
		slog.Error("Validation Error", err)
		panel := panel.GetPanelView(creds, misc.Accounts, err.Error())
		c.HTML(http.StatusBadRequest, "accounts.html", GetAccountsView(dbx, panel))
		return
	}

	acc := form.DbView(creds)
	slog.Info("Account Info", acc)
	c.Redirect(http.StatusFound, "/panel/accounts/uid")

}
