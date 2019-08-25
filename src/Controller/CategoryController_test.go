package Controller

import (
	"bytes"
	"go_books/src/Fixtures"
	"gotest.tools/assert"
	"net/http"
	"testing"
)

func TestCreateCategoryAction(t *testing.T) {
	Fixtures.LoadFixture()
	values := []byte(`{"name": "testowa"}`)
	req, _ := http.NewRequest("POST","http://127.0.0.1:8080/api/category/", bytes.NewBuffer(values))
	client := &http.Client{}
	resp, _ := client.Do(req)
	assert.Equal(t, resp.StatusCode, 200)
}

func TestDeleteCategoryAction(t *testing.T) {
	Fixtures.LoadFixture()
	req, _ := http.NewRequest("DELETE","http://127.0.0.1:8080/api/category/1", bytes.NewBuffer([]byte(`{}`)))
	client := &http.Client{}
	resp, _ := client.Do(req)
	assert.Equal(t, resp.StatusCode, 200)
}

func TestEditCategoryAction(t *testing.T) {
	Fixtures.LoadFixture()
	values := []byte(`{"name": "nowa testowa"}`)
	req, _ := http.NewRequest("PATCH","http://127.0.0.1:8080/api/category/1", bytes.NewBuffer(values))
	client := &http.Client{}
	resp, _ := client.Do(req)
	assert.Equal(t, resp.StatusCode, 200)
}