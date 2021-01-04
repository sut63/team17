package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.String("course").NotEmpty(),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cour_facu", Faculty.Type).
			Ref("facu_cour").Unique(),
		edge.From("cour_degr", Degree.Type).
			Ref("degr_cour").Unique(),
		edge.From("cour_inst", Institution.Type).
			Ref("inst_cour").Unique(),
	}
}
