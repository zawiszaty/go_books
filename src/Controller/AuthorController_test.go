package Controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
	"net/http"
	"testing"
)

func TestCreateAuthorAction(t *testing.T) {
	values := map[string]string{"name": "asdasd"}

	jsonValue, _ := json.Marshal(values)

	resp, _ := http.Post("http://127.0.0.1:8080/api/author/", "application/json", bytes.NewBuffer(jsonValue))
	assert.Equal(t, resp.Status, 200)
}

func TestDeleteAuthorAction(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestEditAuthorAction(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}