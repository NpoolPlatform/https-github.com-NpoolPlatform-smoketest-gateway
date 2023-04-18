// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/plantestcase"
	"github.com/google/uuid"
)

// PlanTestCase is the model entity for the PlanTestCase schema.
type PlanTestCase struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// TestPlanID holds the value of the "test_plan_id" field.
	TestPlanID uuid.UUID `json:"test_plan_id,omitempty"`
	// TestCaseID holds the value of the "test_case_id" field.
	TestCaseID uuid.UUID `json:"test_case_id,omitempty"`
	// TestCaseOutput holds the value of the "test_case_output" field.
	TestCaseOutput string `json:"test_case_output,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// TestUserID holds the value of the "test_user_id" field.
	TestUserID uuid.UUID `json:"test_user_id,omitempty"`
	// RunDuration holds the value of the "run_duration" field.
	RunDuration uint32 `json:"run_duration,omitempty"`
	// Result holds the value of the "result" field.
	Result string `json:"result,omitempty"`
	// Index holds the value of the "index" field.
	Index uint32 `json:"index,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PlanTestCase) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case plantestcase.FieldCreatedAt, plantestcase.FieldUpdatedAt, plantestcase.FieldDeletedAt, plantestcase.FieldRunDuration, plantestcase.FieldIndex:
			values[i] = new(sql.NullInt64)
		case plantestcase.FieldTestCaseOutput, plantestcase.FieldDescription, plantestcase.FieldResult:
			values[i] = new(sql.NullString)
		case plantestcase.FieldID, plantestcase.FieldTestPlanID, plantestcase.FieldTestCaseID, plantestcase.FieldTestUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PlanTestCase", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PlanTestCase fields.
func (ptc *PlanTestCase) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case plantestcase.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ptc.ID = *value
			}
		case plantestcase.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ptc.CreatedAt = uint32(value.Int64)
			}
		case plantestcase.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ptc.UpdatedAt = uint32(value.Int64)
			}
		case plantestcase.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ptc.DeletedAt = uint32(value.Int64)
			}
		case plantestcase.FieldTestPlanID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field test_plan_id", values[i])
			} else if value != nil {
				ptc.TestPlanID = *value
			}
		case plantestcase.FieldTestCaseID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field test_case_id", values[i])
			} else if value != nil {
				ptc.TestCaseID = *value
			}
		case plantestcase.FieldTestCaseOutput:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field test_case_output", values[i])
			} else if value.Valid {
				ptc.TestCaseOutput = value.String
			}
		case plantestcase.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ptc.Description = value.String
			}
		case plantestcase.FieldTestUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field test_user_id", values[i])
			} else if value != nil {
				ptc.TestUserID = *value
			}
		case plantestcase.FieldRunDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field run_duration", values[i])
			} else if value.Valid {
				ptc.RunDuration = uint32(value.Int64)
			}
		case plantestcase.FieldResult:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				ptc.Result = value.String
			}
		case plantestcase.FieldIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field index", values[i])
			} else if value.Valid {
				ptc.Index = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this PlanTestCase.
// Note that you need to call PlanTestCase.Unwrap() before calling this method if this PlanTestCase
// was returned from a transaction, and the transaction was committed or rolled back.
func (ptc *PlanTestCase) Update() *PlanTestCaseUpdateOne {
	return (&PlanTestCaseClient{config: ptc.config}).UpdateOne(ptc)
}

// Unwrap unwraps the PlanTestCase entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ptc *PlanTestCase) Unwrap() *PlanTestCase {
	_tx, ok := ptc.config.driver.(*txDriver)
	if !ok {
		panic("ent: PlanTestCase is not a transactional entity")
	}
	ptc.config.driver = _tx.drv
	return ptc
}

// String implements the fmt.Stringer.
func (ptc *PlanTestCase) String() string {
	var builder strings.Builder
	builder.WriteString("PlanTestCase(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ptc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ptc.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ptc.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ptc.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("test_plan_id=")
	builder.WriteString(fmt.Sprintf("%v", ptc.TestPlanID))
	builder.WriteString(", ")
	builder.WriteString("test_case_id=")
	builder.WriteString(fmt.Sprintf("%v", ptc.TestCaseID))
	builder.WriteString(", ")
	builder.WriteString("test_case_output=")
	builder.WriteString(ptc.TestCaseOutput)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(ptc.Description)
	builder.WriteString(", ")
	builder.WriteString("test_user_id=")
	builder.WriteString(fmt.Sprintf("%v", ptc.TestUserID))
	builder.WriteString(", ")
	builder.WriteString("run_duration=")
	builder.WriteString(fmt.Sprintf("%v", ptc.RunDuration))
	builder.WriteString(", ")
	builder.WriteString("result=")
	builder.WriteString(ptc.Result)
	builder.WriteString(", ")
	builder.WriteString("index=")
	builder.WriteString(fmt.Sprintf("%v", ptc.Index))
	builder.WriteByte(')')
	return builder.String()
}

// PlanTestCases is a parsable slice of PlanTestCase.
type PlanTestCases []*PlanTestCase

func (ptc PlanTestCases) config(cfg config) {
	for _i := range ptc {
		ptc[_i].config = cfg
	}
}