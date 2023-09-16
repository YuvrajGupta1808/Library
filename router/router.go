package router

import (
	"myproject/controllers"

	"github.com/gin-gonic/gin"
)

// SetupAPIRouter initializes and configures the API router.
func SetupAPIRouter(r *gin.RouterGroup) {
	// Add a middleware to handle errors
	r.Use(gin.Recovery())

	// Add your API routes here
	r.GET("/books", controllers.GetBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.GetBookByID)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
}

// SetupViewRouter initializes and configures the view router.
func SetupViewRouter(r *gin.RouterGroup) {
	// Serve the "index.html" file from the "views" directory
	r.StaticFile("/", "./views/index.html")
}
