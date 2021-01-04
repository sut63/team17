package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Province holds the schema definition for the Province entity.
type Province struct {
	ent.Schema
}

// Fields of the Province.
func (Province) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the Province.
func (Province) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("prov_regi", Region.Type).Ref("regi_prov").Unique(),
		edge.From("prov_dist", District.Type).Ref("dist_prov").Unique(),
		edge.To("prov_stud", Student.Type),
	}
}
