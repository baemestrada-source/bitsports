package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Categories struct {
    ent.Schema
}

// Fields of the Categories.
func (Categories) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").
            Unique().
            NotEmpty(),
    }
}

// Edges of the Categories.
func (Categories) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("products", Products.Type).
            Ref("union"),
    }
}