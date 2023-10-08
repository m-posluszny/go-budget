package panel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/auth"
)

func SetRoutes(r *gin.Engine) {
	userPanel := r.Group("/panel")
	userPanel.Use(auth.AuthRequired)
	{
		userPanel.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusUnauthorized, "panel.html", gin.H{"error": "Bad credentials"})
		})

	}
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/panel")
	})

}
