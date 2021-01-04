package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Faculty holds the schema definition for the Faculty entity.

type Faculty struct {
	ent.Schema
}

// Fields of the Faculty.
func (Faculty) Fields() []ent.Field {
	return []ent.Field{

		field.String("faculty").NotEmpty(),
	}
}

// Edges of the Faculty.
func (Faculty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("facu_cour", Course.Type),
		edge.To("facu_prof", Professor.Type),
	}
}
