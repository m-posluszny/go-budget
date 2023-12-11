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

func MustGetCreds(c *gin.Context) *auth.Credentials {
	dbx := db.GetDbRead()
	creds, err := auth.GetCredsFromSession(dbx, c)
	if err != nil {
		panic(err)
	}
	return creds
}

func RenderPanelWithErr(c *gin.Context, errMsg string, creds *auth.Credentials) {
	c.HTML(500, "panel.html", GetPanelView(creds, misc.Panel, errMsg))
}

func RenderPanel(c *gin.Context) {
	creds := MustGetCreds(c)
	dbx := db.GetDbRead()
	creds, err := auth.GetUserFromUid(dbx, creds.Uid)
	if err != nil {
		panic(err)
	}
	c.HTML(http.StatusOK, "panel.html", GetPanelView(creds, misc.Panel, ""))
}
