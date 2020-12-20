package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Activity holds the schema definition for the Activity entity.
type Activity struct {
	ent.Schema
}

// Fields of the Activity.
func (Activity) Fields() []ent.Field {
	return []ent.Field{
		field.String("ACTIVITYNAME"),
		field.Time("ADDED"),
	}
}

// Edges of the Activity.
func (Activity) Edges() []ent.Edge {
	return nil
}
