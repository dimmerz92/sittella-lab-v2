package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"sittellalab.com.au/internal/pages"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env not found")
	}

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8000"
	}

	e := echo.New()

	// static file servers
	e.Static("/static/public", "static/public")

	// page servers
	e.GET("/", pages.Home)

	e.Logger.Fatal(e.Start(":" + port))
}
