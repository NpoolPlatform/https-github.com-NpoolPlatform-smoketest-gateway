// Code generated by ent, DO NOT EDIT.

package testcase

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// ModuleID applies equality check predicate on the "module_id" field. It's identical to ModuleIDEQ.
func ModuleID(v uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModuleID), v))
	})
}

// APIID applies equality check predicate on the "api_id" field. It's identical to APIIDEQ.
func APIID(v uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAPIID), v))
	})
}

// Input applies equality check predicate on the "input" field. It's identical to InputEQ.
func Input(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInput), v))
	})
}

// InputDesc applies equality check predicate on the "input_desc" field. It's identical to InputDescEQ.
func InputDesc(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInputDesc), v))
	})
}

// Expectation applies equality check predicate on the "expectation" field. It's identical to ExpectationEQ.
func Expectation(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpectation), v))
	})
}

// TestCaseType applies equality check predicate on the "test_case_type" field. It's identical to TestCaseTypeEQ.
func TestCaseType(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTestCaseType), v))
	})
}

// Deprecated applies equality check predicate on the "deprecated" field. It's identical to DeprecatedEQ.
func Deprecated(v bool) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeprecated), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldName)))
	})
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldName)))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// ModuleIDEQ applies the EQ predicate on the "module_id" field.
func ModuleIDEQ(v uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModuleID), v))
	})
}

// ModuleIDNEQ applies the NEQ predicate on the "module_id" field.
func ModuleIDNEQ(v uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldModuleID), v))
	})
}

// ModuleIDIn applies the In predicate on the "module_id" field.
func ModuleIDIn(vs ...uuid.UUID) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldModuleID), v...))
	})
}

// ModuleIDNotIn applies the NotIn predicate on the "module_id" field.
func ModuleIDNotIn(vs ...uuid.UUID) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldModuleID), v...))
	})
}

// ModuleIDGT applies the GT predicate on the "module_id" field.
func ModuleIDGT(v uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldModuleID), v))
	})
}

// ModuleIDGTE applies the GTE predicate on the "module_id" field.
func ModuleIDGTE(v uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldModuleID), v))
	})
}

// ModuleIDLT applies the LT predicate on the "module_id" field.
func ModuleIDLT(v uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldModuleID), v))
	})
}

// ModuleIDLTE applies the LTE predicate on the "module_id" field.
func ModuleIDLTE(v uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldModuleID), v))
	})
}

// ModuleIDIsNil applies the IsNil predicate on the "module_id" field.
func ModuleIDIsNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldModuleID)))
	})
}

// ModuleIDNotNil applies the NotNil predicate on the "module_id" field.
func ModuleIDNotNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldModuleID)))
	})
}

// APIIDEQ applies the EQ predicate on the "api_id" field.
func APIIDEQ(v uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAPIID), v))
	})
}

// APIIDNEQ applies the NEQ predicate on the "api_id" field.
func APIIDNEQ(v uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAPIID), v))
	})
}

// APIIDIn applies the In predicate on the "api_id" field.
func APIIDIn(vs ...uuid.UUID) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAPIID), v...))
	})
}

// APIIDNotIn applies the NotIn predicate on the "api_id" field.
func APIIDNotIn(vs ...uuid.UUID) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAPIID), v...))
	})
}

// APIIDGT applies the GT predicate on the "api_id" field.
func APIIDGT(v uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAPIID), v))
	})
}

// APIIDGTE applies the GTE predicate on the "api_id" field.
func APIIDGTE(v uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAPIID), v))
	})
}

// APIIDLT applies the LT predicate on the "api_id" field.
func APIIDLT(v uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAPIID), v))
	})
}

// APIIDLTE applies the LTE predicate on the "api_id" field.
func APIIDLTE(v uuid.UUID) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAPIID), v))
	})
}

// APIIDIsNil applies the IsNil predicate on the "api_id" field.
func APIIDIsNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAPIID)))
	})
}

// APIIDNotNil applies the NotNil predicate on the "api_id" field.
func APIIDNotNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAPIID)))
	})
}

// InputEQ applies the EQ predicate on the "input" field.
func InputEQ(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInput), v))
	})
}

// InputNEQ applies the NEQ predicate on the "input" field.
func InputNEQ(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInput), v))
	})
}

// InputIn applies the In predicate on the "input" field.
func InputIn(vs ...string) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInput), v...))
	})
}

// InputNotIn applies the NotIn predicate on the "input" field.
func InputNotIn(vs ...string) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInput), v...))
	})
}

// InputGT applies the GT predicate on the "input" field.
func InputGT(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInput), v))
	})
}

// InputGTE applies the GTE predicate on the "input" field.
func InputGTE(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInput), v))
	})
}

// InputLT applies the LT predicate on the "input" field.
func InputLT(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInput), v))
	})
}

// InputLTE applies the LTE predicate on the "input" field.
func InputLTE(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInput), v))
	})
}

// InputContains applies the Contains predicate on the "input" field.
func InputContains(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInput), v))
	})
}

// InputHasPrefix applies the HasPrefix predicate on the "input" field.
func InputHasPrefix(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInput), v))
	})
}

// InputHasSuffix applies the HasSuffix predicate on the "input" field.
func InputHasSuffix(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInput), v))
	})
}

// InputIsNil applies the IsNil predicate on the "input" field.
func InputIsNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInput)))
	})
}

// InputNotNil applies the NotNil predicate on the "input" field.
func InputNotNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInput)))
	})
}

// InputEqualFold applies the EqualFold predicate on the "input" field.
func InputEqualFold(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInput), v))
	})
}

// InputContainsFold applies the ContainsFold predicate on the "input" field.
func InputContainsFold(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInput), v))
	})
}

// InputDescEQ applies the EQ predicate on the "input_desc" field.
func InputDescEQ(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInputDesc), v))
	})
}

// InputDescNEQ applies the NEQ predicate on the "input_desc" field.
func InputDescNEQ(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInputDesc), v))
	})
}

// InputDescIn applies the In predicate on the "input_desc" field.
func InputDescIn(vs ...string) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInputDesc), v...))
	})
}

// InputDescNotIn applies the NotIn predicate on the "input_desc" field.
func InputDescNotIn(vs ...string) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInputDesc), v...))
	})
}

// InputDescGT applies the GT predicate on the "input_desc" field.
func InputDescGT(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInputDesc), v))
	})
}

// InputDescGTE applies the GTE predicate on the "input_desc" field.
func InputDescGTE(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInputDesc), v))
	})
}

// InputDescLT applies the LT predicate on the "input_desc" field.
func InputDescLT(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInputDesc), v))
	})
}

// InputDescLTE applies the LTE predicate on the "input_desc" field.
func InputDescLTE(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInputDesc), v))
	})
}

// InputDescContains applies the Contains predicate on the "input_desc" field.
func InputDescContains(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInputDesc), v))
	})
}

// InputDescHasPrefix applies the HasPrefix predicate on the "input_desc" field.
func InputDescHasPrefix(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInputDesc), v))
	})
}

// InputDescHasSuffix applies the HasSuffix predicate on the "input_desc" field.
func InputDescHasSuffix(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInputDesc), v))
	})
}

// InputDescIsNil applies the IsNil predicate on the "input_desc" field.
func InputDescIsNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInputDesc)))
	})
}

// InputDescNotNil applies the NotNil predicate on the "input_desc" field.
func InputDescNotNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInputDesc)))
	})
}

// InputDescEqualFold applies the EqualFold predicate on the "input_desc" field.
func InputDescEqualFold(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInputDesc), v))
	})
}

// InputDescContainsFold applies the ContainsFold predicate on the "input_desc" field.
func InputDescContainsFold(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInputDesc), v))
	})
}

// ExpectationEQ applies the EQ predicate on the "expectation" field.
func ExpectationEQ(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpectation), v))
	})
}

// ExpectationNEQ applies the NEQ predicate on the "expectation" field.
func ExpectationNEQ(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpectation), v))
	})
}

// ExpectationIn applies the In predicate on the "expectation" field.
func ExpectationIn(vs ...string) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExpectation), v...))
	})
}

// ExpectationNotIn applies the NotIn predicate on the "expectation" field.
func ExpectationNotIn(vs ...string) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExpectation), v...))
	})
}

// ExpectationGT applies the GT predicate on the "expectation" field.
func ExpectationGT(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpectation), v))
	})
}

// ExpectationGTE applies the GTE predicate on the "expectation" field.
func ExpectationGTE(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpectation), v))
	})
}

// ExpectationLT applies the LT predicate on the "expectation" field.
func ExpectationLT(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpectation), v))
	})
}

// ExpectationLTE applies the LTE predicate on the "expectation" field.
func ExpectationLTE(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpectation), v))
	})
}

// ExpectationContains applies the Contains predicate on the "expectation" field.
func ExpectationContains(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExpectation), v))
	})
}

// ExpectationHasPrefix applies the HasPrefix predicate on the "expectation" field.
func ExpectationHasPrefix(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExpectation), v))
	})
}

// ExpectationHasSuffix applies the HasSuffix predicate on the "expectation" field.
func ExpectationHasSuffix(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExpectation), v))
	})
}

// ExpectationIsNil applies the IsNil predicate on the "expectation" field.
func ExpectationIsNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExpectation)))
	})
}

// ExpectationNotNil applies the NotNil predicate on the "expectation" field.
func ExpectationNotNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExpectation)))
	})
}

// ExpectationEqualFold applies the EqualFold predicate on the "expectation" field.
func ExpectationEqualFold(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExpectation), v))
	})
}

// ExpectationContainsFold applies the ContainsFold predicate on the "expectation" field.
func ExpectationContainsFold(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExpectation), v))
	})
}

// TestCaseTypeEQ applies the EQ predicate on the "test_case_type" field.
func TestCaseTypeEQ(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTestCaseType), v))
	})
}

// TestCaseTypeNEQ applies the NEQ predicate on the "test_case_type" field.
func TestCaseTypeNEQ(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTestCaseType), v))
	})
}

// TestCaseTypeIn applies the In predicate on the "test_case_type" field.
func TestCaseTypeIn(vs ...string) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTestCaseType), v...))
	})
}

// TestCaseTypeNotIn applies the NotIn predicate on the "test_case_type" field.
func TestCaseTypeNotIn(vs ...string) predicate.TestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTestCaseType), v...))
	})
}

// TestCaseTypeGT applies the GT predicate on the "test_case_type" field.
func TestCaseTypeGT(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTestCaseType), v))
	})
}

// TestCaseTypeGTE applies the GTE predicate on the "test_case_type" field.
func TestCaseTypeGTE(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTestCaseType), v))
	})
}

// TestCaseTypeLT applies the LT predicate on the "test_case_type" field.
func TestCaseTypeLT(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTestCaseType), v))
	})
}

// TestCaseTypeLTE applies the LTE predicate on the "test_case_type" field.
func TestCaseTypeLTE(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTestCaseType), v))
	})
}

// TestCaseTypeContains applies the Contains predicate on the "test_case_type" field.
func TestCaseTypeContains(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTestCaseType), v))
	})
}

// TestCaseTypeHasPrefix applies the HasPrefix predicate on the "test_case_type" field.
func TestCaseTypeHasPrefix(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTestCaseType), v))
	})
}

// TestCaseTypeHasSuffix applies the HasSuffix predicate on the "test_case_type" field.
func TestCaseTypeHasSuffix(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTestCaseType), v))
	})
}

// TestCaseTypeIsNil applies the IsNil predicate on the "test_case_type" field.
func TestCaseTypeIsNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTestCaseType)))
	})
}

// TestCaseTypeNotNil applies the NotNil predicate on the "test_case_type" field.
func TestCaseTypeNotNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTestCaseType)))
	})
}

// TestCaseTypeEqualFold applies the EqualFold predicate on the "test_case_type" field.
func TestCaseTypeEqualFold(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTestCaseType), v))
	})
}

// TestCaseTypeContainsFold applies the ContainsFold predicate on the "test_case_type" field.
func TestCaseTypeContainsFold(v string) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTestCaseType), v))
	})
}

// DeprecatedEQ applies the EQ predicate on the "deprecated" field.
func DeprecatedEQ(v bool) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeprecated), v))
	})
}

// DeprecatedNEQ applies the NEQ predicate on the "deprecated" field.
func DeprecatedNEQ(v bool) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeprecated), v))
	})
}

// DeprecatedIsNil applies the IsNil predicate on the "deprecated" field.
func DeprecatedIsNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeprecated)))
	})
}

// DeprecatedNotNil applies the NotNil predicate on the "deprecated" field.
func DeprecatedNotNil() predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeprecated)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TestCase) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TestCase) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TestCase) predicate.TestCase {
	return predicate.TestCase(func(s *sql.Selector) {
		p(s.Not())
	})
}
