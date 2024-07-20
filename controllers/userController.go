package controllers

import (
	"fmt"
	"net/http"

	"github.com/MohammedAbdulJabbar23/book-management-api/models"
	"github.com/gin-gonic/gin"
)


func Register(c *gin.Context) {
	var user models.User;
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error});
		return;
	}
	if err := user.Register(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()});
		return;
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"});
}

func Login(c *gin.Context) {
	var user models.User;
	if err := c.BindJSON(&user); err != nil {
		fmt.Println(user);
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
		return;
	}
	token, err := user.Authenticate();
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":err.Error()});
		return;
	}
	c.JSON(http.StatusOK, gin.H{"token":token});
}


func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            c.Abort()
            return
        }

        claims, err := models.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("userID", claims.UserID)
        c.Next()
    }
}