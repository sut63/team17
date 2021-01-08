// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/sut63/team17/app/ent"
)

// The ActivityFunc type is an adapter to allow the use of ordinary
// function as Activity mutator.
type ActivityFunc func(context.Context, *ent.ActivityMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ActivityFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ActivityMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ActivityMutation", m)
	}
	return f(ctx, mv)
}

// The AgencyFunc type is an adapter to allow the use of ordinary
// function as Agency mutator.
type AgencyFunc func(context.Context, *ent.AgencyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AgencyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AgencyMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AgencyMutation", m)
	}
	return f(ctx, mv)
}

// The CourseFunc type is an adapter to allow the use of ordinary
// function as Course mutator.
type CourseFunc func(context.Context, *ent.CourseMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CourseFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CourseMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CourseMutation", m)
	}
	return f(ctx, mv)
}

// The DegreeFunc type is an adapter to allow the use of ordinary
// function as Degree mutator.
type DegreeFunc func(context.Context, *ent.DegreeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DegreeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DegreeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DegreeMutation", m)
	}
	return f(ctx, mv)
}

// The DistrictFunc type is an adapter to allow the use of ordinary
// function as District mutator.
type DistrictFunc func(context.Context, *ent.DistrictMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DistrictFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DistrictMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DistrictMutation", m)
	}
	return f(ctx, mv)
}

// The FacultyFunc type is an adapter to allow the use of ordinary
// function as Faculty mutator.
type FacultyFunc func(context.Context, *ent.FacultyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FacultyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FacultyMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FacultyMutation", m)
	}
	return f(ctx, mv)
}

// The GenderFunc type is an adapter to allow the use of ordinary
// function as Gender mutator.
type GenderFunc func(context.Context, *ent.GenderMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GenderFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GenderMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GenderMutation", m)
	}
	return f(ctx, mv)
}

// The InstitutionFunc type is an adapter to allow the use of ordinary
// function as Institution mutator.
type InstitutionFunc func(context.Context, *ent.InstitutionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f InstitutionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.InstitutionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.InstitutionMutation", m)
	}
	return f(ctx, mv)
}

// The PlaceFunc type is an adapter to allow the use of ordinary
// function as Place mutator.
type PlaceFunc func(context.Context, *ent.PlaceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PlaceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PlaceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PlaceMutation", m)
	}
	return f(ctx, mv)
}

// The PostalFunc type is an adapter to allow the use of ordinary
// function as Postal mutator.
type PostalFunc func(context.Context, *ent.PostalMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PostalFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PostalMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PostalMutation", m)
	}
	return f(ctx, mv)
}

// The PrefixFunc type is an adapter to allow the use of ordinary
// function as Prefix mutator.
type PrefixFunc func(context.Context, *ent.PrefixMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PrefixFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PrefixMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PrefixMutation", m)
	}
	return f(ctx, mv)
}

// The ProfessorFunc type is an adapter to allow the use of ordinary
// function as Professor mutator.
type ProfessorFunc func(context.Context, *ent.ProfessorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProfessorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProfessorMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProfessorMutation", m)
	}
	return f(ctx, mv)
}

// The ProfessorshipFunc type is an adapter to allow the use of ordinary
// function as Professorship mutator.
type ProfessorshipFunc func(context.Context, *ent.ProfessorshipMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProfessorshipFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProfessorshipMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProfessorshipMutation", m)
	}
	return f(ctx, mv)
}

// The ProvinceFunc type is an adapter to allow the use of ordinary
// function as Province mutator.
type ProvinceFunc func(context.Context, *ent.ProvinceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProvinceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProvinceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProvinceMutation", m)
	}
	return f(ctx, mv)
}

// The RegionFunc type is an adapter to allow the use of ordinary
// function as Region mutator.
type RegionFunc func(context.Context, *ent.RegionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RegionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RegionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RegionMutation", m)
	}
	return f(ctx, mv)
}

// The ResultsFunc type is an adapter to allow the use of ordinary
// function as Results mutator.
type ResultsFunc func(context.Context, *ent.ResultsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ResultsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ResultsMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ResultsMutation", m)
	}
	return f(ctx, mv)
}

// The StudentFunc type is an adapter to allow the use of ordinary
// function as Student mutator.
type StudentFunc func(context.Context, *ent.StudentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StudentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StudentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StudentMutation", m)
	}
	return f(ctx, mv)
}

// The SubdistrictFunc type is an adapter to allow the use of ordinary
// function as Subdistrict mutator.
type SubdistrictFunc func(context.Context, *ent.SubdistrictMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SubdistrictFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SubdistrictMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SubdistrictMutation", m)
	}
	return f(ctx, mv)
}

// The SubjectFunc type is an adapter to allow the use of ordinary
// function as Subject mutator.
type SubjectFunc func(context.Context, *ent.SubjectMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SubjectFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SubjectMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SubjectMutation", m)
	}
	return f(ctx, mv)
}

// The TermFunc type is an adapter to allow the use of ordinary
// function as Term mutator.
type TermFunc func(context.Context, *ent.TermMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TermFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TermMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TermMutation", m)
	}
	return f(ctx, mv)
}

// The YearFunc type is an adapter to allow the use of ordinary
// function as Year mutator.
type YearFunc func(context.Context, *ent.YearMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f YearFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.YearMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.YearMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	Hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(_ context.Context, m ent.Mutation) (ent.Value, error) {
			return nil, fmt.Errorf("%s operation is not allowed", m.Op())
		})
	}
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
