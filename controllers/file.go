package controllers

import (
	"apuroda/repositories"
	"bytes"
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

	extraHeaders := map[string]string{
		"Content-Disposition": fmt.Sprintf("attachment; filename=%s", file.Name),
	}
	reader, err := file.Get()
	if err != nil {
		panic(reader)
	}

	c.DataFromReader(http.StatusOK, int64(getSize(reader)), "application/octet-stream", reader, extraHeaders)
}

// func getSize(stream io.Reader) int64 {
// 	buf := &bytes.Buffer{}
// 	nRead, err := io.Copy(buf, stream)
// 	if err != nil {
// 		panic(stream)
// 	}
// 	return nRead
// }

func getSize(stream io.Reader) int {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	fmt.Println(buf.Len())
	return buf.Len()
}
