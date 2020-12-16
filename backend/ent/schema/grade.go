package schema

import "github.com/facebookincubator/ent"

// Grade holds the schema definition for the Grade entity.
type Grade struct {
	ent.Schema
}

// Fields of the Grade.
func (Grade) Fields() []ent.Field {
	return []ent.Field{
		field.Float("grade").NotEmpty(),
		field.String("gradesymbol").Positive(),
	}
}

// Edges of the Grade.
func (Grade) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("grade1", Term.Type),
	}
}
