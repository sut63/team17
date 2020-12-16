package schema

import "github.com/facebookincubator/ent"

// Credit holds the schema definition for the Credit entity.
type Credit struct {
	ent.Schema
}

// Fields of the Credit.
func (Credit) Fields() []ent.Field {
	return []ent.Field{
		field.Int("creditpoint").NotEmpty(),
	}
}

// Edges of the Credit.
func (Credit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("credit1", Term.Type),
	}
}
