package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WelcomeController WelcomeController
type WelcomeController struct{}

// Welcome Welcome
func (w WelcomeController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}
