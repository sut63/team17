package schema

import "github.com/facebookincubator/ent"
import "github.com/facebookincubator/ent/schema/field"
import "github.com/facebookincubator/ent/schema/edge"
import (
	"regexp"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("fname").NotEmpty().Match(regexp.MustCompile("[A-Za-z]")),
		field.String("lname").NotEmpty().Match(regexp.MustCompile("[A-Za-z]")),
		field.String("schoolname").NotEmpty().Match(regexp.MustCompile("[A-Za-z]")),
		field.String("recent_address").NotEmpty().Match(regexp.MustCompile("[A-Za-z0-9\\/]")),
		field.Int("telephone"),
		field.String("email").Match(regexp.MustCompile("[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}")),
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("stud_gend", Gender.Type).Ref("gend_stud").Unique(),
		edge.To("stud_acti", Activity.Type),
		edge.To("stud_resu", Results.Type),
		edge.From("stud_prov", Province.Type).Ref("prov_stud").Unique(),
		edge.From("stud_dist", Province.Type).Ref("dist_stud").Unique(),
		edge.From("stud_subd", Province.Type).Ref("subd_stud").Unique(),
		edge.From("stud_post", Province.Type).Ref("post_stud").Unique(),
		edge.From("stud_pref", Prefix.Type).Ref("pref_stud").Unique(),
		edge.From("stud_degr", Degree.Type).Ref("degr_stud").Unique(),
	}
}
