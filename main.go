package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"myproject/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize the web server
	r := gin.Default()

	// Set up your routes using the router package
	r.GET("/", func(c *gin.Context) {
		viewsPath := "./views"
		// Parse the template files from the views folder
		tmpl := template.Must(template.ParseGlob(viewsPath + "/index.html"))
		tmpl.Execute(c.Writer, nil)
	})

	// Set up API routes
	apiGroup := r.Group("/api")
	router.SetupAPIRouter(apiGroup)

	// Start the web server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":8080")

	http.ListenAndServe(":8080", r)
	//if err != nil {
	//	log.Fatalf("Error starting server: %v", err)
	//}
}
