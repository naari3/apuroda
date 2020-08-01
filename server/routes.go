package server

import (
	"apuroda/controllers"

	"github.com/gin-gonic/gin"
)

func applyRoutes(r *gin.Engine) *gin.Engine {
	r.LoadHTMLGlob("templates/*.tmpl")

	ping := new(controllers.PingController)
	// welcome := new(controllers.WelcomeController)

	// r.GET("/", welcome.Welcome)
	r.GET("/ping", ping.Ping)

	user := new(controllers.UserController)
	r.GET("/users", user.Index)
	r.GET("/new_user", user.New)
	r.POST("/users", user.Create)

	file := new(controllers.FileController)
	r.GET("/", file.Index)
	r.POST("/files", file.Create)
	r.GET("/new_file", file.New)
	r.GET("/files/:id", file.Show)
	r.GET("/files/:id/download", file.Download)

	return r
}
