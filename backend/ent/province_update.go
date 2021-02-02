// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team17/app/ent/continent"
	"github.com/sut63/team17/app/ent/country"
	"github.com/sut63/team17/app/ent/predicate"
	"github.com/sut63/team17/app/ent/province"
	"github.com/sut63/team17/app/ent/region"
	"github.com/sut63/team17/app/ent/student"
)

// ProvinceUpdate is the builder for updating Province entities.
type ProvinceUpdate struct {
	config
	hooks      []Hook
	mutation   *ProvinceMutation
	predicates []predicate.Province
}

// Where adds a new predicate for the builder.
func (pu *ProvinceUpdate) Where(ps ...predicate.Province) *ProvinceUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetProvince sets the province field.
func (pu *ProvinceUpdate) SetProvince(s string) *ProvinceUpdate {
	pu.mutation.SetProvince(s)
	return pu
}

// SetDistrict sets the district field.
func (pu *ProvinceUpdate) SetDistrict(s string) *ProvinceUpdate {
	pu.mutation.SetDistrict(s)
	return pu
}

// SetSubdistrict sets the subdistrict field.
func (pu *ProvinceUpdate) SetSubdistrict(s string) *ProvinceUpdate {
	pu.mutation.SetSubdistrict(s)
	return pu
}

// SetPostal sets the postal field.
func (pu *ProvinceUpdate) SetPostal(s string) *ProvinceUpdate {
	pu.mutation.SetPostal(s)
	return pu
}

// SetProvRegiID sets the prov_regi edge to Region by id.
func (pu *ProvinceUpdate) SetProvRegiID(id int) *ProvinceUpdate {
	pu.mutation.SetProvRegiID(id)
	return pu
}

// SetNillableProvRegiID sets the prov_regi edge to Region by id if the given value is not nil.
func (pu *ProvinceUpdate) SetNillableProvRegiID(id *int) *ProvinceUpdate {
	if id != nil {
		pu = pu.SetProvRegiID(*id)
	}
	return pu
}

// SetProvRegi sets the prov_regi edge to Region.
func (pu *ProvinceUpdate) SetProvRegi(r *Region) *ProvinceUpdate {
	return pu.SetProvRegiID(r.ID)
}

// SetProvCounID sets the prov_coun edge to Country by id.
func (pu *ProvinceUpdate) SetProvCounID(id int) *ProvinceUpdate {
	pu.mutation.SetProvCounID(id)
	return pu
}

// SetNillableProvCounID sets the prov_coun edge to Country by id if the given value is not nil.
func (pu *ProvinceUpdate) SetNillableProvCounID(id *int) *ProvinceUpdate {
	if id != nil {
		pu = pu.SetProvCounID(*id)
	}
	return pu
}

// SetProvCoun sets the prov_coun edge to Country.
func (pu *ProvinceUpdate) SetProvCoun(c *Country) *ProvinceUpdate {
	return pu.SetProvCounID(c.ID)
}

// SetProvContID sets the prov_cont edge to Continent by id.
func (pu *ProvinceUpdate) SetProvContID(id int) *ProvinceUpdate {
	pu.mutation.SetProvContID(id)
	return pu
}

// SetNillableProvContID sets the prov_cont edge to Continent by id if the given value is not nil.
func (pu *ProvinceUpdate) SetNillableProvContID(id *int) *ProvinceUpdate {
	if id != nil {
		pu = pu.SetProvContID(*id)
	}
	return pu
}

// SetProvCont sets the prov_cont edge to Continent.
func (pu *ProvinceUpdate) SetProvCont(c *Continent) *ProvinceUpdate {
	return pu.SetProvContID(c.ID)
}

// AddProvStudIDs adds the prov_stud edge to Student by ids.
func (pu *ProvinceUpdate) AddProvStudIDs(ids ...int) *ProvinceUpdate {
	pu.mutation.AddProvStudIDs(ids...)
	return pu
}

// AddProvStud adds the prov_stud edges to Student.
func (pu *ProvinceUpdate) AddProvStud(s ...*Student) *ProvinceUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddProvStudIDs(ids...)
}

// AddDistStudIDs adds the dist_stud edge to Student by ids.
func (pu *ProvinceUpdate) AddDistStudIDs(ids ...int) *ProvinceUpdate {
	pu.mutation.AddDistStudIDs(ids...)
	return pu
}

// AddDistStud adds the dist_stud edges to Student.
func (pu *ProvinceUpdate) AddDistStud(s ...*Student) *ProvinceUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddDistStudIDs(ids...)
}

// AddSubdStudIDs adds the subd_stud edge to Student by ids.
func (pu *ProvinceUpdate) AddSubdStudIDs(ids ...int) *ProvinceUpdate {
	pu.mutation.AddSubdStudIDs(ids...)
	return pu
}

// AddSubdStud adds the subd_stud edges to Student.
func (pu *ProvinceUpdate) AddSubdStud(s ...*Student) *ProvinceUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddSubdStudIDs(ids...)
}

// AddPostStudIDs adds the post_stud edge to Student by ids.
func (pu *ProvinceUpdate) AddPostStudIDs(ids ...int) *ProvinceUpdate {
	pu.mutation.AddPostStudIDs(ids...)
	return pu
}

// AddPostStud adds the post_stud edges to Student.
func (pu *ProvinceUpdate) AddPostStud(s ...*Student) *ProvinceUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddPostStudIDs(ids...)
}

// Mutation returns the ProvinceMutation object of the builder.
func (pu *ProvinceUpdate) Mutation() *ProvinceMutation {
	return pu.mutation
}

// ClearProvRegi clears the prov_regi edge to Region.
func (pu *ProvinceUpdate) ClearProvRegi() *ProvinceUpdate {
	pu.mutation.ClearProvRegi()
	return pu
}

// ClearProvCoun clears the prov_coun edge to Country.
func (pu *ProvinceUpdate) ClearProvCoun() *ProvinceUpdate {
	pu.mutation.ClearProvCoun()
	return pu
}

// ClearProvCont clears the prov_cont edge to Continent.
func (pu *ProvinceUpdate) ClearProvCont() *ProvinceUpdate {
	pu.mutation.ClearProvCont()
	return pu
}

// RemoveProvStudIDs removes the prov_stud edge to Student by ids.
func (pu *ProvinceUpdate) RemoveProvStudIDs(ids ...int) *ProvinceUpdate {
	pu.mutation.RemoveProvStudIDs(ids...)
	return pu
}

// RemoveProvStud removes prov_stud edges to Student.
func (pu *ProvinceUpdate) RemoveProvStud(s ...*Student) *ProvinceUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemoveProvStudIDs(ids...)
}

// RemoveDistStudIDs removes the dist_stud edge to Student by ids.
func (pu *ProvinceUpdate) RemoveDistStudIDs(ids ...int) *ProvinceUpdate {
	pu.mutation.RemoveDistStudIDs(ids...)
	return pu
}

// RemoveDistStud removes dist_stud edges to Student.
func (pu *ProvinceUpdate) RemoveDistStud(s ...*Student) *ProvinceUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemoveDistStudIDs(ids...)
}

// RemoveSubdStudIDs removes the subd_stud edge to Student by ids.
func (pu *ProvinceUpdate) RemoveSubdStudIDs(ids ...int) *ProvinceUpdate {
	pu.mutation.RemoveSubdStudIDs(ids...)
	return pu
}

// RemoveSubdStud removes subd_stud edges to Student.
func (pu *ProvinceUpdate) RemoveSubdStud(s ...*Student) *ProvinceUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemoveSubdStudIDs(ids...)
}

// RemovePostStudIDs removes the post_stud edge to Student by ids.
func (pu *ProvinceUpdate) RemovePostStudIDs(ids ...int) *ProvinceUpdate {
	pu.mutation.RemovePostStudIDs(ids...)
	return pu
}

// RemovePostStud removes post_stud edges to Student.
func (pu *ProvinceUpdate) RemovePostStud(s ...*Student) *ProvinceUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemovePostStudIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *ProvinceUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.Province(); ok {
		if err := province.ProvinceValidator(v); err != nil {
			return 0, &ValidationError{Name: "province", err: fmt.Errorf("ent: validator failed for field \"province\": %w", err)}
		}
	}
	if v, ok := pu.mutation.District(); ok {
		if err := province.DistrictValidator(v); err != nil {
			return 0, &ValidationError{Name: "district", err: fmt.Errorf("ent: validator failed for field \"district\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Subdistrict(); ok {
		if err := province.SubdistrictValidator(v); err != nil {
			return 0, &ValidationError{Name: "subdistrict", err: fmt.Errorf("ent: validator failed for field \"subdistrict\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Postal(); ok {
		if err := province.PostalValidator(v); err != nil {
			return 0, &ValidationError{Name: "postal", err: fmt.Errorf("ent: validator failed for field \"postal\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProvinceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProvinceUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProvinceUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProvinceUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProvinceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   province.Table,
			Columns: province.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: province.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Province(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: province.FieldProvince,
		})
	}
	if value, ok := pu.mutation.District(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: province.FieldDistrict,
		})
	}
	if value, ok := pu.mutation.Subdistrict(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: province.FieldSubdistrict,
		})
	}
	if value, ok := pu.mutation.Postal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: province.FieldPostal,
		})
	}
	if pu.mutation.ProvRegiCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   province.ProvRegiTable,
			Columns: []string{province.ProvRegiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: region.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProvRegiIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   province.ProvRegiTable,
			Columns: []string{province.ProvRegiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: region.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ProvCounCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   province.ProvCounTable,
			Columns: []string{province.ProvCounColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: country.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProvCounIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   province.ProvCounTable,
			Columns: []string{province.ProvCounColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: country.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ProvContCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   province.ProvContTable,
			Columns: []string{province.ProvContColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: continent.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProvContIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   province.ProvContTable,
			Columns: []string{province.ProvContColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: continent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := pu.mutation.RemovedProvStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.ProvStudTable,
			Columns: []string{province.ProvStudColumn},
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
	if nodes := pu.mutation.ProvStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.ProvStudTable,
			Columns: []string{province.ProvStudColumn},
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
	if nodes := pu.mutation.RemovedDistStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.DistStudTable,
			Columns: []string{province.DistStudColumn},
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
	if nodes := pu.mutation.DistStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.DistStudTable,
			Columns: []string{province.DistStudColumn},
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
	if nodes := pu.mutation.RemovedSubdStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.SubdStudTable,
			Columns: []string{province.SubdStudColumn},
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
	if nodes := pu.mutation.SubdStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.SubdStudTable,
			Columns: []string{province.SubdStudColumn},
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
	if nodes := pu.mutation.RemovedPostStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.PostStudTable,
			Columns: []string{province.PostStudColumn},
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
	if nodes := pu.mutation.PostStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.PostStudTable,
			Columns: []string{province.PostStudColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{province.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProvinceUpdateOne is the builder for updating a single Province entity.
type ProvinceUpdateOne struct {
	config
	hooks    []Hook
	mutation *ProvinceMutation
}

// SetProvince sets the province field.
func (puo *ProvinceUpdateOne) SetProvince(s string) *ProvinceUpdateOne {
	puo.mutation.SetProvince(s)
	return puo
}

// SetDistrict sets the district field.
func (puo *ProvinceUpdateOne) SetDistrict(s string) *ProvinceUpdateOne {
	puo.mutation.SetDistrict(s)
	return puo
}

// SetSubdistrict sets the subdistrict field.
func (puo *ProvinceUpdateOne) SetSubdistrict(s string) *ProvinceUpdateOne {
	puo.mutation.SetSubdistrict(s)
	return puo
}

// SetPostal sets the postal field.
func (puo *ProvinceUpdateOne) SetPostal(s string) *ProvinceUpdateOne {
	puo.mutation.SetPostal(s)
	return puo
}

// SetProvRegiID sets the prov_regi edge to Region by id.
func (puo *ProvinceUpdateOne) SetProvRegiID(id int) *ProvinceUpdateOne {
	puo.mutation.SetProvRegiID(id)
	return puo
}

// SetNillableProvRegiID sets the prov_regi edge to Region by id if the given value is not nil.
func (puo *ProvinceUpdateOne) SetNillableProvRegiID(id *int) *ProvinceUpdateOne {
	if id != nil {
		puo = puo.SetProvRegiID(*id)
	}
	return puo
}

// SetProvRegi sets the prov_regi edge to Region.
func (puo *ProvinceUpdateOne) SetProvRegi(r *Region) *ProvinceUpdateOne {
	return puo.SetProvRegiID(r.ID)
}

// SetProvCounID sets the prov_coun edge to Country by id.
func (puo *ProvinceUpdateOne) SetProvCounID(id int) *ProvinceUpdateOne {
	puo.mutation.SetProvCounID(id)
	return puo
}

// SetNillableProvCounID sets the prov_coun edge to Country by id if the given value is not nil.
func (puo *ProvinceUpdateOne) SetNillableProvCounID(id *int) *ProvinceUpdateOne {
	if id != nil {
		puo = puo.SetProvCounID(*id)
	}
	return puo
}

// SetProvCoun sets the prov_coun edge to Country.
func (puo *ProvinceUpdateOne) SetProvCoun(c *Country) *ProvinceUpdateOne {
	return puo.SetProvCounID(c.ID)
}

// SetProvContID sets the prov_cont edge to Continent by id.
func (puo *ProvinceUpdateOne) SetProvContID(id int) *ProvinceUpdateOne {
	puo.mutation.SetProvContID(id)
	return puo
}

// SetNillableProvContID sets the prov_cont edge to Continent by id if the given value is not nil.
func (puo *ProvinceUpdateOne) SetNillableProvContID(id *int) *ProvinceUpdateOne {
	if id != nil {
		puo = puo.SetProvContID(*id)
	}
	return puo
}

// SetProvCont sets the prov_cont edge to Continent.
func (puo *ProvinceUpdateOne) SetProvCont(c *Continent) *ProvinceUpdateOne {
	return puo.SetProvContID(c.ID)
}

// AddProvStudIDs adds the prov_stud edge to Student by ids.
func (puo *ProvinceUpdateOne) AddProvStudIDs(ids ...int) *ProvinceUpdateOne {
	puo.mutation.AddProvStudIDs(ids...)
	return puo
}

// AddProvStud adds the prov_stud edges to Student.
func (puo *ProvinceUpdateOne) AddProvStud(s ...*Student) *ProvinceUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddProvStudIDs(ids...)
}

// AddDistStudIDs adds the dist_stud edge to Student by ids.
func (puo *ProvinceUpdateOne) AddDistStudIDs(ids ...int) *ProvinceUpdateOne {
	puo.mutation.AddDistStudIDs(ids...)
	return puo
}

// AddDistStud adds the dist_stud edges to Student.
func (puo *ProvinceUpdateOne) AddDistStud(s ...*Student) *ProvinceUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddDistStudIDs(ids...)
}

// AddSubdStudIDs adds the subd_stud edge to Student by ids.
func (puo *ProvinceUpdateOne) AddSubdStudIDs(ids ...int) *ProvinceUpdateOne {
	puo.mutation.AddSubdStudIDs(ids...)
	return puo
}

// AddSubdStud adds the subd_stud edges to Student.
func (puo *ProvinceUpdateOne) AddSubdStud(s ...*Student) *ProvinceUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddSubdStudIDs(ids...)
}

// AddPostStudIDs adds the post_stud edge to Student by ids.
func (puo *ProvinceUpdateOne) AddPostStudIDs(ids ...int) *ProvinceUpdateOne {
	puo.mutation.AddPostStudIDs(ids...)
	return puo
}

// AddPostStud adds the post_stud edges to Student.
func (puo *ProvinceUpdateOne) AddPostStud(s ...*Student) *ProvinceUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddPostStudIDs(ids...)
}

// Mutation returns the ProvinceMutation object of the builder.
func (puo *ProvinceUpdateOne) Mutation() *ProvinceMutation {
	return puo.mutation
}

// ClearProvRegi clears the prov_regi edge to Region.
func (puo *ProvinceUpdateOne) ClearProvRegi() *ProvinceUpdateOne {
	puo.mutation.ClearProvRegi()
	return puo
}

// ClearProvCoun clears the prov_coun edge to Country.
func (puo *ProvinceUpdateOne) ClearProvCoun() *ProvinceUpdateOne {
	puo.mutation.ClearProvCoun()
	return puo
}

// ClearProvCont clears the prov_cont edge to Continent.
func (puo *ProvinceUpdateOne) ClearProvCont() *ProvinceUpdateOne {
	puo.mutation.ClearProvCont()
	return puo
}

// RemoveProvStudIDs removes the prov_stud edge to Student by ids.
func (puo *ProvinceUpdateOne) RemoveProvStudIDs(ids ...int) *ProvinceUpdateOne {
	puo.mutation.RemoveProvStudIDs(ids...)
	return puo
}

// RemoveProvStud removes prov_stud edges to Student.
func (puo *ProvinceUpdateOne) RemoveProvStud(s ...*Student) *ProvinceUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemoveProvStudIDs(ids...)
}

// RemoveDistStudIDs removes the dist_stud edge to Student by ids.
func (puo *ProvinceUpdateOne) RemoveDistStudIDs(ids ...int) *ProvinceUpdateOne {
	puo.mutation.RemoveDistStudIDs(ids...)
	return puo
}

// RemoveDistStud removes dist_stud edges to Student.
func (puo *ProvinceUpdateOne) RemoveDistStud(s ...*Student) *ProvinceUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemoveDistStudIDs(ids...)
}

// RemoveSubdStudIDs removes the subd_stud edge to Student by ids.
func (puo *ProvinceUpdateOne) RemoveSubdStudIDs(ids ...int) *ProvinceUpdateOne {
	puo.mutation.RemoveSubdStudIDs(ids...)
	return puo
}

// RemoveSubdStud removes subd_stud edges to Student.
func (puo *ProvinceUpdateOne) RemoveSubdStud(s ...*Student) *ProvinceUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemoveSubdStudIDs(ids...)
}

// RemovePostStudIDs removes the post_stud edge to Student by ids.
func (puo *ProvinceUpdateOne) RemovePostStudIDs(ids ...int) *ProvinceUpdateOne {
	puo.mutation.RemovePostStudIDs(ids...)
	return puo
}

// RemovePostStud removes post_stud edges to Student.
func (puo *ProvinceUpdateOne) RemovePostStud(s ...*Student) *ProvinceUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemovePostStudIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *ProvinceUpdateOne) Save(ctx context.Context) (*Province, error) {
	if v, ok := puo.mutation.Province(); ok {
		if err := province.ProvinceValidator(v); err != nil {
			return nil, &ValidationError{Name: "province", err: fmt.Errorf("ent: validator failed for field \"province\": %w", err)}
		}
	}
	if v, ok := puo.mutation.District(); ok {
		if err := province.DistrictValidator(v); err != nil {
			return nil, &ValidationError{Name: "district", err: fmt.Errorf("ent: validator failed for field \"district\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Subdistrict(); ok {
		if err := province.SubdistrictValidator(v); err != nil {
			return nil, &ValidationError{Name: "subdistrict", err: fmt.Errorf("ent: validator failed for field \"subdistrict\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Postal(); ok {
		if err := province.PostalValidator(v); err != nil {
			return nil, &ValidationError{Name: "postal", err: fmt.Errorf("ent: validator failed for field \"postal\": %w", err)}
		}
	}

	var (
		err  error
		node *Province
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProvinceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProvinceUpdateOne) SaveX(ctx context.Context) *Province {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *ProvinceUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProvinceUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProvinceUpdateOne) sqlSave(ctx context.Context) (pr *Province, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   province.Table,
			Columns: province.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: province.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Province.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Province(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: province.FieldProvince,
		})
	}
	if value, ok := puo.mutation.District(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: province.FieldDistrict,
		})
	}
	if value, ok := puo.mutation.Subdistrict(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: province.FieldSubdistrict,
		})
	}
	if value, ok := puo.mutation.Postal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: province.FieldPostal,
		})
	}
	if puo.mutation.ProvRegiCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   province.ProvRegiTable,
			Columns: []string{province.ProvRegiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: region.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProvRegiIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   province.ProvRegiTable,
			Columns: []string{province.ProvRegiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: region.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ProvCounCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   province.ProvCounTable,
			Columns: []string{province.ProvCounColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: country.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProvCounIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   province.ProvCounTable,
			Columns: []string{province.ProvCounColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: country.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ProvContCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   province.ProvContTable,
			Columns: []string{province.ProvContColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: continent.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProvContIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   province.ProvContTable,
			Columns: []string{province.ProvContColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: continent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := puo.mutation.RemovedProvStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.ProvStudTable,
			Columns: []string{province.ProvStudColumn},
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
	if nodes := puo.mutation.ProvStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.ProvStudTable,
			Columns: []string{province.ProvStudColumn},
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
	if nodes := puo.mutation.RemovedDistStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.DistStudTable,
			Columns: []string{province.DistStudColumn},
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
	if nodes := puo.mutation.DistStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.DistStudTable,
			Columns: []string{province.DistStudColumn},
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
	if nodes := puo.mutation.RemovedSubdStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.SubdStudTable,
			Columns: []string{province.SubdStudColumn},
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
	if nodes := puo.mutation.SubdStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.SubdStudTable,
			Columns: []string{province.SubdStudColumn},
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
	if nodes := puo.mutation.RemovedPostStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.PostStudTable,
			Columns: []string{province.PostStudColumn},
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
	if nodes := puo.mutation.PostStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.PostStudTable,
			Columns: []string{province.PostStudColumn},
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
	pr = &Province{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{province.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
