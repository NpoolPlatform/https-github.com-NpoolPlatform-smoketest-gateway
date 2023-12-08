// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/cond"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/module"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/plantestcase"
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
				Type:   field.TypeUint32,
				Column: cond.FieldID,
			},
		},
		Type: "Cond",
		Fields: map[string]*sqlgraph.FieldSpec{
			cond.FieldCreatedAt:      {Type: field.TypeUint32, Column: cond.FieldCreatedAt},
			cond.FieldUpdatedAt:      {Type: field.TypeUint32, Column: cond.FieldUpdatedAt},
			cond.FieldDeletedAt:      {Type: field.TypeUint32, Column: cond.FieldDeletedAt},
			cond.FieldEntID:          {Type: field.TypeUUID, Column: cond.FieldEntID},
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
				Type:   field.TypeUint32,
				Column: module.FieldID,
			},
		},
		Type: "Module",
		Fields: map[string]*sqlgraph.FieldSpec{
			module.FieldCreatedAt:   {Type: field.TypeUint32, Column: module.FieldCreatedAt},
			module.FieldUpdatedAt:   {Type: field.TypeUint32, Column: module.FieldUpdatedAt},
			module.FieldDeletedAt:   {Type: field.TypeUint32, Column: module.FieldDeletedAt},
			module.FieldEntID:       {Type: field.TypeUUID, Column: module.FieldEntID},
			module.FieldName:        {Type: field.TypeString, Column: module.FieldName},
			module.FieldDescription: {Type: field.TypeString, Column: module.FieldDescription},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   plantestcase.Table,
			Columns: plantestcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: plantestcase.FieldID,
			},
		},
		Type: "PlanTestCase",
		Fields: map[string]*sqlgraph.FieldSpec{
			plantestcase.FieldCreatedAt:   {Type: field.TypeUint32, Column: plantestcase.FieldCreatedAt},
			plantestcase.FieldUpdatedAt:   {Type: field.TypeUint32, Column: plantestcase.FieldUpdatedAt},
			plantestcase.FieldDeletedAt:   {Type: field.TypeUint32, Column: plantestcase.FieldDeletedAt},
			plantestcase.FieldEntID:       {Type: field.TypeUUID, Column: plantestcase.FieldEntID},
			plantestcase.FieldTestPlanID:  {Type: field.TypeUUID, Column: plantestcase.FieldTestPlanID},
			plantestcase.FieldTestCaseID:  {Type: field.TypeUUID, Column: plantestcase.FieldTestCaseID},
			plantestcase.FieldInput:       {Type: field.TypeString, Column: plantestcase.FieldInput},
			plantestcase.FieldOutput:      {Type: field.TypeString, Column: plantestcase.FieldOutput},
			plantestcase.FieldDescription: {Type: field.TypeString, Column: plantestcase.FieldDescription},
			plantestcase.FieldTestUserID:  {Type: field.TypeUUID, Column: plantestcase.FieldTestUserID},
			plantestcase.FieldRunDuration: {Type: field.TypeUint32, Column: plantestcase.FieldRunDuration},
			plantestcase.FieldResult:      {Type: field.TypeString, Column: plantestcase.FieldResult},
			plantestcase.FieldIndex:       {Type: field.TypeUint32, Column: plantestcase.FieldIndex},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   testcase.Table,
			Columns: testcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: testcase.FieldID,
			},
		},
		Type: "TestCase",
		Fields: map[string]*sqlgraph.FieldSpec{
			testcase.FieldCreatedAt:     {Type: field.TypeUint32, Column: testcase.FieldCreatedAt},
			testcase.FieldUpdatedAt:     {Type: field.TypeUint32, Column: testcase.FieldUpdatedAt},
			testcase.FieldDeletedAt:     {Type: field.TypeUint32, Column: testcase.FieldDeletedAt},
			testcase.FieldEntID:         {Type: field.TypeUUID, Column: testcase.FieldEntID},
			testcase.FieldName:          {Type: field.TypeString, Column: testcase.FieldName},
			testcase.FieldDescription:   {Type: field.TypeString, Column: testcase.FieldDescription},
			testcase.FieldModuleID:      {Type: field.TypeUUID, Column: testcase.FieldModuleID},
			testcase.FieldAPIID:         {Type: field.TypeUUID, Column: testcase.FieldAPIID},
			testcase.FieldInput:         {Type: field.TypeString, Column: testcase.FieldInput},
			testcase.FieldInputDesc:     {Type: field.TypeString, Column: testcase.FieldInputDesc},
			testcase.FieldExpectation:   {Type: field.TypeString, Column: testcase.FieldExpectation},
			testcase.FieldOutputDesc:    {Type: field.TypeString, Column: testcase.FieldOutputDesc},
			testcase.FieldTestCaseType:  {Type: field.TypeString, Column: testcase.FieldTestCaseType},
			testcase.FieldTestCaseClass: {Type: field.TypeString, Column: testcase.FieldTestCaseClass},
			testcase.FieldDeprecated:    {Type: field.TypeBool, Column: testcase.FieldDeprecated},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   testplan.Table,
			Columns: testplan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: testplan.FieldID,
			},
		},
		Type: "TestPlan",
		Fields: map[string]*sqlgraph.FieldSpec{
			testplan.FieldCreatedAt:   {Type: field.TypeUint32, Column: testplan.FieldCreatedAt},
			testplan.FieldUpdatedAt:   {Type: field.TypeUint32, Column: testplan.FieldUpdatedAt},
			testplan.FieldDeletedAt:   {Type: field.TypeUint32, Column: testplan.FieldDeletedAt},
			testplan.FieldEntID:       {Type: field.TypeUUID, Column: testplan.FieldEntID},
			testplan.FieldName:        {Type: field.TypeString, Column: testplan.FieldName},
			testplan.FieldState:       {Type: field.TypeString, Column: testplan.FieldState},
			testplan.FieldCreatedBy:   {Type: field.TypeUUID, Column: testplan.FieldCreatedBy},
			testplan.FieldExecutor:    {Type: field.TypeUUID, Column: testplan.FieldExecutor},
			testplan.FieldFails:       {Type: field.TypeUint32, Column: testplan.FieldFails},
			testplan.FieldPasses:      {Type: field.TypeUint32, Column: testplan.FieldPasses},
			testplan.FieldSkips:       {Type: field.TypeUint32, Column: testplan.FieldSkips},
			testplan.FieldRunDuration: {Type: field.TypeUint32, Column: testplan.FieldRunDuration},
			testplan.FieldDeadline:    {Type: field.TypeUint32, Column: testplan.FieldDeadline},
			testplan.FieldResult:      {Type: field.TypeString, Column: testplan.FieldResult},
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

// WhereID applies the entql uint32 predicate on the id field.
func (f *CondFilter) WhereID(p entql.Uint32P) {
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

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *CondFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(cond.FieldEntID))
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

// WhereID applies the entql uint32 predicate on the id field.
func (f *ModuleFilter) WhereID(p entql.Uint32P) {
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

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *ModuleFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(module.FieldEntID))
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
func (ptcq *PlanTestCaseQuery) addPredicate(pred func(s *sql.Selector)) {
	ptcq.predicates = append(ptcq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PlanTestCaseQuery builder.
func (ptcq *PlanTestCaseQuery) Filter() *PlanTestCaseFilter {
	return &PlanTestCaseFilter{config: ptcq.config, predicateAdder: ptcq}
}

// addPredicate implements the predicateAdder interface.
func (m *PlanTestCaseMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PlanTestCaseMutation builder.
func (m *PlanTestCaseMutation) Filter() *PlanTestCaseFilter {
	return &PlanTestCaseFilter{config: m.config, predicateAdder: m}
}

// PlanTestCaseFilter provides a generic filtering capability at runtime for PlanTestCaseQuery.
type PlanTestCaseFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *PlanTestCaseFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *PlanTestCaseFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(plantestcase.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *PlanTestCaseFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(plantestcase.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *PlanTestCaseFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(plantestcase.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *PlanTestCaseFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(plantestcase.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *PlanTestCaseFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(plantestcase.FieldEntID))
}

// WhereTestPlanID applies the entql [16]byte predicate on the test_plan_id field.
func (f *PlanTestCaseFilter) WhereTestPlanID(p entql.ValueP) {
	f.Where(p.Field(plantestcase.FieldTestPlanID))
}

// WhereTestCaseID applies the entql [16]byte predicate on the test_case_id field.
func (f *PlanTestCaseFilter) WhereTestCaseID(p entql.ValueP) {
	f.Where(p.Field(plantestcase.FieldTestCaseID))
}

// WhereInput applies the entql string predicate on the input field.
func (f *PlanTestCaseFilter) WhereInput(p entql.StringP) {
	f.Where(p.Field(plantestcase.FieldInput))
}

// WhereOutput applies the entql string predicate on the output field.
func (f *PlanTestCaseFilter) WhereOutput(p entql.StringP) {
	f.Where(p.Field(plantestcase.FieldOutput))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *PlanTestCaseFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(plantestcase.FieldDescription))
}

// WhereTestUserID applies the entql [16]byte predicate on the test_user_id field.
func (f *PlanTestCaseFilter) WhereTestUserID(p entql.ValueP) {
	f.Where(p.Field(plantestcase.FieldTestUserID))
}

// WhereRunDuration applies the entql uint32 predicate on the run_duration field.
func (f *PlanTestCaseFilter) WhereRunDuration(p entql.Uint32P) {
	f.Where(p.Field(plantestcase.FieldRunDuration))
}

// WhereResult applies the entql string predicate on the result field.
func (f *PlanTestCaseFilter) WhereResult(p entql.StringP) {
	f.Where(p.Field(plantestcase.FieldResult))
}

// WhereIndex applies the entql uint32 predicate on the index field.
func (f *PlanTestCaseFilter) WhereIndex(p entql.Uint32P) {
	f.Where(p.Field(plantestcase.FieldIndex))
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

// WhereID applies the entql uint32 predicate on the id field.
func (f *TestCaseFilter) WhereID(p entql.Uint32P) {
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

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *TestCaseFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(testcase.FieldEntID))
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

// WhereInput applies the entql string predicate on the input field.
func (f *TestCaseFilter) WhereInput(p entql.StringP) {
	f.Where(p.Field(testcase.FieldInput))
}

// WhereInputDesc applies the entql string predicate on the input_desc field.
func (f *TestCaseFilter) WhereInputDesc(p entql.StringP) {
	f.Where(p.Field(testcase.FieldInputDesc))
}

// WhereExpectation applies the entql string predicate on the expectation field.
func (f *TestCaseFilter) WhereExpectation(p entql.StringP) {
	f.Where(p.Field(testcase.FieldExpectation))
}

// WhereOutputDesc applies the entql string predicate on the output_desc field.
func (f *TestCaseFilter) WhereOutputDesc(p entql.StringP) {
	f.Where(p.Field(testcase.FieldOutputDesc))
}

// WhereTestCaseType applies the entql string predicate on the test_case_type field.
func (f *TestCaseFilter) WhereTestCaseType(p entql.StringP) {
	f.Where(p.Field(testcase.FieldTestCaseType))
}

// WhereTestCaseClass applies the entql string predicate on the test_case_class field.
func (f *TestCaseFilter) WhereTestCaseClass(p entql.StringP) {
	f.Where(p.Field(testcase.FieldTestCaseClass))
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

// WhereID applies the entql uint32 predicate on the id field.
func (f *TestPlanFilter) WhereID(p entql.Uint32P) {
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

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *TestPlanFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(testplan.FieldEntID))
}

// WhereName applies the entql string predicate on the name field.
func (f *TestPlanFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(testplan.FieldName))
}

// WhereState applies the entql string predicate on the state field.
func (f *TestPlanFilter) WhereState(p entql.StringP) {
	f.Where(p.Field(testplan.FieldState))
}

// WhereCreatedBy applies the entql [16]byte predicate on the created_by field.
func (f *TestPlanFilter) WhereCreatedBy(p entql.ValueP) {
	f.Where(p.Field(testplan.FieldCreatedBy))
}

// WhereExecutor applies the entql [16]byte predicate on the executor field.
func (f *TestPlanFilter) WhereExecutor(p entql.ValueP) {
	f.Where(p.Field(testplan.FieldExecutor))
}

// WhereFails applies the entql uint32 predicate on the fails field.
func (f *TestPlanFilter) WhereFails(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldFails))
}

// WherePasses applies the entql uint32 predicate on the passes field.
func (f *TestPlanFilter) WherePasses(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldPasses))
}

// WhereSkips applies the entql uint32 predicate on the skips field.
func (f *TestPlanFilter) WhereSkips(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldSkips))
}

// WhereRunDuration applies the entql uint32 predicate on the run_duration field.
func (f *TestPlanFilter) WhereRunDuration(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldRunDuration))
}

// WhereDeadline applies the entql uint32 predicate on the deadline field.
func (f *TestPlanFilter) WhereDeadline(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldDeadline))
}

// WhereResult applies the entql string predicate on the result field.
func (f *TestPlanFilter) WhereResult(p entql.StringP) {
	f.Where(p.Field(testplan.FieldResult))
}
