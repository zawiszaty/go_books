package Controller

import (
	"github.com/gin-gonic/gin"
	"go_books/src/Command"
	"go_books/src/Service"
	"net/http"
	"strconv"
)

func CreateCategoryAction(c *gin.Context) {
	var categoryCommand Command.CategoryCommand
	err := c.ShouldBindJSON(&categoryCommand)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = Service.AddCategory(categoryCommand)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"name": categoryCommand.Name})
	}
}

func DeleteCategoryAction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	err = Service.DeleteCategory(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func EditCategoryAction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	var categoryCommand Command.CategoryCommand
	err = c.ShouldBindJSON(&categoryCommand)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = Service.EditCategory(id, categoryCommand)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	} else {
		c.JSON(http.StatusOK, "success")
	}
}