package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Emp holds the schema definition for the Emp entity.
type Emp struct {
	ent.Schema
}

// Fields of the Emp.
func (Emp) Fields() []ent.Field {
	return []ent.Field{
		field.String("user").NotEmpty(),
		field.String("pass").NotEmpty(),
		field.String("role"),
	}
}

// Edges of the Emp.
func (Emp) Edges() []ent.Edge {
	return nil
}
