package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Term holds the schema definition for the Term entity.
type Term struct {
	ent.Schema
}

// Fields of the Term.
func (Term) Fields() []ent.Field {
	return []ent.Field{
		field.Int("semester").Positive(),
	}
}

// Edges of the Term.
func (Term) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("term_resu", Results.Type),
		edge.To("term_acti", Activity.Type),
	}
}
