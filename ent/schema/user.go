package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(20)",
		}),
		field.String("email").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)",
		}),
		field.Int("age").SchemaType(map[string]string{
			dialect.MySQL: "int",
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
