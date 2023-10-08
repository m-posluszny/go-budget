package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/config"
	"github.com/m-posluszny/go-ynab/src/panel"
)

func Init(cfg config.Config) {
	gin.SetMode(cfg.Server.Mode)
	server := gin.Default()

	server.LoadHTMLGlob("./src/templates/**/*")
	server.StaticFS("/static", http.Dir("./src/static"))
	authSess := auth.InitAuthSession(cfg.Auth, cfg.Redis)
	server.Use(authSess)
	server.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		auth.RenderLogin(c, "Unknown Server Error", http.StatusFound)
	}))
	loadRoutes(server)

	server.Run(cfg.Server.Host)
}

func loadRoutes(r *gin.Engine) {
	auth.SetRouters(r)
	panel.SetRoutes(r)

}
