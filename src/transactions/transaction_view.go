package transactions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/accounts"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/budgets"
	"github.com/m-posluszny/go-ynab/src/dates"
	"github.com/m-posluszny/go-ynab/src/db"
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

func GetTransactionView(creds *auth.Credentials, accUid string, d time.Time) TransactionView {
	account := accounts.GetAccountsView(creds)
	acc, _ := accounts.GetAccount(accUid)
	transacts, _ := GetTransactions(creds.Uid, accUid)
	payees, _ := GetPayees(creds.Uid)
	memos, _ := GetMemos(creds.Uid)
	categories, _ := budgets.GetCategories(creds.Uid)
	return TransactionView{account, acc, transacts, categories, memos, payees, dates.GetMonthSet(d)}
}

func RenderPanel(c *gin.Context) {
	accUid := c.Param("uid")
	var q TransactionQuery
	fmt.Print(c.Query("month"))
	c.Bind(&q)
	now := dates.MustDateFromString(q.Month)
	uid, err := auth.GetUIDFromSession(c)
	if err != nil {
		panic(err)
	}
	dbx := db.GetDbRead()
	creds, err := auth.GetUserFromUid(dbx, uid)
	if err != nil {
		panic(err)
	}

	c.HTML(http.StatusOK, "transactions.html", GetTransactionView(creds, accUid, now))

}
