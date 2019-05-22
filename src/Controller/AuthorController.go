package Controller

import (
	"github.com/gin-gonic/gin"
	"go_books/src/Command"
	"go_books/src/Service"
	"net/http"
	"strconv"
)

func CreateAuthorAction(c *gin.Context) {
	var authorCommand Command.AuthorCommand
	err := c.ShouldBindJSON(&authorCommand)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = Service.AddAuthor(authorCommand)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"name": authorCommand.Name})
	}
}

func DeleteAuthorAction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	err = Service.DeleteAuthor(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func EditAuthorAction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	var authorCommand Command.AuthorCommand
	err = c.ShouldBindJSON(&authorCommand)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = Service.EditAuthor(id, authorCommand)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	} else {
		c.JSON(http.StatusOK, "success")
	}
}