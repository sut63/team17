package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Agency holds the schema definition for the Agency entity.
type Agency struct {
	ent.Schema
}

// Fields of the Agency.
func (Agency) Fields() []ent.Field {
	return []ent.Field{

		field.String("AGENCY").NotEmpty(),
	}
}

// Edges of the Agency.
func (Agency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("agen_acti", Activity.Type),
	}
}
