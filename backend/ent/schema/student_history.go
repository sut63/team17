package schema

import "github.com/facebookincubator/ent"

// Student_history holds the schema definition for the Student_history entity.
type Student_history struct {
	ent.Schema
}

// Fields of the Student_history.
func (Student_history) Fields() []ent.Field {
	return nil
}

// Edges of the Student_history.
func (Student_history) Edges() []ent.Edge {
	return nil
}
