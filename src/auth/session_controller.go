package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/db"

	"github.com/gin-contrib/sessions"
)

const userKey = "uid"

func InitAuthSession(store sessions.Store) gin.HandlerFunc {
	return sessions.Sessions("auth_session", store)
}

func DeauthRedirect(c *gin.Context) {
	err := DeleteSession(c)
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect(http.StatusFound, "/login")
}

func AuthRequired(c *gin.Context) {

	dbx := db.GetDbRead()
	uid, err := GetUIDFromSession(c)
	_, userErr := GetUserFromUid(dbx, uid)
	if err != nil || userErr != nil {
		fmt.Println("deauth")
		DeauthRedirect(c)
		c.Abort()
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

func GetUIDFromSession(c *gin.Context) (string, error) {
	session := sessions.Default(c)
	uid := session.Get(userKey)
	if uid == nil {
		return "", errors.New("empty user uid")
	}
	return uid.(string), nil
}
