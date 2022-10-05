package main

import (
	"github.com/umermasood/gin-gorm-crud/initializers"
	"github.com/umermasood/gin-gorm-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
