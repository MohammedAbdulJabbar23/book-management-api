package controllers

import (
	"net/http"

	"github.com/MohammedAbdulJabbar23/book-management-api/config"
	"github.com/MohammedAbdulJabbar23/book-management-api/models"
	"github.com/gin-gonic/gin"
)


func FindBooks(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, title, author, year FROM books");
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": "Something went wrong from our side"});
		return;
	}
	defer rows.Close();
	books := []models.Book{};
	for rows.Next() {
		var book models.Book;
		if err := rows.Scan(&book.ID,&book.Title,&book.Author,&book.Year); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return;
		}
		books = append(books,book);
	}
	c.JSON(http.StatusOK,gin.H{"books":books});
}

func CreateBook(c *gin.Context) {
	var input models.Book;
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return;
	}
	query := "INSERT INTO books (title, author, year) VALUES ($1,$2,$3) RETURNING id";
	err := config.DB.QueryRow(query, input.Title, input.Author, input.Year).Scan(&input.ID);
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusOK, gin.H{"book": input});
}