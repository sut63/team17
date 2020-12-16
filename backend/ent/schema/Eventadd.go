package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)
// Eventadd holds the schema definition for the Eventadd entity.
type Eventadd struct {
	ent.Schema
}
// Fields of the Eventadd.
func (Eventadd) Fields() []ent.Field {
	return []ent.Field{
		field.String("EVENTNAME"),
		field.Time("ADDED"),

		
	}
}
// Edges of the Eventadd.
func (Eventadd) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Institution", Institution.Type).Ref("eventadd1").Unique(),
		edge.From("Place", Place.Type).Ref("eventadd2").Unique(),
		edge.From("User", User.Type).Ref("eventadd3").Unique(),
	}
}