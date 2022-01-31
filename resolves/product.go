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

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"info": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"categorie_id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var updateProductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UpdateProductType",
	Fields: graphql.Fields{
		"modifiedCount": &graphql.Field{Type: graphql.Int},
		"result":        &graphql.Field{Type: productType},
	},
})

var createProductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreateProductType",
	Fields: graphql.Fields{
		"createCount":   &graphql.Field{Type: graphql.Int},
		"result":        &graphql.Field{Type: productType},
	},
})

var deleteproductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DeleteProductType",
	Fields: graphql.Fields{
		"deletedCount": &graphql.Field{Type: graphql.Int},
	},
})

func GetProductByKey() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}
	return &graphql.Field{
		Name:        "Obtiene el producto buscado",
		Type:        productType,
		Description: "Get Product by id",
		Args:        args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(int)
			if ok {
				// Find product
				err, respuesta := db.FindProduct(id)
				if err == nil {
					return respuesta, nil
				}
			}
			return nil, nil
		},
	}
}

//InsProduct registra usuario
func InsProduct() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"info": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"price": &graphql.ArgumentConfig{
			Type: graphql.Float,
		},
		"categorie_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}
	return &graphql.Field{
		Name:        "InsProduct",
		Type:        createProductType,
		Description: "Ingresa un producto a la base de datos",
		Args:        args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			product := models.Product{}
			
			body, err := json.Marshal(p.Args)
			if err != nil {
				return models.CreateProduct{CreateCount:0}, err
			}
			
			err = json.Unmarshal(body, &product)
			if err != nil {
				return models.CreateProduct{CreateCount:0}, err
			}

			log.Println("product: ",product)
			
			err, createProduct := db.CreateProduct(product)
			
			log.Println("respuesta: ",createProduct)

			if err != nil {
				return models.CreateProduct{CreateCount:0}, err
			}

			return createProduct, nil
		},
	}
}

func DelProduct() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}
	return &graphql.Field{
		Name:        "Delete Producto",
		Description: "delete product by id",
		Type:        deleteproductType,
		Args:        args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(int)

			err, deletedCount := db.DeleteProduct(id)
			if err != nil {
				return deletedCount, err
			}

			return deletedCount,nil
		},
	}
}

func UpdProduct() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"info": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"price": &graphql.ArgumentConfig{
			Type: graphql.Float,
		},
		"categorie_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}
	return &graphql.Field{
		Name:        "Update Product",
		Description: "update producto by id",
		Type:        updateProductType,
		Args:        args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			product := models.Product{}
			
			body, err := json.Marshal(p.Args)
			if err != nil {
				return models.UpdateProduct{ModifiedCount:0}, err
			}
			
			err = json.Unmarshal(body, &product)
			if err != nil {
				return models.UpdateProduct{ModifiedCount:0}, err
			}

			//log.Println("product: ",product)
			
			err, updateProduct := db.UpdateProduct(product)
			
			if err != nil {
				return models.UpdateProduct{ModifiedCount:0}, err
			}

			return updateProduct, nil
		},
	}
}


func GetProducts() *graphql.Field {
	return &graphql.Field{
		Name:        "list_products",
		Type:        graphql.NewList(productType),
		Description: "Get product list",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {

			err, array_products := db.GetProducts()
			if err != nil {
				return nil, err
			} 
			//log.Println("resolver retorno: ", array_products)

			return array_products, nil
		},
	}
}