package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", GetEnv("DB_URL"))

	if err != nil {
		fmt.Println(err)
	}

	return db
}
