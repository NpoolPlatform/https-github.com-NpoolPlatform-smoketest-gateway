// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/planrelatedtestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// PlanRelatedTestCaseUpdate is the builder for updating PlanRelatedTestCase entities.
type PlanRelatedTestCaseUpdate struct {
	config
	hooks     []Hook
	mutation  *PlanRelatedTestCaseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PlanRelatedTestCaseUpdate builder.
func (prtcu *PlanRelatedTestCaseUpdate) Where(ps ...predicate.PlanRelatedTestCase) *PlanRelatedTestCaseUpdate {
	prtcu.mutation.Where(ps...)
	return prtcu
}

// SetCreatedAt sets the "created_at" field.
func (prtcu *PlanRelatedTestCaseUpdate) SetCreatedAt(u uint32) *PlanRelatedTestCaseUpdate {
	prtcu.mutation.ResetCreatedAt()
	prtcu.mutation.SetCreatedAt(u)
	return prtcu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (prtcu *PlanRelatedTestCaseUpdate) SetNillableCreatedAt(u *uint32) *PlanRelatedTestCaseUpdate {
	if u != nil {
		prtcu.SetCreatedAt(*u)
	}
	return prtcu
}

// AddCreatedAt adds u to the "created_at" field.
func (prtcu *PlanRelatedTestCaseUpdate) AddCreatedAt(u int32) *PlanRelatedTestCaseUpdate {
	prtcu.mutation.AddCreatedAt(u)
	return prtcu
}

// SetUpdatedAt sets the "updated_at" field.
func (prtcu *PlanRelatedTestCaseUpdate) SetUpdatedAt(u uint32) *PlanRelatedTestCaseUpdate {
	prtcu.mutation.ResetUpdatedAt()
	prtcu.mutation.SetUpdatedAt(u)
	return prtcu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (prtcu *PlanRelatedTestCaseUpdate) AddUpdatedAt(u int32) *PlanRelatedTestCaseUpdate {
	prtcu.mutation.AddUpdatedAt(u)
	return prtcu
}

// SetDeletedAt sets the "deleted_at" field.
func (prtcu *PlanRelatedTestCaseUpdate) SetDeletedAt(u uint32) *PlanRelatedTestCaseUpdate {
	prtcu.mutation.ResetDeletedAt()
	prtcu.mutation.SetDeletedAt(u)
	return prtcu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (prtcu *PlanRelatedTestCaseUpdate) SetNillableDeletedAt(u *uint32) *PlanRelatedTestCaseUpdate {
	if u != nil {
		prtcu.SetDeletedAt(*u)
	}
	return prtcu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (prtcu *PlanRelatedTestCaseUpdate) AddDeletedAt(u int32) *PlanRelatedTestCaseUpdate {
	prtcu.mutation.AddDeletedAt(u)
	return prtcu
}

// SetTestPlanID sets the "test_plan_id" field.
func (prtcu *PlanRelatedTestCaseUpdate) SetTestPlanID(u uuid.UUID) *PlanRelatedTestCaseUpdate {
	prtcu.mutation.SetTestPlanID(u)
	return prtcu
}

// SetNillableTestPlanID sets the "test_plan_id" field if the given value is not nil.
func (prtcu *PlanRelatedTestCaseUpdate) SetNillableTestPlanID(u *uuid.UUID) *PlanRelatedTestCaseUpdate {
	if u != nil {
		prtcu.SetTestPlanID(*u)
	}
	return prtcu
}

// ClearTestPlanID clears the value of the "test_plan_id" field.
func (prtcu *PlanRelatedTestCaseUpdate) ClearTestPlanID() *PlanRelatedTestCaseUpdate {
	prtcu.mutation.ClearTestPlanID()
	return prtcu
}

// SetTestCaseID sets the "test_case_id" field.
func (prtcu *PlanRelatedTestCaseUpdate) SetTestCaseID(u uuid.UUID) *PlanRelatedTestCaseUpdate {
	prtcu.mutation.SetTestCaseID(u)
	return prtcu
}

// SetNillableTestCaseID sets the "test_case_id" field if the given value is not nil.
func (prtcu *PlanRelatedTestCaseUpdate) SetNillableTestCaseID(u *uuid.UUID) *PlanRelatedTestCaseUpdate {
	if u != nil {
		prtcu.SetTestCaseID(*u)
	}
	return prtcu
}

// ClearTestCaseID clears the value of the "test_case_id" field.
func (prtcu *PlanRelatedTestCaseUpdate) ClearTestCaseID() *PlanRelatedTestCaseUpdate {
	prtcu.mutation.ClearTestCaseID()
	return prtcu
}

// SetTestCaseOutput sets the "test_case_output" field.
func (prtcu *PlanRelatedTestCaseUpdate) SetTestCaseOutput(s string) *PlanRelatedTestCaseUpdate {
	prtcu.mutation.SetTestCaseOutput(s)
	return prtcu
}

// SetNillableTestCaseOutput sets the "test_case_output" field if the given value is not nil.
func (prtcu *PlanRelatedTestCaseUpdate) SetNillableTestCaseOutput(s *string) *PlanRelatedTestCaseUpdate {
	if s != nil {
		prtcu.SetTestCaseOutput(*s)
	}
	return prtcu
}

// ClearTestCaseOutput clears the value of the "test_case_output" field.
func (prtcu *PlanRelatedTestCaseUpdate) ClearTestCaseOutput() *PlanRelatedTestCaseUpdate {
	prtcu.mutation.ClearTestCaseOutput()
	return prtcu
}

// SetDescription sets the "description" field.
func (prtcu *PlanRelatedTestCaseUpdate) SetDescription(s string) *PlanRelatedTestCaseUpdate {
	prtcu.mutation.SetDescription(s)
	return prtcu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (prtcu *PlanRelatedTestCaseUpdate) SetNillableDescription(s *string) *PlanRelatedTestCaseUpdate {
	if s != nil {
		prtcu.SetDescription(*s)
	}
	return prtcu
}

// ClearDescription clears the value of the "description" field.
func (prtcu *PlanRelatedTestCaseUpdate) ClearDescription() *PlanRelatedTestCaseUpdate {
	prtcu.mutation.ClearDescription()
	return prtcu
}

// SetTestUserID sets the "test_user_id" field.
func (prtcu *PlanRelatedTestCaseUpdate) SetTestUserID(u uuid.UUID) *PlanRelatedTestCaseUpdate {
	prtcu.mutation.SetTestUserID(u)
	return prtcu
}

// SetNillableTestUserID sets the "test_user_id" field if the given value is not nil.
func (prtcu *PlanRelatedTestCaseUpdate) SetNillableTestUserID(u *uuid.UUID) *PlanRelatedTestCaseUpdate {
	if u != nil {
		prtcu.SetTestUserID(*u)
	}
	return prtcu
}

// ClearTestUserID clears the value of the "test_user_id" field.
func (prtcu *PlanRelatedTestCaseUpdate) ClearTestUserID() *PlanRelatedTestCaseUpdate {
	prtcu.mutation.ClearTestUserID()
	return prtcu
}

// SetRunDuration sets the "run_duration" field.
func (prtcu *PlanRelatedTestCaseUpdate) SetRunDuration(u uint32) *PlanRelatedTestCaseUpdate {
	prtcu.mutation.ResetRunDuration()
	prtcu.mutation.SetRunDuration(u)
	return prtcu
}

// SetNillableRunDuration sets the "run_duration" field if the given value is not nil.
func (prtcu *PlanRelatedTestCaseUpdate) SetNillableRunDuration(u *uint32) *PlanRelatedTestCaseUpdate {
	if u != nil {
		prtcu.SetRunDuration(*u)
	}
	return prtcu
}

// AddRunDuration adds u to the "run_duration" field.
func (prtcu *PlanRelatedTestCaseUpdate) AddRunDuration(u int32) *PlanRelatedTestCaseUpdate {
	prtcu.mutation.AddRunDuration(u)
	return prtcu
}

// ClearRunDuration clears the value of the "run_duration" field.
func (prtcu *PlanRelatedTestCaseUpdate) ClearRunDuration() *PlanRelatedTestCaseUpdate {
	prtcu.mutation.ClearRunDuration()
	return prtcu
}

// SetTestCaseResult sets the "test_case_result" field.
func (prtcu *PlanRelatedTestCaseUpdate) SetTestCaseResult(s string) *PlanRelatedTestCaseUpdate {
	prtcu.mutation.SetTestCaseResult(s)
	return prtcu
}

// SetNillableTestCaseResult sets the "test_case_result" field if the given value is not nil.
func (prtcu *PlanRelatedTestCaseUpdate) SetNillableTestCaseResult(s *string) *PlanRelatedTestCaseUpdate {
	if s != nil {
		prtcu.SetTestCaseResult(*s)
	}
	return prtcu
}

// ClearTestCaseResult clears the value of the "test_case_result" field.
func (prtcu *PlanRelatedTestCaseUpdate) ClearTestCaseResult() *PlanRelatedTestCaseUpdate {
	prtcu.mutation.ClearTestCaseResult()
	return prtcu
}

// Mutation returns the PlanRelatedTestCaseMutation object of the builder.
func (prtcu *PlanRelatedTestCaseUpdate) Mutation() *PlanRelatedTestCaseMutation {
	return prtcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (prtcu *PlanRelatedTestCaseUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := prtcu.defaults(); err != nil {
		return 0, err
	}
	if len(prtcu.hooks) == 0 {
		affected, err = prtcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlanRelatedTestCaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			prtcu.mutation = mutation
			affected, err = prtcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(prtcu.hooks) - 1; i >= 0; i-- {
			if prtcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = prtcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, prtcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (prtcu *PlanRelatedTestCaseUpdate) SaveX(ctx context.Context) int {
	affected, err := prtcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (prtcu *PlanRelatedTestCaseUpdate) Exec(ctx context.Context) error {
	_, err := prtcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (prtcu *PlanRelatedTestCaseUpdate) ExecX(ctx context.Context) {
	if err := prtcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (prtcu *PlanRelatedTestCaseUpdate) defaults() error {
	if _, ok := prtcu.mutation.UpdatedAt(); !ok {
		if planrelatedtestcase.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized planrelatedtestcase.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := planrelatedtestcase.UpdateDefaultUpdatedAt()
		prtcu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (prtcu *PlanRelatedTestCaseUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PlanRelatedTestCaseUpdate {
	prtcu.modifiers = append(prtcu.modifiers, modifiers...)
	return prtcu
}

func (prtcu *PlanRelatedTestCaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   planrelatedtestcase.Table,
			Columns: planrelatedtestcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: planrelatedtestcase.FieldID,
			},
		},
	}
	if ps := prtcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := prtcu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: planrelatedtestcase.FieldCreatedAt,
		})
	}
	if value, ok := prtcu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: planrelatedtestcase.FieldCreatedAt,
		})
	}
	if value, ok := prtcu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: planrelatedtestcase.FieldUpdatedAt,
		})
	}
	if value, ok := prtcu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: planrelatedtestcase.FieldUpdatedAt,
		})
	}
	if value, ok := prtcu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: planrelatedtestcase.FieldDeletedAt,
		})
	}
	if value, ok := prtcu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: planrelatedtestcase.FieldDeletedAt,
		})
	}
	if value, ok := prtcu.mutation.TestPlanID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: planrelatedtestcase.FieldTestPlanID,
		})
	}
	if prtcu.mutation.TestPlanIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: planrelatedtestcase.FieldTestPlanID,
		})
	}
	if value, ok := prtcu.mutation.TestCaseID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: planrelatedtestcase.FieldTestCaseID,
		})
	}
	if prtcu.mutation.TestCaseIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: planrelatedtestcase.FieldTestCaseID,
		})
	}
	if value, ok := prtcu.mutation.TestCaseOutput(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: planrelatedtestcase.FieldTestCaseOutput,
		})
	}
	if prtcu.mutation.TestCaseOutputCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: planrelatedtestcase.FieldTestCaseOutput,
		})
	}
	if value, ok := prtcu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: planrelatedtestcase.FieldDescription,
		})
	}
	if prtcu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: planrelatedtestcase.FieldDescription,
		})
	}
	if value, ok := prtcu.mutation.TestUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: planrelatedtestcase.FieldTestUserID,
		})
	}
	if prtcu.mutation.TestUserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: planrelatedtestcase.FieldTestUserID,
		})
	}
	if value, ok := prtcu.mutation.RunDuration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: planrelatedtestcase.FieldRunDuration,
		})
	}
	if value, ok := prtcu.mutation.AddedRunDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: planrelatedtestcase.FieldRunDuration,
		})
	}
	if prtcu.mutation.RunDurationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: planrelatedtestcase.FieldRunDuration,
		})
	}
	if value, ok := prtcu.mutation.TestCaseResult(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: planrelatedtestcase.FieldTestCaseResult,
		})
	}
	if prtcu.mutation.TestCaseResultCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: planrelatedtestcase.FieldTestCaseResult,
		})
	}
	_spec.Modifiers = prtcu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, prtcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{planrelatedtestcase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PlanRelatedTestCaseUpdateOne is the builder for updating a single PlanRelatedTestCase entity.
type PlanRelatedTestCaseUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PlanRelatedTestCaseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetCreatedAt(u uint32) *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.ResetCreatedAt()
	prtcuo.mutation.SetCreatedAt(u)
	return prtcuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetNillableCreatedAt(u *uint32) *PlanRelatedTestCaseUpdateOne {
	if u != nil {
		prtcuo.SetCreatedAt(*u)
	}
	return prtcuo
}

// AddCreatedAt adds u to the "created_at" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) AddCreatedAt(u int32) *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.AddCreatedAt(u)
	return prtcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetUpdatedAt(u uint32) *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.ResetUpdatedAt()
	prtcuo.mutation.SetUpdatedAt(u)
	return prtcuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) AddUpdatedAt(u int32) *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.AddUpdatedAt(u)
	return prtcuo
}

// SetDeletedAt sets the "deleted_at" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetDeletedAt(u uint32) *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.ResetDeletedAt()
	prtcuo.mutation.SetDeletedAt(u)
	return prtcuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetNillableDeletedAt(u *uint32) *PlanRelatedTestCaseUpdateOne {
	if u != nil {
		prtcuo.SetDeletedAt(*u)
	}
	return prtcuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) AddDeletedAt(u int32) *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.AddDeletedAt(u)
	return prtcuo
}

// SetTestPlanID sets the "test_plan_id" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetTestPlanID(u uuid.UUID) *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.SetTestPlanID(u)
	return prtcuo
}

// SetNillableTestPlanID sets the "test_plan_id" field if the given value is not nil.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetNillableTestPlanID(u *uuid.UUID) *PlanRelatedTestCaseUpdateOne {
	if u != nil {
		prtcuo.SetTestPlanID(*u)
	}
	return prtcuo
}

// ClearTestPlanID clears the value of the "test_plan_id" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) ClearTestPlanID() *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.ClearTestPlanID()
	return prtcuo
}

// SetTestCaseID sets the "test_case_id" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetTestCaseID(u uuid.UUID) *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.SetTestCaseID(u)
	return prtcuo
}

// SetNillableTestCaseID sets the "test_case_id" field if the given value is not nil.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetNillableTestCaseID(u *uuid.UUID) *PlanRelatedTestCaseUpdateOne {
	if u != nil {
		prtcuo.SetTestCaseID(*u)
	}
	return prtcuo
}

// ClearTestCaseID clears the value of the "test_case_id" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) ClearTestCaseID() *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.ClearTestCaseID()
	return prtcuo
}

// SetTestCaseOutput sets the "test_case_output" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetTestCaseOutput(s string) *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.SetTestCaseOutput(s)
	return prtcuo
}

// SetNillableTestCaseOutput sets the "test_case_output" field if the given value is not nil.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetNillableTestCaseOutput(s *string) *PlanRelatedTestCaseUpdateOne {
	if s != nil {
		prtcuo.SetTestCaseOutput(*s)
	}
	return prtcuo
}

// ClearTestCaseOutput clears the value of the "test_case_output" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) ClearTestCaseOutput() *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.ClearTestCaseOutput()
	return prtcuo
}

// SetDescription sets the "description" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetDescription(s string) *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.SetDescription(s)
	return prtcuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetNillableDescription(s *string) *PlanRelatedTestCaseUpdateOne {
	if s != nil {
		prtcuo.SetDescription(*s)
	}
	return prtcuo
}

// ClearDescription clears the value of the "description" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) ClearDescription() *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.ClearDescription()
	return prtcuo
}

// SetTestUserID sets the "test_user_id" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetTestUserID(u uuid.UUID) *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.SetTestUserID(u)
	return prtcuo
}

// SetNillableTestUserID sets the "test_user_id" field if the given value is not nil.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetNillableTestUserID(u *uuid.UUID) *PlanRelatedTestCaseUpdateOne {
	if u != nil {
		prtcuo.SetTestUserID(*u)
	}
	return prtcuo
}

// ClearTestUserID clears the value of the "test_user_id" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) ClearTestUserID() *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.ClearTestUserID()
	return prtcuo
}

// SetRunDuration sets the "run_duration" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetRunDuration(u uint32) *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.ResetRunDuration()
	prtcuo.mutation.SetRunDuration(u)
	return prtcuo
}

// SetNillableRunDuration sets the "run_duration" field if the given value is not nil.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetNillableRunDuration(u *uint32) *PlanRelatedTestCaseUpdateOne {
	if u != nil {
		prtcuo.SetRunDuration(*u)
	}
	return prtcuo
}

// AddRunDuration adds u to the "run_duration" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) AddRunDuration(u int32) *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.AddRunDuration(u)
	return prtcuo
}

// ClearRunDuration clears the value of the "run_duration" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) ClearRunDuration() *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.ClearRunDuration()
	return prtcuo
}

// SetTestCaseResult sets the "test_case_result" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetTestCaseResult(s string) *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.SetTestCaseResult(s)
	return prtcuo
}

// SetNillableTestCaseResult sets the "test_case_result" field if the given value is not nil.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SetNillableTestCaseResult(s *string) *PlanRelatedTestCaseUpdateOne {
	if s != nil {
		prtcuo.SetTestCaseResult(*s)
	}
	return prtcuo
}

// ClearTestCaseResult clears the value of the "test_case_result" field.
func (prtcuo *PlanRelatedTestCaseUpdateOne) ClearTestCaseResult() *PlanRelatedTestCaseUpdateOne {
	prtcuo.mutation.ClearTestCaseResult()
	return prtcuo
}

// Mutation returns the PlanRelatedTestCaseMutation object of the builder.
func (prtcuo *PlanRelatedTestCaseUpdateOne) Mutation() *PlanRelatedTestCaseMutation {
	return prtcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (prtcuo *PlanRelatedTestCaseUpdateOne) Select(field string, fields ...string) *PlanRelatedTestCaseUpdateOne {
	prtcuo.fields = append([]string{field}, fields...)
	return prtcuo
}

// Save executes the query and returns the updated PlanRelatedTestCase entity.
func (prtcuo *PlanRelatedTestCaseUpdateOne) Save(ctx context.Context) (*PlanRelatedTestCase, error) {
	var (
		err  error
		node *PlanRelatedTestCase
	)
	if err := prtcuo.defaults(); err != nil {
		return nil, err
	}
	if len(prtcuo.hooks) == 0 {
		node, err = prtcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlanRelatedTestCaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			prtcuo.mutation = mutation
			node, err = prtcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(prtcuo.hooks) - 1; i >= 0; i-- {
			if prtcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = prtcuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, prtcuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PlanRelatedTestCase)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PlanRelatedTestCaseMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (prtcuo *PlanRelatedTestCaseUpdateOne) SaveX(ctx context.Context) *PlanRelatedTestCase {
	node, err := prtcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (prtcuo *PlanRelatedTestCaseUpdateOne) Exec(ctx context.Context) error {
	_, err := prtcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (prtcuo *PlanRelatedTestCaseUpdateOne) ExecX(ctx context.Context) {
	if err := prtcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (prtcuo *PlanRelatedTestCaseUpdateOne) defaults() error {
	if _, ok := prtcuo.mutation.UpdatedAt(); !ok {
		if planrelatedtestcase.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized planrelatedtestcase.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := planrelatedtestcase.UpdateDefaultUpdatedAt()
		prtcuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (prtcuo *PlanRelatedTestCaseUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PlanRelatedTestCaseUpdateOne {
	prtcuo.modifiers = append(prtcuo.modifiers, modifiers...)
	return prtcuo
}

func (prtcuo *PlanRelatedTestCaseUpdateOne) sqlSave(ctx context.Context) (_node *PlanRelatedTestCase, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   planrelatedtestcase.Table,
			Columns: planrelatedtestcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: planrelatedtestcase.FieldID,
			},
		},
	}
	id, ok := prtcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PlanRelatedTestCase.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := prtcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, planrelatedtestcase.FieldID)
		for _, f := range fields {
			if !planrelatedtestcase.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != planrelatedtestcase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := prtcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := prtcuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: planrelatedtestcase.FieldCreatedAt,
		})
	}
	if value, ok := prtcuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: planrelatedtestcase.FieldCreatedAt,
		})
	}
	if value, ok := prtcuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: planrelatedtestcase.FieldUpdatedAt,
		})
	}
	if value, ok := prtcuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: planrelatedtestcase.FieldUpdatedAt,
		})
	}
	if value, ok := prtcuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: planrelatedtestcase.FieldDeletedAt,
		})
	}
	if value, ok := prtcuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: planrelatedtestcase.FieldDeletedAt,
		})
	}
	if value, ok := prtcuo.mutation.TestPlanID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: planrelatedtestcase.FieldTestPlanID,
		})
	}
	if prtcuo.mutation.TestPlanIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: planrelatedtestcase.FieldTestPlanID,
		})
	}
	if value, ok := prtcuo.mutation.TestCaseID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: planrelatedtestcase.FieldTestCaseID,
		})
	}
	if prtcuo.mutation.TestCaseIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: planrelatedtestcase.FieldTestCaseID,
		})
	}
	if value, ok := prtcuo.mutation.TestCaseOutput(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: planrelatedtestcase.FieldTestCaseOutput,
		})
	}
	if prtcuo.mutation.TestCaseOutputCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: planrelatedtestcase.FieldTestCaseOutput,
		})
	}
	if value, ok := prtcuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: planrelatedtestcase.FieldDescription,
		})
	}
	if prtcuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: planrelatedtestcase.FieldDescription,
		})
	}
	if value, ok := prtcuo.mutation.TestUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: planrelatedtestcase.FieldTestUserID,
		})
	}
	if prtcuo.mutation.TestUserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: planrelatedtestcase.FieldTestUserID,
		})
	}
	if value, ok := prtcuo.mutation.RunDuration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: planrelatedtestcase.FieldRunDuration,
		})
	}
	if value, ok := prtcuo.mutation.AddedRunDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: planrelatedtestcase.FieldRunDuration,
		})
	}
	if prtcuo.mutation.RunDurationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: planrelatedtestcase.FieldRunDuration,
		})
	}
	if value, ok := prtcuo.mutation.TestCaseResult(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: planrelatedtestcase.FieldTestCaseResult,
		})
	}
	if prtcuo.mutation.TestCaseResultCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: planrelatedtestcase.FieldTestCaseResult,
		})
	}
	_spec.Modifiers = prtcuo.modifiers
	_node = &PlanRelatedTestCase{config: prtcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, prtcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{planrelatedtestcase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}