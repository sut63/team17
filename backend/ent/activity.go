// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team17/app/ent/activity"
	"github.com/sut63/team17/app/ent/agency"
	"github.com/sut63/team17/app/ent/place"
	"github.com/sut63/team17/app/ent/student"
	"github.com/sut63/team17/app/ent/term"
	"github.com/sut63/team17/app/ent/year"
)

// Activity is the model entity for the Activity schema.
type Activity struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ACTIVITYNAME holds the value of the "ACTIVITYNAME" field.
	ACTIVITYNAME string `json:"ACTIVITYNAME,omitempty"`
	// Added holds the value of the "added" field.
	Added time.Time `json:"added,omitempty"`
	// Hours holds the value of the "hours" field.
	Hours string `json:"hours,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ActivityQuery when eager-loading is set.
	Edges             ActivityEdges `json:"edges"`
	agency_agen_acti  *int
	place_place_acti  *int
	student_stud_acti *int
	term_term_acti    *int
	year_year_acti    *int
}

// ActivityEdges holds the relations/edges for other nodes in the graph.
type ActivityEdges struct {
	// ActiStud holds the value of the acti_stud edge.
	ActiStud *Student
	// ActiPlace holds the value of the acti_place edge.
	ActiPlace *Place
	// ActiAgen holds the value of the acti_agen edge.
	ActiAgen *Agency
	// ActiYear holds the value of the acti_year edge.
	ActiYear *Year
	// ActiTerm holds the value of the acti_term edge.
	ActiTerm *Term
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// ActiStudOrErr returns the ActiStud value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ActivityEdges) ActiStudOrErr() (*Student, error) {
	if e.loadedTypes[0] {
		if e.ActiStud == nil {
			// The edge acti_stud was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: student.Label}
		}
		return e.ActiStud, nil
	}
	return nil, &NotLoadedError{edge: "acti_stud"}
}

// ActiPlaceOrErr returns the ActiPlace value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ActivityEdges) ActiPlaceOrErr() (*Place, error) {
	if e.loadedTypes[1] {
		if e.ActiPlace == nil {
			// The edge acti_place was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: place.Label}
		}
		return e.ActiPlace, nil
	}
	return nil, &NotLoadedError{edge: "acti_place"}
}

// ActiAgenOrErr returns the ActiAgen value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ActivityEdges) ActiAgenOrErr() (*Agency, error) {
	if e.loadedTypes[2] {
		if e.ActiAgen == nil {
			// The edge acti_agen was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: agency.Label}
		}
		return e.ActiAgen, nil
	}
	return nil, &NotLoadedError{edge: "acti_agen"}
}

// ActiYearOrErr returns the ActiYear value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ActivityEdges) ActiYearOrErr() (*Year, error) {
	if e.loadedTypes[3] {
		if e.ActiYear == nil {
			// The edge acti_year was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: year.Label}
		}
		return e.ActiYear, nil
	}
	return nil, &NotLoadedError{edge: "acti_year"}
}

// ActiTermOrErr returns the ActiTerm value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ActivityEdges) ActiTermOrErr() (*Term, error) {
	if e.loadedTypes[4] {
		if e.ActiTerm == nil {
			// The edge acti_term was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: term.Label}
		}
		return e.ActiTerm, nil
	}
	return nil, &NotLoadedError{edge: "acti_term"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Activity) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // ACTIVITYNAME
		&sql.NullTime{},   // added
		&sql.NullString{}, // hours
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Activity) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // agency_agen_acti
		&sql.NullInt64{}, // place_place_acti
		&sql.NullInt64{}, // student_stud_acti
		&sql.NullInt64{}, // term_term_acti
		&sql.NullInt64{}, // year_year_acti
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Activity fields.
func (a *Activity) assignValues(values ...interface{}) error {
	if m, n := len(values), len(activity.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	a.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field ACTIVITYNAME", values[0])
	} else if value.Valid {
		a.ACTIVITYNAME = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field added", values[1])
	} else if value.Valid {
		a.Added = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field hours", values[2])
	} else if value.Valid {
		a.Hours = value.String
	}
	values = values[3:]
	if len(values) == len(activity.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field agency_agen_acti", value)
		} else if value.Valid {
			a.agency_agen_acti = new(int)
			*a.agency_agen_acti = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field place_place_acti", value)
		} else if value.Valid {
			a.place_place_acti = new(int)
			*a.place_place_acti = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field student_stud_acti", value)
		} else if value.Valid {
			a.student_stud_acti = new(int)
			*a.student_stud_acti = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field term_term_acti", value)
		} else if value.Valid {
			a.term_term_acti = new(int)
			*a.term_term_acti = int(value.Int64)
		}
		if value, ok := values[4].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field year_year_acti", value)
		} else if value.Valid {
			a.year_year_acti = new(int)
			*a.year_year_acti = int(value.Int64)
		}
	}
	return nil
}

// QueryActiStud queries the acti_stud edge of the Activity.
func (a *Activity) QueryActiStud() *StudentQuery {
	return (&ActivityClient{config: a.config}).QueryActiStud(a)
}

// QueryActiPlace queries the acti_place edge of the Activity.
func (a *Activity) QueryActiPlace() *PlaceQuery {
	return (&ActivityClient{config: a.config}).QueryActiPlace(a)
}

// QueryActiAgen queries the acti_agen edge of the Activity.
func (a *Activity) QueryActiAgen() *AgencyQuery {
	return (&ActivityClient{config: a.config}).QueryActiAgen(a)
}

// QueryActiYear queries the acti_year edge of the Activity.
func (a *Activity) QueryActiYear() *YearQuery {
	return (&ActivityClient{config: a.config}).QueryActiYear(a)
}

// QueryActiTerm queries the acti_term edge of the Activity.
func (a *Activity) QueryActiTerm() *TermQuery {
	return (&ActivityClient{config: a.config}).QueryActiTerm(a)
}

// Update returns a builder for updating this Activity.
// Note that, you need to call Activity.Unwrap() before calling this method, if this Activity
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Activity) Update() *ActivityUpdateOne {
	return (&ActivityClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Activity) Unwrap() *Activity {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Activity is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Activity) String() string {
	var builder strings.Builder
	builder.WriteString("Activity(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", ACTIVITYNAME=")
	builder.WriteString(a.ACTIVITYNAME)
	builder.WriteString(", added=")
	builder.WriteString(a.Added.Format(time.ANSIC))
	builder.WriteString(", hours=")
	builder.WriteString(a.Hours)
	builder.WriteByte(')')
	return builder.String()
}

// Activities is a parsable slice of Activity.
type Activities []*Activity

func (a Activities) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
