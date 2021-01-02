package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"

)

// Place holds the schema definition for the Place entity.
type Place struct {
	ent.Schema
}

// Fields of the Place.
func (Place) Fields() []ent.Field {
	return []ent.Field{

		field.String("PLACE").NotEmpty(),
	}
}

// Edges of the Place.
func (Place) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("place_acti", Activity.Type),
	}
}
