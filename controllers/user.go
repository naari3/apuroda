package controllers

import (
	"apuroda/repositories"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserController UserController
type UserController struct{}

// Index Index
func (u UserController) Index(c *gin.Context) {
	repo := repositories.UserRepository{}
	users := repo.All()
	c.HTML(http.StatusOK, "user_index.tmpl", gin.H{
		"users": users,
	})
}

// New New
func (u UserController) New(c *gin.Context) {
	next := c.Query("next")
	if next == "" {
		next = "/"
	}
	c.HTML(http.StatusOK, "user_new.tmpl", gin.H{"next": next})
}

// Create Create
func (u UserController) Create(c *gin.Context) {
	repo := repositories.UserRepository{}
	user, err := repo.Create(c.PostForm("name"))
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusTemporaryRedirect, "/users")
	}
	session := sessions.Default(c)
	session.Set("user_id", user.ID.String())
	session.Save()

	next := c.PostForm("next")
	c.Redirect(http.StatusFound, next)
}
