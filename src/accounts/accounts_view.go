package accounts

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/db"
	"github.com/m-posluszny/go-ynab/src/misc"
	"github.com/m-posluszny/go-ynab/src/panel"
)

type AccountsView struct {
	panel.PanelView
	Budget    []Account
	Offbudget []Account
}

func GetAccountsView(dbx *db.DBRead, panel panel.PanelView) AccountsView {
	q := DefaultQuery(panel.UserUid)
	q.offBudget = false
	onAccs, _ := GetAccountsFromUserUid(dbx, q)
	q.offBudget = false
	offAccs, _ := GetAccountsFromUserUid(dbx, q)
	return AccountsView{panel, onAccs, offAccs}
}

func RenderPanel(c *gin.Context) {
	uid, err := auth.GetUIDFromSession(c)
	if err != nil {
		panic(err)
	}
	dbx := db.GetDbRead()
	creds, err := auth.GetUserFromUid(dbx, uid)
	if err != nil {
		panic(err)
	}
	panel := panel.GetPanelView(creds, misc.Accounts, "")
	c.HTML(http.StatusOK, "accounts.html", GetAccountsView(dbx, panel))
}

func validateForm(c *gin.Context, form *AccountForm) error {
	if err := c.ShouldBind(form); err != nil {
		return err
	}
	if !misc.ValidateLength(form.Name, 4, 24) {
		return errors.New("account name has to have between 4 and 24 characters")
	}
	return nil
}

func PostCreateAccount(c *gin.Context) {
	var form AccountForm
	uid, err := auth.GetUIDFromSession(c)
	if err != nil {
		panic(err)
	}
	dbx := db.GetDbWrite()
	_, err = auth.GetUserFromUid(dbx, uid)
	if err != nil {
		panic(err)
	}

	if err := validateForm(c, &form); err != nil {
		slog.Error("Validation Error")
		slog.Error(err.Error())
		RenderPanel(c)
		return
	}
	slog.Error(strconv.FormatBool(form.Offbudget.Bool()))

	c.Redirect(http.StatusFound, "/panel/accounts/uid")

}
