package main

import (
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"it1shka.com/my-recipes/database"
	"it1shka.com/my-recipes/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	database.Connect()

	server := gin.Default()

	secret := os.Getenv("SESSION_SECRET")
	store := gormsessions.NewStore(database.DB, true, []byte(secret))
	server.Use(sessions.Sessions("_mysessions", store))

	server.LoadHTMLGlob("templates/*.html")
	server.Static("/assets", "./assets")

	handlers.Setup(server)

	APP_PORT := os.Getenv("APP_PORT")
	server.Run(APP_PORT)
}
