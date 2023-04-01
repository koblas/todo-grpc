package messagev1connect

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/koblas/grpc-todo/gen/core/message/v1/messagev1connect.MessageServiceHandler -o ./message_service_handler_mock.go -n MessageServiceHandlerMock

import (
	context "context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	connect_go "github.com/bufbuild/connect-go"
	"github.com/gojuno/minimock/v3"
	v1 "github.com/koblas/grpc-todo/gen/core/message/v1"
)

// MessageServiceHandlerMock implements MessageServiceHandler
type MessageServiceHandlerMock struct {
	t minimock.Tester

	funcAdd          func(ctx context.Context, pp1 *connect_go.Request[v1.AddRequest]) (pp2 *connect_go.Response[v1.AddResponse], err error)
	inspectFuncAdd   func(ctx context.Context, pp1 *connect_go.Request[v1.AddRequest])
	afterAddCounter  uint64
	beforeAddCounter uint64
	AddMock          mMessageServiceHandlerMockAdd

	funcDelete          func(ctx context.Context, pp1 *connect_go.Request[v1.DeleteRequest]) (pp2 *connect_go.Response[v1.DeleteResponse], err error)
	inspectFuncDelete   func(ctx context.Context, pp1 *connect_go.Request[v1.DeleteRequest])
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mMessageServiceHandlerMockDelete

	funcList          func(ctx context.Context, pp1 *connect_go.Request[v1.ListRequest]) (pp2 *connect_go.Response[v1.ListResponse], err error)
	inspectFuncList   func(ctx context.Context, pp1 *connect_go.Request[v1.ListRequest])
	afterListCounter  uint64
	beforeListCounter uint64
	ListMock          mMessageServiceHandlerMockList
}

// NewMessageServiceHandlerMock returns a mock for MessageServiceHandler
func NewMessageServiceHandlerMock(t minimock.Tester) *MessageServiceHandlerMock {
	m := &MessageServiceHandlerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddMock = mMessageServiceHandlerMockAdd{mock: m}
	m.AddMock.callArgs = []*MessageServiceHandlerMockAddParams{}

	m.DeleteMock = mMessageServiceHandlerMockDelete{mock: m}
	m.DeleteMock.callArgs = []*MessageServiceHandlerMockDeleteParams{}

	m.ListMock = mMessageServiceHandlerMockList{mock: m}
	m.ListMock.callArgs = []*MessageServiceHandlerMockListParams{}

	return m
}

type mMessageServiceHandlerMockAdd struct {
	mock               *MessageServiceHandlerMock
	defaultExpectation *MessageServiceHandlerMockAddExpectation
	expectations       []*MessageServiceHandlerMockAddExpectation

	callArgs []*MessageServiceHandlerMockAddParams
	mutex    sync.RWMutex
}

// MessageServiceHandlerMockAddExpectation specifies expectation struct of the MessageServiceHandler.Add
type MessageServiceHandlerMockAddExpectation struct {
	mock    *MessageServiceHandlerMock
	params  *MessageServiceHandlerMockAddParams
	results *MessageServiceHandlerMockAddResults
	Counter uint64
}

// MessageServiceHandlerMockAddParams contains parameters of the MessageServiceHandler.Add
type MessageServiceHandlerMockAddParams struct {
	ctx context.Context
	pp1 *connect_go.Request[v1.AddRequest]
}

// MessageServiceHandlerMockAddResults contains results of the MessageServiceHandler.Add
type MessageServiceHandlerMockAddResults struct {
	pp2 *connect_go.Response[v1.AddResponse]
	err error
}

// Expect sets up expected params for MessageServiceHandler.Add
func (mmAdd *mMessageServiceHandlerMockAdd) Expect(ctx context.Context, pp1 *connect_go.Request[v1.AddRequest]) *mMessageServiceHandlerMockAdd {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("MessageServiceHandlerMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &MessageServiceHandlerMockAddExpectation{}
	}

	mmAdd.defaultExpectation.params = &MessageServiceHandlerMockAddParams{ctx, pp1}
	for _, e := range mmAdd.expectations {
		if minimock.Equal(e.params, mmAdd.defaultExpectation.params) {
			mmAdd.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAdd.defaultExpectation.params)
		}
	}

	return mmAdd
}

// Inspect accepts an inspector function that has same arguments as the MessageServiceHandler.Add
func (mmAdd *mMessageServiceHandlerMockAdd) Inspect(f func(ctx context.Context, pp1 *connect_go.Request[v1.AddRequest])) *mMessageServiceHandlerMockAdd {
	if mmAdd.mock.inspectFuncAdd != nil {
		mmAdd.mock.t.Fatalf("Inspect function is already set for MessageServiceHandlerMock.Add")
	}

	mmAdd.mock.inspectFuncAdd = f

	return mmAdd
}

// Return sets up results that will be returned by MessageServiceHandler.Add
func (mmAdd *mMessageServiceHandlerMockAdd) Return(pp2 *connect_go.Response[v1.AddResponse], err error) *MessageServiceHandlerMock {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("MessageServiceHandlerMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &MessageServiceHandlerMockAddExpectation{mock: mmAdd.mock}
	}
	mmAdd.defaultExpectation.results = &MessageServiceHandlerMockAddResults{pp2, err}
	return mmAdd.mock
}

// Set uses given function f to mock the MessageServiceHandler.Add method
func (mmAdd *mMessageServiceHandlerMockAdd) Set(f func(ctx context.Context, pp1 *connect_go.Request[v1.AddRequest]) (pp2 *connect_go.Response[v1.AddResponse], err error)) *MessageServiceHandlerMock {
	if mmAdd.defaultExpectation != nil {
		mmAdd.mock.t.Fatalf("Default expectation is already set for the MessageServiceHandler.Add method")
	}

	if len(mmAdd.expectations) > 0 {
		mmAdd.mock.t.Fatalf("Some expectations are already set for the MessageServiceHandler.Add method")
	}

	mmAdd.mock.funcAdd = f
	return mmAdd.mock
}

// When sets expectation for the MessageServiceHandler.Add which will trigger the result defined by the following
// Then helper
func (mmAdd *mMessageServiceHandlerMockAdd) When(ctx context.Context, pp1 *connect_go.Request[v1.AddRequest]) *MessageServiceHandlerMockAddExpectation {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("MessageServiceHandlerMock.Add mock is already set by Set")
	}

	expectation := &MessageServiceHandlerMockAddExpectation{
		mock:   mmAdd.mock,
		params: &MessageServiceHandlerMockAddParams{ctx, pp1},
	}
	mmAdd.expectations = append(mmAdd.expectations, expectation)
	return expectation
}

// Then sets up MessageServiceHandler.Add return parameters for the expectation previously defined by the When method
func (e *MessageServiceHandlerMockAddExpectation) Then(pp2 *connect_go.Response[v1.AddResponse], err error) *MessageServiceHandlerMock {
	e.results = &MessageServiceHandlerMockAddResults{pp2, err}
	return e.mock
}

// Add implements MessageServiceHandler
func (mmAdd *MessageServiceHandlerMock) Add(ctx context.Context, pp1 *connect_go.Request[v1.AddRequest]) (pp2 *connect_go.Response[v1.AddResponse], err error) {
	mm_atomic.AddUint64(&mmAdd.beforeAddCounter, 1)
	defer mm_atomic.AddUint64(&mmAdd.afterAddCounter, 1)

	if mmAdd.inspectFuncAdd != nil {
		mmAdd.inspectFuncAdd(ctx, pp1)
	}

	mm_params := &MessageServiceHandlerMockAddParams{ctx, pp1}

	// Record call args
	mmAdd.AddMock.mutex.Lock()
	mmAdd.AddMock.callArgs = append(mmAdd.AddMock.callArgs, mm_params)
	mmAdd.AddMock.mutex.Unlock()

	for _, e := range mmAdd.AddMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.pp2, e.results.err
		}
	}

	if mmAdd.AddMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAdd.AddMock.defaultExpectation.Counter, 1)
		mm_want := mmAdd.AddMock.defaultExpectation.params
		mm_got := MessageServiceHandlerMockAddParams{ctx, pp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAdd.t.Errorf("MessageServiceHandlerMock.Add got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAdd.AddMock.defaultExpectation.results
		if mm_results == nil {
			mmAdd.t.Fatal("No results are set for the MessageServiceHandlerMock.Add")
		}
		return (*mm_results).pp2, (*mm_results).err
	}
	if mmAdd.funcAdd != nil {
		return mmAdd.funcAdd(ctx, pp1)
	}
	mmAdd.t.Fatalf("Unexpected call to MessageServiceHandlerMock.Add. %v %v", ctx, pp1)
	return
}

// AddAfterCounter returns a count of finished MessageServiceHandlerMock.Add invocations
func (mmAdd *MessageServiceHandlerMock) AddAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAdd.afterAddCounter)
}

// AddBeforeCounter returns a count of MessageServiceHandlerMock.Add invocations
func (mmAdd *MessageServiceHandlerMock) AddBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAdd.beforeAddCounter)
}

// Calls returns a list of arguments used in each call to MessageServiceHandlerMock.Add.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAdd *mMessageServiceHandlerMockAdd) Calls() []*MessageServiceHandlerMockAddParams {
	mmAdd.mutex.RLock()

	argCopy := make([]*MessageServiceHandlerMockAddParams, len(mmAdd.callArgs))
	copy(argCopy, mmAdd.callArgs)

	mmAdd.mutex.RUnlock()

	return argCopy
}

// MinimockAddDone returns true if the count of the Add invocations corresponds
// the number of defined expectations
func (m *MessageServiceHandlerMock) MinimockAddDone() bool {
	for _, e := range m.AddMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAdd != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		return false
	}
	return true
}

// MinimockAddInspect logs each unmet expectation
func (m *MessageServiceHandlerMock) MinimockAddInspect() {
	for _, e := range m.AddMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessageServiceHandlerMock.Add with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		if m.AddMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MessageServiceHandlerMock.Add")
		} else {
			m.t.Errorf("Expected call to MessageServiceHandlerMock.Add with params: %#v", *m.AddMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAdd != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		m.t.Error("Expected call to MessageServiceHandlerMock.Add")
	}
}

type mMessageServiceHandlerMockDelete struct {
	mock               *MessageServiceHandlerMock
	defaultExpectation *MessageServiceHandlerMockDeleteExpectation
	expectations       []*MessageServiceHandlerMockDeleteExpectation

	callArgs []*MessageServiceHandlerMockDeleteParams
	mutex    sync.RWMutex
}

// MessageServiceHandlerMockDeleteExpectation specifies expectation struct of the MessageServiceHandler.Delete
type MessageServiceHandlerMockDeleteExpectation struct {
	mock    *MessageServiceHandlerMock
	params  *MessageServiceHandlerMockDeleteParams
	results *MessageServiceHandlerMockDeleteResults
	Counter uint64
}

// MessageServiceHandlerMockDeleteParams contains parameters of the MessageServiceHandler.Delete
type MessageServiceHandlerMockDeleteParams struct {
	ctx context.Context
	pp1 *connect_go.Request[v1.DeleteRequest]
}

// MessageServiceHandlerMockDeleteResults contains results of the MessageServiceHandler.Delete
type MessageServiceHandlerMockDeleteResults struct {
	pp2 *connect_go.Response[v1.DeleteResponse]
	err error
}

// Expect sets up expected params for MessageServiceHandler.Delete
func (mmDelete *mMessageServiceHandlerMockDelete) Expect(ctx context.Context, pp1 *connect_go.Request[v1.DeleteRequest]) *mMessageServiceHandlerMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("MessageServiceHandlerMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &MessageServiceHandlerMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &MessageServiceHandlerMockDeleteParams{ctx, pp1}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the MessageServiceHandler.Delete
func (mmDelete *mMessageServiceHandlerMockDelete) Inspect(f func(ctx context.Context, pp1 *connect_go.Request[v1.DeleteRequest])) *mMessageServiceHandlerMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for MessageServiceHandlerMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by MessageServiceHandler.Delete
func (mmDelete *mMessageServiceHandlerMockDelete) Return(pp2 *connect_go.Response[v1.DeleteResponse], err error) *MessageServiceHandlerMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("MessageServiceHandlerMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &MessageServiceHandlerMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &MessageServiceHandlerMockDeleteResults{pp2, err}
	return mmDelete.mock
}

// Set uses given function f to mock the MessageServiceHandler.Delete method
func (mmDelete *mMessageServiceHandlerMockDelete) Set(f func(ctx context.Context, pp1 *connect_go.Request[v1.DeleteRequest]) (pp2 *connect_go.Response[v1.DeleteResponse], err error)) *MessageServiceHandlerMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the MessageServiceHandler.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the MessageServiceHandler.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the MessageServiceHandler.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mMessageServiceHandlerMockDelete) When(ctx context.Context, pp1 *connect_go.Request[v1.DeleteRequest]) *MessageServiceHandlerMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("MessageServiceHandlerMock.Delete mock is already set by Set")
	}

	expectation := &MessageServiceHandlerMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &MessageServiceHandlerMockDeleteParams{ctx, pp1},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up MessageServiceHandler.Delete return parameters for the expectation previously defined by the When method
func (e *MessageServiceHandlerMockDeleteExpectation) Then(pp2 *connect_go.Response[v1.DeleteResponse], err error) *MessageServiceHandlerMock {
	e.results = &MessageServiceHandlerMockDeleteResults{pp2, err}
	return e.mock
}

// Delete implements MessageServiceHandler
func (mmDelete *MessageServiceHandlerMock) Delete(ctx context.Context, pp1 *connect_go.Request[v1.DeleteRequest]) (pp2 *connect_go.Response[v1.DeleteResponse], err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ctx, pp1)
	}

	mm_params := &MessageServiceHandlerMockDeleteParams{ctx, pp1}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.pp2, e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := MessageServiceHandlerMockDeleteParams{ctx, pp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("MessageServiceHandlerMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the MessageServiceHandlerMock.Delete")
		}
		return (*mm_results).pp2, (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ctx, pp1)
	}
	mmDelete.t.Fatalf("Unexpected call to MessageServiceHandlerMock.Delete. %v %v", ctx, pp1)
	return
}

// DeleteAfterCounter returns a count of finished MessageServiceHandlerMock.Delete invocations
func (mmDelete *MessageServiceHandlerMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of MessageServiceHandlerMock.Delete invocations
func (mmDelete *MessageServiceHandlerMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to MessageServiceHandlerMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mMessageServiceHandlerMockDelete) Calls() []*MessageServiceHandlerMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*MessageServiceHandlerMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *MessageServiceHandlerMock) MinimockDeleteDone() bool {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteInspect logs each unmet expectation
func (m *MessageServiceHandlerMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessageServiceHandlerMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MessageServiceHandlerMock.Delete")
		} else {
			m.t.Errorf("Expected call to MessageServiceHandlerMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to MessageServiceHandlerMock.Delete")
	}
}

type mMessageServiceHandlerMockList struct {
	mock               *MessageServiceHandlerMock
	defaultExpectation *MessageServiceHandlerMockListExpectation
	expectations       []*MessageServiceHandlerMockListExpectation

	callArgs []*MessageServiceHandlerMockListParams
	mutex    sync.RWMutex
}

// MessageServiceHandlerMockListExpectation specifies expectation struct of the MessageServiceHandler.List
type MessageServiceHandlerMockListExpectation struct {
	mock    *MessageServiceHandlerMock
	params  *MessageServiceHandlerMockListParams
	results *MessageServiceHandlerMockListResults
	Counter uint64
}

// MessageServiceHandlerMockListParams contains parameters of the MessageServiceHandler.List
type MessageServiceHandlerMockListParams struct {
	ctx context.Context
	pp1 *connect_go.Request[v1.ListRequest]
}

// MessageServiceHandlerMockListResults contains results of the MessageServiceHandler.List
type MessageServiceHandlerMockListResults struct {
	pp2 *connect_go.Response[v1.ListResponse]
	err error
}

// Expect sets up expected params for MessageServiceHandler.List
func (mmList *mMessageServiceHandlerMockList) Expect(ctx context.Context, pp1 *connect_go.Request[v1.ListRequest]) *mMessageServiceHandlerMockList {
	if mmList.mock.funcList != nil {
		mmList.mock.t.Fatalf("MessageServiceHandlerMock.List mock is already set by Set")
	}

	if mmList.defaultExpectation == nil {
		mmList.defaultExpectation = &MessageServiceHandlerMockListExpectation{}
	}

	mmList.defaultExpectation.params = &MessageServiceHandlerMockListParams{ctx, pp1}
	for _, e := range mmList.expectations {
		if minimock.Equal(e.params, mmList.defaultExpectation.params) {
			mmList.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmList.defaultExpectation.params)
		}
	}

	return mmList
}

// Inspect accepts an inspector function that has same arguments as the MessageServiceHandler.List
func (mmList *mMessageServiceHandlerMockList) Inspect(f func(ctx context.Context, pp1 *connect_go.Request[v1.ListRequest])) *mMessageServiceHandlerMockList {
	if mmList.mock.inspectFuncList != nil {
		mmList.mock.t.Fatalf("Inspect function is already set for MessageServiceHandlerMock.List")
	}

	mmList.mock.inspectFuncList = f

	return mmList
}

// Return sets up results that will be returned by MessageServiceHandler.List
func (mmList *mMessageServiceHandlerMockList) Return(pp2 *connect_go.Response[v1.ListResponse], err error) *MessageServiceHandlerMock {
	if mmList.mock.funcList != nil {
		mmList.mock.t.Fatalf("MessageServiceHandlerMock.List mock is already set by Set")
	}

	if mmList.defaultExpectation == nil {
		mmList.defaultExpectation = &MessageServiceHandlerMockListExpectation{mock: mmList.mock}
	}
	mmList.defaultExpectation.results = &MessageServiceHandlerMockListResults{pp2, err}
	return mmList.mock
}

// Set uses given function f to mock the MessageServiceHandler.List method
func (mmList *mMessageServiceHandlerMockList) Set(f func(ctx context.Context, pp1 *connect_go.Request[v1.ListRequest]) (pp2 *connect_go.Response[v1.ListResponse], err error)) *MessageServiceHandlerMock {
	if mmList.defaultExpectation != nil {
		mmList.mock.t.Fatalf("Default expectation is already set for the MessageServiceHandler.List method")
	}

	if len(mmList.expectations) > 0 {
		mmList.mock.t.Fatalf("Some expectations are already set for the MessageServiceHandler.List method")
	}

	mmList.mock.funcList = f
	return mmList.mock
}

// When sets expectation for the MessageServiceHandler.List which will trigger the result defined by the following
// Then helper
func (mmList *mMessageServiceHandlerMockList) When(ctx context.Context, pp1 *connect_go.Request[v1.ListRequest]) *MessageServiceHandlerMockListExpectation {
	if mmList.mock.funcList != nil {
		mmList.mock.t.Fatalf("MessageServiceHandlerMock.List mock is already set by Set")
	}

	expectation := &MessageServiceHandlerMockListExpectation{
		mock:   mmList.mock,
		params: &MessageServiceHandlerMockListParams{ctx, pp1},
	}
	mmList.expectations = append(mmList.expectations, expectation)
	return expectation
}

// Then sets up MessageServiceHandler.List return parameters for the expectation previously defined by the When method
func (e *MessageServiceHandlerMockListExpectation) Then(pp2 *connect_go.Response[v1.ListResponse], err error) *MessageServiceHandlerMock {
	e.results = &MessageServiceHandlerMockListResults{pp2, err}
	return e.mock
}

// List implements MessageServiceHandler
func (mmList *MessageServiceHandlerMock) List(ctx context.Context, pp1 *connect_go.Request[v1.ListRequest]) (pp2 *connect_go.Response[v1.ListResponse], err error) {
	mm_atomic.AddUint64(&mmList.beforeListCounter, 1)
	defer mm_atomic.AddUint64(&mmList.afterListCounter, 1)

	if mmList.inspectFuncList != nil {
		mmList.inspectFuncList(ctx, pp1)
	}

	mm_params := &MessageServiceHandlerMockListParams{ctx, pp1}

	// Record call args
	mmList.ListMock.mutex.Lock()
	mmList.ListMock.callArgs = append(mmList.ListMock.callArgs, mm_params)
	mmList.ListMock.mutex.Unlock()

	for _, e := range mmList.ListMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.pp2, e.results.err
		}
	}

	if mmList.ListMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmList.ListMock.defaultExpectation.Counter, 1)
		mm_want := mmList.ListMock.defaultExpectation.params
		mm_got := MessageServiceHandlerMockListParams{ctx, pp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmList.t.Errorf("MessageServiceHandlerMock.List got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmList.ListMock.defaultExpectation.results
		if mm_results == nil {
			mmList.t.Fatal("No results are set for the MessageServiceHandlerMock.List")
		}
		return (*mm_results).pp2, (*mm_results).err
	}
	if mmList.funcList != nil {
		return mmList.funcList(ctx, pp1)
	}
	mmList.t.Fatalf("Unexpected call to MessageServiceHandlerMock.List. %v %v", ctx, pp1)
	return
}

// ListAfterCounter returns a count of finished MessageServiceHandlerMock.List invocations
func (mmList *MessageServiceHandlerMock) ListAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmList.afterListCounter)
}

// ListBeforeCounter returns a count of MessageServiceHandlerMock.List invocations
func (mmList *MessageServiceHandlerMock) ListBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmList.beforeListCounter)
}

// Calls returns a list of arguments used in each call to MessageServiceHandlerMock.List.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmList *mMessageServiceHandlerMockList) Calls() []*MessageServiceHandlerMockListParams {
	mmList.mutex.RLock()

	argCopy := make([]*MessageServiceHandlerMockListParams, len(mmList.callArgs))
	copy(argCopy, mmList.callArgs)

	mmList.mutex.RUnlock()

	return argCopy
}

// MinimockListDone returns true if the count of the List invocations corresponds
// the number of defined expectations
func (m *MessageServiceHandlerMock) MinimockListDone() bool {
	for _, e := range m.ListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcList != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		return false
	}
	return true
}

// MinimockListInspect logs each unmet expectation
func (m *MessageServiceHandlerMock) MinimockListInspect() {
	for _, e := range m.ListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessageServiceHandlerMock.List with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		if m.ListMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MessageServiceHandlerMock.List")
		} else {
			m.t.Errorf("Expected call to MessageServiceHandlerMock.List with params: %#v", *m.ListMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcList != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		m.t.Error("Expected call to MessageServiceHandlerMock.List")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MessageServiceHandlerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAddInspect()

		m.MinimockDeleteInspect()

		m.MinimockListInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MessageServiceHandlerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *MessageServiceHandlerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddDone() &&
		m.MinimockDeleteDone() &&
		m.MinimockListDone()
}
