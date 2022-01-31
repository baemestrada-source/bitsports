package resolves

import (	
	"github.com/graphql-go/graphql" 
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"Product": 	 		 GetProductByKey(),
		"Products":  	     GetProducts(),
		"User":     	     GetUserByKey(),
		"Categories":     	 GetCategories(),
		"Categorie":         GetCategorieByKey(),
	},
})

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"checkin": CheckIn(),
		"login": Login(),
		"createProduct": InsProduct(),
		"deleteProduct": DelProduct(),
		"updateProduct": UpdProduct(),
		"createCategorie": InsCategorie(),
		"deleteCategorie": DelCategorie(),
		"updateCategorie": UpdCategorie(),

	},
})

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    RootQuery,
		Mutation: RootMutation,
	},
)