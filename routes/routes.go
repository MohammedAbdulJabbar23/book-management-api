package routes

import (
	"github.com/MohammedAbdulJabbar23/book-management-api/controllers"
	"github.com/gin-gonic/gin"
)



func SetupRoutes(router *gin.Engine) {
	router.GET("/books", controllers.FindBooks);
	router.POST("/books", controllers.CreateBook);
}