package schema

import "github.com/facebookincubator/ent"
import "github.com/facebookincubator/ent/schema/field"
import "github.com/facebookincubator/ent/schema/edge"

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("fname").NotEmpty(),
		field.String("lname").NotEmpty(),
		field.String("institution").NotEmpty(),
		field.String("recent_address").NotEmpty(),
		field.Int("telephone"),
		field.String("email"),
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("stud_gend", Gender.Type).Ref("gend_stud").Unique(),
		edge.To("stud_acti", Activity.Type),
		edge.From("stud_prov", Province.Type).Ref("prov_stud").Unique(),
		edge.To("stud_resu", Result.Type),
		edge.From("stud_pref", Prefix.Type).Ref("pref_stud").Unique(),
		edge.From("stud_degr", Degree.Type).Ref("degr_stud").Unique(),
	}
}