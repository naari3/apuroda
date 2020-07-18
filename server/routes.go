package server

import (
	"apuroda/controllers"

	"github.com/gin-gonic/gin"
)

func applyRoutes(r *gin.Engine) *gin.Engine {
	r.LoadHTMLGlob("templates/*.tmpl")

	ping := new(controllers.PingController)
	welcome := new(controllers.WelcomeController)

	r.GET("/", welcome.Welcome)
	r.GET("/ping", ping.Ping)

	user := new(controllers.UserController)
	r.GET("/users", user.Index)
	r.GET("/users/new", user.New)
	r.POST("/users", user.Create)

	return r
}
