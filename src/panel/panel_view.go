package panel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/db"
	"github.com/m-posluszny/go-ynab/src/misc"
)

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
	c.HTML(http.StatusOK, "panel.html", gin.H{"username": creds.Username, "category": misc.Panel})
}
