// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team17/app/ent/emp"
)

// EmpCreate is the builder for creating a Emp entity.
type EmpCreate struct {
	config
	mutation *EmpMutation
	hooks    []Hook
}

// SetUser sets the user field.
func (ec *EmpCreate) SetUser(s string) *EmpCreate {
	ec.mutation.SetUser(s)
	return ec
}

// SetPass sets the pass field.
func (ec *EmpCreate) SetPass(s string) *EmpCreate {
	ec.mutation.SetPass(s)
	return ec
}

// SetRole sets the role field.
func (ec *EmpCreate) SetRole(s string) *EmpCreate {
	ec.mutation.SetRole(s)
	return ec
}

// Mutation returns the EmpMutation object of the builder.
func (ec *EmpCreate) Mutation() *EmpMutation {
	return ec.mutation
}

// Save creates the Emp in the database.
func (ec *EmpCreate) Save(ctx context.Context) (*Emp, error) {
	if _, ok := ec.mutation.User(); !ok {
		return nil, &ValidationError{Name: "user", err: errors.New("ent: missing required field \"user\"")}
	}
	if v, ok := ec.mutation.User(); ok {
		if err := emp.UserValidator(v); err != nil {
			return nil, &ValidationError{Name: "user", err: fmt.Errorf("ent: validator failed for field \"user\": %w", err)}
		}
	}
	if _, ok := ec.mutation.Pass(); !ok {
		return nil, &ValidationError{Name: "pass", err: errors.New("ent: missing required field \"pass\"")}
	}
	if v, ok := ec.mutation.Pass(); ok {
		if err := emp.PassValidator(v); err != nil {
			return nil, &ValidationError{Name: "pass", err: fmt.Errorf("ent: validator failed for field \"pass\": %w", err)}
		}
	}
	if _, ok := ec.mutation.Role(); !ok {
		return nil, &ValidationError{Name: "role", err: errors.New("ent: missing required field \"role\"")}
	}
	var (
		err  error
		node *Emp
	)
	if len(ec.hooks) == 0 {
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmpMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ec.mutation = mutation
			node, err = ec.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EmpCreate) SaveX(ctx context.Context) *Emp {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ec *EmpCreate) sqlSave(ctx context.Context) (*Emp, error) {
	e, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	e.ID = int(id)
	return e, nil
}

func (ec *EmpCreate) createSpec() (*Emp, *sqlgraph.CreateSpec) {
	var (
		e     = &Emp{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: emp.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: emp.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.User(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emp.FieldUser,
		})
		e.User = value
	}
	if value, ok := ec.mutation.Pass(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emp.FieldPass,
		})
		e.Pass = value
	}
	if value, ok := ec.mutation.Role(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emp.FieldRole,
		})
		e.Role = value
	}
	return e, _spec
}
