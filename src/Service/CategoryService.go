package Service

import (
	"errors"
	"go_books/src/Command"
	"go_books/src/Config"
	"go_books/src/Domain"
)

func AddCategory(command Command.CategoryCommand) error {
	db := Config.GetDB()
	var category Domain.Category
	db.Where(&Domain.Category{Name: command.Name}).First(&category)

	if category.Name != "" {
		return errors.New("Category Name Exist")
	}

	if dbc := db.Create(&Domain.Category{Name: command.Name}); dbc.Error != nil {
		return errors.New("Category Name Exist")
	}

	return nil
}

func DeleteCategory(id int) error {
	db := Config.GetDB()
	var category Domain.Category
	db.Where("id = ?", id).First(&category)

	if category.Name == "" {
		return errors.New("Category Not Exist")
	}

	if dbc := db.Delete(&category); dbc.Error != nil {
		return dbc.Error
	}

	return nil
}

func EditCategory(id int, command Command.CategoryCommand) error  {
	db := Config.GetDB()
	var category Domain.Category
	db.Where("id = ?", id).First(&category)

	if category.Name == "" {
		return errors.New("Category Not Exist")
	}

	if category.Name != command.Name {
		var checkedCategory Domain.Category
		db.Where("name = ?", command.Name).First(&checkedCategory)

		if checkedCategory.Name != "" {
			return errors.New("Author Name Exist")
		}
	}
	category.Name = command.Name
	db.Save(&category)

	return nil
}
