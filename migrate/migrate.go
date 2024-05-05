package main

import (
	"API/initializers"
	"API/models"
)

func init() {
	initializers.LoadEnvVariabels()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.Post{},
	)
}
