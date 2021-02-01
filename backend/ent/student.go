// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team17/app/ent/degree"
	"github.com/sut63/team17/app/ent/gender"
	"github.com/sut63/team17/app/ent/prefix"
	"github.com/sut63/team17/app/ent/province"
	"github.com/sut63/team17/app/ent/student"
)

// Student is the model entity for the Student schema.
type Student struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Fname holds the value of the "fname" field.
	Fname string `json:"fname,omitempty"`
	// Lname holds the value of the "lname" field.
	Lname string `json:"lname,omitempty"`
	// Schoolname holds the value of the "schoolname" field.
	Schoolname string `json:"schoolname,omitempty"`
	// RecentAddress holds the value of the "recent_address" field.
	RecentAddress string `json:"recent_address,omitempty"`
	// Telephone holds the value of the "telephone" field.
	Telephone string `json:"telephone,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudentQuery when eager-loading is set.
	Edges              StudentEdges `json:"edges"`
	degree_degr_stud   *int
	gender_gend_stud   *int
	prefix_pref_stud   *int
	province_prov_stud *int
	province_dist_stud *int
	province_subd_stud *int
	province_post_stud *int
}

// StudentEdges holds the relations/edges for other nodes in the graph.
type StudentEdges struct {
	// StudGend holds the value of the stud_gend edge.
	StudGend *Gender
	// StudActi holds the value of the stud_acti edge.
	StudActi []*Activity
	// StudResu holds the value of the stud_resu edge.
	StudResu []*Results
	// StudProv holds the value of the stud_prov edge.
	StudProv *Province
	// StudDist holds the value of the stud_dist edge.
	StudDist *Province
	// StudSubd holds the value of the stud_subd edge.
	StudSubd *Province
	// StudPost holds the value of the stud_post edge.
	StudPost *Province
	// StudPref holds the value of the stud_pref edge.
	StudPref *Prefix
	// StudDegr holds the value of the stud_degr edge.
	StudDegr *Degree
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
}

// StudGendOrErr returns the StudGend value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentEdges) StudGendOrErr() (*Gender, error) {
	if e.loadedTypes[0] {
		if e.StudGend == nil {
			// The edge stud_gend was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: gender.Label}
		}
		return e.StudGend, nil
	}
	return nil, &NotLoadedError{edge: "stud_gend"}
}

// StudActiOrErr returns the StudActi value or an error if the edge
// was not loaded in eager-loading.
func (e StudentEdges) StudActiOrErr() ([]*Activity, error) {
	if e.loadedTypes[1] {
		return e.StudActi, nil
	}
	return nil, &NotLoadedError{edge: "stud_acti"}
}

// StudResuOrErr returns the StudResu value or an error if the edge
// was not loaded in eager-loading.
func (e StudentEdges) StudResuOrErr() ([]*Results, error) {
	if e.loadedTypes[2] {
		return e.StudResu, nil
	}
	return nil, &NotLoadedError{edge: "stud_resu"}
}

// StudProvOrErr returns the StudProv value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentEdges) StudProvOrErr() (*Province, error) {
	if e.loadedTypes[3] {
		if e.StudProv == nil {
			// The edge stud_prov was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: province.Label}
		}
		return e.StudProv, nil
	}
	return nil, &NotLoadedError{edge: "stud_prov"}
}

// StudDistOrErr returns the StudDist value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentEdges) StudDistOrErr() (*Province, error) {
	if e.loadedTypes[4] {
		if e.StudDist == nil {
			// The edge stud_dist was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: province.Label}
		}
		return e.StudDist, nil
	}
	return nil, &NotLoadedError{edge: "stud_dist"}
}

// StudSubdOrErr returns the StudSubd value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentEdges) StudSubdOrErr() (*Province, error) {
	if e.loadedTypes[5] {
		if e.StudSubd == nil {
			// The edge stud_subd was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: province.Label}
		}
		return e.StudSubd, nil
	}
	return nil, &NotLoadedError{edge: "stud_subd"}
}

// StudPostOrErr returns the StudPost value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentEdges) StudPostOrErr() (*Province, error) {
	if e.loadedTypes[6] {
		if e.StudPost == nil {
			// The edge stud_post was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: province.Label}
		}
		return e.StudPost, nil
	}
	return nil, &NotLoadedError{edge: "stud_post"}
}

// StudPrefOrErr returns the StudPref value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentEdges) StudPrefOrErr() (*Prefix, error) {
	if e.loadedTypes[7] {
		if e.StudPref == nil {
			// The edge stud_pref was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: prefix.Label}
		}
		return e.StudPref, nil
	}
	return nil, &NotLoadedError{edge: "stud_pref"}
}

// StudDegrOrErr returns the StudDegr value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentEdges) StudDegrOrErr() (*Degree, error) {
	if e.loadedTypes[8] {
		if e.StudDegr == nil {
			// The edge stud_degr was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: degree.Label}
		}
		return e.StudDegr, nil
	}
	return nil, &NotLoadedError{edge: "stud_degr"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Student) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // fname
		&sql.NullString{}, // lname
		&sql.NullString{}, // schoolname
		&sql.NullString{}, // recent_address
		&sql.NullString{}, // telephone
		&sql.NullString{}, // email
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Student) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // degree_degr_stud
		&sql.NullInt64{}, // gender_gend_stud
		&sql.NullInt64{}, // prefix_pref_stud
		&sql.NullInt64{}, // province_prov_stud
		&sql.NullInt64{}, // province_dist_stud
		&sql.NullInt64{}, // province_subd_stud
		&sql.NullInt64{}, // province_post_stud
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Student fields.
func (s *Student) assignValues(values ...interface{}) error {
	if m, n := len(values), len(student.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field fname", values[0])
	} else if value.Valid {
		s.Fname = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field lname", values[1])
	} else if value.Valid {
		s.Lname = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field schoolname", values[2])
	} else if value.Valid {
		s.Schoolname = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field recent_address", values[3])
	} else if value.Valid {
		s.RecentAddress = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field telephone", values[4])
	} else if value.Valid {
		s.Telephone = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[5])
	} else if value.Valid {
		s.Email = value.String
	}
	values = values[6:]
	if len(values) == len(student.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field degree_degr_stud", value)
		} else if value.Valid {
			s.degree_degr_stud = new(int)
			*s.degree_degr_stud = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field gender_gend_stud", value)
		} else if value.Valid {
			s.gender_gend_stud = new(int)
			*s.gender_gend_stud = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field prefix_pref_stud", value)
		} else if value.Valid {
			s.prefix_pref_stud = new(int)
			*s.prefix_pref_stud = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field province_prov_stud", value)
		} else if value.Valid {
			s.province_prov_stud = new(int)
			*s.province_prov_stud = int(value.Int64)
		}
		if value, ok := values[4].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field province_dist_stud", value)
		} else if value.Valid {
			s.province_dist_stud = new(int)
			*s.province_dist_stud = int(value.Int64)
		}
		if value, ok := values[5].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field province_subd_stud", value)
		} else if value.Valid {
			s.province_subd_stud = new(int)
			*s.province_subd_stud = int(value.Int64)
		}
		if value, ok := values[6].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field province_post_stud", value)
		} else if value.Valid {
			s.province_post_stud = new(int)
			*s.province_post_stud = int(value.Int64)
		}
	}
	return nil
}

// QueryStudGend queries the stud_gend edge of the Student.
func (s *Student) QueryStudGend() *GenderQuery {
	return (&StudentClient{config: s.config}).QueryStudGend(s)
}

// QueryStudActi queries the stud_acti edge of the Student.
func (s *Student) QueryStudActi() *ActivityQuery {
	return (&StudentClient{config: s.config}).QueryStudActi(s)
}

// QueryStudResu queries the stud_resu edge of the Student.
func (s *Student) QueryStudResu() *ResultsQuery {
	return (&StudentClient{config: s.config}).QueryStudResu(s)
}

// QueryStudProv queries the stud_prov edge of the Student.
func (s *Student) QueryStudProv() *ProvinceQuery {
	return (&StudentClient{config: s.config}).QueryStudProv(s)
}

// QueryStudDist queries the stud_dist edge of the Student.
func (s *Student) QueryStudDist() *ProvinceQuery {
	return (&StudentClient{config: s.config}).QueryStudDist(s)
}

// QueryStudSubd queries the stud_subd edge of the Student.
func (s *Student) QueryStudSubd() *ProvinceQuery {
	return (&StudentClient{config: s.config}).QueryStudSubd(s)
}

// QueryStudPost queries the stud_post edge of the Student.
func (s *Student) QueryStudPost() *ProvinceQuery {
	return (&StudentClient{config: s.config}).QueryStudPost(s)
}

// QueryStudPref queries the stud_pref edge of the Student.
func (s *Student) QueryStudPref() *PrefixQuery {
	return (&StudentClient{config: s.config}).QueryStudPref(s)
}

// QueryStudDegr queries the stud_degr edge of the Student.
func (s *Student) QueryStudDegr() *DegreeQuery {
	return (&StudentClient{config: s.config}).QueryStudDegr(s)
}

// Update returns a builder for updating this Student.
// Note that, you need to call Student.Unwrap() before calling this method, if this Student
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Student) Update() *StudentUpdateOne {
	return (&StudentClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Student) Unwrap() *Student {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Student is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Student) String() string {
	var builder strings.Builder
	builder.WriteString("Student(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", fname=")
	builder.WriteString(s.Fname)
	builder.WriteString(", lname=")
	builder.WriteString(s.Lname)
	builder.WriteString(", schoolname=")
	builder.WriteString(s.Schoolname)
	builder.WriteString(", recent_address=")
	builder.WriteString(s.RecentAddress)
	builder.WriteString(", telephone=")
	builder.WriteString(s.Telephone)
	builder.WriteString(", email=")
	builder.WriteString(s.Email)
	builder.WriteByte(')')
	return builder.String()
}

// Students is a parsable slice of Student.
type Students []*Student

func (s Students) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
