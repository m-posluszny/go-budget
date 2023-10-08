package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRegisterForm(c *gin.Context) {
	var form RegisterForm
	if err := c.ShouldBind(&form); err != nil {
		RenderLogin(c, "Bad credentials", http.StatusForbidden)
		return
	}
	if form.Password != form.RePassword {
		RenderLogin(c, "Passwords doesn't match", http.StatusBadRequest)
		return
	}
	creds, err := CreateUser(form)
	if err != nil {
		RenderLogin(c, "Bad credentials", http.StatusForbidden)
	}
	CreateSession(c, creds.Uid)
	c.Redirect(http.StatusFound, "/panel")
}
