package transactions

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/accounts"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/budgets"
	"github.com/m-posluszny/go-ynab/src/dates"
	"github.com/m-posluszny/go-ynab/src/db"
	"github.com/m-posluszny/go-ynab/src/misc"
	"github.com/m-posluszny/go-ynab/src/panel"
)

type TransactionView struct {
	accounts.AccountsView
	Account      accounts.Account
	Transactions []Transaction
	Categories   []budgets.Category
	Memos        []string
	Payees       []string
	Months       dates.MonthSet
}

type TransactionQuery struct {
	Month string `form:"month"`
}

func GetTransactionView(dbx *db.DBRead, panel panel.PanelView, accUid string, d time.Time) TransactionView {
	account := accounts.GetAccountsView(dbx, panel)
	acc, _ := accounts.GetAccountFromUid(dbx, accUid)
	transacts, _ := GetTransactions(panel.UserUid, accUid)
	payees, _ := GetPayees(panel.UserUid)
	memos, _ := GetMemos(panel.UserUid)
	categories, _ := budgets.GetCategories(panel.UserUid)
	return TransactionView{account, acc, transacts, categories, memos, payees, dates.GetMonthSet(d)}
}

func GetTransactionQuery(c *gin.Context) TransactionQuery {
	var q TransactionQuery
	c.Bind(&q)
	return q
}

func RenderPanel(c *gin.Context) {
	accUid := c.Param("uid")
	dbx := db.GetDbRead()
	q := GetTransactionQuery(c)
	creds, err := auth.GetCredsFromSession(dbx, c)
	panel := panel.GetPanelView(creds, misc.Transactions, "")
	if err != nil {
		panic(err)
	}
	now := dates.MustDateFromString(q.Month)
	c.HTML(http.StatusOK, "transactions.html", GetTransactionView(dbx, panel, accUid, now))

}
