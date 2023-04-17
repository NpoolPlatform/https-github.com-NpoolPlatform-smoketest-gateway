// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/cond"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/module"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/planrelatedtestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testplan"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 5)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   cond.Table,
			Columns: cond.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: cond.FieldID,
			},
		},
		Type: "Cond",
		Fields: map[string]*sqlgraph.FieldSpec{
			cond.FieldCreatedAt:      {Type: field.TypeUint32, Column: cond.FieldCreatedAt},
			cond.FieldUpdatedAt:      {Type: field.TypeUint32, Column: cond.FieldUpdatedAt},
			cond.FieldDeletedAt:      {Type: field.TypeUint32, Column: cond.FieldDeletedAt},
			cond.FieldCondType:       {Type: field.TypeString, Column: cond.FieldCondType},
			cond.FieldTestCaseID:     {Type: field.TypeUUID, Column: cond.FieldTestCaseID},
			cond.FieldCondTestCaseID: {Type: field.TypeUUID, Column: cond.FieldCondTestCaseID},
			cond.FieldArgumentMap:    {Type: field.TypeString, Column: cond.FieldArgumentMap},
			cond.FieldIndex:          {Type: field.TypeUint32, Column: cond.FieldIndex},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   module.Table,
			Columns: module.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: module.FieldID,
			},
		},
		Type: "Module",
		Fields: map[string]*sqlgraph.FieldSpec{
			module.FieldCreatedAt:   {Type: field.TypeUint32, Column: module.FieldCreatedAt},
			module.FieldUpdatedAt:   {Type: field.TypeUint32, Column: module.FieldUpdatedAt},
			module.FieldDeletedAt:   {Type: field.TypeUint32, Column: module.FieldDeletedAt},
			module.FieldName:        {Type: field.TypeString, Column: module.FieldName},
			module.FieldDescription: {Type: field.TypeString, Column: module.FieldDescription},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   planrelatedtestcase.Table,
			Columns: planrelatedtestcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: planrelatedtestcase.FieldID,
			},
		},
		Type: "PlanRelatedTestCase",
		Fields: map[string]*sqlgraph.FieldSpec{
			planrelatedtestcase.FieldCreatedAt:      {Type: field.TypeUint32, Column: planrelatedtestcase.FieldCreatedAt},
			planrelatedtestcase.FieldUpdatedAt:      {Type: field.TypeUint32, Column: planrelatedtestcase.FieldUpdatedAt},
			planrelatedtestcase.FieldDeletedAt:      {Type: field.TypeUint32, Column: planrelatedtestcase.FieldDeletedAt},
			planrelatedtestcase.FieldTestPlanID:     {Type: field.TypeUUID, Column: planrelatedtestcase.FieldTestPlanID},
			planrelatedtestcase.FieldTestCaseID:     {Type: field.TypeUUID, Column: planrelatedtestcase.FieldTestCaseID},
			planrelatedtestcase.FieldTestCaseOutput: {Type: field.TypeString, Column: planrelatedtestcase.FieldTestCaseOutput},
			planrelatedtestcase.FieldDescription:    {Type: field.TypeString, Column: planrelatedtestcase.FieldDescription},
			planrelatedtestcase.FieldTestUserID:     {Type: field.TypeUUID, Column: planrelatedtestcase.FieldTestUserID},
			planrelatedtestcase.FieldRunDuration:    {Type: field.TypeUint32, Column: planrelatedtestcase.FieldRunDuration},
			planrelatedtestcase.FieldResult:         {Type: field.TypeString, Column: planrelatedtestcase.FieldResult},
			planrelatedtestcase.FieldIndex:          {Type: field.TypeUint32, Column: planrelatedtestcase.FieldIndex},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   testcase.Table,
			Columns: testcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: testcase.FieldID,
			},
		},
		Type: "TestCase",
		Fields: map[string]*sqlgraph.FieldSpec{
			testcase.FieldCreatedAt:          {Type: field.TypeUint32, Column: testcase.FieldCreatedAt},
			testcase.FieldUpdatedAt:          {Type: field.TypeUint32, Column: testcase.FieldUpdatedAt},
			testcase.FieldDeletedAt:          {Type: field.TypeUint32, Column: testcase.FieldDeletedAt},
			testcase.FieldName:               {Type: field.TypeString, Column: testcase.FieldName},
			testcase.FieldDescription:        {Type: field.TypeString, Column: testcase.FieldDescription},
			testcase.FieldModuleID:           {Type: field.TypeUUID, Column: testcase.FieldModuleID},
			testcase.FieldAPIID:              {Type: field.TypeUUID, Column: testcase.FieldAPIID},
			testcase.FieldArguments:          {Type: field.TypeString, Column: testcase.FieldArguments},
			testcase.FieldArgTypeDescription: {Type: field.TypeString, Column: testcase.FieldArgTypeDescription},
			testcase.FieldExpectationResult:  {Type: field.TypeString, Column: testcase.FieldExpectationResult},
			testcase.FieldTestCaseType:       {Type: field.TypeString, Column: testcase.FieldTestCaseType},
			testcase.FieldDeprecated:         {Type: field.TypeBool, Column: testcase.FieldDeprecated},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   testplan.Table,
			Columns: testplan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: testplan.FieldID,
			},
		},
		Type: "TestPlan",
		Fields: map[string]*sqlgraph.FieldSpec{
			testplan.FieldCreatedAt:             {Type: field.TypeUint32, Column: testplan.FieldCreatedAt},
			testplan.FieldUpdatedAt:             {Type: field.TypeUint32, Column: testplan.FieldUpdatedAt},
			testplan.FieldDeletedAt:             {Type: field.TypeUint32, Column: testplan.FieldDeletedAt},
			testplan.FieldName:                  {Type: field.TypeString, Column: testplan.FieldName},
			testplan.FieldState:                 {Type: field.TypeString, Column: testplan.FieldState},
			testplan.FieldOwnerID:               {Type: field.TypeUUID, Column: testplan.FieldOwnerID},
			testplan.FieldResponsibleUserID:     {Type: field.TypeUUID, Column: testplan.FieldResponsibleUserID},
			testplan.FieldFailedTestCasesCount:  {Type: field.TypeUint32, Column: testplan.FieldFailedTestCasesCount},
			testplan.FieldPassedTestCasesCount:  {Type: field.TypeUint32, Column: testplan.FieldPassedTestCasesCount},
			testplan.FieldSkippedTestCasesCount: {Type: field.TypeUint32, Column: testplan.FieldSkippedTestCasesCount},
			testplan.FieldRunDuration:           {Type: field.TypeUint32, Column: testplan.FieldRunDuration},
			testplan.FieldDeadline:              {Type: field.TypeUint32, Column: testplan.FieldDeadline},
			testplan.FieldTestResult:            {Type: field.TypeString, Column: testplan.FieldTestResult},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (cq *CondQuery) addPredicate(pred func(s *sql.Selector)) {
	cq.predicates = append(cq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CondQuery builder.
func (cq *CondQuery) Filter() *CondFilter {
	return &CondFilter{config: cq.config, predicateAdder: cq}
}

// addPredicate implements the predicateAdder interface.
func (m *CondMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CondMutation builder.
func (m *CondMutation) Filter() *CondFilter {
	return &CondFilter{config: m.config, predicateAdder: m}
}

// CondFilter provides a generic filtering capability at runtime for CondQuery.
type CondFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CondFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *CondFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(cond.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *CondFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(cond.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *CondFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(cond.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *CondFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(cond.FieldDeletedAt))
}

// WhereCondType applies the entql string predicate on the cond_type field.
func (f *CondFilter) WhereCondType(p entql.StringP) {
	f.Where(p.Field(cond.FieldCondType))
}

// WhereTestCaseID applies the entql [16]byte predicate on the test_case_id field.
func (f *CondFilter) WhereTestCaseID(p entql.ValueP) {
	f.Where(p.Field(cond.FieldTestCaseID))
}

// WhereCondTestCaseID applies the entql [16]byte predicate on the cond_test_case_id field.
func (f *CondFilter) WhereCondTestCaseID(p entql.ValueP) {
	f.Where(p.Field(cond.FieldCondTestCaseID))
}

// WhereArgumentMap applies the entql string predicate on the argument_map field.
func (f *CondFilter) WhereArgumentMap(p entql.StringP) {
	f.Where(p.Field(cond.FieldArgumentMap))
}

// WhereIndex applies the entql uint32 predicate on the index field.
func (f *CondFilter) WhereIndex(p entql.Uint32P) {
	f.Where(p.Field(cond.FieldIndex))
}

// addPredicate implements the predicateAdder interface.
func (mq *ModuleQuery) addPredicate(pred func(s *sql.Selector)) {
	mq.predicates = append(mq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ModuleQuery builder.
func (mq *ModuleQuery) Filter() *ModuleFilter {
	return &ModuleFilter{config: mq.config, predicateAdder: mq}
}

// addPredicate implements the predicateAdder interface.
func (m *ModuleMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ModuleMutation builder.
func (m *ModuleMutation) Filter() *ModuleFilter {
	return &ModuleFilter{config: m.config, predicateAdder: m}
}

// ModuleFilter provides a generic filtering capability at runtime for ModuleQuery.
type ModuleFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ModuleFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *ModuleFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(module.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *ModuleFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(module.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *ModuleFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(module.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *ModuleFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(module.FieldDeletedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *ModuleFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(module.FieldName))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *ModuleFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(module.FieldDescription))
}

// addPredicate implements the predicateAdder interface.
func (prtcq *PlanRelatedTestCaseQuery) addPredicate(pred func(s *sql.Selector)) {
	prtcq.predicates = append(prtcq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PlanRelatedTestCaseQuery builder.
func (prtcq *PlanRelatedTestCaseQuery) Filter() *PlanRelatedTestCaseFilter {
	return &PlanRelatedTestCaseFilter{config: prtcq.config, predicateAdder: prtcq}
}

// addPredicate implements the predicateAdder interface.
func (m *PlanRelatedTestCaseMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PlanRelatedTestCaseMutation builder.
func (m *PlanRelatedTestCaseMutation) Filter() *PlanRelatedTestCaseFilter {
	return &PlanRelatedTestCaseFilter{config: m.config, predicateAdder: m}
}

// PlanRelatedTestCaseFilter provides a generic filtering capability at runtime for PlanRelatedTestCaseQuery.
type PlanRelatedTestCaseFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *PlanRelatedTestCaseFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *PlanRelatedTestCaseFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(planrelatedtestcase.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *PlanRelatedTestCaseFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(planrelatedtestcase.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *PlanRelatedTestCaseFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(planrelatedtestcase.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *PlanRelatedTestCaseFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(planrelatedtestcase.FieldDeletedAt))
}

// WhereTestPlanID applies the entql [16]byte predicate on the test_plan_id field.
func (f *PlanRelatedTestCaseFilter) WhereTestPlanID(p entql.ValueP) {
	f.Where(p.Field(planrelatedtestcase.FieldTestPlanID))
}

// WhereTestCaseID applies the entql [16]byte predicate on the test_case_id field.
func (f *PlanRelatedTestCaseFilter) WhereTestCaseID(p entql.ValueP) {
	f.Where(p.Field(planrelatedtestcase.FieldTestCaseID))
}

// WhereTestCaseOutput applies the entql string predicate on the test_case_output field.
func (f *PlanRelatedTestCaseFilter) WhereTestCaseOutput(p entql.StringP) {
	f.Where(p.Field(planrelatedtestcase.FieldTestCaseOutput))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *PlanRelatedTestCaseFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(planrelatedtestcase.FieldDescription))
}

// WhereTestUserID applies the entql [16]byte predicate on the test_user_id field.
func (f *PlanRelatedTestCaseFilter) WhereTestUserID(p entql.ValueP) {
	f.Where(p.Field(planrelatedtestcase.FieldTestUserID))
}

// WhereRunDuration applies the entql uint32 predicate on the run_duration field.
func (f *PlanRelatedTestCaseFilter) WhereRunDuration(p entql.Uint32P) {
	f.Where(p.Field(planrelatedtestcase.FieldRunDuration))
}

// WhereResult applies the entql string predicate on the result field.
func (f *PlanRelatedTestCaseFilter) WhereResult(p entql.StringP) {
	f.Where(p.Field(planrelatedtestcase.FieldResult))
}

// WhereIndex applies the entql uint32 predicate on the index field.
func (f *PlanRelatedTestCaseFilter) WhereIndex(p entql.Uint32P) {
	f.Where(p.Field(planrelatedtestcase.FieldIndex))
}

// addPredicate implements the predicateAdder interface.
func (tcq *TestCaseQuery) addPredicate(pred func(s *sql.Selector)) {
	tcq.predicates = append(tcq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TestCaseQuery builder.
func (tcq *TestCaseQuery) Filter() *TestCaseFilter {
	return &TestCaseFilter{config: tcq.config, predicateAdder: tcq}
}

// addPredicate implements the predicateAdder interface.
func (m *TestCaseMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TestCaseMutation builder.
func (m *TestCaseMutation) Filter() *TestCaseFilter {
	return &TestCaseFilter{config: m.config, predicateAdder: m}
}

// TestCaseFilter provides a generic filtering capability at runtime for TestCaseQuery.
type TestCaseFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *TestCaseFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *TestCaseFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(testcase.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *TestCaseFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(testcase.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *TestCaseFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(testcase.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *TestCaseFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(testcase.FieldDeletedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *TestCaseFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(testcase.FieldName))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *TestCaseFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(testcase.FieldDescription))
}

// WhereModuleID applies the entql [16]byte predicate on the module_id field.
func (f *TestCaseFilter) WhereModuleID(p entql.ValueP) {
	f.Where(p.Field(testcase.FieldModuleID))
}

// WhereAPIID applies the entql [16]byte predicate on the api_id field.
func (f *TestCaseFilter) WhereAPIID(p entql.ValueP) {
	f.Where(p.Field(testcase.FieldAPIID))
}

// WhereArguments applies the entql string predicate on the arguments field.
func (f *TestCaseFilter) WhereArguments(p entql.StringP) {
	f.Where(p.Field(testcase.FieldArguments))
}

// WhereArgTypeDescription applies the entql string predicate on the arg_type_description field.
func (f *TestCaseFilter) WhereArgTypeDescription(p entql.StringP) {
	f.Where(p.Field(testcase.FieldArgTypeDescription))
}

// WhereExpectationResult applies the entql string predicate on the expectation_result field.
func (f *TestCaseFilter) WhereExpectationResult(p entql.StringP) {
	f.Where(p.Field(testcase.FieldExpectationResult))
}

// WhereTestCaseType applies the entql string predicate on the test_case_type field.
func (f *TestCaseFilter) WhereTestCaseType(p entql.StringP) {
	f.Where(p.Field(testcase.FieldTestCaseType))
}

// WhereDeprecated applies the entql bool predicate on the deprecated field.
func (f *TestCaseFilter) WhereDeprecated(p entql.BoolP) {
	f.Where(p.Field(testcase.FieldDeprecated))
}

// addPredicate implements the predicateAdder interface.
func (tpq *TestPlanQuery) addPredicate(pred func(s *sql.Selector)) {
	tpq.predicates = append(tpq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TestPlanQuery builder.
func (tpq *TestPlanQuery) Filter() *TestPlanFilter {
	return &TestPlanFilter{config: tpq.config, predicateAdder: tpq}
}

// addPredicate implements the predicateAdder interface.
func (m *TestPlanMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TestPlanMutation builder.
func (m *TestPlanMutation) Filter() *TestPlanFilter {
	return &TestPlanFilter{config: m.config, predicateAdder: m}
}

// TestPlanFilter provides a generic filtering capability at runtime for TestPlanQuery.
type TestPlanFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *TestPlanFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *TestPlanFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(testplan.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *TestPlanFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *TestPlanFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *TestPlanFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldDeletedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *TestPlanFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(testplan.FieldName))
}

// WhereState applies the entql string predicate on the state field.
func (f *TestPlanFilter) WhereState(p entql.StringP) {
	f.Where(p.Field(testplan.FieldState))
}

// WhereOwnerID applies the entql [16]byte predicate on the owner_id field.
func (f *TestPlanFilter) WhereOwnerID(p entql.ValueP) {
	f.Where(p.Field(testplan.FieldOwnerID))
}

// WhereResponsibleUserID applies the entql [16]byte predicate on the responsible_user_id field.
func (f *TestPlanFilter) WhereResponsibleUserID(p entql.ValueP) {
	f.Where(p.Field(testplan.FieldResponsibleUserID))
}

// WhereFailedTestCasesCount applies the entql uint32 predicate on the failed_test_cases_count field.
func (f *TestPlanFilter) WhereFailedTestCasesCount(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldFailedTestCasesCount))
}

// WherePassedTestCasesCount applies the entql uint32 predicate on the passed_test_cases_count field.
func (f *TestPlanFilter) WherePassedTestCasesCount(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldPassedTestCasesCount))
}

// WhereSkippedTestCasesCount applies the entql uint32 predicate on the skipped_test_cases_count field.
func (f *TestPlanFilter) WhereSkippedTestCasesCount(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldSkippedTestCasesCount))
}

// WhereRunDuration applies the entql uint32 predicate on the run_duration field.
func (f *TestPlanFilter) WhereRunDuration(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldRunDuration))
}

// WhereDeadline applies the entql uint32 predicate on the deadline field.
func (f *TestPlanFilter) WhereDeadline(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldDeadline))
}

// WhereTestResult applies the entql string predicate on the test_result field.
func (f *TestPlanFilter) WhereTestResult(p entql.StringP) {
	f.Where(p.Field(testplan.FieldTestResult))
}
