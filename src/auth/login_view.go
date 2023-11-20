package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/db"
)

func LoginPage(c *gin.Context) {
	RenderLogin(c, "", http.StatusOK)
}

func GetLoginForm(c *gin.Context) {
	var form LoginForm
	if err := c.ShouldBind(&form); err != nil {
		RenderLogin(c, "Bad credentials", http.StatusUnauthorized)
		return
	}
	dbx := db.GetDbRead()
	user, err := GetUserFromName(dbx, form.Username)
	if err == nil && MustMatchPassword(dbx, form) {
		CreateSession(c, user.Uid)
		c.Redirect(http.StatusFound, "/panel")
		return
	} else {
		RenderLogin(c, "Bad credentials", http.StatusUnauthorized)
		return
	}
}

func RenderLogin(c *gin.Context, err string, status int) {
	c.HTML(status, "login.html", gin.H{"error": err})
}
