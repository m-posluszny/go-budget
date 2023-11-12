package transactions

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/accounts"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/db"
	"github.com/m-posluszny/go-ynab/src/misc"
)

func RenderPanel(c *gin.Context) {
	accUid := c.Param("uid")
	uid, err := auth.GetUIDFromSession(c)
	if err != nil {
		panic(err)
	}
	dbx := db.GetDbRead()
	creds, err := auth.GetUserFromUid(dbx, uid)
	if err != nil {
		panic(err)
	}
	accs, offAccs, _ := accounts.GetAccounts(uid)
	var transacts []Transaction
	acc := accounts.Account{Uid: accUid, Name: "Afdasfasd", Balance: 30000, Offbudget: false}
	for i := 0; i < 100; i++ {
		val := -1000 + rand.Float64()*(2000)
		transacts = append(transacts, Transaction{"uid", "Memo", "Payee", "Groceries", "GroUID", time.Now(), val})
	}
	c.HTML(http.StatusOK, "transactions.html", gin.H{"username": creds.Username, "category": misc.Accounts, "transactions": transacts, "account": acc, "accountsBudget": accs, "accountsOffbudget": offAccs})
}
