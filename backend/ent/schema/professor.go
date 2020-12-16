package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Professor holds the schema definition for the Professor entity.
type Professor struct {
	ent.Schema
}

// Fields of the Professor.
func (Professor) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the Professor.
func (Professor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("prefix", Prefix.Type).
			Ref("professors").Unique(),
		edge.From("faculty", Faculty.Type).
			Ref("professors").Unique(),
		edge.From("professorship", Professorship.Type).
			Ref("professors").Unique(),
	}
}