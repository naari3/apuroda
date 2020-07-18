package controllers

import (
	"apuroda/stores"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// WelcomeController WelcomeController
type WelcomeController struct{}

// Welcome Welcome
func (w WelcomeController) Welcome(c *gin.Context) {
	session := sessions.Default(c)

	fmt.Println(session.Get("user_id").(string))
	id, err := uuid.Parse(session.Get("user_id").(string))
	if err == nil {
		user, err := stores.UserStore.GetByID(id)
		if err != nil {
			// panic(err)
		}
		fmt.Println(user)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"user": user,
		})
		return
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
