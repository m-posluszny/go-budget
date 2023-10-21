package stores

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/m-posluszny/go-ynab/src/config"
)

func GetMockSessionStore(conf config.AuthConf) sessions.Store {
	return cookie.NewStore([]byte(conf.Secret))

}
