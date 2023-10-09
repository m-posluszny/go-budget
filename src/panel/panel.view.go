package panel

import (
	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/db"
)

func RenderPanel(c *gin.Context, uid string, status int) {
	dbx := db.GetDbRead()
	creds, err := auth.GetUserFromUid(dbx, uid)
	if err != nil {
		panic(err)
	}
	c.HTML(status, "panel.html", gin.H{"username": creds.Username})
}
