package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Results holds the schema definition for the Results entity.
type Results struct {
	ent.Schema
}

// Fields of the Results.
func (Results) Fields() []ent.Field {
	return []ent.Field{
		field.Float("grade").Positive(),
	}
}

// Edges of the Results.
func (Results) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("YearID", Year.Type).Ref("year1").Unique(),
		edge.From("SubjectID", Subject.Type).Ref("subject1").Unique(),
		edge.From("StudentID", Student.Type).Ref("stud_resu").Unique(),
	}
}
