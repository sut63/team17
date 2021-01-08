// Code generated by entc, DO NOT EDIT.

package student

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team17/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Fname applies equality check predicate on the "fname" field. It's identical to FnameEQ.
func Fname(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFname), v))
	})
}

// Lname applies equality check predicate on the "lname" field. It's identical to LnameEQ.
func Lname(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLname), v))
	})
}

// Schoolname applies equality check predicate on the "schoolname" field. It's identical to SchoolnameEQ.
func Schoolname(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSchoolname), v))
	})
}

// RecentAddress applies equality check predicate on the "recent_address" field. It's identical to RecentAddressEQ.
func RecentAddress(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecentAddress), v))
	})
}

// Telephone applies equality check predicate on the "telephone" field. It's identical to TelephoneEQ.
func Telephone(v int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTelephone), v))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// FnameEQ applies the EQ predicate on the "fname" field.
func FnameEQ(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFname), v))
	})
}

// FnameNEQ applies the NEQ predicate on the "fname" field.
func FnameNEQ(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFname), v))
	})
}

// FnameIn applies the In predicate on the "fname" field.
func FnameIn(vs ...string) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFname), v...))
	})
}

// FnameNotIn applies the NotIn predicate on the "fname" field.
func FnameNotIn(vs ...string) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFname), v...))
	})
}

// FnameGT applies the GT predicate on the "fname" field.
func FnameGT(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFname), v))
	})
}

// FnameGTE applies the GTE predicate on the "fname" field.
func FnameGTE(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFname), v))
	})
}

// FnameLT applies the LT predicate on the "fname" field.
func FnameLT(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFname), v))
	})
}

// FnameLTE applies the LTE predicate on the "fname" field.
func FnameLTE(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFname), v))
	})
}

// FnameContains applies the Contains predicate on the "fname" field.
func FnameContains(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFname), v))
	})
}

// FnameHasPrefix applies the HasPrefix predicate on the "fname" field.
func FnameHasPrefix(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFname), v))
	})
}

// FnameHasSuffix applies the HasSuffix predicate on the "fname" field.
func FnameHasSuffix(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFname), v))
	})
}

// FnameEqualFold applies the EqualFold predicate on the "fname" field.
func FnameEqualFold(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFname), v))
	})
}

// FnameContainsFold applies the ContainsFold predicate on the "fname" field.
func FnameContainsFold(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFname), v))
	})
}

// LnameEQ applies the EQ predicate on the "lname" field.
func LnameEQ(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLname), v))
	})
}

// LnameNEQ applies the NEQ predicate on the "lname" field.
func LnameNEQ(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLname), v))
	})
}

// LnameIn applies the In predicate on the "lname" field.
func LnameIn(vs ...string) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLname), v...))
	})
}

// LnameNotIn applies the NotIn predicate on the "lname" field.
func LnameNotIn(vs ...string) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLname), v...))
	})
}

// LnameGT applies the GT predicate on the "lname" field.
func LnameGT(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLname), v))
	})
}

// LnameGTE applies the GTE predicate on the "lname" field.
func LnameGTE(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLname), v))
	})
}

// LnameLT applies the LT predicate on the "lname" field.
func LnameLT(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLname), v))
	})
}

// LnameLTE applies the LTE predicate on the "lname" field.
func LnameLTE(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLname), v))
	})
}

// LnameContains applies the Contains predicate on the "lname" field.
func LnameContains(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLname), v))
	})
}

// LnameHasPrefix applies the HasPrefix predicate on the "lname" field.
func LnameHasPrefix(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLname), v))
	})
}

// LnameHasSuffix applies the HasSuffix predicate on the "lname" field.
func LnameHasSuffix(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLname), v))
	})
}

// LnameEqualFold applies the EqualFold predicate on the "lname" field.
func LnameEqualFold(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLname), v))
	})
}

// LnameContainsFold applies the ContainsFold predicate on the "lname" field.
func LnameContainsFold(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLname), v))
	})
}

// SchoolnameEQ applies the EQ predicate on the "schoolname" field.
func SchoolnameEQ(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSchoolname), v))
	})
}

// SchoolnameNEQ applies the NEQ predicate on the "schoolname" field.
func SchoolnameNEQ(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSchoolname), v))
	})
}

// SchoolnameIn applies the In predicate on the "schoolname" field.
func SchoolnameIn(vs ...string) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSchoolname), v...))
	})
}

// SchoolnameNotIn applies the NotIn predicate on the "schoolname" field.
func SchoolnameNotIn(vs ...string) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSchoolname), v...))
	})
}

// SchoolnameGT applies the GT predicate on the "schoolname" field.
func SchoolnameGT(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSchoolname), v))
	})
}

// SchoolnameGTE applies the GTE predicate on the "schoolname" field.
func SchoolnameGTE(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSchoolname), v))
	})
}

// SchoolnameLT applies the LT predicate on the "schoolname" field.
func SchoolnameLT(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSchoolname), v))
	})
}

// SchoolnameLTE applies the LTE predicate on the "schoolname" field.
func SchoolnameLTE(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSchoolname), v))
	})
}

// SchoolnameContains applies the Contains predicate on the "schoolname" field.
func SchoolnameContains(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSchoolname), v))
	})
}

// SchoolnameHasPrefix applies the HasPrefix predicate on the "schoolname" field.
func SchoolnameHasPrefix(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSchoolname), v))
	})
}

// SchoolnameHasSuffix applies the HasSuffix predicate on the "schoolname" field.
func SchoolnameHasSuffix(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSchoolname), v))
	})
}

// SchoolnameEqualFold applies the EqualFold predicate on the "schoolname" field.
func SchoolnameEqualFold(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSchoolname), v))
	})
}

// SchoolnameContainsFold applies the ContainsFold predicate on the "schoolname" field.
func SchoolnameContainsFold(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSchoolname), v))
	})
}

// RecentAddressEQ applies the EQ predicate on the "recent_address" field.
func RecentAddressEQ(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecentAddress), v))
	})
}

// RecentAddressNEQ applies the NEQ predicate on the "recent_address" field.
func RecentAddressNEQ(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRecentAddress), v))
	})
}

// RecentAddressIn applies the In predicate on the "recent_address" field.
func RecentAddressIn(vs ...string) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRecentAddress), v...))
	})
}

// RecentAddressNotIn applies the NotIn predicate on the "recent_address" field.
func RecentAddressNotIn(vs ...string) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRecentAddress), v...))
	})
}

// RecentAddressGT applies the GT predicate on the "recent_address" field.
func RecentAddressGT(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRecentAddress), v))
	})
}

// RecentAddressGTE applies the GTE predicate on the "recent_address" field.
func RecentAddressGTE(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRecentAddress), v))
	})
}

// RecentAddressLT applies the LT predicate on the "recent_address" field.
func RecentAddressLT(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRecentAddress), v))
	})
}

// RecentAddressLTE applies the LTE predicate on the "recent_address" field.
func RecentAddressLTE(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRecentAddress), v))
	})
}

// RecentAddressContains applies the Contains predicate on the "recent_address" field.
func RecentAddressContains(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRecentAddress), v))
	})
}

// RecentAddressHasPrefix applies the HasPrefix predicate on the "recent_address" field.
func RecentAddressHasPrefix(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRecentAddress), v))
	})
}

// RecentAddressHasSuffix applies the HasSuffix predicate on the "recent_address" field.
func RecentAddressHasSuffix(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRecentAddress), v))
	})
}

// RecentAddressEqualFold applies the EqualFold predicate on the "recent_address" field.
func RecentAddressEqualFold(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRecentAddress), v))
	})
}

// RecentAddressContainsFold applies the ContainsFold predicate on the "recent_address" field.
func RecentAddressContainsFold(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRecentAddress), v))
	})
}

// TelephoneEQ applies the EQ predicate on the "telephone" field.
func TelephoneEQ(v int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTelephone), v))
	})
}

// TelephoneNEQ applies the NEQ predicate on the "telephone" field.
func TelephoneNEQ(v int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTelephone), v))
	})
}

// TelephoneIn applies the In predicate on the "telephone" field.
func TelephoneIn(vs ...int) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTelephone), v...))
	})
}

// TelephoneNotIn applies the NotIn predicate on the "telephone" field.
func TelephoneNotIn(vs ...int) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTelephone), v...))
	})
}

// TelephoneGT applies the GT predicate on the "telephone" field.
func TelephoneGT(v int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTelephone), v))
	})
}

// TelephoneGTE applies the GTE predicate on the "telephone" field.
func TelephoneGTE(v int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTelephone), v))
	})
}

// TelephoneLT applies the LT predicate on the "telephone" field.
func TelephoneLT(v int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTelephone), v))
	})
}

// TelephoneLTE applies the LTE predicate on the "telephone" field.
func TelephoneLTE(v int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTelephone), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// HasStudGend applies the HasEdge predicate on the "stud_gend" edge.
func HasStudGend() predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudGendTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudGendTable, StudGendColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudGendWith applies the HasEdge predicate on the "stud_gend" edge with a given conditions (other predicates).
func HasStudGendWith(preds ...predicate.Gender) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudGendInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudGendTable, StudGendColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStudActi applies the HasEdge predicate on the "stud_acti" edge.
func HasStudActi() predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudActiTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StudActiTable, StudActiColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudActiWith applies the HasEdge predicate on the "stud_acti" edge with a given conditions (other predicates).
func HasStudActiWith(preds ...predicate.Activity) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudActiInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StudActiTable, StudActiColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStudProv applies the HasEdge predicate on the "stud_prov" edge.
func HasStudProv() predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudProvTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudProvTable, StudProvColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudProvWith applies the HasEdge predicate on the "stud_prov" edge with a given conditions (other predicates).
func HasStudProvWith(preds ...predicate.Province) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudProvInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudProvTable, StudProvColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStudResu applies the HasEdge predicate on the "stud_resu" edge.
func HasStudResu() predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudResuTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StudResuTable, StudResuColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudResuWith applies the HasEdge predicate on the "stud_resu" edge with a given conditions (other predicates).
func HasStudResuWith(preds ...predicate.Results) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudResuInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StudResuTable, StudResuColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStudPref applies the HasEdge predicate on the "stud_pref" edge.
func HasStudPref() predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudPrefTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudPrefTable, StudPrefColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudPrefWith applies the HasEdge predicate on the "stud_pref" edge with a given conditions (other predicates).
func HasStudPrefWith(preds ...predicate.Prefix) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudPrefInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudPrefTable, StudPrefColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStudDegr applies the HasEdge predicate on the "stud_degr" edge.
func HasStudDegr() predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudDegrTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudDegrTable, StudDegrColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudDegrWith applies the HasEdge predicate on the "stud_degr" edge with a given conditions (other predicates).
func HasStudDegrWith(preds ...predicate.Degree) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudDegrInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudDegrTable, StudDegrColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Student) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Student) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
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
func Not(p predicate.Student) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		p(s.Not())
	})
}
