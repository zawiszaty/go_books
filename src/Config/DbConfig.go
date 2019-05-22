package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/go_books?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}

	return db
}