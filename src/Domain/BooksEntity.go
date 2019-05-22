package Domain

import (
	"github.com/jinzhu/gorm"
)

type Books struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	Description string
	Author      uint
	Category    uint
}
