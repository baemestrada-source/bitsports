package models

type Categorie struct {
	ID    int   `json:"id"`
	Name  string  `json:"name"`
}

type DeleteCategorie struct {
	DeletedCount int64 `json:"deletedCount"`
}

type UpdateCategorie struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        Categorie
}

type CreateCategorie struct {
	CreateCount   int64 `json:"createCount"`
	Result        Categorie
}
