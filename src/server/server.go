package server

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/config"
	"github.com/m-posluszny/go-ynab/src/panel"
)

func Init(cfg config.Config, store sessions.Store, templateDir string) *gin.Engine {
	gin.SetMode(cfg.Server.Mode)
	server := gin.Default()

	server.LoadHTMLGlob(templateDir)
	server.StaticFS("/static", http.Dir("./src/static"))

	authSess := auth.InitAuthSession(store)

	server.Use(authSess)
	server.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		auth.RenderLogin(c, "Unknown Server Error", http.StatusFound)
	}))
	loadRoutes(server)
	return server
}

func loadRoutes(r *gin.Engine) {
	auth.SetRouters(r)
	panel.SetRoutes(r)

}
