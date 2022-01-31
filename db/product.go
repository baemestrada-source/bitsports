package db

import (
	"context"
	"log"
//	"github.com/baemestrada-source/bitsports/ent/products"
	"github.com/baemestrada-source/bitsports/models"
)

/*FindProduct realiza una busqueda del producto en la BD */
func FindProduct(id int) (error, models.Product) {
	ctx := context.Background()

	client := PosgresCN

	p, err := client.Products.Get(ctx, id) //
    if err != nil {
        return err, models.Product{}  //si no ecuentra el producto
    }

	log.Println(p)

	retorna := models.Product{ 
    	Name: p.Name,
		Info: p.Info,
		Price: p.Price,
		Categorie_id: p.CategorieID,
	}

	//log.Println("user returned: ", retorna)

	return nil, retorna
}

/*CreateProduct crea el producto en la BD */
func CreateProduct(product models.Product) (error, models.CreateProduct) {
	ctx := context.Background()
	log.Println("entro ", product)

	client := PosgresCN
	
	p, err := client.Products.Create().      
    SetName(product.Name).       
    SetInfo(product.Info).      
	SetPrice(product.Price).      
	SetCategorieID(product.Categorie_id).      
    Save(ctx) 

    if err != nil {
        return err, models.CreateProduct{CreateCount:0}
    }

    log.Println("product was create: ", p)
 
	retorna := models.CreateProduct{
        CreateCount:1, 
        Result: models.Product{
			ID: p.ID,
			Name: p.Name,
			Info: p.Info,
			Price: p.Price,
			Categorie_id: p.CategorieID,
		},
    }

	log.Println("product create: ", retorna)

	return nil, retorna
}


/*DeleteProduct crea el producto en la BD */
func DeleteProduct(id int) (error, models.DeleteProduct) {
	ctx := context.Background()

	client := PosgresCN
	
	err := client.Products.DeleteOneID(id).Exec(ctx)

    if err != nil {
        return err, models.DeleteProduct{0}
    }

	retorna := models.DeleteProduct{1}

	log.Println("id eliminado")

	return nil, retorna
}

/*UpdateProduct actualiza el producto en la BD */
func UpdateProduct(product models.Product) (error, models.UpdateProduct) {
	ctx := context.Background()

	client := PosgresCN
	
	p, err := client.Products.UpdateOneID(product.ID).      
    SetName(product.Name).       
    SetInfo(product.Info).      
	SetPrice(product.Price).      
	SetCategorieID(product.Categorie_id).      
    Save(ctx) 

    if err != nil {
        return err, models.UpdateProduct{ModifiedCount:0}
    }

    log.Println("user was updated: ", p)
 
	retorna := models.UpdateProduct{
        ModifiedCount:1, 
        Result: models.Product{
			Name: p.Name,
			Info: p.Info,
			Price: p.Price,
			Categorie_id: p.CategorieID,
		},
    }

	log.Println("user returned: ", retorna)

	return nil, retorna
}


/*GetProducts realiza una busqueda del producto en la BD */
func GetProducts() (error, []models.Product) {
	ctx := context.Background()

	client := PosgresCN

	items, err := client.Products.Query().All(ctx)
    
	if err != nil {
        log.Fatalf("failed querying todos: %v", err)
    }
	//log.Println("items ", items)
    
	var ProductArray []models.Product

	for _, t := range items {
			ProductArray = append(ProductArray,
				 models.Product{
					ID:    t.ID,
					Name:  t.Name,
					Info:  t.Info,
					Price: t.Price,
					Categorie_id: t.CategorieID,
				})	
    }

	//log.Println("user returned: ", ProductArray)

	return nil, ProductArray
}