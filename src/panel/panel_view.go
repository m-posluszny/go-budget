package panel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/db"
	"github.com/m-posluszny/go-ynab/src/misc"
)

const PanelHtml = "panel.html"

type PanelView struct {
	Username string
	UserUid  string
	Category misc.PanelCategory
	ErrMsg   string
}

func GetPanelView(creds *auth.Credentials, category misc.PanelCategory, errMsg string) PanelView {
	return PanelView{creds.Username, creds.Uid, category, errMsg}
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
	c.HTML(http.StatusOK, "panel.html", GetPanelView(creds, misc.Panel, ""))
}
