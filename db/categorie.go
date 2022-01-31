package db

import (
	"context"
	"log"
//	"github.com/baemestrada-source/bitsports/ent/Categories"
	"github.com/baemestrada-source/bitsports/models"
)

/*FindCategorie realiza una busqueda del Categorieo en la BD */
func FindCategorie(id int) (error, models.Categorie) {
	ctx := context.Background()

	client := PosgresCN

	p, err := client.Categories.Get(ctx, id) //
    if err != nil {
        return err, models.Categorie{}  //si no ecuentra el Categorieo
    }

	log.Println(p)

	retorna := models.Categorie{ 
    	Name: p.Name,
	}

	//log.Println("user returned: ", retorna)

	return nil, retorna
}

/*CreateCategorie crea el Categorieo en la BD */
func CreateCategorie(Categorie models.Categorie) (error, models.CreateCategorie) {
	ctx := context.Background()
	log.Println("entro ", Categorie)

	client := PosgresCN
	
	p, err := client.Categories.Create().      
    SetName(Categorie.Name).       
    Save(ctx) 

    if err != nil {
        return err, models.CreateCategorie{CreateCount:0}
    }

    log.Println("Categorie was create: ", p)
 
	retorna := models.CreateCategorie{
        CreateCount:1, 
        Result: models.Categorie{
			ID: p.ID,
			Name: p.Name,
		},
    }

	log.Println("Categorie create: ", retorna)

	return nil, retorna
}


/*DeleteCategorie crea el Categorieo en la BD */
func DeleteCategorie(id int) (error, models.DeleteCategorie) {
	ctx := context.Background()

	client := PosgresCN
	
	err := client.Categories.DeleteOneID(id).Exec(ctx)

    if err != nil {
        return err, models.DeleteCategorie{0}
    }

	retorna := models.DeleteCategorie{1}

	log.Println("id eliminado")

	return nil, retorna
}

/*UpdateCategorie actualiza el Categorieo en la BD */
func UpdateCategorie(Categorie models.Categorie) (error, models.UpdateCategorie) {
	ctx := context.Background()

	client := PosgresCN
	
	p, err := client.Categories.UpdateOneID(Categorie.ID).      
    SetName(Categorie.Name).       
    Save(ctx) 

    if err != nil {
        return err, models.UpdateCategorie{ModifiedCount:0}
    }

    log.Println("user was updated: ", p)
 
	retorna := models.UpdateCategorie{
        ModifiedCount:1, 
        Result: models.Categorie{
			Name: p.Name,
		},
    }

	log.Println("user returned: ", retorna)

	return nil, retorna
}


/*GetCategories realiza una busqueda del Categorieo en la BD */
func GetCategories() (error, []models.Categorie) {
	ctx := context.Background()

	client := PosgresCN

	items, err := client.Categories.Query().All(ctx)
    
	if err != nil {
        log.Fatalf("failed querying todos: %v", err)
    }
	//log.Println("items ", items)
    
	var CategorieArray []models.Categorie

	for _, t := range items {
			CategorieArray = append(CategorieArray,
				 models.Categorie{
					ID:    t.ID,
					Name:  t.Name,
				})	
    }

	//log.Println("user returned: ", CategorieArray)

	return nil, CategorieArray
}