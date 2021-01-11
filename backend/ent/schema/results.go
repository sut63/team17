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
		edge.From("resu_year", Year.Type).Ref("year_resu").Unique(),
		edge.From("resu_subj", Subject.Type).Ref("subj_resu").Unique(),
		edge.From("resu_stud", Student.Type).Ref("stud_resu").Unique(),
		edge.From("resu_term", Term.Type).Ref("term_resu").Unique(),
	}
}
