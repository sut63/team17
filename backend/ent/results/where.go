// Code generated by entc, DO NOT EDIT.

package results

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team17/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Grade applies equality check predicate on the "grade" field. It's identical to GradeEQ.
func Grade(v float64) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGrade), v))
	})
}

// GradeEQ applies the EQ predicate on the "grade" field.
func GradeEQ(v float64) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGrade), v))
	})
}

// GradeNEQ applies the NEQ predicate on the "grade" field.
func GradeNEQ(v float64) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGrade), v))
	})
}

// GradeIn applies the In predicate on the "grade" field.
func GradeIn(vs ...float64) predicate.Results {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Results(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGrade), v...))
	})
}

// GradeNotIn applies the NotIn predicate on the "grade" field.
func GradeNotIn(vs ...float64) predicate.Results {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Results(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGrade), v...))
	})
}

// GradeGT applies the GT predicate on the "grade" field.
func GradeGT(v float64) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGrade), v))
	})
}

// GradeGTE applies the GTE predicate on the "grade" field.
func GradeGTE(v float64) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGrade), v))
	})
}

// GradeLT applies the LT predicate on the "grade" field.
func GradeLT(v float64) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGrade), v))
	})
}

// GradeLTE applies the LTE predicate on the "grade" field.
func GradeLTE(v float64) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGrade), v))
	})
}

// HasResuYear applies the HasEdge predicate on the "resu_year" edge.
func HasResuYear() predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResuYearTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ResuYearTable, ResuYearColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResuYearWith applies the HasEdge predicate on the "resu_year" edge with a given conditions (other predicates).
func HasResuYearWith(preds ...predicate.Year) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResuYearInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ResuYearTable, ResuYearColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasResuSubj applies the HasEdge predicate on the "resu_subj" edge.
func HasResuSubj() predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResuSubjTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ResuSubjTable, ResuSubjColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResuSubjWith applies the HasEdge predicate on the "resu_subj" edge with a given conditions (other predicates).
func HasResuSubjWith(preds ...predicate.Subject) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResuSubjInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ResuSubjTable, ResuSubjColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasResuStud applies the HasEdge predicate on the "resu_stud" edge.
func HasResuStud() predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResuStudTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ResuStudTable, ResuStudColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResuStudWith applies the HasEdge predicate on the "resu_stud" edge with a given conditions (other predicates).
func HasResuStudWith(preds ...predicate.Student) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResuStudInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ResuStudTable, ResuStudColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Results) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Results) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Results) predicate.Results {
	return predicate.Results(func(s *sql.Selector) {
		p(s.Not())
	})
}
