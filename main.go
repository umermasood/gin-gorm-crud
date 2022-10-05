package main

import (
	"github.com/gin-gonic/gin"
	"github.com/umermasood/gin-gorm-crud/controllers"
	"github.com/umermasood/gin-gorm-crud/initializers"
)

func init() {
	// environment variables are loaded into the ENV
	initializers.LoadEnvVariables()
	// postgres db is initialized and exported for other packages
	initializers.ConnectToDB()
}

func main() {
	// default gin router is created
	r := gin.Default()

	// CRUD API using GORM

	// CREATE post
	r.POST("/posts", controllers.PostsCreate)
	// READ posts
	r.GET("/posts", controllers.PostsRead)
	r.GET("/posts/:id", controllers.PostsShow)
	// UPDATE post
	r.PUT("/posts/:id", controllers.PostsUpdate)
	// DELETE post
	r.DELETE("/posts/:id", controllers.PostsDelete)

	_ = r.Run() // listen and serve on localhost:3000
}
