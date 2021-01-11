// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team17/app/ent/term"
	"github.com/sut63/team17/app/ent/year"
)

// TermCreate is the builder for creating a Term entity.
type TermCreate struct {
	config
	mutation *TermMutation
	hooks    []Hook
}

// SetSemester sets the semester field.
func (tc *TermCreate) SetSemester(i int) *TermCreate {
	tc.mutation.SetSemester(i)
	return tc
}

// SetTermYearID sets the term_year edge to Year by id.
func (tc *TermCreate) SetTermYearID(id int) *TermCreate {
	tc.mutation.SetTermYearID(id)
	return tc
}

// SetNillableTermYearID sets the term_year edge to Year by id if the given value is not nil.
func (tc *TermCreate) SetNillableTermYearID(id *int) *TermCreate {
	if id != nil {
		tc = tc.SetTermYearID(*id)
	}
	return tc
}

// SetTermYear sets the term_year edge to Year.
func (tc *TermCreate) SetTermYear(y *Year) *TermCreate {
	return tc.SetTermYearID(y.ID)
}

// Mutation returns the TermMutation object of the builder.
func (tc *TermCreate) Mutation() *TermMutation {
	return tc.mutation
}

// Save creates the Term in the database.
func (tc *TermCreate) Save(ctx context.Context) (*Term, error) {
	if _, ok := tc.mutation.Semester(); !ok {
		return nil, &ValidationError{Name: "semester", err: errors.New("ent: missing required field \"semester\"")}
	}
	if v, ok := tc.mutation.Semester(); ok {
		if err := term.SemesterValidator(v); err != nil {
			return nil, &ValidationError{Name: "semester", err: fmt.Errorf("ent: validator failed for field \"semester\": %w", err)}
		}
	}
	var (
		err  error
		node *Term
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TermMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TermCreate) SaveX(ctx context.Context) *Term {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TermCreate) sqlSave(ctx context.Context) (*Term, error) {
	t, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}

func (tc *TermCreate) createSpec() (*Term, *sqlgraph.CreateSpec) {
	var (
		t     = &Term{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: term.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: term.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Semester(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: term.FieldSemester,
		})
		t.Semester = value
	}
	if nodes := tc.mutation.TermYearIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   term.TermYearTable,
			Columns: []string{term.TermYearColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: year.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return t, _spec
}
