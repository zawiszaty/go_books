package Service

import (
	"errors"
	"go_books/src/Command"
	"go_books/src/Config"
	"go_books/src/Domain"
)

func AddBooks(command Command.BookCommand) error {
	db := Config.GetDB()
	var book Domain.Books
	db.Where(&Domain.Books{Name: command.Name}).First(&book)
	var author Domain.Author
	db.Where("id = ?", command.Author).First(&author)
	var category Domain.Category
	db.Where("id = ?", command.Category).First(&category)

	if author.Name == "" {
		return errors.New("Author Not Exist")
	}

	if category.Name == "" {
		return errors.New("Category Not Exist")
	}

	if book.Name != "" {
		return errors.New("Book Name Exist")
	}

	if dbc := db.Create(&Domain.Books{Name: command.Name, Description: command.Description, Author: uint(command.Author), Category: uint(command.Category)}); dbc.Error != nil {
		return dbc.Error
	}

	return nil
}
