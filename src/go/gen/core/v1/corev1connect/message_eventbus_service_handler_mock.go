package corev1connect

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/koblas/grpc-todo/gen/core/v1/corev1connect.MessageEventbusServiceHandler -o ./message_eventbus_service_handler_mock.go -n MessageEventbusServiceHandlerMock

import (
	context "context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	connect_go "github.com/bufbuild/connect-go"
	"github.com/gojuno/minimock/v3"
	v11 "github.com/koblas/grpc-todo/gen/core/message/v1"
	v1 "github.com/koblas/grpc-todo/gen/core/v1"
)

// MessageEventbusServiceHandlerMock implements MessageEventbusServiceHandler
type MessageEventbusServiceHandlerMock struct {
	t minimock.Tester

	funcChange          func(ctx context.Context, pp1 *connect_go.Request[v11.MessageChangeEvent]) (pp2 *connect_go.Response[v1.MessageEventbusServiceChangeResponse], err error)
	inspectFuncChange   func(ctx context.Context, pp1 *connect_go.Request[v11.MessageChangeEvent])
	afterChangeCounter  uint64
	beforeChangeCounter uint64
	ChangeMock          mMessageEventbusServiceHandlerMockChange
}

// NewMessageEventbusServiceHandlerMock returns a mock for MessageEventbusServiceHandler
func NewMessageEventbusServiceHandlerMock(t minimock.Tester) *MessageEventbusServiceHandlerMock {
	m := &MessageEventbusServiceHandlerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ChangeMock = mMessageEventbusServiceHandlerMockChange{mock: m}
	m.ChangeMock.callArgs = []*MessageEventbusServiceHandlerMockChangeParams{}

	return m
}

type mMessageEventbusServiceHandlerMockChange struct {
	mock               *MessageEventbusServiceHandlerMock
	defaultExpectation *MessageEventbusServiceHandlerMockChangeExpectation
	expectations       []*MessageEventbusServiceHandlerMockChangeExpectation

	callArgs []*MessageEventbusServiceHandlerMockChangeParams
	mutex    sync.RWMutex
}

// MessageEventbusServiceHandlerMockChangeExpectation specifies expectation struct of the MessageEventbusServiceHandler.Change
type MessageEventbusServiceHandlerMockChangeExpectation struct {
	mock    *MessageEventbusServiceHandlerMock
	params  *MessageEventbusServiceHandlerMockChangeParams
	results *MessageEventbusServiceHandlerMockChangeResults
	Counter uint64
}

// MessageEventbusServiceHandlerMockChangeParams contains parameters of the MessageEventbusServiceHandler.Change
type MessageEventbusServiceHandlerMockChangeParams struct {
	ctx context.Context
	pp1 *connect_go.Request[v11.MessageChangeEvent]
}

// MessageEventbusServiceHandlerMockChangeResults contains results of the MessageEventbusServiceHandler.Change
type MessageEventbusServiceHandlerMockChangeResults struct {
	pp2 *connect_go.Response[v1.MessageEventbusServiceChangeResponse]
	err error
}

// Expect sets up expected params for MessageEventbusServiceHandler.Change
func (mmChange *mMessageEventbusServiceHandlerMockChange) Expect(ctx context.Context, pp1 *connect_go.Request[v11.MessageChangeEvent]) *mMessageEventbusServiceHandlerMockChange {
	if mmChange.mock.funcChange != nil {
		mmChange.mock.t.Fatalf("MessageEventbusServiceHandlerMock.Change mock is already set by Set")
	}

	if mmChange.defaultExpectation == nil {
		mmChange.defaultExpectation = &MessageEventbusServiceHandlerMockChangeExpectation{}
	}

	mmChange.defaultExpectation.params = &MessageEventbusServiceHandlerMockChangeParams{ctx, pp1}
	for _, e := range mmChange.expectations {
		if minimock.Equal(e.params, mmChange.defaultExpectation.params) {
			mmChange.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmChange.defaultExpectation.params)
		}
	}

	return mmChange
}

// Inspect accepts an inspector function that has same arguments as the MessageEventbusServiceHandler.Change
func (mmChange *mMessageEventbusServiceHandlerMockChange) Inspect(f func(ctx context.Context, pp1 *connect_go.Request[v11.MessageChangeEvent])) *mMessageEventbusServiceHandlerMockChange {
	if mmChange.mock.inspectFuncChange != nil {
		mmChange.mock.t.Fatalf("Inspect function is already set for MessageEventbusServiceHandlerMock.Change")
	}

	mmChange.mock.inspectFuncChange = f

	return mmChange
}

// Return sets up results that will be returned by MessageEventbusServiceHandler.Change
func (mmChange *mMessageEventbusServiceHandlerMockChange) Return(pp2 *connect_go.Response[v1.MessageEventbusServiceChangeResponse], err error) *MessageEventbusServiceHandlerMock {
	if mmChange.mock.funcChange != nil {
		mmChange.mock.t.Fatalf("MessageEventbusServiceHandlerMock.Change mock is already set by Set")
	}

	if mmChange.defaultExpectation == nil {
		mmChange.defaultExpectation = &MessageEventbusServiceHandlerMockChangeExpectation{mock: mmChange.mock}
	}
	mmChange.defaultExpectation.results = &MessageEventbusServiceHandlerMockChangeResults{pp2, err}
	return mmChange.mock
}

// Set uses given function f to mock the MessageEventbusServiceHandler.Change method
func (mmChange *mMessageEventbusServiceHandlerMockChange) Set(f func(ctx context.Context, pp1 *connect_go.Request[v11.MessageChangeEvent]) (pp2 *connect_go.Response[v1.MessageEventbusServiceChangeResponse], err error)) *MessageEventbusServiceHandlerMock {
	if mmChange.defaultExpectation != nil {
		mmChange.mock.t.Fatalf("Default expectation is already set for the MessageEventbusServiceHandler.Change method")
	}

	if len(mmChange.expectations) > 0 {
		mmChange.mock.t.Fatalf("Some expectations are already set for the MessageEventbusServiceHandler.Change method")
	}

	mmChange.mock.funcChange = f
	return mmChange.mock
}

// When sets expectation for the MessageEventbusServiceHandler.Change which will trigger the result defined by the following
// Then helper
func (mmChange *mMessageEventbusServiceHandlerMockChange) When(ctx context.Context, pp1 *connect_go.Request[v11.MessageChangeEvent]) *MessageEventbusServiceHandlerMockChangeExpectation {
	if mmChange.mock.funcChange != nil {
		mmChange.mock.t.Fatalf("MessageEventbusServiceHandlerMock.Change mock is already set by Set")
	}

	expectation := &MessageEventbusServiceHandlerMockChangeExpectation{
		mock:   mmChange.mock,
		params: &MessageEventbusServiceHandlerMockChangeParams{ctx, pp1},
	}
	mmChange.expectations = append(mmChange.expectations, expectation)
	return expectation
}

// Then sets up MessageEventbusServiceHandler.Change return parameters for the expectation previously defined by the When method
func (e *MessageEventbusServiceHandlerMockChangeExpectation) Then(pp2 *connect_go.Response[v1.MessageEventbusServiceChangeResponse], err error) *MessageEventbusServiceHandlerMock {
	e.results = &MessageEventbusServiceHandlerMockChangeResults{pp2, err}
	return e.mock
}

// Change implements MessageEventbusServiceHandler
func (mmChange *MessageEventbusServiceHandlerMock) Change(ctx context.Context, pp1 *connect_go.Request[v11.MessageChangeEvent]) (pp2 *connect_go.Response[v1.MessageEventbusServiceChangeResponse], err error) {
	mm_atomic.AddUint64(&mmChange.beforeChangeCounter, 1)
	defer mm_atomic.AddUint64(&mmChange.afterChangeCounter, 1)

	if mmChange.inspectFuncChange != nil {
		mmChange.inspectFuncChange(ctx, pp1)
	}

	mm_params := &MessageEventbusServiceHandlerMockChangeParams{ctx, pp1}

	// Record call args
	mmChange.ChangeMock.mutex.Lock()
	mmChange.ChangeMock.callArgs = append(mmChange.ChangeMock.callArgs, mm_params)
	mmChange.ChangeMock.mutex.Unlock()

	for _, e := range mmChange.ChangeMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.pp2, e.results.err
		}
	}

	if mmChange.ChangeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmChange.ChangeMock.defaultExpectation.Counter, 1)
		mm_want := mmChange.ChangeMock.defaultExpectation.params
		mm_got := MessageEventbusServiceHandlerMockChangeParams{ctx, pp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmChange.t.Errorf("MessageEventbusServiceHandlerMock.Change got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmChange.ChangeMock.defaultExpectation.results
		if mm_results == nil {
			mmChange.t.Fatal("No results are set for the MessageEventbusServiceHandlerMock.Change")
		}
		return (*mm_results).pp2, (*mm_results).err
	}
	if mmChange.funcChange != nil {
		return mmChange.funcChange(ctx, pp1)
	}
	mmChange.t.Fatalf("Unexpected call to MessageEventbusServiceHandlerMock.Change. %v %v", ctx, pp1)
	return
}

// ChangeAfterCounter returns a count of finished MessageEventbusServiceHandlerMock.Change invocations
func (mmChange *MessageEventbusServiceHandlerMock) ChangeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmChange.afterChangeCounter)
}

// ChangeBeforeCounter returns a count of MessageEventbusServiceHandlerMock.Change invocations
func (mmChange *MessageEventbusServiceHandlerMock) ChangeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmChange.beforeChangeCounter)
}

// Calls returns a list of arguments used in each call to MessageEventbusServiceHandlerMock.Change.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmChange *mMessageEventbusServiceHandlerMockChange) Calls() []*MessageEventbusServiceHandlerMockChangeParams {
	mmChange.mutex.RLock()

	argCopy := make([]*MessageEventbusServiceHandlerMockChangeParams, len(mmChange.callArgs))
	copy(argCopy, mmChange.callArgs)

	mmChange.mutex.RUnlock()

	return argCopy
}

// MinimockChangeDone returns true if the count of the Change invocations corresponds
// the number of defined expectations
func (m *MessageEventbusServiceHandlerMock) MinimockChangeDone() bool {
	for _, e := range m.ChangeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ChangeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterChangeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcChange != nil && mm_atomic.LoadUint64(&m.afterChangeCounter) < 1 {
		return false
	}
	return true
}

// MinimockChangeInspect logs each unmet expectation
func (m *MessageEventbusServiceHandlerMock) MinimockChangeInspect() {
	for _, e := range m.ChangeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessageEventbusServiceHandlerMock.Change with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ChangeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterChangeCounter) < 1 {
		if m.ChangeMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MessageEventbusServiceHandlerMock.Change")
		} else {
			m.t.Errorf("Expected call to MessageEventbusServiceHandlerMock.Change with params: %#v", *m.ChangeMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcChange != nil && mm_atomic.LoadUint64(&m.afterChangeCounter) < 1 {
		m.t.Error("Expected call to MessageEventbusServiceHandlerMock.Change")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MessageEventbusServiceHandlerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockChangeInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MessageEventbusServiceHandlerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *MessageEventbusServiceHandlerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockChangeDone()
}
