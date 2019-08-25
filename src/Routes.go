package main

import (
	"github.com/gin-gonic/gin"
	"go_books/src/Controller"
)

func AuthorRoutes(router *gin.RouterGroup) {
	router.POST("/", Controller.CreateAuthorAction)
	router.DELETE("/:id", Controller.DeleteAuthorAction)
	router.PATCH("/:id", Controller.EditAuthorAction)
}

func BookRoutes(router *gin.RouterGroup) {
	router.POST("/", Controller.CreateBooksAction)
}

func CategoryRoutes(router *gin.RouterGroup) {
	router.POST("/", Controller.CreateCategoryAction)
	router.DELETE("/:id", Controller.DeleteCategoryAction)
	router.PATCH("/:id", Controller.EditCategoryAction)
}
