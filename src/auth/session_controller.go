package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/config"
	"github.com/m-posluszny/go-ynab/src/db"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

const userKey = "uid"

func InitAuthSession(authCfg config.AuthConf, redisConf config.RedisConf) gin.HandlerFunc {
	session, err := redis.NewStore(redisConf.Size, "tcp", redisConf.Host, redisConf.Password, []byte(authCfg.Secret))
	if err != nil {
		panic(err)
	}
	return sessions.Sessions("auth_session", session)
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
		DeauthRedirect(c)
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
