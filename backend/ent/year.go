// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team17/app/ent/year"
)

// Year is the model entity for the Year schema.
type Year struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Years holds the value of the "years" field.
	Years int `json:"years,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the YearQuery when eager-loading is set.
	Edges YearEdges `json:"edges"`
}

// YearEdges holds the relations/edges for other nodes in the graph.
type YearEdges struct {
	// YearTerm holds the value of the year_term edge.
	YearTerm []*Term
	// YearResu holds the value of the year_resu edge.
	YearResu []*Results
	// YearActi holds the value of the year_acti edge.
	YearActi []*Activity
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// YearTermOrErr returns the YearTerm value or an error if the edge
// was not loaded in eager-loading.
func (e YearEdges) YearTermOrErr() ([]*Term, error) {
	if e.loadedTypes[0] {
		return e.YearTerm, nil
	}
	return nil, &NotLoadedError{edge: "year_term"}
}

// YearResuOrErr returns the YearResu value or an error if the edge
// was not loaded in eager-loading.
func (e YearEdges) YearResuOrErr() ([]*Results, error) {
	if e.loadedTypes[1] {
		return e.YearResu, nil
	}
	return nil, &NotLoadedError{edge: "year_resu"}
}

// YearActiOrErr returns the YearActi value or an error if the edge
// was not loaded in eager-loading.
func (e YearEdges) YearActiOrErr() ([]*Activity, error) {
	if e.loadedTypes[2] {
		return e.YearActi, nil
	}
	return nil, &NotLoadedError{edge: "year_acti"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Year) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullInt64{}, // years
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Year fields.
func (y *Year) assignValues(values ...interface{}) error {
	if m, n := len(values), len(year.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	y.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field years", values[0])
	} else if value.Valid {
		y.Years = int(value.Int64)
	}
	return nil
}

// QueryYearTerm queries the year_term edge of the Year.
func (y *Year) QueryYearTerm() *TermQuery {
	return (&YearClient{config: y.config}).QueryYearTerm(y)
}

// QueryYearResu queries the year_resu edge of the Year.
func (y *Year) QueryYearResu() *ResultsQuery {
	return (&YearClient{config: y.config}).QueryYearResu(y)
}

// QueryYearActi queries the year_acti edge of the Year.
func (y *Year) QueryYearActi() *ActivityQuery {
	return (&YearClient{config: y.config}).QueryYearActi(y)
}

// Update returns a builder for updating this Year.
// Note that, you need to call Year.Unwrap() before calling this method, if this Year
// was returned from a transaction, and the transaction was committed or rolled back.
func (y *Year) Update() *YearUpdateOne {
	return (&YearClient{config: y.config}).UpdateOne(y)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (y *Year) Unwrap() *Year {
	tx, ok := y.config.driver.(*txDriver)
	if !ok {
		panic("ent: Year is not a transactional entity")
	}
	y.config.driver = tx.drv
	return y
}

// String implements the fmt.Stringer.
func (y *Year) String() string {
	var builder strings.Builder
	builder.WriteString("Year(")
	builder.WriteString(fmt.Sprintf("id=%v", y.ID))
	builder.WriteString(", years=")
	builder.WriteString(fmt.Sprintf("%v", y.Years))
	builder.WriteByte(')')
	return builder.String()
}

// Years is a parsable slice of Year.
type Years []*Year

func (y Years) config(cfg config) {
	for _i := range y {
		y[_i].config = cfg
	}
}
