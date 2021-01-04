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
		edge.From("dist_subd", Subdistrict.Type).Ref("subd_dist").Unique(),
		edge.From("dist_post", Postal_number.Type).Ref("post_dist").Unique(),
		edge.To("dist_prov", Province.Type),
	}
}
