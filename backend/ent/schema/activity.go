package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Activity holds the schema definition for the Activity entity.
type Activity struct {
	ent.Schema
}

// Fields of the Activity.
func (Activity) Fields() []ent.Field {
	return []ent.Field{
		field.String("ACTIVITYNAME").MinLen(5),
		field.Time("added"),
		field.Int("hours").Min(1).Positive(),
	}
}

// Edges of the Activity.
func (Activity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("acti_stud", Student.Type).Ref("stud_acti").Unique(),
		edge.From("acti_place", Place.Type).Ref("place_acti").Unique(),
		edge.From("acti_agen", Agency.Type).Ref("agen_acti").Unique(),
		edge.From("acti_year", Year.Type).Ref("year_acti").Unique(),
		edge.From("acti_term", Term.Type).Ref("term_acti").Unique(),
	}
}
