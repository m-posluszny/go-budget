package accounts

import (
	"net/http"

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

func GetAccountsView(creds *auth.Credentials) AccountsView {
	panel := panel.GetPanelView(creds, misc.Accounts)
	accs, offAccs, _ := GetAccounts(creds.Uid)
	return AccountsView{panel, accs, offAccs}
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

	c.HTML(http.StatusOK, "accounts.html", GetAccountsView(creds))

}
