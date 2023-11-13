package panel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/db"
	"github.com/m-posluszny/go-ynab/src/misc"
)

type PanelView struct {
	Username string
	Category misc.PanelCategory
}

func GetPanelView(creds *auth.Credentials, category misc.PanelCategory) PanelView {
	return PanelView{creds.Username, category}
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
	c.HTML(http.StatusOK, "panel.html", GetPanelView(creds, misc.Panel))
}
