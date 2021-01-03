package schema

import ("github.com/facebookincubator/ent"
		"github.com/facebookincubator/ent/schema/edge"
		"github.com/facebookincubator/ent/schema/field"
)

// Prefix holds the schema definition for the Prefix entity.
type Prefix struct {
	ent.Schema
}

// Fields of the Prefix.
func (Prefix) Fields() []ent.Field {
	return []ent.Field{
        field.String("prefix").NotEmpty(),
    }
}

// Edges of the Prefix.
func (Prefix) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pre_prof", Professor.Type).StorageKey(edge.Column("prefix_ID")),
		edge.To("pref_stud", Student.Type),
	}

}