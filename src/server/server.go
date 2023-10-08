package server

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/config"
	"github.com/m-posluszny/go-ynab/src/panel"
)

func Init(srvCfg config.ServerConf, authCfg config.AuthConf) {
	gin.SetMode(srvCfg.Mode)
	server := gin.Default()

	server.LoadHTMLGlob("./src/templates/**/*")
	server.StaticFS("/static", http.Dir("./src/static"))
	server.Use(sessions.Sessions("auth_session", cookie.NewStore([]byte(authCfg.Secret))))

	loadRoutes(server)

	server.Run(srvCfg.Host)
}

func loadRoutes(r *gin.Engine) {
	auth.SetRouters(r)
	panel.SetRoutes(r)

}
