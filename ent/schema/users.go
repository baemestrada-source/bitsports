package schema

import ( 
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Users struct {
	ent.Schema
}

// Fields of the User.
func (Users) Fields() []ent.Field {
    return []ent.Field{
		field.String("username").
		Unique(),
		field.String("name"),
        field.String("password"),
		field.String("email"),
    }
}