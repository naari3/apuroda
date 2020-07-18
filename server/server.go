package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Start Start
func Start() {
	r := newRouter()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Run()
}
