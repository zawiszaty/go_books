package Controller

import (
	"github.com/gin-gonic/gin"
	"go_books/src/Command"
	"go_books/src/Service"
	"net/http"
)

func CreateBooksAction(c *gin.Context) {
	var bookCommand Command.BookCommand
	err := c.ShouldBindJSON(&bookCommand)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = Service.AddBooks(bookCommand)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Name":      bookCommand.Name,
			"Description": bookCommand.Description,
			"Author": bookCommand.Author,
			"Category": bookCommand.Category,
		})
	}
}
