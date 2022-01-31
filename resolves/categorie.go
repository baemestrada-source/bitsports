package resolves

import 
(
	"log"
//	"errors"
	"encoding/json"
	"github.com/baemestrada-source/bitsports/models" 
	"github.com/baemestrada-source/bitsports/db" 
	"github.com/graphql-go/graphql"
)

var categorieType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Categorie",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var updateCategorieType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UpdateCategorieType",
	Fields: graphql.Fields{
		"modifiedCount": &graphql.Field{Type: graphql.Int},
		"result":        &graphql.Field{Type: categorieType},
	},
})

var createCategorieType = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreateCategorieType",
	Fields: graphql.Fields{
		"createCount":   &graphql.Field{Type: graphql.Int},
		"result":        &graphql.Field{Type: categorieType},
	},
})

var deleteCategorieType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DeleteCategorieType",
	Fields: graphql.Fields{
		"deletedCount": &graphql.Field{Type: graphql.Int},
	},
})

func GetCategorieByKey() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}
	return &graphql.Field{
		Name:        "Obtiene el Categorieo buscado",
		Type:        categorieType,
		Description: "Get Categorie by id",
		Args:        args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(int)
			if ok {
				// Find Categorie
				err, respuesta := db.FindCategorie(id)
				if err == nil {
					return respuesta, nil
				}
			}
			return nil, nil
		},
	}
}

//InsCategorie registra usuario
func InsCategorie() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
	return &graphql.Field{
		Name:        "InsCategorie",
		Type:        createCategorieType,
		Description: "Ingresa un Categorieo a la base de datos",
		Args:        args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			Categorie := models.Categorie{}
			
			body, err := json.Marshal(p.Args)
			if err != nil {
				return models.CreateCategorie{CreateCount:0}, err
			}
			
			err = json.Unmarshal(body, &Categorie)
			if err != nil {
				return models.CreateCategorie{CreateCount:0}, err
			}

			log.Println("Categorie: ",Categorie)
			
			err, createCategorie := db.CreateCategorie(Categorie)
			
			log.Println("respuesta: ",createCategorie)

			if err != nil {
				return models.CreateCategorie{CreateCount:0}, err
			}

			return createCategorie, nil
		},
	}
}

func DelCategorie() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}
	return &graphql.Field{
		Name:        "Delete Categorieo",
		Description: "delete Categorie by id",
		Type:        deleteCategorieType,
		Args:        args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(int)

			err, deletedCount := db.DeleteCategorie(id)
			if err != nil {
				return deletedCount, err
			}

			return deletedCount,nil
		},
	}
}

func UpdCategorie() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
	return &graphql.Field{
		Name:        "Update Categorie",
		Description: "update Categorieo by id",
		Type:        updateCategorieType,
		Args:        args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			Categorie := models.Categorie{}
			
			body, err := json.Marshal(p.Args)
			if err != nil {
				return models.UpdateCategorie{ModifiedCount:0}, err
			}
			
			err = json.Unmarshal(body, &Categorie)
			if err != nil {
				return models.UpdateCategorie{ModifiedCount:0}, err
			}

			//log.Println("Categorie: ",Categorie)
			
			err, updateCategorie := db.UpdateCategorie(Categorie)
			
			if err != nil {
				return models.UpdateCategorie{ModifiedCount:0}, err
			}

			return updateCategorie, nil
		},
	}
}


func GetCategories() *graphql.Field {
	return &graphql.Field{
		Name:        "list_Categories",
		Type:        graphql.NewList(categorieType),
		Description: "Get Categorie list",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {

			err, array_Categories := db.GetCategories()
			if err != nil {
				return nil, err
			} 
			//log.Println("resolver retorno: ", array_Categories)

			return array_Categories, nil
		},
	}
}