package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func MatchPassword(c *gin.Context) {

}

func GetLoginForm(c *gin.Context) {
	var form LoginForm
	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Bad credentials"})
	}
	if form.Username == "admin" && form.Password == "admin" {
		c.Redirect(http.StatusFound, "/home")
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Bad credentials"})
	}
}
