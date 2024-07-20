package routes

import (
	"github.com/MohammedAbdulJabbar23/book-management-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    authorized := r.Group("/")
    authorized.Use(controllers.AuthMiddleware())
    {
        authorized.POST("/books", controllers.CreateBook);
        authorized.GET("/books", controllers.FindBooks);
        authorized.GET("/books/:id",controllers.FindBook);
        authorized.DELETE("/books/:id",controllers.DeleteBook);
        // authorized.POST("/books/:id/upload", controllers.UploadPDF)
        // authorized.GET("/books/:id/download", controllers.DownloadPDF)
    }
    return r
}
