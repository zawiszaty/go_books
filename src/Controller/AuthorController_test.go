package Controller

import (
	"bytes"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/takashabe/go-fixture/mysql"
	"go_books/src/Fixtures"
	"gotest.tools/assert"
	"net/http"
	"testing"
)

var (
	db       *sql.DB
)



func TestCreateAuthorAction(t *testing.T) {
	Fixtures.LoadFixture()
	values := []byte(`{"name": "testowa"}`)
	req, _ := http.NewRequest("POST","http://127.0.0.1:8080/api/author/", bytes.NewBuffer(values))
	client := &http.Client{}
	resp, _ := client.Do(req)
	assert.Equal(t, resp.StatusCode, 200)
}

func TestDeleteAuthorAction(t *testing.T) {
	Fixtures.LoadFixture()
	req, _ := http.NewRequest("DELETE","http://127.0.0.1:8080/api/author/1", bytes.NewBuffer([]byte(`{}`)))
	client := &http.Client{}
	resp, _ := client.Do(req)
	assert.Equal(t, resp.StatusCode, 200)
}

func TestEditAuthorAction(t *testing.T) {
	Fixtures.LoadFixture()
	values := []byte(`{"name": "nowa testowa"}`)
	req, _ := http.NewRequest("PATCH","http://127.0.0.1:8080/api/author/1", bytes.NewBuffer(values))
	client := &http.Client{}
	resp, _ := client.Do(req)
	assert.Equal(t, resp.StatusCode, 200)
}
