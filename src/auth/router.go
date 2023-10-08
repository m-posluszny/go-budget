package auth

import "github.com/gin-gonic/gin"

func SetRouters(r *gin.Engine) {
	r.GET("/login", LoginPage)
	r.POST("/login", GetLoginForm)
	r.POST("/register", GetRegisterForm)

}
