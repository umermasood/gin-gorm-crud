package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/umermasood/gin-gorm-crud/initializers"
	"github.com/umermasood/gin-gorm-crud/models"
	"net/http"
)

func PostsCreate(c *gin.Context) {
	// struct that holds request body data
	var body struct {
		Title string
		Body  string
	}
	// bind the request body json to the body struct
	if err := c.Bind(&body); err != nil {
		_ = errors.New(err.Error())
	}

	// create a new post to store in db
	post := models.Post{Title: body.Title, Body: body.Body}

	// store post in db
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	// return the post as a response
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func PostsRead(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)
	if post.Title == "" && post.Body == "" {
		msg := fmt.Sprintf("post with id=%v doesn't exist", id)
		c.JSON(http.StatusBadRequest, gin.H{"message": msg})
	} else {
		c.JSON(http.StatusOK, gin.H{"post": post})
	}

}

func PostsUpdate(c *gin.Context) {
	// Get the post id from URL
	id := c.Param("id")
	// struct that holds request body data
	var body struct {
		Title string
		Body  string
	}
	// bind the request body json to the body struct
	if err := c.Bind(&body); err != nil {
		_ = errors.New(err.Error())
	}

	// find the post that we have to update
	var post models.Post
	initializers.DB.First(&post, id)
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// return the response
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(new(models.Post), id)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("post with id=%v deleted", id)})
}
