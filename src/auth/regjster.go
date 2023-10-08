package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterForm struct {
	Username   string `form:"username" binding:"required"`
	Password   string `form:"password" binding:"required"`
	RePassword string `form:"repassword" binding:"required"`
}

func GetRegisterForm(c *gin.Context) {
	var form RegisterForm
	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusForbidden, "login.html", gin.H{"error": "Bad credentials"})
		fmt.Println(err)
		return
	}
	if form.Password != form.RePassword {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Passwords doesn't match"})
		return
	}
	c.Redirect(http.StatusFound, "/home")
}
