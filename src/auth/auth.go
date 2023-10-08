package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
)

const userKey = "user"

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get(userKey)
	if uid == nil {
		c.Redirect(http.StatusPermanentRedirect, "/login")
		return
	}
	c.Next()
}

func CreateSession(c *gin.Context, userId string) error {
	session := sessions.Default(c)
	session.Set(userKey, userId)
	if err := session.Save(); err != nil {
		return errors.New("cannot login")
	}
	return nil

}

func DeleteSession(c *gin.Context) error {
	session := sessions.Default(c)
	user := session.Get(userKey)
	if user == nil {
		return errors.New("session invalid")
	}
	session.Delete(userKey)
	if err := session.Save(); err != nil {
		return errors.New("cannot logout")
	}
	return nil

}

func GetUID(c *gin.Context) string {
	session := sessions.Default(c)
	rawUid := session.Get(userKey)
	if uid, ok := rawUid.(string); ok {
		return uid
	}
	return ""
}
