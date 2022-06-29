package main

import (
	// "errors"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

//! dummy json slice
var bookSample = []book{
	{ID: "0", Title: "Lost In The Ocean", Author: "Frank Mercer", Quantity: 2},
	{ID: "1", Title: "Simple Love", Author: "Edward Drule", Quantity: 3},
	{ID: "2", Title: "Lust of the Ancients", Author: "Merle Queen", Quantity: 7},
	{ID: "3", Title: "Crew List", Author: "Poper Jenkins", Quantity: 4},
}

//! Add stuff to the json Slice
func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		//catch error
		return
	}
	bookSample = append(bookSample, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bookSample)
}

func main() {

	router := gin.Default()
	//? name a route to fetch the dummy sample

	router.GET("/bookSample", getBooks)
	//?router to add new book

	router.POST("/bookSample", createBook)
	router.Run("localhost:8080")
}
