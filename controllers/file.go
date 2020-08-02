package controllers

import (
	"apuroda/repositories"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// FileController FileController
type FileController struct{}

// Index Index
func (f FileController) Index(c *gin.Context) {
	user, _ := c.Get("user")
	repo := repositories.FileRepository{}
	files := repo.All()
	c.HTML(http.StatusOK, "file_index.tmpl", gin.H{
		"files": files,
		"user":  user,
	})
}

// New New
func (f FileController) New(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.Redirect(http.StatusTemporaryRedirect, "/new_user")
		return
	}

	c.HTML(http.StatusOK, "file_new.tmpl", gin.H{"user": user})
}

// Create Create
func (f FileController) Create(c *gin.Context) {
	repo := repositories.FileRepository{}
	c.Request.ParseForm()
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusTemporaryRedirect, "/files")
	}

	fileModel, err := repo.Create(header.Filename, file)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusTemporaryRedirect, "/files")
	}
	fmt.Println(fileModel)
	c.Redirect(http.StatusFound, "/files/"+fileModel.ID.String())
}

// Show Show
func (f FileController) Show(c *gin.Context) {
	user, _ := c.Get("user")
	repo := repositories.FileRepository{}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		panic(err)
	}
	file, err := repo.GetByID(id)
	if err != nil && err.Error() == "not found" {
		c.Status(http.StatusNotFound)
		return
	}
	c.HTML(http.StatusOK, "file_show.tmpl", gin.H{"file": file, "user": user})
}

// Download Download
func (f FileController) Download(c *gin.Context) {
	_, ok := c.Get("user")
	if !ok {
		c.Status(http.StatusForbidden)
		return
	}
	repo := repositories.FileRepository{}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		panic(err)
	}
	file, err := repo.GetByID(id)

	reader, err := file.Get()
	if err != nil {
		panic(reader)
	}

	c.Writer.Header().Set("Content-type", "application/octet-stream")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Name))
	c.Stream(func(w io.Writer) bool {
		_, err := io.Copy(w, reader)
		if err != nil {
			panic(err)
		}
		return false
	})
}
