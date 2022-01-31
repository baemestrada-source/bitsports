package main

import (
	"fmt"
	"context"
	"log"
	"github.com/baemestrada-source/bitsports/db"
	_ "github.com/lib/pq"
	"net/http"
	"github.com/graphql-go/handler"
	"github.com/baemestrada-source/bitsports/handlers"
	"github.com/baemestrada-source/bitsports/resolves"
	"github.com/baemestrada-source/bitsports/middlew"
)

func main() {

	client := db.PosgresCN

	defer client.Close()
	ctx := context.Background()
	// Ejecuta migracion de base de datos en base a esquemas
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	
	h := handler.New(&handler.Config{
		Schema: &resolves.Schema,
		Pretty: true,
		GraphiQL: true,
	})
	//inicio el path para mi graphql siempre que tenga el Token
	http.Handle("/graphql", middlew.HttpHeaderMiddleware(h))
	
	//rutas directas por cada uno
	http.HandleFunc("/product", middlew.ValidoJWT(handlers.Products))
	http.HandleFunc("/categorie", middlew.ValidoJWT(handlers.Categories))
	http.HandleFunc("/user", handlers.Users)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)	
}
