package panel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/accounts"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/transactions"
)

func SetRoutes(r *gin.Engine) {
	userPanel := r.Group("/panel")
	userPanel.Use(auth.AuthRequired)
	{
		userPanel.GET("/", RenderPanel)
		userPanel.GET("/accounts", accounts.RenderPanel)
		userPanel.GET("/accounts/:uid", transactions.RenderPanel)
	}
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/panel")
	})

}
