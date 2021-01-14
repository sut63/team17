// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"errors"
	"fmt"

	"github.com/sut63/team17/app/ent"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = errors.New("ent/privacy: allow rule")

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = errors.New("ent/privacy: deny rule")

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = errors.New("ent/privacy: skip rule")
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

type decisionCtxKey struct{}

// DecisionContext creates a decision context.
func DecisionContext(parent context.Context, decision error) context.Context {
	if decision == nil || errors.Is(decision, Skip) {
		return parent
	}
	return context.WithValue(parent, decisionCtxKey{}, decision)
}

func decisionFromContext(ctx context.Context) (error, bool) {
	decision, ok := ctx.Value(decisionCtxKey{}).(error)
	if ok && errors.Is(decision, Allow) {
		decision = nil
	}
	return decision, ok
}

type (
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy []QueryRule

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule interface {
		EvalQuery(context.Context, ent.Query) error
	}
)

// EvalQuery evaluates a query against a query policy.
func (policy QueryPolicy) EvalQuery(ctx context.Context, q ent.Query) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalQuery(ctx, q); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

type (
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy []MutationRule

	// MutationRule defines the interface deciding whether a
	// mutation is allowed and optionally modify it.
	MutationRule interface {
		EvalMutation(context.Context, ent.Mutation) error
	}
)

// EvalMutation evaluates a mutation against a mutation policy.
func (policy MutationPolicy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalMutation(ctx, m); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// MutationRuleFunc type is an adapter to allow the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// Policy groups query and mutation policies.
type Policy struct {
	Query    QueryPolicy
	Mutation MutationPolicy
}

// EvalQuery forwards evaluation to query policy.
func (policy Policy) EvalQuery(ctx context.Context, q ent.Query) error {
	return policy.Query.EvalQuery(ctx, q)
}

// EvalMutation forwards evaluation to mutation policy.
func (policy Policy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return policy.Mutation.EvalMutation(ctx, m)
}

// QueryMutationRule is the interface that groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The ActivityQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ActivityQueryRuleFunc func(context.Context, *ent.ActivityQuery) error

// EvalQuery return f(ctx, q).
func (f ActivityQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ActivityQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ActivityQuery", q)
}

// The ActivityMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ActivityMutationRuleFunc func(context.Context, *ent.ActivityMutation) error

// EvalMutation calls f(ctx, m).
func (f ActivityMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ActivityMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ActivityMutation", m)
}

// The AgencyQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AgencyQueryRuleFunc func(context.Context, *ent.AgencyQuery) error

// EvalQuery return f(ctx, q).
func (f AgencyQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AgencyQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AgencyQuery", q)
}

// The AgencyMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AgencyMutationRuleFunc func(context.Context, *ent.AgencyMutation) error

// EvalMutation calls f(ctx, m).
func (f AgencyMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AgencyMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AgencyMutation", m)
}

// The ContinentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ContinentQueryRuleFunc func(context.Context, *ent.ContinentQuery) error

// EvalQuery return f(ctx, q).
func (f ContinentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ContinentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ContinentQuery", q)
}

// The ContinentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ContinentMutationRuleFunc func(context.Context, *ent.ContinentMutation) error

// EvalMutation calls f(ctx, m).
func (f ContinentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ContinentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ContinentMutation", m)
}

// The CountryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CountryQueryRuleFunc func(context.Context, *ent.CountryQuery) error

// EvalQuery return f(ctx, q).
func (f CountryQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CountryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CountryQuery", q)
}

// The CountryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CountryMutationRuleFunc func(context.Context, *ent.CountryMutation) error

// EvalMutation calls f(ctx, m).
func (f CountryMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CountryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CountryMutation", m)
}

// The CourseQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CourseQueryRuleFunc func(context.Context, *ent.CourseQuery) error

// EvalQuery return f(ctx, q).
func (f CourseQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CourseQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CourseQuery", q)
}

// The CourseMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CourseMutationRuleFunc func(context.Context, *ent.CourseMutation) error

// EvalMutation calls f(ctx, m).
func (f CourseMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CourseMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CourseMutation", m)
}

// The DegreeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DegreeQueryRuleFunc func(context.Context, *ent.DegreeQuery) error

// EvalQuery return f(ctx, q).
func (f DegreeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DegreeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DegreeQuery", q)
}

// The DegreeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DegreeMutationRuleFunc func(context.Context, *ent.DegreeMutation) error

// EvalMutation calls f(ctx, m).
func (f DegreeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DegreeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DegreeMutation", m)
}

// The EmpQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EmpQueryRuleFunc func(context.Context, *ent.EmpQuery) error

// EvalQuery return f(ctx, q).
func (f EmpQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EmpQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.EmpQuery", q)
}

// The EmpMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EmpMutationRuleFunc func(context.Context, *ent.EmpMutation) error

// EvalMutation calls f(ctx, m).
func (f EmpMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.EmpMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.EmpMutation", m)
}

// The FacultyQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type FacultyQueryRuleFunc func(context.Context, *ent.FacultyQuery) error

// EvalQuery return f(ctx, q).
func (f FacultyQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FacultyQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.FacultyQuery", q)
}

// The FacultyMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type FacultyMutationRuleFunc func(context.Context, *ent.FacultyMutation) error

// EvalMutation calls f(ctx, m).
func (f FacultyMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.FacultyMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.FacultyMutation", m)
}

// The GenderQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GenderQueryRuleFunc func(context.Context, *ent.GenderQuery) error

// EvalQuery return f(ctx, q).
func (f GenderQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GenderQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GenderQuery", q)
}

// The GenderMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GenderMutationRuleFunc func(context.Context, *ent.GenderMutation) error

// EvalMutation calls f(ctx, m).
func (f GenderMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GenderMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GenderMutation", m)
}

// The InstitutionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type InstitutionQueryRuleFunc func(context.Context, *ent.InstitutionQuery) error

// EvalQuery return f(ctx, q).
func (f InstitutionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.InstitutionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.InstitutionQuery", q)
}

// The InstitutionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type InstitutionMutationRuleFunc func(context.Context, *ent.InstitutionMutation) error

// EvalMutation calls f(ctx, m).
func (f InstitutionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.InstitutionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.InstitutionMutation", m)
}

// The PlaceQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PlaceQueryRuleFunc func(context.Context, *ent.PlaceQuery) error

// EvalQuery return f(ctx, q).
func (f PlaceQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PlaceQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PlaceQuery", q)
}

// The PlaceMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PlaceMutationRuleFunc func(context.Context, *ent.PlaceMutation) error

// EvalMutation calls f(ctx, m).
func (f PlaceMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PlaceMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PlaceMutation", m)
}

// The PrefixQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PrefixQueryRuleFunc func(context.Context, *ent.PrefixQuery) error

// EvalQuery return f(ctx, q).
func (f PrefixQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PrefixQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PrefixQuery", q)
}

// The PrefixMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PrefixMutationRuleFunc func(context.Context, *ent.PrefixMutation) error

// EvalMutation calls f(ctx, m).
func (f PrefixMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PrefixMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PrefixMutation", m)
}

// The ProfessorQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ProfessorQueryRuleFunc func(context.Context, *ent.ProfessorQuery) error

// EvalQuery return f(ctx, q).
func (f ProfessorQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ProfessorQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ProfessorQuery", q)
}

// The ProfessorMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ProfessorMutationRuleFunc func(context.Context, *ent.ProfessorMutation) error

// EvalMutation calls f(ctx, m).
func (f ProfessorMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ProfessorMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ProfessorMutation", m)
}

// The ProfessorshipQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ProfessorshipQueryRuleFunc func(context.Context, *ent.ProfessorshipQuery) error

// EvalQuery return f(ctx, q).
func (f ProfessorshipQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ProfessorshipQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ProfessorshipQuery", q)
}

// The ProfessorshipMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ProfessorshipMutationRuleFunc func(context.Context, *ent.ProfessorshipMutation) error

// EvalMutation calls f(ctx, m).
func (f ProfessorshipMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ProfessorshipMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ProfessorshipMutation", m)
}

// The ProvinceQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ProvinceQueryRuleFunc func(context.Context, *ent.ProvinceQuery) error

// EvalQuery return f(ctx, q).
func (f ProvinceQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ProvinceQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ProvinceQuery", q)
}

// The ProvinceMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ProvinceMutationRuleFunc func(context.Context, *ent.ProvinceMutation) error

// EvalMutation calls f(ctx, m).
func (f ProvinceMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ProvinceMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ProvinceMutation", m)
}

// The RegionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RegionQueryRuleFunc func(context.Context, *ent.RegionQuery) error

// EvalQuery return f(ctx, q).
func (f RegionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RegionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RegionQuery", q)
}

// The RegionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RegionMutationRuleFunc func(context.Context, *ent.RegionMutation) error

// EvalMutation calls f(ctx, m).
func (f RegionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RegionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RegionMutation", m)
}

// The ResultsQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ResultsQueryRuleFunc func(context.Context, *ent.ResultsQuery) error

// EvalQuery return f(ctx, q).
func (f ResultsQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ResultsQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ResultsQuery", q)
}

// The ResultsMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ResultsMutationRuleFunc func(context.Context, *ent.ResultsMutation) error

// EvalMutation calls f(ctx, m).
func (f ResultsMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ResultsMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ResultsMutation", m)
}

// The StudentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type StudentQueryRuleFunc func(context.Context, *ent.StudentQuery) error

// EvalQuery return f(ctx, q).
func (f StudentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.StudentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.StudentQuery", q)
}

// The StudentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type StudentMutationRuleFunc func(context.Context, *ent.StudentMutation) error

// EvalMutation calls f(ctx, m).
func (f StudentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.StudentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.StudentMutation", m)
}

// The SubjectQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SubjectQueryRuleFunc func(context.Context, *ent.SubjectQuery) error

// EvalQuery return f(ctx, q).
func (f SubjectQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SubjectQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SubjectQuery", q)
}

// The SubjectMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SubjectMutationRuleFunc func(context.Context, *ent.SubjectMutation) error

// EvalMutation calls f(ctx, m).
func (f SubjectMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SubjectMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SubjectMutation", m)
}

// The TermQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TermQueryRuleFunc func(context.Context, *ent.TermQuery) error

// EvalQuery return f(ctx, q).
func (f TermQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TermQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TermQuery", q)
}

// The TermMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TermMutationRuleFunc func(context.Context, *ent.TermMutation) error

// EvalMutation calls f(ctx, m).
func (f TermMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TermMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TermMutation", m)
}

// The YearQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type YearQueryRuleFunc func(context.Context, *ent.YearQuery) error

// EvalQuery return f(ctx, q).
func (f YearQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.YearQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.YearQuery", q)
}

// The YearMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type YearMutationRuleFunc func(context.Context, *ent.YearMutation) error

// EvalMutation calls f(ctx, m).
func (f YearMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.YearMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.YearMutation", m)
}
