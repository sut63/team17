// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team17/app/ent/district"
	"github.com/sut63/team17/app/ent/postal"
)

// PostalCreate is the builder for creating a Postal entity.
type PostalCreate struct {
	config
	mutation *PostalMutation
	hooks    []Hook
}

// SetPostal sets the postal field.
func (pc *PostalCreate) SetPostal(s string) *PostalCreate {
	pc.mutation.SetPostal(s)
	return pc
}

// AddPostDistIDs adds the post_dist edge to District by ids.
func (pc *PostalCreate) AddPostDistIDs(ids ...int) *PostalCreate {
	pc.mutation.AddPostDistIDs(ids...)
	return pc
}

// AddPostDist adds the post_dist edges to District.
func (pc *PostalCreate) AddPostDist(d ...*District) *PostalCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pc.AddPostDistIDs(ids...)
}

// Mutation returns the PostalMutation object of the builder.
func (pc *PostalCreate) Mutation() *PostalMutation {
	return pc.mutation
}

// Save creates the Postal in the database.
func (pc *PostalCreate) Save(ctx context.Context) (*Postal, error) {
	if _, ok := pc.mutation.Postal(); !ok {
		return nil, &ValidationError{Name: "postal", err: errors.New("ent: missing required field \"postal\"")}
	}
	if v, ok := pc.mutation.Postal(); ok {
		if err := postal.PostalValidator(v); err != nil {
			return nil, &ValidationError{Name: "postal", err: fmt.Errorf("ent: validator failed for field \"postal\": %w", err)}
		}
	}
	var (
		err  error
		node *Postal
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostalCreate) SaveX(ctx context.Context) *Postal {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PostalCreate) sqlSave(ctx context.Context) (*Postal, error) {
	po, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	po.ID = int(id)
	return po, nil
}

func (pc *PostalCreate) createSpec() (*Postal, *sqlgraph.CreateSpec) {
	var (
		po    = &Postal{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: postal.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: postal.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Postal(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postal.FieldPostal,
		})
		po.Postal = value
	}
	if nodes := pc.mutation.PostDistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   postal.PostDistTable,
			Columns: []string{postal.PostDistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: district.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return po, _spec
}