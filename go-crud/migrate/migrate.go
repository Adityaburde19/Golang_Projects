package main

import (
	"go-curd/initializers"
	"go-curd/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.POST{})
}
