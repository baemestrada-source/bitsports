package models

type Product struct {
	ID    int   `json:"id"`
	Name  string  `json:"name"`
	Info  string  `json:"info,omitempty"`
	Price float64 `json:"price"`
	Categorie_id int `json:"categorie_id"`
}

type DeleteProduct struct {
	DeletedCount int64 `json:"deletedCount"`
}

type UpdateProduct struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        Product
}

type CreateProduct struct {
	CreateCount   int64 `json:"createCount"`
	Result        Product
}