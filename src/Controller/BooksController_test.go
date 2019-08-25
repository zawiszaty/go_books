package Controller

import (
	"bytes"
	"go_books/src/Fixtures"
	"gotest.tools/assert"
	"net/http"
	"testing"
)

func TestCreateBooksAction(t *testing.T) {
	Fixtures.LoadFixture()
	values := []byte(`{"name": "testowa", "description": "test", "author": 1, "category": 1}`)
	req, _ := http.NewRequest("POST","http://127.0.0.1:8080/api/book/", bytes.NewBuffer(values))
	client := &http.Client{}
	resp, _ := client.Do(req)
	assert.Equal(t, resp.StatusCode, 200)
}