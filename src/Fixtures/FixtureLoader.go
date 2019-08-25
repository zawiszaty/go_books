package Fixtures

import (
	"database/sql"
	"github.com/takashabe/go-fixture"
	"go_books/src/Config"
	"log"
)

func LoadFixture() {
	db, _ := sql.Open("mysql", Config.GetEnv("DB_URL"))
	f, _ := fixture.NewFixture(db, "mysql")
	err := f.Load("../Fixtures/Author.sql")
	if err != nil {
		log.Fatal(err)
	}

	err = f.Load("../Fixtures/Category.sql")
	if err != nil {
		log.Fatal(err)
	}

	err = f.Load("../Fixtures/Book.sql")
	if err != nil {
		log.Fatal(err)
	}
}
