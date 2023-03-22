package corev1connect

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/koblas/grpc-todo/gen/core/v1/corev1connect.TodoServiceClient -o ./todo_service_client_mock_test.go -n TodoServiceClientMock

import (
	context "context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	connect_go "github.com/bufbuild/connect-go"
	"github.com/gojuno/minimock/v3"
	v1 "github.com/koblas/grpc-todo/gen/core/v1"
)

// TodoServiceClientMock implements TodoServiceClient
type TodoServiceClientMock struct {
	t minimock.Tester

	funcTodoAdd          func(ctx context.Context, pp1 *connect_go.Request[v1.TodoAddRequest]) (pp2 *connect_go.Response[v1.TodoAddResponse], err error)
	inspectFuncTodoAdd   func(ctx context.Context, pp1 *connect_go.Request[v1.TodoAddRequest])
	afterTodoAddCounter  uint64
	beforeTodoAddCounter uint64
	TodoAddMock          mTodoServiceClientMockTodoAdd

	funcTodoDelete          func(ctx context.Context, pp1 *connect_go.Request[v1.TodoDeleteRequest]) (pp2 *connect_go.Response[v1.TodoDeleteResponse], err error)
	inspectFuncTodoDelete   func(ctx context.Context, pp1 *connect_go.Request[v1.TodoDeleteRequest])
	afterTodoDeleteCounter  uint64
	beforeTodoDeleteCounter uint64
	TodoDeleteMock          mTodoServiceClientMockTodoDelete

	funcTodoList          func(ctx context.Context, pp1 *connect_go.Request[v1.TodoListRequest]) (pp2 *connect_go.Response[v1.TodoListResponse], err error)
	inspectFuncTodoList   func(ctx context.Context, pp1 *connect_go.Request[v1.TodoListRequest])
	afterTodoListCounter  uint64
	beforeTodoListCounter uint64
	TodoListMock          mTodoServiceClientMockTodoList
}

// NewTodoServiceClientMock returns a mock for TodoServiceClient
func NewTodoServiceClientMock(t minimock.Tester) *TodoServiceClientMock {
	m := &TodoServiceClientMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.TodoAddMock = mTodoServiceClientMockTodoAdd{mock: m}
	m.TodoAddMock.callArgs = []*TodoServiceClientMockTodoAddParams{}

	m.TodoDeleteMock = mTodoServiceClientMockTodoDelete{mock: m}
	m.TodoDeleteMock.callArgs = []*TodoServiceClientMockTodoDeleteParams{}

	m.TodoListMock = mTodoServiceClientMockTodoList{mock: m}
	m.TodoListMock.callArgs = []*TodoServiceClientMockTodoListParams{}

	return m
}

type mTodoServiceClientMockTodoAdd struct {
	mock               *TodoServiceClientMock
	defaultExpectation *TodoServiceClientMockTodoAddExpectation
	expectations       []*TodoServiceClientMockTodoAddExpectation

	callArgs []*TodoServiceClientMockTodoAddParams
	mutex    sync.RWMutex
}

// TodoServiceClientMockTodoAddExpectation specifies expectation struct of the TodoServiceClient.TodoAdd
type TodoServiceClientMockTodoAddExpectation struct {
	mock    *TodoServiceClientMock
	params  *TodoServiceClientMockTodoAddParams
	results *TodoServiceClientMockTodoAddResults
	Counter uint64
}

// TodoServiceClientMockTodoAddParams contains parameters of the TodoServiceClient.TodoAdd
type TodoServiceClientMockTodoAddParams struct {
	ctx context.Context
	pp1 *connect_go.Request[v1.TodoAddRequest]
}

// TodoServiceClientMockTodoAddResults contains results of the TodoServiceClient.TodoAdd
type TodoServiceClientMockTodoAddResults struct {
	pp2 *connect_go.Response[v1.TodoAddResponse]
	err error
}

// Expect sets up expected params for TodoServiceClient.TodoAdd
func (mmTodoAdd *mTodoServiceClientMockTodoAdd) Expect(ctx context.Context, pp1 *connect_go.Request[v1.TodoAddRequest]) *mTodoServiceClientMockTodoAdd {
	if mmTodoAdd.mock.funcTodoAdd != nil {
		mmTodoAdd.mock.t.Fatalf("TodoServiceClientMock.TodoAdd mock is already set by Set")
	}

	if mmTodoAdd.defaultExpectation == nil {
		mmTodoAdd.defaultExpectation = &TodoServiceClientMockTodoAddExpectation{}
	}

	mmTodoAdd.defaultExpectation.params = &TodoServiceClientMockTodoAddParams{ctx, pp1}
	for _, e := range mmTodoAdd.expectations {
		if minimock.Equal(e.params, mmTodoAdd.defaultExpectation.params) {
			mmTodoAdd.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmTodoAdd.defaultExpectation.params)
		}
	}

	return mmTodoAdd
}

// Inspect accepts an inspector function that has same arguments as the TodoServiceClient.TodoAdd
func (mmTodoAdd *mTodoServiceClientMockTodoAdd) Inspect(f func(ctx context.Context, pp1 *connect_go.Request[v1.TodoAddRequest])) *mTodoServiceClientMockTodoAdd {
	if mmTodoAdd.mock.inspectFuncTodoAdd != nil {
		mmTodoAdd.mock.t.Fatalf("Inspect function is already set for TodoServiceClientMock.TodoAdd")
	}

	mmTodoAdd.mock.inspectFuncTodoAdd = f

	return mmTodoAdd
}

// Return sets up results that will be returned by TodoServiceClient.TodoAdd
func (mmTodoAdd *mTodoServiceClientMockTodoAdd) Return(pp2 *connect_go.Response[v1.TodoAddResponse], err error) *TodoServiceClientMock {
	if mmTodoAdd.mock.funcTodoAdd != nil {
		mmTodoAdd.mock.t.Fatalf("TodoServiceClientMock.TodoAdd mock is already set by Set")
	}

	if mmTodoAdd.defaultExpectation == nil {
		mmTodoAdd.defaultExpectation = &TodoServiceClientMockTodoAddExpectation{mock: mmTodoAdd.mock}
	}
	mmTodoAdd.defaultExpectation.results = &TodoServiceClientMockTodoAddResults{pp2, err}
	return mmTodoAdd.mock
}

// Set uses given function f to mock the TodoServiceClient.TodoAdd method
func (mmTodoAdd *mTodoServiceClientMockTodoAdd) Set(f func(ctx context.Context, pp1 *connect_go.Request[v1.TodoAddRequest]) (pp2 *connect_go.Response[v1.TodoAddResponse], err error)) *TodoServiceClientMock {
	if mmTodoAdd.defaultExpectation != nil {
		mmTodoAdd.mock.t.Fatalf("Default expectation is already set for the TodoServiceClient.TodoAdd method")
	}

	if len(mmTodoAdd.expectations) > 0 {
		mmTodoAdd.mock.t.Fatalf("Some expectations are already set for the TodoServiceClient.TodoAdd method")
	}

	mmTodoAdd.mock.funcTodoAdd = f
	return mmTodoAdd.mock
}

// When sets expectation for the TodoServiceClient.TodoAdd which will trigger the result defined by the following
// Then helper
func (mmTodoAdd *mTodoServiceClientMockTodoAdd) When(ctx context.Context, pp1 *connect_go.Request[v1.TodoAddRequest]) *TodoServiceClientMockTodoAddExpectation {
	if mmTodoAdd.mock.funcTodoAdd != nil {
		mmTodoAdd.mock.t.Fatalf("TodoServiceClientMock.TodoAdd mock is already set by Set")
	}

	expectation := &TodoServiceClientMockTodoAddExpectation{
		mock:   mmTodoAdd.mock,
		params: &TodoServiceClientMockTodoAddParams{ctx, pp1},
	}
	mmTodoAdd.expectations = append(mmTodoAdd.expectations, expectation)
	return expectation
}

// Then sets up TodoServiceClient.TodoAdd return parameters for the expectation previously defined by the When method
func (e *TodoServiceClientMockTodoAddExpectation) Then(pp2 *connect_go.Response[v1.TodoAddResponse], err error) *TodoServiceClientMock {
	e.results = &TodoServiceClientMockTodoAddResults{pp2, err}
	return e.mock
}

// TodoAdd implements TodoServiceClient
func (mmTodoAdd *TodoServiceClientMock) TodoAdd(ctx context.Context, pp1 *connect_go.Request[v1.TodoAddRequest]) (pp2 *connect_go.Response[v1.TodoAddResponse], err error) {
	mm_atomic.AddUint64(&mmTodoAdd.beforeTodoAddCounter, 1)
	defer mm_atomic.AddUint64(&mmTodoAdd.afterTodoAddCounter, 1)

	if mmTodoAdd.inspectFuncTodoAdd != nil {
		mmTodoAdd.inspectFuncTodoAdd(ctx, pp1)
	}

	mm_params := &TodoServiceClientMockTodoAddParams{ctx, pp1}

	// Record call args
	mmTodoAdd.TodoAddMock.mutex.Lock()
	mmTodoAdd.TodoAddMock.callArgs = append(mmTodoAdd.TodoAddMock.callArgs, mm_params)
	mmTodoAdd.TodoAddMock.mutex.Unlock()

	for _, e := range mmTodoAdd.TodoAddMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.pp2, e.results.err
		}
	}

	if mmTodoAdd.TodoAddMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmTodoAdd.TodoAddMock.defaultExpectation.Counter, 1)
		mm_want := mmTodoAdd.TodoAddMock.defaultExpectation.params
		mm_got := TodoServiceClientMockTodoAddParams{ctx, pp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmTodoAdd.t.Errorf("TodoServiceClientMock.TodoAdd got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmTodoAdd.TodoAddMock.defaultExpectation.results
		if mm_results == nil {
			mmTodoAdd.t.Fatal("No results are set for the TodoServiceClientMock.TodoAdd")
		}
		return (*mm_results).pp2, (*mm_results).err
	}
	if mmTodoAdd.funcTodoAdd != nil {
		return mmTodoAdd.funcTodoAdd(ctx, pp1)
	}
	mmTodoAdd.t.Fatalf("Unexpected call to TodoServiceClientMock.TodoAdd. %v %v", ctx, pp1)
	return
}

// TodoAddAfterCounter returns a count of finished TodoServiceClientMock.TodoAdd invocations
func (mmTodoAdd *TodoServiceClientMock) TodoAddAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmTodoAdd.afterTodoAddCounter)
}

// TodoAddBeforeCounter returns a count of TodoServiceClientMock.TodoAdd invocations
func (mmTodoAdd *TodoServiceClientMock) TodoAddBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmTodoAdd.beforeTodoAddCounter)
}

// Calls returns a list of arguments used in each call to TodoServiceClientMock.TodoAdd.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmTodoAdd *mTodoServiceClientMockTodoAdd) Calls() []*TodoServiceClientMockTodoAddParams {
	mmTodoAdd.mutex.RLock()

	argCopy := make([]*TodoServiceClientMockTodoAddParams, len(mmTodoAdd.callArgs))
	copy(argCopy, mmTodoAdd.callArgs)

	mmTodoAdd.mutex.RUnlock()

	return argCopy
}

// MinimockTodoAddDone returns true if the count of the TodoAdd invocations corresponds
// the number of defined expectations
func (m *TodoServiceClientMock) MinimockTodoAddDone() bool {
	for _, e := range m.TodoAddMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.TodoAddMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterTodoAddCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcTodoAdd != nil && mm_atomic.LoadUint64(&m.afterTodoAddCounter) < 1 {
		return false
	}
	return true
}

// MinimockTodoAddInspect logs each unmet expectation
func (m *TodoServiceClientMock) MinimockTodoAddInspect() {
	for _, e := range m.TodoAddMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TodoServiceClientMock.TodoAdd with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.TodoAddMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterTodoAddCounter) < 1 {
		if m.TodoAddMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TodoServiceClientMock.TodoAdd")
		} else {
			m.t.Errorf("Expected call to TodoServiceClientMock.TodoAdd with params: %#v", *m.TodoAddMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcTodoAdd != nil && mm_atomic.LoadUint64(&m.afterTodoAddCounter) < 1 {
		m.t.Error("Expected call to TodoServiceClientMock.TodoAdd")
	}
}

type mTodoServiceClientMockTodoDelete struct {
	mock               *TodoServiceClientMock
	defaultExpectation *TodoServiceClientMockTodoDeleteExpectation
	expectations       []*TodoServiceClientMockTodoDeleteExpectation

	callArgs []*TodoServiceClientMockTodoDeleteParams
	mutex    sync.RWMutex
}

// TodoServiceClientMockTodoDeleteExpectation specifies expectation struct of the TodoServiceClient.TodoDelete
type TodoServiceClientMockTodoDeleteExpectation struct {
	mock    *TodoServiceClientMock
	params  *TodoServiceClientMockTodoDeleteParams
	results *TodoServiceClientMockTodoDeleteResults
	Counter uint64
}

// TodoServiceClientMockTodoDeleteParams contains parameters of the TodoServiceClient.TodoDelete
type TodoServiceClientMockTodoDeleteParams struct {
	ctx context.Context
	pp1 *connect_go.Request[v1.TodoDeleteRequest]
}

// TodoServiceClientMockTodoDeleteResults contains results of the TodoServiceClient.TodoDelete
type TodoServiceClientMockTodoDeleteResults struct {
	pp2 *connect_go.Response[v1.TodoDeleteResponse]
	err error
}

// Expect sets up expected params for TodoServiceClient.TodoDelete
func (mmTodoDelete *mTodoServiceClientMockTodoDelete) Expect(ctx context.Context, pp1 *connect_go.Request[v1.TodoDeleteRequest]) *mTodoServiceClientMockTodoDelete {
	if mmTodoDelete.mock.funcTodoDelete != nil {
		mmTodoDelete.mock.t.Fatalf("TodoServiceClientMock.TodoDelete mock is already set by Set")
	}

	if mmTodoDelete.defaultExpectation == nil {
		mmTodoDelete.defaultExpectation = &TodoServiceClientMockTodoDeleteExpectation{}
	}

	mmTodoDelete.defaultExpectation.params = &TodoServiceClientMockTodoDeleteParams{ctx, pp1}
	for _, e := range mmTodoDelete.expectations {
		if minimock.Equal(e.params, mmTodoDelete.defaultExpectation.params) {
			mmTodoDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmTodoDelete.defaultExpectation.params)
		}
	}

	return mmTodoDelete
}

// Inspect accepts an inspector function that has same arguments as the TodoServiceClient.TodoDelete
func (mmTodoDelete *mTodoServiceClientMockTodoDelete) Inspect(f func(ctx context.Context, pp1 *connect_go.Request[v1.TodoDeleteRequest])) *mTodoServiceClientMockTodoDelete {
	if mmTodoDelete.mock.inspectFuncTodoDelete != nil {
		mmTodoDelete.mock.t.Fatalf("Inspect function is already set for TodoServiceClientMock.TodoDelete")
	}

	mmTodoDelete.mock.inspectFuncTodoDelete = f

	return mmTodoDelete
}

// Return sets up results that will be returned by TodoServiceClient.TodoDelete
func (mmTodoDelete *mTodoServiceClientMockTodoDelete) Return(pp2 *connect_go.Response[v1.TodoDeleteResponse], err error) *TodoServiceClientMock {
	if mmTodoDelete.mock.funcTodoDelete != nil {
		mmTodoDelete.mock.t.Fatalf("TodoServiceClientMock.TodoDelete mock is already set by Set")
	}

	if mmTodoDelete.defaultExpectation == nil {
		mmTodoDelete.defaultExpectation = &TodoServiceClientMockTodoDeleteExpectation{mock: mmTodoDelete.mock}
	}
	mmTodoDelete.defaultExpectation.results = &TodoServiceClientMockTodoDeleteResults{pp2, err}
	return mmTodoDelete.mock
}

// Set uses given function f to mock the TodoServiceClient.TodoDelete method
func (mmTodoDelete *mTodoServiceClientMockTodoDelete) Set(f func(ctx context.Context, pp1 *connect_go.Request[v1.TodoDeleteRequest]) (pp2 *connect_go.Response[v1.TodoDeleteResponse], err error)) *TodoServiceClientMock {
	if mmTodoDelete.defaultExpectation != nil {
		mmTodoDelete.mock.t.Fatalf("Default expectation is already set for the TodoServiceClient.TodoDelete method")
	}

	if len(mmTodoDelete.expectations) > 0 {
		mmTodoDelete.mock.t.Fatalf("Some expectations are already set for the TodoServiceClient.TodoDelete method")
	}

	mmTodoDelete.mock.funcTodoDelete = f
	return mmTodoDelete.mock
}

// When sets expectation for the TodoServiceClient.TodoDelete which will trigger the result defined by the following
// Then helper
func (mmTodoDelete *mTodoServiceClientMockTodoDelete) When(ctx context.Context, pp1 *connect_go.Request[v1.TodoDeleteRequest]) *TodoServiceClientMockTodoDeleteExpectation {
	if mmTodoDelete.mock.funcTodoDelete != nil {
		mmTodoDelete.mock.t.Fatalf("TodoServiceClientMock.TodoDelete mock is already set by Set")
	}

	expectation := &TodoServiceClientMockTodoDeleteExpectation{
		mock:   mmTodoDelete.mock,
		params: &TodoServiceClientMockTodoDeleteParams{ctx, pp1},
	}
	mmTodoDelete.expectations = append(mmTodoDelete.expectations, expectation)
	return expectation
}

// Then sets up TodoServiceClient.TodoDelete return parameters for the expectation previously defined by the When method
func (e *TodoServiceClientMockTodoDeleteExpectation) Then(pp2 *connect_go.Response[v1.TodoDeleteResponse], err error) *TodoServiceClientMock {
	e.results = &TodoServiceClientMockTodoDeleteResults{pp2, err}
	return e.mock
}

// TodoDelete implements TodoServiceClient
func (mmTodoDelete *TodoServiceClientMock) TodoDelete(ctx context.Context, pp1 *connect_go.Request[v1.TodoDeleteRequest]) (pp2 *connect_go.Response[v1.TodoDeleteResponse], err error) {
	mm_atomic.AddUint64(&mmTodoDelete.beforeTodoDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmTodoDelete.afterTodoDeleteCounter, 1)

	if mmTodoDelete.inspectFuncTodoDelete != nil {
		mmTodoDelete.inspectFuncTodoDelete(ctx, pp1)
	}

	mm_params := &TodoServiceClientMockTodoDeleteParams{ctx, pp1}

	// Record call args
	mmTodoDelete.TodoDeleteMock.mutex.Lock()
	mmTodoDelete.TodoDeleteMock.callArgs = append(mmTodoDelete.TodoDeleteMock.callArgs, mm_params)
	mmTodoDelete.TodoDeleteMock.mutex.Unlock()

	for _, e := range mmTodoDelete.TodoDeleteMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.pp2, e.results.err
		}
	}

	if mmTodoDelete.TodoDeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmTodoDelete.TodoDeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmTodoDelete.TodoDeleteMock.defaultExpectation.params
		mm_got := TodoServiceClientMockTodoDeleteParams{ctx, pp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmTodoDelete.t.Errorf("TodoServiceClientMock.TodoDelete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmTodoDelete.TodoDeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmTodoDelete.t.Fatal("No results are set for the TodoServiceClientMock.TodoDelete")
		}
		return (*mm_results).pp2, (*mm_results).err
	}
	if mmTodoDelete.funcTodoDelete != nil {
		return mmTodoDelete.funcTodoDelete(ctx, pp1)
	}
	mmTodoDelete.t.Fatalf("Unexpected call to TodoServiceClientMock.TodoDelete. %v %v", ctx, pp1)
	return
}

// TodoDeleteAfterCounter returns a count of finished TodoServiceClientMock.TodoDelete invocations
func (mmTodoDelete *TodoServiceClientMock) TodoDeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmTodoDelete.afterTodoDeleteCounter)
}

// TodoDeleteBeforeCounter returns a count of TodoServiceClientMock.TodoDelete invocations
func (mmTodoDelete *TodoServiceClientMock) TodoDeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmTodoDelete.beforeTodoDeleteCounter)
}

// Calls returns a list of arguments used in each call to TodoServiceClientMock.TodoDelete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmTodoDelete *mTodoServiceClientMockTodoDelete) Calls() []*TodoServiceClientMockTodoDeleteParams {
	mmTodoDelete.mutex.RLock()

	argCopy := make([]*TodoServiceClientMockTodoDeleteParams, len(mmTodoDelete.callArgs))
	copy(argCopy, mmTodoDelete.callArgs)

	mmTodoDelete.mutex.RUnlock()

	return argCopy
}

// MinimockTodoDeleteDone returns true if the count of the TodoDelete invocations corresponds
// the number of defined expectations
func (m *TodoServiceClientMock) MinimockTodoDeleteDone() bool {
	for _, e := range m.TodoDeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.TodoDeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterTodoDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcTodoDelete != nil && mm_atomic.LoadUint64(&m.afterTodoDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockTodoDeleteInspect logs each unmet expectation
func (m *TodoServiceClientMock) MinimockTodoDeleteInspect() {
	for _, e := range m.TodoDeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TodoServiceClientMock.TodoDelete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.TodoDeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterTodoDeleteCounter) < 1 {
		if m.TodoDeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TodoServiceClientMock.TodoDelete")
		} else {
			m.t.Errorf("Expected call to TodoServiceClientMock.TodoDelete with params: %#v", *m.TodoDeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcTodoDelete != nil && mm_atomic.LoadUint64(&m.afterTodoDeleteCounter) < 1 {
		m.t.Error("Expected call to TodoServiceClientMock.TodoDelete")
	}
}

type mTodoServiceClientMockTodoList struct {
	mock               *TodoServiceClientMock
	defaultExpectation *TodoServiceClientMockTodoListExpectation
	expectations       []*TodoServiceClientMockTodoListExpectation

	callArgs []*TodoServiceClientMockTodoListParams
	mutex    sync.RWMutex
}

// TodoServiceClientMockTodoListExpectation specifies expectation struct of the TodoServiceClient.TodoList
type TodoServiceClientMockTodoListExpectation struct {
	mock    *TodoServiceClientMock
	params  *TodoServiceClientMockTodoListParams
	results *TodoServiceClientMockTodoListResults
	Counter uint64
}

// TodoServiceClientMockTodoListParams contains parameters of the TodoServiceClient.TodoList
type TodoServiceClientMockTodoListParams struct {
	ctx context.Context
	pp1 *connect_go.Request[v1.TodoListRequest]
}

// TodoServiceClientMockTodoListResults contains results of the TodoServiceClient.TodoList
type TodoServiceClientMockTodoListResults struct {
	pp2 *connect_go.Response[v1.TodoListResponse]
	err error
}

// Expect sets up expected params for TodoServiceClient.TodoList
func (mmTodoList *mTodoServiceClientMockTodoList) Expect(ctx context.Context, pp1 *connect_go.Request[v1.TodoListRequest]) *mTodoServiceClientMockTodoList {
	if mmTodoList.mock.funcTodoList != nil {
		mmTodoList.mock.t.Fatalf("TodoServiceClientMock.TodoList mock is already set by Set")
	}

	if mmTodoList.defaultExpectation == nil {
		mmTodoList.defaultExpectation = &TodoServiceClientMockTodoListExpectation{}
	}

	mmTodoList.defaultExpectation.params = &TodoServiceClientMockTodoListParams{ctx, pp1}
	for _, e := range mmTodoList.expectations {
		if minimock.Equal(e.params, mmTodoList.defaultExpectation.params) {
			mmTodoList.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmTodoList.defaultExpectation.params)
		}
	}

	return mmTodoList
}

// Inspect accepts an inspector function that has same arguments as the TodoServiceClient.TodoList
func (mmTodoList *mTodoServiceClientMockTodoList) Inspect(f func(ctx context.Context, pp1 *connect_go.Request[v1.TodoListRequest])) *mTodoServiceClientMockTodoList {
	if mmTodoList.mock.inspectFuncTodoList != nil {
		mmTodoList.mock.t.Fatalf("Inspect function is already set for TodoServiceClientMock.TodoList")
	}

	mmTodoList.mock.inspectFuncTodoList = f

	return mmTodoList
}

// Return sets up results that will be returned by TodoServiceClient.TodoList
func (mmTodoList *mTodoServiceClientMockTodoList) Return(pp2 *connect_go.Response[v1.TodoListResponse], err error) *TodoServiceClientMock {
	if mmTodoList.mock.funcTodoList != nil {
		mmTodoList.mock.t.Fatalf("TodoServiceClientMock.TodoList mock is already set by Set")
	}

	if mmTodoList.defaultExpectation == nil {
		mmTodoList.defaultExpectation = &TodoServiceClientMockTodoListExpectation{mock: mmTodoList.mock}
	}
	mmTodoList.defaultExpectation.results = &TodoServiceClientMockTodoListResults{pp2, err}
	return mmTodoList.mock
}

// Set uses given function f to mock the TodoServiceClient.TodoList method
func (mmTodoList *mTodoServiceClientMockTodoList) Set(f func(ctx context.Context, pp1 *connect_go.Request[v1.TodoListRequest]) (pp2 *connect_go.Response[v1.TodoListResponse], err error)) *TodoServiceClientMock {
	if mmTodoList.defaultExpectation != nil {
		mmTodoList.mock.t.Fatalf("Default expectation is already set for the TodoServiceClient.TodoList method")
	}

	if len(mmTodoList.expectations) > 0 {
		mmTodoList.mock.t.Fatalf("Some expectations are already set for the TodoServiceClient.TodoList method")
	}

	mmTodoList.mock.funcTodoList = f
	return mmTodoList.mock
}

// When sets expectation for the TodoServiceClient.TodoList which will trigger the result defined by the following
// Then helper
func (mmTodoList *mTodoServiceClientMockTodoList) When(ctx context.Context, pp1 *connect_go.Request[v1.TodoListRequest]) *TodoServiceClientMockTodoListExpectation {
	if mmTodoList.mock.funcTodoList != nil {
		mmTodoList.mock.t.Fatalf("TodoServiceClientMock.TodoList mock is already set by Set")
	}

	expectation := &TodoServiceClientMockTodoListExpectation{
		mock:   mmTodoList.mock,
		params: &TodoServiceClientMockTodoListParams{ctx, pp1},
	}
	mmTodoList.expectations = append(mmTodoList.expectations, expectation)
	return expectation
}

// Then sets up TodoServiceClient.TodoList return parameters for the expectation previously defined by the When method
func (e *TodoServiceClientMockTodoListExpectation) Then(pp2 *connect_go.Response[v1.TodoListResponse], err error) *TodoServiceClientMock {
	e.results = &TodoServiceClientMockTodoListResults{pp2, err}
	return e.mock
}

// TodoList implements TodoServiceClient
func (mmTodoList *TodoServiceClientMock) TodoList(ctx context.Context, pp1 *connect_go.Request[v1.TodoListRequest]) (pp2 *connect_go.Response[v1.TodoListResponse], err error) {
	mm_atomic.AddUint64(&mmTodoList.beforeTodoListCounter, 1)
	defer mm_atomic.AddUint64(&mmTodoList.afterTodoListCounter, 1)

	if mmTodoList.inspectFuncTodoList != nil {
		mmTodoList.inspectFuncTodoList(ctx, pp1)
	}

	mm_params := &TodoServiceClientMockTodoListParams{ctx, pp1}

	// Record call args
	mmTodoList.TodoListMock.mutex.Lock()
	mmTodoList.TodoListMock.callArgs = append(mmTodoList.TodoListMock.callArgs, mm_params)
	mmTodoList.TodoListMock.mutex.Unlock()

	for _, e := range mmTodoList.TodoListMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.pp2, e.results.err
		}
	}

	if mmTodoList.TodoListMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmTodoList.TodoListMock.defaultExpectation.Counter, 1)
		mm_want := mmTodoList.TodoListMock.defaultExpectation.params
		mm_got := TodoServiceClientMockTodoListParams{ctx, pp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmTodoList.t.Errorf("TodoServiceClientMock.TodoList got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmTodoList.TodoListMock.defaultExpectation.results
		if mm_results == nil {
			mmTodoList.t.Fatal("No results are set for the TodoServiceClientMock.TodoList")
		}
		return (*mm_results).pp2, (*mm_results).err
	}
	if mmTodoList.funcTodoList != nil {
		return mmTodoList.funcTodoList(ctx, pp1)
	}
	mmTodoList.t.Fatalf("Unexpected call to TodoServiceClientMock.TodoList. %v %v", ctx, pp1)
	return
}

// TodoListAfterCounter returns a count of finished TodoServiceClientMock.TodoList invocations
func (mmTodoList *TodoServiceClientMock) TodoListAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmTodoList.afterTodoListCounter)
}

// TodoListBeforeCounter returns a count of TodoServiceClientMock.TodoList invocations
func (mmTodoList *TodoServiceClientMock) TodoListBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmTodoList.beforeTodoListCounter)
}

// Calls returns a list of arguments used in each call to TodoServiceClientMock.TodoList.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmTodoList *mTodoServiceClientMockTodoList) Calls() []*TodoServiceClientMockTodoListParams {
	mmTodoList.mutex.RLock()

	argCopy := make([]*TodoServiceClientMockTodoListParams, len(mmTodoList.callArgs))
	copy(argCopy, mmTodoList.callArgs)

	mmTodoList.mutex.RUnlock()

	return argCopy
}

// MinimockTodoListDone returns true if the count of the TodoList invocations corresponds
// the number of defined expectations
func (m *TodoServiceClientMock) MinimockTodoListDone() bool {
	for _, e := range m.TodoListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.TodoListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterTodoListCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcTodoList != nil && mm_atomic.LoadUint64(&m.afterTodoListCounter) < 1 {
		return false
	}
	return true
}

// MinimockTodoListInspect logs each unmet expectation
func (m *TodoServiceClientMock) MinimockTodoListInspect() {
	for _, e := range m.TodoListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TodoServiceClientMock.TodoList with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.TodoListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterTodoListCounter) < 1 {
		if m.TodoListMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TodoServiceClientMock.TodoList")
		} else {
			m.t.Errorf("Expected call to TodoServiceClientMock.TodoList with params: %#v", *m.TodoListMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcTodoList != nil && mm_atomic.LoadUint64(&m.afterTodoListCounter) < 1 {
		m.t.Error("Expected call to TodoServiceClientMock.TodoList")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *TodoServiceClientMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockTodoAddInspect()

		m.MinimockTodoDeleteInspect()

		m.MinimockTodoListInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *TodoServiceClientMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *TodoServiceClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockTodoAddDone() &&
		m.MinimockTodoDeleteDone() &&
		m.MinimockTodoListDone()
}
