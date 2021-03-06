// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team17/app/ent/activity"
	"github.com/sut63/team17/app/ent/degree"
	"github.com/sut63/team17/app/ent/gender"
	"github.com/sut63/team17/app/ent/prefix"
	"github.com/sut63/team17/app/ent/province"
	"github.com/sut63/team17/app/ent/results"
	"github.com/sut63/team17/app/ent/student"
)

// StudentCreate is the builder for creating a Student entity.
type StudentCreate struct {
	config
	mutation *StudentMutation
	hooks    []Hook
}

// SetFname sets the fname field.
func (sc *StudentCreate) SetFname(s string) *StudentCreate {
	sc.mutation.SetFname(s)
	return sc
}

// SetLname sets the lname field.
func (sc *StudentCreate) SetLname(s string) *StudentCreate {
	sc.mutation.SetLname(s)
	return sc
}

// SetSchoolname sets the schoolname field.
func (sc *StudentCreate) SetSchoolname(s string) *StudentCreate {
	sc.mutation.SetSchoolname(s)
	return sc
}

// SetRecentAddress sets the recent_address field.
func (sc *StudentCreate) SetRecentAddress(s string) *StudentCreate {
	sc.mutation.SetRecentAddress(s)
	return sc
}

// SetTelephone sets the telephone field.
func (sc *StudentCreate) SetTelephone(s string) *StudentCreate {
	sc.mutation.SetTelephone(s)
	return sc
}

// SetEmail sets the email field.
func (sc *StudentCreate) SetEmail(s string) *StudentCreate {
	sc.mutation.SetEmail(s)
	return sc
}

// SetStudGendID sets the stud_gend edge to Gender by id.
func (sc *StudentCreate) SetStudGendID(id int) *StudentCreate {
	sc.mutation.SetStudGendID(id)
	return sc
}

// SetNillableStudGendID sets the stud_gend edge to Gender by id if the given value is not nil.
func (sc *StudentCreate) SetNillableStudGendID(id *int) *StudentCreate {
	if id != nil {
		sc = sc.SetStudGendID(*id)
	}
	return sc
}

// SetStudGend sets the stud_gend edge to Gender.
func (sc *StudentCreate) SetStudGend(g *Gender) *StudentCreate {
	return sc.SetStudGendID(g.ID)
}

// AddStudActiIDs adds the stud_acti edge to Activity by ids.
func (sc *StudentCreate) AddStudActiIDs(ids ...int) *StudentCreate {
	sc.mutation.AddStudActiIDs(ids...)
	return sc
}

// AddStudActi adds the stud_acti edges to Activity.
func (sc *StudentCreate) AddStudActi(a ...*Activity) *StudentCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sc.AddStudActiIDs(ids...)
}

// AddStudResuIDs adds the stud_resu edge to Results by ids.
func (sc *StudentCreate) AddStudResuIDs(ids ...int) *StudentCreate {
	sc.mutation.AddStudResuIDs(ids...)
	return sc
}

// AddStudResu adds the stud_resu edges to Results.
func (sc *StudentCreate) AddStudResu(r ...*Results) *StudentCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return sc.AddStudResuIDs(ids...)
}

// SetStudProvID sets the stud_prov edge to Province by id.
func (sc *StudentCreate) SetStudProvID(id int) *StudentCreate {
	sc.mutation.SetStudProvID(id)
	return sc
}

// SetNillableStudProvID sets the stud_prov edge to Province by id if the given value is not nil.
func (sc *StudentCreate) SetNillableStudProvID(id *int) *StudentCreate {
	if id != nil {
		sc = sc.SetStudProvID(*id)
	}
	return sc
}

// SetStudProv sets the stud_prov edge to Province.
func (sc *StudentCreate) SetStudProv(p *Province) *StudentCreate {
	return sc.SetStudProvID(p.ID)
}

// SetStudDistID sets the stud_dist edge to Province by id.
func (sc *StudentCreate) SetStudDistID(id int) *StudentCreate {
	sc.mutation.SetStudDistID(id)
	return sc
}

// SetNillableStudDistID sets the stud_dist edge to Province by id if the given value is not nil.
func (sc *StudentCreate) SetNillableStudDistID(id *int) *StudentCreate {
	if id != nil {
		sc = sc.SetStudDistID(*id)
	}
	return sc
}

// SetStudDist sets the stud_dist edge to Province.
func (sc *StudentCreate) SetStudDist(p *Province) *StudentCreate {
	return sc.SetStudDistID(p.ID)
}

// SetStudSubdID sets the stud_subd edge to Province by id.
func (sc *StudentCreate) SetStudSubdID(id int) *StudentCreate {
	sc.mutation.SetStudSubdID(id)
	return sc
}

// SetNillableStudSubdID sets the stud_subd edge to Province by id if the given value is not nil.
func (sc *StudentCreate) SetNillableStudSubdID(id *int) *StudentCreate {
	if id != nil {
		sc = sc.SetStudSubdID(*id)
	}
	return sc
}

// SetStudSubd sets the stud_subd edge to Province.
func (sc *StudentCreate) SetStudSubd(p *Province) *StudentCreate {
	return sc.SetStudSubdID(p.ID)
}

// SetStudPostID sets the stud_post edge to Province by id.
func (sc *StudentCreate) SetStudPostID(id int) *StudentCreate {
	sc.mutation.SetStudPostID(id)
	return sc
}

// SetNillableStudPostID sets the stud_post edge to Province by id if the given value is not nil.
func (sc *StudentCreate) SetNillableStudPostID(id *int) *StudentCreate {
	if id != nil {
		sc = sc.SetStudPostID(*id)
	}
	return sc
}

// SetStudPost sets the stud_post edge to Province.
func (sc *StudentCreate) SetStudPost(p *Province) *StudentCreate {
	return sc.SetStudPostID(p.ID)
}

// SetStudPrefID sets the stud_pref edge to Prefix by id.
func (sc *StudentCreate) SetStudPrefID(id int) *StudentCreate {
	sc.mutation.SetStudPrefID(id)
	return sc
}

// SetNillableStudPrefID sets the stud_pref edge to Prefix by id if the given value is not nil.
func (sc *StudentCreate) SetNillableStudPrefID(id *int) *StudentCreate {
	if id != nil {
		sc = sc.SetStudPrefID(*id)
	}
	return sc
}

// SetStudPref sets the stud_pref edge to Prefix.
func (sc *StudentCreate) SetStudPref(p *Prefix) *StudentCreate {
	return sc.SetStudPrefID(p.ID)
}

// SetStudDegrID sets the stud_degr edge to Degree by id.
func (sc *StudentCreate) SetStudDegrID(id int) *StudentCreate {
	sc.mutation.SetStudDegrID(id)
	return sc
}

// SetNillableStudDegrID sets the stud_degr edge to Degree by id if the given value is not nil.
func (sc *StudentCreate) SetNillableStudDegrID(id *int) *StudentCreate {
	if id != nil {
		sc = sc.SetStudDegrID(*id)
	}
	return sc
}

// SetStudDegr sets the stud_degr edge to Degree.
func (sc *StudentCreate) SetStudDegr(d *Degree) *StudentCreate {
	return sc.SetStudDegrID(d.ID)
}

// Mutation returns the StudentMutation object of the builder.
func (sc *StudentCreate) Mutation() *StudentMutation {
	return sc.mutation
}

// Save creates the Student in the database.
func (sc *StudentCreate) Save(ctx context.Context) (*Student, error) {
	if _, ok := sc.mutation.Fname(); !ok {
		return nil, &ValidationError{Name: "fname", err: errors.New("ent: missing required field \"fname\"")}
	}
	if v, ok := sc.mutation.Fname(); ok {
		if err := student.FnameValidator(v); err != nil {
			return nil, &ValidationError{Name: "fname", err: fmt.Errorf("ent: validator failed for field \"fname\": %w", err)}
		}
	}
	if _, ok := sc.mutation.Lname(); !ok {
		return nil, &ValidationError{Name: "lname", err: errors.New("ent: missing required field \"lname\"")}
	}
	if v, ok := sc.mutation.Lname(); ok {
		if err := student.LnameValidator(v); err != nil {
			return nil, &ValidationError{Name: "lname", err: fmt.Errorf("ent: validator failed for field \"lname\": %w", err)}
		}
	}
	if _, ok := sc.mutation.Schoolname(); !ok {
		return nil, &ValidationError{Name: "schoolname", err: errors.New("ent: missing required field \"schoolname\"")}
	}
	if v, ok := sc.mutation.Schoolname(); ok {
		if err := student.SchoolnameValidator(v); err != nil {
			return nil, &ValidationError{Name: "schoolname", err: fmt.Errorf("ent: validator failed for field \"schoolname\": %w", err)}
		}
	}
	if _, ok := sc.mutation.RecentAddress(); !ok {
		return nil, &ValidationError{Name: "recent_address", err: errors.New("ent: missing required field \"recent_address\"")}
	}
	if v, ok := sc.mutation.RecentAddress(); ok {
		if err := student.RecentAddressValidator(v); err != nil {
			return nil, &ValidationError{Name: "recent_address", err: fmt.Errorf("ent: validator failed for field \"recent_address\": %w", err)}
		}
	}
	if _, ok := sc.mutation.Telephone(); !ok {
		return nil, &ValidationError{Name: "telephone", err: errors.New("ent: missing required field \"telephone\"")}
	}
	if v, ok := sc.mutation.Telephone(); ok {
		if err := student.TelephoneValidator(v); err != nil {
			return nil, &ValidationError{Name: "telephone", err: fmt.Errorf("ent: validator failed for field \"telephone\": %w", err)}
		}
	}
	if _, ok := sc.mutation.Email(); !ok {
		return nil, &ValidationError{Name: "email", err: errors.New("ent: missing required field \"email\"")}
	}
	if v, ok := sc.mutation.Email(); ok {
		if err := student.EmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	var (
		err  error
		node *Student
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StudentCreate) SaveX(ctx context.Context) *Student {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *StudentCreate) sqlSave(ctx context.Context) (*Student, error) {
	s, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}

func (sc *StudentCreate) createSpec() (*Student, *sqlgraph.CreateSpec) {
	var (
		s     = &Student{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: student.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: student.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Fname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldFname,
		})
		s.Fname = value
	}
	if value, ok := sc.mutation.Lname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldLname,
		})
		s.Lname = value
	}
	if value, ok := sc.mutation.Schoolname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldSchoolname,
		})
		s.Schoolname = value
	}
	if value, ok := sc.mutation.RecentAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldRecentAddress,
		})
		s.RecentAddress = value
	}
	if value, ok := sc.mutation.Telephone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldTelephone,
		})
		s.Telephone = value
	}
	if value, ok := sc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldEmail,
		})
		s.Email = value
	}
	if nodes := sc.mutation.StudGendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.StudGendTable,
			Columns: []string{student.StudGendColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StudActiIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.StudActiTable,
			Columns: []string{student.StudActiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: activity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StudResuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.StudResuTable,
			Columns: []string{student.StudResuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: results.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StudProvIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.StudProvTable,
			Columns: []string{student.StudProvColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: province.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StudDistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.StudDistTable,
			Columns: []string{student.StudDistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: province.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StudSubdIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.StudSubdTable,
			Columns: []string{student.StudSubdColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: province.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StudPostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.StudPostTable,
			Columns: []string{student.StudPostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: province.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StudPrefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.StudPrefTable,
			Columns: []string{student.StudPrefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prefix.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StudDegrIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.StudDegrTable,
			Columns: []string{student.StudDegrColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: degree.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return s, _spec
}
