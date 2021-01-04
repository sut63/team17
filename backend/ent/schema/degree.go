package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Degree holds the schema definition for the  Degree entity.

type Degree struct {
	ent.Schema
}

// Fields of the  Degree.
func (Degree) Fields() []ent.Field {
	return []ent.Field{

		field.String("degree").NotEmpty(),
	}
}

// Edges of the  Degree.
func (Degree) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("course", Course.Type),
		edge.From("degr_stud", Student.Type).Ref("stud_degr").Unique(),
	}
}
