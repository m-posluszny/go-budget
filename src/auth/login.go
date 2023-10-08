package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/db"
)

func LoginPage(c *gin.Context) {
	RenderLogin(c, "", http.StatusOK)
}

func LogoutAction(c *gin.Context) {
	err := DeleteSession(c)
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect(http.StatusFound, "/login")

}

func GetLoginForm(c *gin.Context) {
	var form LoginForm
	if err := c.ShouldBind(&form); err != nil {
		RenderLogin(c, "Bad credentials", http.StatusUnauthorized)
	}
	dbx := db.GetDbRead()
	user, err := GetUserFromName(dbx, form.Username)
	if err == nil && MatchPassword(dbx, form) {
		CreateSession(c, user.Uid)
		c.Redirect(http.StatusFound, "/panel")
		return
	} else {
		RenderLogin(c, "Bad credentials", http.StatusUnauthorized)
	}
}

func RenderLogin(c *gin.Context, err string, status int) {
	c.HTML(status, "login.html", gin.H{"error": err})
}
