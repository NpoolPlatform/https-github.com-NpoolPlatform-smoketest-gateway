// Code generated by ent, DO NOT EDIT.

package testcase

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the testcase type in the database.
	Label = "test_case"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldModuleID holds the string denoting the module_id field in the database.
	FieldModuleID = "module_id"
	// FieldAPIID holds the string denoting the api_id field in the database.
	FieldAPIID = "api_id"
	// FieldInput holds the string denoting the input field in the database.
	FieldInput = "input"
	// FieldInputDesc holds the string denoting the input_desc field in the database.
	FieldInputDesc = "input_desc"
	// FieldExpectation holds the string denoting the expectation field in the database.
	FieldExpectation = "expectation"
	// FieldTestCaseType holds the string denoting the test_case_type field in the database.
	FieldTestCaseType = "test_case_type"
	// FieldDeprecated holds the string denoting the deprecated field in the database.
	FieldDeprecated = "deprecated"
	// Table holds the table name of the testcase in the database.
	Table = "test_cases"
)

// Columns holds all SQL columns for testcase fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldDescription,
	FieldModuleID,
	FieldAPIID,
	FieldInput,
	FieldInputDesc,
	FieldExpectation,
	FieldTestCaseType,
	FieldDeprecated,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultModuleID holds the default value on creation for the "module_id" field.
	DefaultModuleID func() uuid.UUID
	// DefaultAPIID holds the default value on creation for the "api_id" field.
	DefaultAPIID func() uuid.UUID
	// DefaultInput holds the default value on creation for the "input" field.
	DefaultInput string
	// DefaultInputDesc holds the default value on creation for the "input_desc" field.
	DefaultInputDesc string
	// DefaultExpectation holds the default value on creation for the "expectation" field.
	DefaultExpectation string
	// DefaultTestCaseType holds the default value on creation for the "test_case_type" field.
	DefaultTestCaseType string
	// DefaultDeprecated holds the default value on creation for the "deprecated" field.
	DefaultDeprecated bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
