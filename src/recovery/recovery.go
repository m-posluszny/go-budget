package recovery

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/db"
	"github.com/m-posluszny/go-ynab/src/panel"
)

func Recover(c *gin.Context, err any) {
	dbx := db.GetDbRead()
	slog.Info("gin", "c", c)
	creds, authErr := auth.GetCredsFromSession(dbx, c)
	if authErr != nil {
		auth.RenderLogin(c, "Unknown Server Error", http.StatusFound)
		return
	}
	panel.RenderPanelWithErr(c, "Unknown Server Error", creds)

}
