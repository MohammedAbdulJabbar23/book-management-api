package main

import (
	"github.com/MohammedAbdulJabbar23/book-management-api/config"
	"github.com/MohammedAbdulJabbar23/book-management-api/routes"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default();
	config.ConnectDatabase();
	routes.SetupRoutes(router);
	router.Run(":8080");
}