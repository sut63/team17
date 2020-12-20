package schema

import "github.com/facebookincubator/ent"

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("student_id").NotEmpty(),
		field.String("fname").NotEmpty(),
		field.String("lname").NotEmpty(),
		field.String("institution").NotEmpty(),
		field.String("recent_address").NotEmpty(),
		field.Int("telephone").NotEmpty(),
		field.String("email").NotEmpty(),
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return nil
}
