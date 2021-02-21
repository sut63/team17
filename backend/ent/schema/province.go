package schema

import (
	"regexp"

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
		field.String("province").NotEmpty(),
		field.String("district").NotEmpty(),
		field.String("subdistrict").NotEmpty(),
		field.String("postal").Match(regexp.MustCompile("[0-9]")).MaxLen(5).MinLen(5),
	}
}

// Edges of the Province.
func (Province) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("prov_regi", Region.Type).Ref("regi_prov").Unique(),
		edge.From("prov_coun", Country.Type).Ref("coun_prov").Unique(),
		edge.From("prov_cont", Continent.Type).Ref("cont_prov").Unique(),
		edge.To("prov_stud", Student.Type),
		edge.To("dist_stud", Student.Type),
		edge.To("subd_stud", Student.Type),
		edge.To("post_stud", Student.Type),
	}
}
