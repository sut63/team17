// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team17/app/ent/subdistrict"
)

// Subdistrict is the model entity for the Subdistrict schema.
type Subdistrict struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Subdistrict holds the value of the "subdistrict" field.
	Subdistrict string `json:"subdistrict,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubdistrictQuery when eager-loading is set.
	Edges SubdistrictEdges `json:"edges"`
}

// SubdistrictEdges holds the relations/edges for other nodes in the graph.
type SubdistrictEdges struct {
	// SubdDist holds the value of the subd_dist edge.
	SubdDist []*District
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SubdDistOrErr returns the SubdDist value or an error if the edge
// was not loaded in eager-loading.
func (e SubdistrictEdges) SubdDistOrErr() ([]*District, error) {
	if e.loadedTypes[0] {
		return e.SubdDist, nil
	}
	return nil, &NotLoadedError{edge: "subd_dist"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subdistrict) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // subdistrict
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subdistrict fields.
func (s *Subdistrict) assignValues(values ...interface{}) error {
	if m, n := len(values), len(subdistrict.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field subdistrict", values[0])
	} else if value.Valid {
		s.Subdistrict = value.String
	}
	return nil
}

// QuerySubdDist queries the subd_dist edge of the Subdistrict.
func (s *Subdistrict) QuerySubdDist() *DistrictQuery {
	return (&SubdistrictClient{config: s.config}).QuerySubdDist(s)
}

// Update returns a builder for updating this Subdistrict.
// Note that, you need to call Subdistrict.Unwrap() before calling this method, if this Subdistrict
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subdistrict) Update() *SubdistrictUpdateOne {
	return (&SubdistrictClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Subdistrict) Unwrap() *Subdistrict {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subdistrict is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subdistrict) String() string {
	var builder strings.Builder
	builder.WriteString("Subdistrict(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", subdistrict=")
	builder.WriteString(s.Subdistrict)
	builder.WriteByte(')')
	return builder.String()
}

// Subdistricts is a parsable slice of Subdistrict.
type Subdistricts []*Subdistrict

func (s Subdistricts) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
