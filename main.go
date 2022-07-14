package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// function for handler with JSON output
func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Albert Afriadi Sigiro",
		"bio":  "Junior Developer",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "Golang API",
		"subtitle": "Learn Golang API",
	})
}

// function using param
func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

// function using query
func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

// struct data
type BookInput struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}

// function for POST
func postBooksHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

			// send error with message
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", postBooksHandler)

	router.Run()
}
