// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/relatedtestcase"
	"github.com/google/uuid"
)

// RelatedTestCaseUpdate is the builder for updating RelatedTestCase entities.
type RelatedTestCaseUpdate struct {
	config
	hooks     []Hook
	mutation  *RelatedTestCaseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the RelatedTestCaseUpdate builder.
func (rtcu *RelatedTestCaseUpdate) Where(ps ...predicate.RelatedTestCase) *RelatedTestCaseUpdate {
	rtcu.mutation.Where(ps...)
	return rtcu
}

// SetCreatedAt sets the "created_at" field.
func (rtcu *RelatedTestCaseUpdate) SetCreatedAt(u uint32) *RelatedTestCaseUpdate {
	rtcu.mutation.ResetCreatedAt()
	rtcu.mutation.SetCreatedAt(u)
	return rtcu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rtcu *RelatedTestCaseUpdate) SetNillableCreatedAt(u *uint32) *RelatedTestCaseUpdate {
	if u != nil {
		rtcu.SetCreatedAt(*u)
	}
	return rtcu
}

// AddCreatedAt adds u to the "created_at" field.
func (rtcu *RelatedTestCaseUpdate) AddCreatedAt(u int32) *RelatedTestCaseUpdate {
	rtcu.mutation.AddCreatedAt(u)
	return rtcu
}

// SetUpdatedAt sets the "updated_at" field.
func (rtcu *RelatedTestCaseUpdate) SetUpdatedAt(u uint32) *RelatedTestCaseUpdate {
	rtcu.mutation.ResetUpdatedAt()
	rtcu.mutation.SetUpdatedAt(u)
	return rtcu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (rtcu *RelatedTestCaseUpdate) AddUpdatedAt(u int32) *RelatedTestCaseUpdate {
	rtcu.mutation.AddUpdatedAt(u)
	return rtcu
}

// SetDeletedAt sets the "deleted_at" field.
func (rtcu *RelatedTestCaseUpdate) SetDeletedAt(u uint32) *RelatedTestCaseUpdate {
	rtcu.mutation.ResetDeletedAt()
	rtcu.mutation.SetDeletedAt(u)
	return rtcu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rtcu *RelatedTestCaseUpdate) SetNillableDeletedAt(u *uint32) *RelatedTestCaseUpdate {
	if u != nil {
		rtcu.SetDeletedAt(*u)
	}
	return rtcu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (rtcu *RelatedTestCaseUpdate) AddDeletedAt(u int32) *RelatedTestCaseUpdate {
	rtcu.mutation.AddDeletedAt(u)
	return rtcu
}

// SetCondType sets the "cond_type" field.
func (rtcu *RelatedTestCaseUpdate) SetCondType(s string) *RelatedTestCaseUpdate {
	rtcu.mutation.SetCondType(s)
	return rtcu
}

// SetNillableCondType sets the "cond_type" field if the given value is not nil.
func (rtcu *RelatedTestCaseUpdate) SetNillableCondType(s *string) *RelatedTestCaseUpdate {
	if s != nil {
		rtcu.SetCondType(*s)
	}
	return rtcu
}

// ClearCondType clears the value of the "cond_type" field.
func (rtcu *RelatedTestCaseUpdate) ClearCondType() *RelatedTestCaseUpdate {
	rtcu.mutation.ClearCondType()
	return rtcu
}

// SetTestCaseID sets the "test_case_id" field.
func (rtcu *RelatedTestCaseUpdate) SetTestCaseID(u uuid.UUID) *RelatedTestCaseUpdate {
	rtcu.mutation.SetTestCaseID(u)
	return rtcu
}

// SetNillableTestCaseID sets the "test_case_id" field if the given value is not nil.
func (rtcu *RelatedTestCaseUpdate) SetNillableTestCaseID(u *uuid.UUID) *RelatedTestCaseUpdate {
	if u != nil {
		rtcu.SetTestCaseID(*u)
	}
	return rtcu
}

// ClearTestCaseID clears the value of the "test_case_id" field.
func (rtcu *RelatedTestCaseUpdate) ClearTestCaseID() *RelatedTestCaseUpdate {
	rtcu.mutation.ClearTestCaseID()
	return rtcu
}

// SetRelatedTestCaseID sets the "related_test_case_id" field.
func (rtcu *RelatedTestCaseUpdate) SetRelatedTestCaseID(u uuid.UUID) *RelatedTestCaseUpdate {
	rtcu.mutation.SetRelatedTestCaseID(u)
	return rtcu
}

// SetNillableRelatedTestCaseID sets the "related_test_case_id" field if the given value is not nil.
func (rtcu *RelatedTestCaseUpdate) SetNillableRelatedTestCaseID(u *uuid.UUID) *RelatedTestCaseUpdate {
	if u != nil {
		rtcu.SetRelatedTestCaseID(*u)
	}
	return rtcu
}

// ClearRelatedTestCaseID clears the value of the "related_test_case_id" field.
func (rtcu *RelatedTestCaseUpdate) ClearRelatedTestCaseID() *RelatedTestCaseUpdate {
	rtcu.mutation.ClearRelatedTestCaseID()
	return rtcu
}

// SetArgumentsTransfer sets the "arguments_transfer" field.
func (rtcu *RelatedTestCaseUpdate) SetArgumentsTransfer(s string) *RelatedTestCaseUpdate {
	rtcu.mutation.SetArgumentsTransfer(s)
	return rtcu
}

// SetNillableArgumentsTransfer sets the "arguments_transfer" field if the given value is not nil.
func (rtcu *RelatedTestCaseUpdate) SetNillableArgumentsTransfer(s *string) *RelatedTestCaseUpdate {
	if s != nil {
		rtcu.SetArgumentsTransfer(*s)
	}
	return rtcu
}

// ClearArgumentsTransfer clears the value of the "arguments_transfer" field.
func (rtcu *RelatedTestCaseUpdate) ClearArgumentsTransfer() *RelatedTestCaseUpdate {
	rtcu.mutation.ClearArgumentsTransfer()
	return rtcu
}

// SetIndex sets the "index" field.
func (rtcu *RelatedTestCaseUpdate) SetIndex(u uint32) *RelatedTestCaseUpdate {
	rtcu.mutation.ResetIndex()
	rtcu.mutation.SetIndex(u)
	return rtcu
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (rtcu *RelatedTestCaseUpdate) SetNillableIndex(u *uint32) *RelatedTestCaseUpdate {
	if u != nil {
		rtcu.SetIndex(*u)
	}
	return rtcu
}

// AddIndex adds u to the "index" field.
func (rtcu *RelatedTestCaseUpdate) AddIndex(u int32) *RelatedTestCaseUpdate {
	rtcu.mutation.AddIndex(u)
	return rtcu
}

// ClearIndex clears the value of the "index" field.
func (rtcu *RelatedTestCaseUpdate) ClearIndex() *RelatedTestCaseUpdate {
	rtcu.mutation.ClearIndex()
	return rtcu
}

// Mutation returns the RelatedTestCaseMutation object of the builder.
func (rtcu *RelatedTestCaseUpdate) Mutation() *RelatedTestCaseMutation {
	return rtcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rtcu *RelatedTestCaseUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := rtcu.defaults(); err != nil {
		return 0, err
	}
	if len(rtcu.hooks) == 0 {
		affected, err = rtcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RelatedTestCaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rtcu.mutation = mutation
			affected, err = rtcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rtcu.hooks) - 1; i >= 0; i-- {
			if rtcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rtcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rtcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rtcu *RelatedTestCaseUpdate) SaveX(ctx context.Context) int {
	affected, err := rtcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rtcu *RelatedTestCaseUpdate) Exec(ctx context.Context) error {
	_, err := rtcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtcu *RelatedTestCaseUpdate) ExecX(ctx context.Context) {
	if err := rtcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rtcu *RelatedTestCaseUpdate) defaults() error {
	if _, ok := rtcu.mutation.UpdatedAt(); !ok {
		if relatedtestcase.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized relatedtestcase.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := relatedtestcase.UpdateDefaultUpdatedAt()
		rtcu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rtcu *RelatedTestCaseUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RelatedTestCaseUpdate {
	rtcu.modifiers = append(rtcu.modifiers, modifiers...)
	return rtcu
}

func (rtcu *RelatedTestCaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   relatedtestcase.Table,
			Columns: relatedtestcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: relatedtestcase.FieldID,
			},
		},
	}
	if ps := rtcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtcu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: relatedtestcase.FieldCreatedAt,
		})
	}
	if value, ok := rtcu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: relatedtestcase.FieldCreatedAt,
		})
	}
	if value, ok := rtcu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: relatedtestcase.FieldUpdatedAt,
		})
	}
	if value, ok := rtcu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: relatedtestcase.FieldUpdatedAt,
		})
	}
	if value, ok := rtcu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: relatedtestcase.FieldDeletedAt,
		})
	}
	if value, ok := rtcu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: relatedtestcase.FieldDeletedAt,
		})
	}
	if value, ok := rtcu.mutation.CondType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: relatedtestcase.FieldCondType,
		})
	}
	if rtcu.mutation.CondTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: relatedtestcase.FieldCondType,
		})
	}
	if value, ok := rtcu.mutation.TestCaseID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: relatedtestcase.FieldTestCaseID,
		})
	}
	if rtcu.mutation.TestCaseIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: relatedtestcase.FieldTestCaseID,
		})
	}
	if value, ok := rtcu.mutation.RelatedTestCaseID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: relatedtestcase.FieldRelatedTestCaseID,
		})
	}
	if rtcu.mutation.RelatedTestCaseIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: relatedtestcase.FieldRelatedTestCaseID,
		})
	}
	if value, ok := rtcu.mutation.ArgumentsTransfer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: relatedtestcase.FieldArgumentsTransfer,
		})
	}
	if rtcu.mutation.ArgumentsTransferCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: relatedtestcase.FieldArgumentsTransfer,
		})
	}
	if value, ok := rtcu.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: relatedtestcase.FieldIndex,
		})
	}
	if value, ok := rtcu.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: relatedtestcase.FieldIndex,
		})
	}
	if rtcu.mutation.IndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: relatedtestcase.FieldIndex,
		})
	}
	_spec.Modifiers = rtcu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, rtcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{relatedtestcase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RelatedTestCaseUpdateOne is the builder for updating a single RelatedTestCase entity.
type RelatedTestCaseUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *RelatedTestCaseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (rtcuo *RelatedTestCaseUpdateOne) SetCreatedAt(u uint32) *RelatedTestCaseUpdateOne {
	rtcuo.mutation.ResetCreatedAt()
	rtcuo.mutation.SetCreatedAt(u)
	return rtcuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rtcuo *RelatedTestCaseUpdateOne) SetNillableCreatedAt(u *uint32) *RelatedTestCaseUpdateOne {
	if u != nil {
		rtcuo.SetCreatedAt(*u)
	}
	return rtcuo
}

// AddCreatedAt adds u to the "created_at" field.
func (rtcuo *RelatedTestCaseUpdateOne) AddCreatedAt(u int32) *RelatedTestCaseUpdateOne {
	rtcuo.mutation.AddCreatedAt(u)
	return rtcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (rtcuo *RelatedTestCaseUpdateOne) SetUpdatedAt(u uint32) *RelatedTestCaseUpdateOne {
	rtcuo.mutation.ResetUpdatedAt()
	rtcuo.mutation.SetUpdatedAt(u)
	return rtcuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (rtcuo *RelatedTestCaseUpdateOne) AddUpdatedAt(u int32) *RelatedTestCaseUpdateOne {
	rtcuo.mutation.AddUpdatedAt(u)
	return rtcuo
}

// SetDeletedAt sets the "deleted_at" field.
func (rtcuo *RelatedTestCaseUpdateOne) SetDeletedAt(u uint32) *RelatedTestCaseUpdateOne {
	rtcuo.mutation.ResetDeletedAt()
	rtcuo.mutation.SetDeletedAt(u)
	return rtcuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rtcuo *RelatedTestCaseUpdateOne) SetNillableDeletedAt(u *uint32) *RelatedTestCaseUpdateOne {
	if u != nil {
		rtcuo.SetDeletedAt(*u)
	}
	return rtcuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (rtcuo *RelatedTestCaseUpdateOne) AddDeletedAt(u int32) *RelatedTestCaseUpdateOne {
	rtcuo.mutation.AddDeletedAt(u)
	return rtcuo
}

// SetCondType sets the "cond_type" field.
func (rtcuo *RelatedTestCaseUpdateOne) SetCondType(s string) *RelatedTestCaseUpdateOne {
	rtcuo.mutation.SetCondType(s)
	return rtcuo
}

// SetNillableCondType sets the "cond_type" field if the given value is not nil.
func (rtcuo *RelatedTestCaseUpdateOne) SetNillableCondType(s *string) *RelatedTestCaseUpdateOne {
	if s != nil {
		rtcuo.SetCondType(*s)
	}
	return rtcuo
}

// ClearCondType clears the value of the "cond_type" field.
func (rtcuo *RelatedTestCaseUpdateOne) ClearCondType() *RelatedTestCaseUpdateOne {
	rtcuo.mutation.ClearCondType()
	return rtcuo
}

// SetTestCaseID sets the "test_case_id" field.
func (rtcuo *RelatedTestCaseUpdateOne) SetTestCaseID(u uuid.UUID) *RelatedTestCaseUpdateOne {
	rtcuo.mutation.SetTestCaseID(u)
	return rtcuo
}

// SetNillableTestCaseID sets the "test_case_id" field if the given value is not nil.
func (rtcuo *RelatedTestCaseUpdateOne) SetNillableTestCaseID(u *uuid.UUID) *RelatedTestCaseUpdateOne {
	if u != nil {
		rtcuo.SetTestCaseID(*u)
	}
	return rtcuo
}

// ClearTestCaseID clears the value of the "test_case_id" field.
func (rtcuo *RelatedTestCaseUpdateOne) ClearTestCaseID() *RelatedTestCaseUpdateOne {
	rtcuo.mutation.ClearTestCaseID()
	return rtcuo
}

// SetRelatedTestCaseID sets the "related_test_case_id" field.
func (rtcuo *RelatedTestCaseUpdateOne) SetRelatedTestCaseID(u uuid.UUID) *RelatedTestCaseUpdateOne {
	rtcuo.mutation.SetRelatedTestCaseID(u)
	return rtcuo
}

// SetNillableRelatedTestCaseID sets the "related_test_case_id" field if the given value is not nil.
func (rtcuo *RelatedTestCaseUpdateOne) SetNillableRelatedTestCaseID(u *uuid.UUID) *RelatedTestCaseUpdateOne {
	if u != nil {
		rtcuo.SetRelatedTestCaseID(*u)
	}
	return rtcuo
}

// ClearRelatedTestCaseID clears the value of the "related_test_case_id" field.
func (rtcuo *RelatedTestCaseUpdateOne) ClearRelatedTestCaseID() *RelatedTestCaseUpdateOne {
	rtcuo.mutation.ClearRelatedTestCaseID()
	return rtcuo
}

// SetArgumentsTransfer sets the "arguments_transfer" field.
func (rtcuo *RelatedTestCaseUpdateOne) SetArgumentsTransfer(s string) *RelatedTestCaseUpdateOne {
	rtcuo.mutation.SetArgumentsTransfer(s)
	return rtcuo
}

// SetNillableArgumentsTransfer sets the "arguments_transfer" field if the given value is not nil.
func (rtcuo *RelatedTestCaseUpdateOne) SetNillableArgumentsTransfer(s *string) *RelatedTestCaseUpdateOne {
	if s != nil {
		rtcuo.SetArgumentsTransfer(*s)
	}
	return rtcuo
}

// ClearArgumentsTransfer clears the value of the "arguments_transfer" field.
func (rtcuo *RelatedTestCaseUpdateOne) ClearArgumentsTransfer() *RelatedTestCaseUpdateOne {
	rtcuo.mutation.ClearArgumentsTransfer()
	return rtcuo
}

// SetIndex sets the "index" field.
func (rtcuo *RelatedTestCaseUpdateOne) SetIndex(u uint32) *RelatedTestCaseUpdateOne {
	rtcuo.mutation.ResetIndex()
	rtcuo.mutation.SetIndex(u)
	return rtcuo
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (rtcuo *RelatedTestCaseUpdateOne) SetNillableIndex(u *uint32) *RelatedTestCaseUpdateOne {
	if u != nil {
		rtcuo.SetIndex(*u)
	}
	return rtcuo
}

// AddIndex adds u to the "index" field.
func (rtcuo *RelatedTestCaseUpdateOne) AddIndex(u int32) *RelatedTestCaseUpdateOne {
	rtcuo.mutation.AddIndex(u)
	return rtcuo
}

// ClearIndex clears the value of the "index" field.
func (rtcuo *RelatedTestCaseUpdateOne) ClearIndex() *RelatedTestCaseUpdateOne {
	rtcuo.mutation.ClearIndex()
	return rtcuo
}

// Mutation returns the RelatedTestCaseMutation object of the builder.
func (rtcuo *RelatedTestCaseUpdateOne) Mutation() *RelatedTestCaseMutation {
	return rtcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rtcuo *RelatedTestCaseUpdateOne) Select(field string, fields ...string) *RelatedTestCaseUpdateOne {
	rtcuo.fields = append([]string{field}, fields...)
	return rtcuo
}

// Save executes the query and returns the updated RelatedTestCase entity.
func (rtcuo *RelatedTestCaseUpdateOne) Save(ctx context.Context) (*RelatedTestCase, error) {
	var (
		err  error
		node *RelatedTestCase
	)
	if err := rtcuo.defaults(); err != nil {
		return nil, err
	}
	if len(rtcuo.hooks) == 0 {
		node, err = rtcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RelatedTestCaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rtcuo.mutation = mutation
			node, err = rtcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rtcuo.hooks) - 1; i >= 0; i-- {
			if rtcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rtcuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rtcuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*RelatedTestCase)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RelatedTestCaseMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rtcuo *RelatedTestCaseUpdateOne) SaveX(ctx context.Context) *RelatedTestCase {
	node, err := rtcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rtcuo *RelatedTestCaseUpdateOne) Exec(ctx context.Context) error {
	_, err := rtcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtcuo *RelatedTestCaseUpdateOne) ExecX(ctx context.Context) {
	if err := rtcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rtcuo *RelatedTestCaseUpdateOne) defaults() error {
	if _, ok := rtcuo.mutation.UpdatedAt(); !ok {
		if relatedtestcase.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized relatedtestcase.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := relatedtestcase.UpdateDefaultUpdatedAt()
		rtcuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rtcuo *RelatedTestCaseUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RelatedTestCaseUpdateOne {
	rtcuo.modifiers = append(rtcuo.modifiers, modifiers...)
	return rtcuo
}

func (rtcuo *RelatedTestCaseUpdateOne) sqlSave(ctx context.Context) (_node *RelatedTestCase, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   relatedtestcase.Table,
			Columns: relatedtestcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: relatedtestcase.FieldID,
			},
		},
	}
	id, ok := rtcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RelatedTestCase.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rtcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, relatedtestcase.FieldID)
		for _, f := range fields {
			if !relatedtestcase.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != relatedtestcase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rtcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtcuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: relatedtestcase.FieldCreatedAt,
		})
	}
	if value, ok := rtcuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: relatedtestcase.FieldCreatedAt,
		})
	}
	if value, ok := rtcuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: relatedtestcase.FieldUpdatedAt,
		})
	}
	if value, ok := rtcuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: relatedtestcase.FieldUpdatedAt,
		})
	}
	if value, ok := rtcuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: relatedtestcase.FieldDeletedAt,
		})
	}
	if value, ok := rtcuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: relatedtestcase.FieldDeletedAt,
		})
	}
	if value, ok := rtcuo.mutation.CondType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: relatedtestcase.FieldCondType,
		})
	}
	if rtcuo.mutation.CondTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: relatedtestcase.FieldCondType,
		})
	}
	if value, ok := rtcuo.mutation.TestCaseID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: relatedtestcase.FieldTestCaseID,
		})
	}
	if rtcuo.mutation.TestCaseIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: relatedtestcase.FieldTestCaseID,
		})
	}
	if value, ok := rtcuo.mutation.RelatedTestCaseID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: relatedtestcase.FieldRelatedTestCaseID,
		})
	}
	if rtcuo.mutation.RelatedTestCaseIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: relatedtestcase.FieldRelatedTestCaseID,
		})
	}
	if value, ok := rtcuo.mutation.ArgumentsTransfer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: relatedtestcase.FieldArgumentsTransfer,
		})
	}
	if rtcuo.mutation.ArgumentsTransferCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: relatedtestcase.FieldArgumentsTransfer,
		})
	}
	if value, ok := rtcuo.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: relatedtestcase.FieldIndex,
		})
	}
	if value, ok := rtcuo.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: relatedtestcase.FieldIndex,
		})
	}
	if rtcuo.mutation.IndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: relatedtestcase.FieldIndex,
		})
	}
	_spec.Modifiers = rtcuo.modifiers
	_node = &RelatedTestCase{config: rtcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rtcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{relatedtestcase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
