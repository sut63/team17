// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team17/app/ent/professor"
	"github.com/sut63/team17/app/ent/professorship"
)

// ProfessorshipCreate is the builder for creating a Professorship entity.
type ProfessorshipCreate struct {
	config
	mutation *ProfessorshipMutation
	hooks    []Hook
}

// SetProfessorship sets the professorship field.
func (pc *ProfessorshipCreate) SetProfessorship(s string) *ProfessorshipCreate {
	pc.mutation.SetProfessorship(s)
	return pc
}

// AddProsProfIDs adds the pros_prof edge to Professor by ids.
func (pc *ProfessorshipCreate) AddProsProfIDs(ids ...int) *ProfessorshipCreate {
	pc.mutation.AddProsProfIDs(ids...)
	return pc
}

// AddProsProf adds the pros_prof edges to Professor.
func (pc *ProfessorshipCreate) AddProsProf(p ...*Professor) *ProfessorshipCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProsProfIDs(ids...)
}

// Mutation returns the ProfessorshipMutation object of the builder.
func (pc *ProfessorshipCreate) Mutation() *ProfessorshipMutation {
	return pc.mutation
}

// Save creates the Professorship in the database.
func (pc *ProfessorshipCreate) Save(ctx context.Context) (*Professorship, error) {
	if _, ok := pc.mutation.Professorship(); !ok {
		return nil, &ValidationError{Name: "professorship", err: errors.New("ent: missing required field \"professorship\"")}
	}
	if v, ok := pc.mutation.Professorship(); ok {
		if err := professorship.ProfessorshipValidator(v); err != nil {
			return nil, &ValidationError{Name: "professorship", err: fmt.Errorf("ent: validator failed for field \"professorship\": %w", err)}
		}
	}
	var (
		err  error
		node *Professorship
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfessorshipMutation)
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
func (pc *ProfessorshipCreate) SaveX(ctx context.Context) *Professorship {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *ProfessorshipCreate) sqlSave(ctx context.Context) (*Professorship, error) {
	pr, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pr.ID = int(id)
	return pr, nil
}

func (pc *ProfessorshipCreate) createSpec() (*Professorship, *sqlgraph.CreateSpec) {
	var (
		pr    = &Professorship{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: professorship.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: professorship.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Professorship(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: professorship.FieldProfessorship,
		})
		pr.Professorship = value
	}
	if nodes := pc.mutation.ProsProfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   professorship.ProsProfTable,
			Columns: []string{professorship.ProsProfColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: professor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pr, _spec
}