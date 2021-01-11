package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Year holds the schema definition for the Year entity.
type Year struct {
	ent.Schema
}

// Fields of the Year.
func (Year) Fields() []ent.Field {
	return []ent.Field{
		field.Int("years").Positive(),
	}
}

// Edges of the Year.
func (Year) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("year_term", Term.Type),
		edge.To("year_resu", Results.Type),
		edge.To("year_acti", Activity.Type),
	}
}
