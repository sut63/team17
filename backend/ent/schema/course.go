package schema

import (
		"github.com/facebookincubator/ent"
		"github.com/facebookincubator/ent/schema/edge"
		"github.com/facebookincubator/ent/schema/field")


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
		edge.From("course_fac", Faculty.Type).
			Ref("fac_course").Unique(),
		edge.From("course_deg", Degree.Type).
			Ref("deg_course").Unique(),
		edge.From("course_ins", Institution.Type).
			Ref("ins_course").Unique(),
    }
}
