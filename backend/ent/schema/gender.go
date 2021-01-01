package schema

import "github.com/facebookincubator/ent"
import "github.com/facebookincubator/ent/schema/field"
import "github.com/facebookincubator/ent/schema/edge"

// Gender holds the schema definition for the Gender entity.
type Gender struct {
	ent.Schema
}

// Fields of the Gender.
func (Gender) Fields() []ent.Field {
	return []ent.Field{
		field.String("gender").Unique(),
	}
}

// Edges of the Gender.
func (Gender) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("gend_stud", Student.Type),
	}
}
