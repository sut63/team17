package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// District holds the schema definition for the District entity.
type District struct {
	ent.Schema
}

// Fields of the District.
func (District) Fields() []ent.Field {
	return []ent.Field{
		field.String("district").NotEmpty(),
	}
}

// Edges of the District.
func (District) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dist_subd", Subdistrict.Type),
		edge.To("dist_post", Postal_number.Type),
		edge.From("dist_prov", Province.Type).Ref("prov_dist").Unique(),
	}
}
