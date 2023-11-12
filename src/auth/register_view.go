package auth

import (
	"errors"
	"fmt"
	"net/http"

	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/db"
)

var isUsernameValid = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
var isPasswordValid = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

func validateLength(s string, min int, max int) bool {
	l := len(s)
	return min <= l && l <= max
}
func RegisterPage(c *gin.Context) {
	RenderRegister(c, "", http.StatusOK)
}

func RenderRegister(c *gin.Context, err string, status int) {
	c.HTML(status, "register.html", gin.H{"error": err})
}

func validateForm(c *gin.Context, form *RegisterForm) error {
	if err := c.ShouldBind(form); err != nil {
		fmt.Println(err)
		return errors.New("invalid input")
	}
	if !validateLength(form.Username, 4, 24) {
		return errors.New("username has to have between 4 and 24 characters")
	}
	if !validateLength(form.Password, 8, 64) {
		return errors.New("password has to have between 8 and 64 characters")
	}
	if !isUsernameValid(form.Username) {
		return errors.New("username can only contain alphanumeric characters")
	}
	if !isPasswordValid(form.Password) {
		return errors.New("password can only contain alphanumeric characters and special characters")
	}
	if form.Password != form.RePassword {
		return errors.New("passwords doesn't match")
	}
	return nil
}

func GetRegisterForm(c *gin.Context) {
	var form RegisterForm
	if err := validateForm(c, &form); err != nil {
		fmt.Println(err)
		RenderRegister(c, err.Error(), http.StatusBadRequest)
		return
	}
	dbx := db.GetDbWrite()
	creds, err := CreateUser(dbx, form.LoginForm.DbView())
	if err != nil {
		RenderRegister(c, "Username already taken", http.StatusForbidden)
		fmt.Println(err)
		return
	}
	CreateSession(c, creds.Uid)
	c.Redirect(http.StatusFound, "/panel")
}
