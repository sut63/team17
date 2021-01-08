// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team17/app/ent/place"
)

// Place is the model entity for the Place schema.
type Place struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PLACE holds the value of the "PLACE" field.
	PLACE string `json:"PLACE,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlaceQuery when eager-loading is set.
	Edges PlaceEdges `json:"edges"`
}

// PlaceEdges holds the relations/edges for other nodes in the graph.
type PlaceEdges struct {
	// PlaceActi holds the value of the place_acti edge.
	PlaceActi []*Activity
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PlaceActiOrErr returns the PlaceActi value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) PlaceActiOrErr() ([]*Activity, error) {
	if e.loadedTypes[0] {
		return e.PlaceActi, nil
	}
	return nil, &NotLoadedError{edge: "place_acti"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Place) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // PLACE
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Place fields.
func (pl *Place) assignValues(values ...interface{}) error {
	if m, n := len(values), len(place.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pl.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field PLACE", values[0])
	} else if value.Valid {
		pl.PLACE = value.String
	}
	return nil
}

// QueryPlaceActi queries the place_acti edge of the Place.
func (pl *Place) QueryPlaceActi() *ActivityQuery {
	return (&PlaceClient{config: pl.config}).QueryPlaceActi(pl)
}

// Update returns a builder for updating this Place.
// Note that, you need to call Place.Unwrap() before calling this method, if this Place
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Place) Update() *PlaceUpdateOne {
	return (&PlaceClient{config: pl.config}).UpdateOne(pl)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pl *Place) Unwrap() *Place {
	tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Place is not a transactional entity")
	}
	pl.config.driver = tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Place) String() string {
	var builder strings.Builder
	builder.WriteString("Place(")
	builder.WriteString(fmt.Sprintf("id=%v", pl.ID))
	builder.WriteString(", PLACE=")
	builder.WriteString(pl.PLACE)
	builder.WriteByte(')')
	return builder.String()
}

// Places is a parsable slice of Place.
type Places []*Place

func (pl Places) config(cfg config) {
	for _i := range pl {
		pl[_i].config = cfg
	}
}