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
	c.HTML(http.StatusOK, "user_new.tmpl", gin.H{})
}

// Create Create
func (u UserController) Create(c *gin.Context) {
	repo := repositories.UserRepository{}
	c.Request.ParseForm()
	user, err := repo.Create(c.Request.Form["name"][0])
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusTemporaryRedirect, "/users")
	}
	session := sessions.Default(c)
	session.Set("user_id", user.ID.String())
	session.Save()
	c.Redirect(http.StatusFound, "/")
}
