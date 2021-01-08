// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team17/app/ent/course"
	"github.com/sut63/team17/app/ent/degree"
	"github.com/sut63/team17/app/ent/predicate"
	"github.com/sut63/team17/app/ent/student"
)

// DegreeUpdate is the builder for updating Degree entities.
type DegreeUpdate struct {
	config
	hooks      []Hook
	mutation   *DegreeMutation
	predicates []predicate.Degree
}

// Where adds a new predicate for the builder.
func (du *DegreeUpdate) Where(ps ...predicate.Degree) *DegreeUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetDegree sets the degree field.
func (du *DegreeUpdate) SetDegree(s string) *DegreeUpdate {
	du.mutation.SetDegree(s)
	return du
}

// AddDegrStudIDs adds the degr_stud edge to Student by ids.
func (du *DegreeUpdate) AddDegrStudIDs(ids ...int) *DegreeUpdate {
	du.mutation.AddDegrStudIDs(ids...)
	return du
}

// AddDegrStud adds the degr_stud edges to Student.
func (du *DegreeUpdate) AddDegrStud(s ...*Student) *DegreeUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return du.AddDegrStudIDs(ids...)
}

// AddDegrCourIDs adds the degr_cour edge to Course by ids.
func (du *DegreeUpdate) AddDegrCourIDs(ids ...int) *DegreeUpdate {
	du.mutation.AddDegrCourIDs(ids...)
	return du
}

// AddDegrCour adds the degr_cour edges to Course.
func (du *DegreeUpdate) AddDegrCour(c ...*Course) *DegreeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return du.AddDegrCourIDs(ids...)
}

// Mutation returns the DegreeMutation object of the builder.
func (du *DegreeUpdate) Mutation() *DegreeMutation {
	return du.mutation
}

// RemoveDegrStudIDs removes the degr_stud edge to Student by ids.
func (du *DegreeUpdate) RemoveDegrStudIDs(ids ...int) *DegreeUpdate {
	du.mutation.RemoveDegrStudIDs(ids...)
	return du
}

// RemoveDegrStud removes degr_stud edges to Student.
func (du *DegreeUpdate) RemoveDegrStud(s ...*Student) *DegreeUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return du.RemoveDegrStudIDs(ids...)
}

// RemoveDegrCourIDs removes the degr_cour edge to Course by ids.
func (du *DegreeUpdate) RemoveDegrCourIDs(ids ...int) *DegreeUpdate {
	du.mutation.RemoveDegrCourIDs(ids...)
	return du
}

// RemoveDegrCour removes degr_cour edges to Course.
func (du *DegreeUpdate) RemoveDegrCour(c ...*Course) *DegreeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return du.RemoveDegrCourIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DegreeUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.Degree(); ok {
		if err := degree.DegreeValidator(v); err != nil {
			return 0, &ValidationError{Name: "degree", err: fmt.Errorf("ent: validator failed for field \"degree\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DegreeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DegreeUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DegreeUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DegreeUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DegreeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   degree.Table,
			Columns: degree.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: degree.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Degree(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: degree.FieldDegree,
		})
	}
	if nodes := du.mutation.RemovedDegrStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   degree.DegrStudTable,
			Columns: []string{degree.DegrStudColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DegrStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   degree.DegrStudTable,
			Columns: []string{degree.DegrStudColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := du.mutation.RemovedDegrCourIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   degree.DegrCourTable,
			Columns: []string{degree.DegrCourColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DegrCourIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   degree.DegrCourTable,
			Columns: []string{degree.DegrCourColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{degree.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DegreeUpdateOne is the builder for updating a single Degree entity.
type DegreeUpdateOne struct {
	config
	hooks    []Hook
	mutation *DegreeMutation
}

// SetDegree sets the degree field.
func (duo *DegreeUpdateOne) SetDegree(s string) *DegreeUpdateOne {
	duo.mutation.SetDegree(s)
	return duo
}

// AddDegrStudIDs adds the degr_stud edge to Student by ids.
func (duo *DegreeUpdateOne) AddDegrStudIDs(ids ...int) *DegreeUpdateOne {
	duo.mutation.AddDegrStudIDs(ids...)
	return duo
}

// AddDegrStud adds the degr_stud edges to Student.
func (duo *DegreeUpdateOne) AddDegrStud(s ...*Student) *DegreeUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return duo.AddDegrStudIDs(ids...)
}

// AddDegrCourIDs adds the degr_cour edge to Course by ids.
func (duo *DegreeUpdateOne) AddDegrCourIDs(ids ...int) *DegreeUpdateOne {
	duo.mutation.AddDegrCourIDs(ids...)
	return duo
}

// AddDegrCour adds the degr_cour edges to Course.
func (duo *DegreeUpdateOne) AddDegrCour(c ...*Course) *DegreeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return duo.AddDegrCourIDs(ids...)
}

// Mutation returns the DegreeMutation object of the builder.
func (duo *DegreeUpdateOne) Mutation() *DegreeMutation {
	return duo.mutation
}

// RemoveDegrStudIDs removes the degr_stud edge to Student by ids.
func (duo *DegreeUpdateOne) RemoveDegrStudIDs(ids ...int) *DegreeUpdateOne {
	duo.mutation.RemoveDegrStudIDs(ids...)
	return duo
}

// RemoveDegrStud removes degr_stud edges to Student.
func (duo *DegreeUpdateOne) RemoveDegrStud(s ...*Student) *DegreeUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return duo.RemoveDegrStudIDs(ids...)
}

// RemoveDegrCourIDs removes the degr_cour edge to Course by ids.
func (duo *DegreeUpdateOne) RemoveDegrCourIDs(ids ...int) *DegreeUpdateOne {
	duo.mutation.RemoveDegrCourIDs(ids...)
	return duo
}

// RemoveDegrCour removes degr_cour edges to Course.
func (duo *DegreeUpdateOne) RemoveDegrCour(c ...*Course) *DegreeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return duo.RemoveDegrCourIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DegreeUpdateOne) Save(ctx context.Context) (*Degree, error) {
	if v, ok := duo.mutation.Degree(); ok {
		if err := degree.DegreeValidator(v); err != nil {
			return nil, &ValidationError{Name: "degree", err: fmt.Errorf("ent: validator failed for field \"degree\": %w", err)}
		}
	}

	var (
		err  error
		node *Degree
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DegreeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DegreeUpdateOne) SaveX(ctx context.Context) *Degree {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DegreeUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DegreeUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DegreeUpdateOne) sqlSave(ctx context.Context) (d *Degree, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   degree.Table,
			Columns: degree.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: degree.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Degree.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Degree(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: degree.FieldDegree,
		})
	}
	if nodes := duo.mutation.RemovedDegrStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   degree.DegrStudTable,
			Columns: []string{degree.DegrStudColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DegrStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   degree.DegrStudTable,
			Columns: []string{degree.DegrStudColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := duo.mutation.RemovedDegrCourIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   degree.DegrCourTable,
			Columns: []string{degree.DegrCourColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DegrCourIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   degree.DegrCourTable,
			Columns: []string{degree.DegrCourColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	d = &Degree{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{degree.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}
