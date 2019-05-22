package Service

import (
	"errors"
	"go_books/src/Command"
	"go_books/src/Config"
	"go_books/src/Domain"
)

func AddAuthor(command Command.AuthorCommand) error {
	db := Config.GetDB()
	var author Domain.Author
	db.Where(&Domain.Author{Name: command.Name}).First(&author)

	if author.Name != "" {
		return errors.New("Author Name Exist")
	}

	if dbc := db.Create(&Domain.Author{Name: command.Name}); dbc.Error != nil {
		return dbc.Error
	}

	return nil
}

func DeleteAuthor(id int) error {
	db := Config.GetDB()
	var author Domain.Author
	db.Where("id = ?", id).First(&author)

	if author.Name == "" {
		return errors.New("Author Not Exist")
	}

	if dbc := db.Delete(&author); dbc.Error != nil {
		return dbc.Error
	}

	return nil
}

func EditAuthor(id int, command Command.AuthorCommand) error  {
	db := Config.GetDB()
	var author Domain.Author
	db.Where("id = ?", id).First(&author)

	if author.Name == "" {
		return errors.New("Author Not Exist")
	}

	if author.Name != command.Name {
		var checkedAuthor Domain.Author
		db.Where("name = ?", command.Name).First(&checkedAuthor)

		if checkedAuthor.Name != "" {
			return errors.New("Author Name Exist")
		}
	}
	author.Name = command.Name
	db.Save(&author)

	return nil
}