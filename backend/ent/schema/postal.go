package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Postal holds the schema definition for the Subdistrict entity.
type Postal struct {
	ent.Schema
}

// Fields of the Postal.
func (Postal) Fields() []ent.Field {
	return []ent.Field{
		field.String("postal").NotEmpty(),
	}
}

// Edges of the Postal.
func (Postal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("post_dist", District.Type),
	}
}
