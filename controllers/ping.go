package controllers

import (
	"github.com/gin-gonic/gin"
)

// PingController PingController
type PingController struct{}

// Ping Ping
func (p PingController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
