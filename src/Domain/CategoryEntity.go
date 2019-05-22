package Domain

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name  string  `gorm:"unique;not null"`
	Books []Books `gorm:"foreignkey:Category"`
}
