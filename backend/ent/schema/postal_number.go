package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Postal_number holds the schema definition for the Subdistrict entity.
type Postal_number struct {
	ent.Schema
}

// Fields of the Postal_number.
func (Postal_number) Fields() []ent.Field {
	return []ent.Field{
		field.String("postal_number").NotEmpty(),
	}
}

// Edges of the Postal_number.
func (Postal_number) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("province", Province.Type).StorageKey(edge.Column("postal_number")),
	}
}
