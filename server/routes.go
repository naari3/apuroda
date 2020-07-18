package server

import (
	"apuroda/controllers"

	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	r := gin.New()
	r.LoadHTMLGlob("templates/*.tmpl")

	ping := new(controllers.PingController)
	welcome := new(controllers.WelcomeController)

	r.GET("/", welcome.Welcome)
	r.GET("/ping", ping.Ping)

	return r
}
