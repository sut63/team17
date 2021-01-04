package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Institution holds the schema definition for the Institution entity.

type Institution struct {
	ent.Schema
}

// Fields of the Institution.
func (Institution) Fields() []ent.Field {
	return []ent.Field{

		field.String("institution").NotEmpty(),
	}
}

// Edges of the Institution.
func (Institution) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("inst_cour", Course.Type),
	}
}
