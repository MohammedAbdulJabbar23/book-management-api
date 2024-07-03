package routes

import (
	"github.com/MohammedAbdulJabbar23/book-management-api/controllers"
	"github.com/gin-gonic/gin"
)



func SetupRoutes(router *gin.Engine) {
	router.GET("/books", controllers.FindBooks);
	router.POST("/books", controllers.CreateBook);
	router.GET("/books/:id",controllers.FindBook);
	router.PUT("/books/:id", controllers.UpdateBook);
	router.DELETE("/books/:id",controllers.DeleteBook);
}