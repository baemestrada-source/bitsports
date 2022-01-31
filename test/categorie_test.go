package test

import (
	"testing"
	"net/http"
	"bytes"
	"net/http/httptest"
	"github.com/baemestrada-source/bitsports/handlers"
	"github.com/stretchr/testify/assert"
)

func TestCategories(t *testing.T) {
	
	var query = []byte(`{
		mutation {
			createCategorie(
				   name: "categoria nueva", 
				) 
			{
			  createCount
			  result{
				id
				name
			   }
			}
		  }
		}`)
	
	req, err := http.NewRequest("POST", "/graphql", bytes.NewBuffer(query))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(handlers.Categories)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	status := response.Code
	assert.Equal(t, status, http.StatusOK)
}