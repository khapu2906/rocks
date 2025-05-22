package main

import (
	"log"
	"os"
	"rocks/internal/user"
	"rocks/pkg/core"
	"rocks/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.Connect()

	r := gin.Default()
	r.Use(core.ErrorHandler())

    api := r.Group("/api")

    user.RegisterModule(api, db)

	r.Run(":" + os.Getenv("PORT"))
}
