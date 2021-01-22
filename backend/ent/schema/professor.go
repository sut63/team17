package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"regexp"
)

// Professor holds the schema definition for the Professor entity.
type Professor struct {
	ent.Schema
}

// Fields of the Professor.
func (Professor) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("tel").MaxLen(10).MinLen(10),
		field.String("email").Match(regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")),
	}
}

// Edges of the Professor.
func (Professor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("prof_pre", Prefix.Type).
			Ref("pre_prof").Unique(),
		edge.From("prof_facu", Faculty.Type).
			Ref("facu_prof").Unique(),
		edge.From("prof_pros", Professorship.Type).
			Ref("pros_prof").Unique(),
	}
}