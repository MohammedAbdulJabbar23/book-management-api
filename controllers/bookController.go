package controllers

import (
	"database/sql"
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

func UpdateBook(c *gin.Context) {
	var input models.Book;
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"please provide all the information"});
		return;
	}
	query := "UPDATE books SET title=$1, author=$2, year=$3 WHERE id=$4";
	_, err := config.DB.Exec(query,input.Title,input.Author,input.Year, c.Param("id"));
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error});
		return;
	}
	c.JSON(http.StatusOK, gin.H{"book":input});
}


func FindBook(c *gin.Context) {
	var book models.Book;
	query := "SELECT * FROM books WHERE id=$1";
	row := config.DB.QueryRow(query, c.Param("id"));
	err := row.Scan(&book.ID,&book.Title,&book.Author,&book.Year);
	if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"});
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()});
        }
        return;
    }
	c.JSON(http.StatusOK, gin.H{"book":book});
}


func DeleteBook(c *gin.Context) {
	query := "DELETE FROM books WHERE id=$1";
	_, err := config.DB.Exec(query,c.Param("id"));
	if err != nil {
		if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"});
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()});
        }
        return;
	}
	c.JSON(http.StatusOK, gin.H{"info":"book successfully deleted"});
}