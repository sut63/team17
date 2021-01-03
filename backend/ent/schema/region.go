package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Province holds the schema definition for the Province entity.
type Region struct {
	ent.Schema
}

// Fields of the Region.
func (Region) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the Region.
func (Region) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("regi_prov", Province.Type),
	}
}
