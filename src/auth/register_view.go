package auth

import (
	"errors"
	"net/http"

	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/db"
)

var isUsernameValid = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
var isPasswordValid = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func validateLength(s string, min int, max int) bool {
	l := len(s)
	return min <= l && l <= max
}

func validateForm(c *gin.Context, form *RegisterForm) error {
	if err := c.ShouldBind(form); err != nil {
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
		RenderLogin(c, err.Error(), http.StatusBadRequest)
		return
	}
	dbx := db.GetDbWrite()
	creds, err := CreateUser(dbx, form)
	if err != nil {
		RenderLogin(c, "Username already taken", http.StatusForbidden)
		return
	}
	CreateSession(c, creds.Uid)
	c.Redirect(http.StatusFound, "/panel")
}
