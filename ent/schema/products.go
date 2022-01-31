package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Products struct {
    ent.Schema
}

func (Products) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").
            NotEmpty(),
		field.String("info"),
		field.Float("price"),
        field.Int("categorie_id"), // <-- explictly add the field we want to contain the FK
    }
}

// Edges of the Products.
func (Products) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("union", Categories.Type).
            Field("categorie_id"). // <-- tell ent which field holds the reference to the owner
            Unique().
            Required(),
    }
}