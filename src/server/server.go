package server

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/config"
	"github.com/m-posluszny/go-ynab/src/menu"
)

func Init(cfg config.Config, store sessions.Store, templateDir string) *gin.Engine {
	gin.SetMode(cfg.Server.Mode)
	srv := gin.Default()

	srv.SetFuncMap(FuncMap)
	srv.LoadHTMLGlob(templateDir)
	srv.StaticFS("/static", http.Dir("./src/static"))

	authSess := auth.InitAuthSession(store)

	srv.Use(authSess)
	srv.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		auth.RenderLogin(c, "Unknown Server Error", http.StatusFound)
	}))
	loadRoutes(srv)
	return srv
}

func loadRoutes(r *gin.Engine) {
	auth.SetRouters(r)
	menu.SetRoutes(r)

}
