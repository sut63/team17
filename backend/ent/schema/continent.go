package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Continent holds the schema definition for the Continent entity.
type Continent struct {
	ent.Schema
}

// Fields of the Continent.
func (Continent) Fields() []ent.Field {
	return []ent.Field{
		field.String("continent").NotEmpty().Unique(),
	}
}

// Edges of the Continent.
func (Continent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cont_prov", Province.Type),
	}
}
