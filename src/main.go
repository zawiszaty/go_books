package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_books/src/Config"
	"go_books/src/Domain"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Domain.Author{})
	db.AutoMigrate(&Domain.Category{})
	db.AutoMigrate(&Domain.Books{})
}

func main() {
	db := Config.GetDB()
	Migrate(db)
	defer db.Close()
	r := gin.Default()
	v1 := r.Group("/api")
	AuthorRoutes(v1.Group("author"))
	CategoryRoutes(v1.Group("category"))
	BookRoutes(v1.Group("book"))
	r.Run()
	fmt.Println("port 8080")
}
