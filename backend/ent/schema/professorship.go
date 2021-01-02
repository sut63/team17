package schema

import ("github.com/facebookincubator/ent"
		"github.com/facebookincubator/ent/schema/edge"
		"github.com/facebookincubator/ent/schema/field"
)

// Professorship holds the schema definition for the Professorship entity.
type Professorship struct {
	ent.Schema
}

// Fields of the Professorship.
func (Professorship) Fields() []ent.Field {
	return []ent.Field{
        field.String("professorship").NotEmpty(),
    }
}

// Edges of the Professorship.
func (Professorship) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pros_prof", Professor.Type).StorageKey(edge.Column("professorship_ID")),
	}
}