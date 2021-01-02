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
		edge.From("degree", Degree.Type).Ref("course").Unique(),
		edge.From("faculty", Faculty.Type).Ref("course").Unique(),
		edge.From("institution", School.Type).Ref("course").Unique(),
		edge.From("creditfaculty", Creditfaculty.Type).Ref("course").Unique(),
    }
}