package handlers

import (
	"net/http"
	"github.com/graphql-go/graphql"
	"fmt"
	"github.com/baemestrada-source/bitsports/resolves"
	"encoding/json"
)


func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        resolves.Schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func Products(w http.ResponseWriter, r *http.Request) {
	result := executeQuery(r.URL.Query().Get("query"), resolves.Schema)
	json.NewEncoder(w).Encode(result)
}

func Users(w http.ResponseWriter, r *http.Request) {
	result := executeQuery(r.URL.Query().Get("query"), resolves.Schema)
	json.NewEncoder(w).Encode(result)
}

func Categories(w http.ResponseWriter, r *http.Request) {
	result := executeQuery(r.URL.Query().Get("query"), resolves.Schema)
	json.NewEncoder(w).Encode(result)
}