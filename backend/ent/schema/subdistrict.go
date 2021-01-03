package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Subdistrict holds the schema definition for the Subdistrict entity.
type Subdistrict struct {
	ent.Schema
}

// Fields of the Subdistrict.
func (Subdistrict) Fields() []ent.Field {
	return []ent.Field{
		field.String("subdistrict").NotEmpty(),
	}
}

// Edges of the Subdistrict.
func (Subdistrict) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subd_dist", District.Type),
	}
}
