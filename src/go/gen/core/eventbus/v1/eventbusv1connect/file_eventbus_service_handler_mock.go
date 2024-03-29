package eventbusv1connect

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect.FileEventbusServiceHandler -o ./file_eventbus_service_handler_mock.go -n FileEventbusServiceHandlerMock

import (
	context "context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	connect_go "github.com/bufbuild/connect-go"
	"github.com/gojuno/minimock/v3"
	v11 "github.com/koblas/grpc-todo/gen/core/eventbus/v1"
	v14 "github.com/koblas/grpc-todo/gen/core/file/v1"
)

// FileEventbusServiceHandlerMock implements FileEventbusServiceHandler
type FileEventbusServiceHandlerMock struct {
	t minimock.Tester

	funcFileComplete          func(ctx context.Context, pp1 *connect_go.Request[v14.FileServiceCompleteEvent]) (pp2 *connect_go.Response[v11.FileEventbusFileCompleteResponse], err error)
	inspectFuncFileComplete   func(ctx context.Context, pp1 *connect_go.Request[v14.FileServiceCompleteEvent])
	afterFileCompleteCounter  uint64
	beforeFileCompleteCounter uint64
	FileCompleteMock          mFileEventbusServiceHandlerMockFileComplete

	funcFileUploaded          func(ctx context.Context, pp1 *connect_go.Request[v14.FileServiceUploadEvent]) (pp2 *connect_go.Response[v11.FileEventbusFileUploadedResponse], err error)
	inspectFuncFileUploaded   func(ctx context.Context, pp1 *connect_go.Request[v14.FileServiceUploadEvent])
	afterFileUploadedCounter  uint64
	beforeFileUploadedCounter uint64
	FileUploadedMock          mFileEventbusServiceHandlerMockFileUploaded
}

// NewFileEventbusServiceHandlerMock returns a mock for FileEventbusServiceHandler
func NewFileEventbusServiceHandlerMock(t minimock.Tester) *FileEventbusServiceHandlerMock {
	m := &FileEventbusServiceHandlerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.FileCompleteMock = mFileEventbusServiceHandlerMockFileComplete{mock: m}
	m.FileCompleteMock.callArgs = []*FileEventbusServiceHandlerMockFileCompleteParams{}

	m.FileUploadedMock = mFileEventbusServiceHandlerMockFileUploaded{mock: m}
	m.FileUploadedMock.callArgs = []*FileEventbusServiceHandlerMockFileUploadedParams{}

	return m
}

type mFileEventbusServiceHandlerMockFileComplete struct {
	mock               *FileEventbusServiceHandlerMock
	defaultExpectation *FileEventbusServiceHandlerMockFileCompleteExpectation
	expectations       []*FileEventbusServiceHandlerMockFileCompleteExpectation

	callArgs []*FileEventbusServiceHandlerMockFileCompleteParams
	mutex    sync.RWMutex
}

// FileEventbusServiceHandlerMockFileCompleteExpectation specifies expectation struct of the FileEventbusServiceHandler.FileComplete
type FileEventbusServiceHandlerMockFileCompleteExpectation struct {
	mock    *FileEventbusServiceHandlerMock
	params  *FileEventbusServiceHandlerMockFileCompleteParams
	results *FileEventbusServiceHandlerMockFileCompleteResults
	Counter uint64
}

// FileEventbusServiceHandlerMockFileCompleteParams contains parameters of the FileEventbusServiceHandler.FileComplete
type FileEventbusServiceHandlerMockFileCompleteParams struct {
	ctx context.Context
	pp1 *connect_go.Request[v14.FileServiceCompleteEvent]
}

// FileEventbusServiceHandlerMockFileCompleteResults contains results of the FileEventbusServiceHandler.FileComplete
type FileEventbusServiceHandlerMockFileCompleteResults struct {
	pp2 *connect_go.Response[v11.FileEventbusFileCompleteResponse]
	err error
}

// Expect sets up expected params for FileEventbusServiceHandler.FileComplete
func (mmFileComplete *mFileEventbusServiceHandlerMockFileComplete) Expect(ctx context.Context, pp1 *connect_go.Request[v14.FileServiceCompleteEvent]) *mFileEventbusServiceHandlerMockFileComplete {
	if mmFileComplete.mock.funcFileComplete != nil {
		mmFileComplete.mock.t.Fatalf("FileEventbusServiceHandlerMock.FileComplete mock is already set by Set")
	}

	if mmFileComplete.defaultExpectation == nil {
		mmFileComplete.defaultExpectation = &FileEventbusServiceHandlerMockFileCompleteExpectation{}
	}

	mmFileComplete.defaultExpectation.params = &FileEventbusServiceHandlerMockFileCompleteParams{ctx, pp1}
	for _, e := range mmFileComplete.expectations {
		if minimock.Equal(e.params, mmFileComplete.defaultExpectation.params) {
			mmFileComplete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmFileComplete.defaultExpectation.params)
		}
	}

	return mmFileComplete
}

// Inspect accepts an inspector function that has same arguments as the FileEventbusServiceHandler.FileComplete
func (mmFileComplete *mFileEventbusServiceHandlerMockFileComplete) Inspect(f func(ctx context.Context, pp1 *connect_go.Request[v14.FileServiceCompleteEvent])) *mFileEventbusServiceHandlerMockFileComplete {
	if mmFileComplete.mock.inspectFuncFileComplete != nil {
		mmFileComplete.mock.t.Fatalf("Inspect function is already set for FileEventbusServiceHandlerMock.FileComplete")
	}

	mmFileComplete.mock.inspectFuncFileComplete = f

	return mmFileComplete
}

// Return sets up results that will be returned by FileEventbusServiceHandler.FileComplete
func (mmFileComplete *mFileEventbusServiceHandlerMockFileComplete) Return(pp2 *connect_go.Response[v11.FileEventbusFileCompleteResponse], err error) *FileEventbusServiceHandlerMock {
	if mmFileComplete.mock.funcFileComplete != nil {
		mmFileComplete.mock.t.Fatalf("FileEventbusServiceHandlerMock.FileComplete mock is already set by Set")
	}

	if mmFileComplete.defaultExpectation == nil {
		mmFileComplete.defaultExpectation = &FileEventbusServiceHandlerMockFileCompleteExpectation{mock: mmFileComplete.mock}
	}
	mmFileComplete.defaultExpectation.results = &FileEventbusServiceHandlerMockFileCompleteResults{pp2, err}
	return mmFileComplete.mock
}

// Set uses given function f to mock the FileEventbusServiceHandler.FileComplete method
func (mmFileComplete *mFileEventbusServiceHandlerMockFileComplete) Set(f func(ctx context.Context, pp1 *connect_go.Request[v14.FileServiceCompleteEvent]) (pp2 *connect_go.Response[v11.FileEventbusFileCompleteResponse], err error)) *FileEventbusServiceHandlerMock {
	if mmFileComplete.defaultExpectation != nil {
		mmFileComplete.mock.t.Fatalf("Default expectation is already set for the FileEventbusServiceHandler.FileComplete method")
	}

	if len(mmFileComplete.expectations) > 0 {
		mmFileComplete.mock.t.Fatalf("Some expectations are already set for the FileEventbusServiceHandler.FileComplete method")
	}

	mmFileComplete.mock.funcFileComplete = f
	return mmFileComplete.mock
}

// When sets expectation for the FileEventbusServiceHandler.FileComplete which will trigger the result defined by the following
// Then helper
func (mmFileComplete *mFileEventbusServiceHandlerMockFileComplete) When(ctx context.Context, pp1 *connect_go.Request[v14.FileServiceCompleteEvent]) *FileEventbusServiceHandlerMockFileCompleteExpectation {
	if mmFileComplete.mock.funcFileComplete != nil {
		mmFileComplete.mock.t.Fatalf("FileEventbusServiceHandlerMock.FileComplete mock is already set by Set")
	}

	expectation := &FileEventbusServiceHandlerMockFileCompleteExpectation{
		mock:   mmFileComplete.mock,
		params: &FileEventbusServiceHandlerMockFileCompleteParams{ctx, pp1},
	}
	mmFileComplete.expectations = append(mmFileComplete.expectations, expectation)
	return expectation
}

// Then sets up FileEventbusServiceHandler.FileComplete return parameters for the expectation previously defined by the When method
func (e *FileEventbusServiceHandlerMockFileCompleteExpectation) Then(pp2 *connect_go.Response[v11.FileEventbusFileCompleteResponse], err error) *FileEventbusServiceHandlerMock {
	e.results = &FileEventbusServiceHandlerMockFileCompleteResults{pp2, err}
	return e.mock
}

// FileComplete implements FileEventbusServiceHandler
func (mmFileComplete *FileEventbusServiceHandlerMock) FileComplete(ctx context.Context, pp1 *connect_go.Request[v14.FileServiceCompleteEvent]) (pp2 *connect_go.Response[v11.FileEventbusFileCompleteResponse], err error) {
	mm_atomic.AddUint64(&mmFileComplete.beforeFileCompleteCounter, 1)
	defer mm_atomic.AddUint64(&mmFileComplete.afterFileCompleteCounter, 1)

	if mmFileComplete.inspectFuncFileComplete != nil {
		mmFileComplete.inspectFuncFileComplete(ctx, pp1)
	}

	mm_params := &FileEventbusServiceHandlerMockFileCompleteParams{ctx, pp1}

	// Record call args
	mmFileComplete.FileCompleteMock.mutex.Lock()
	mmFileComplete.FileCompleteMock.callArgs = append(mmFileComplete.FileCompleteMock.callArgs, mm_params)
	mmFileComplete.FileCompleteMock.mutex.Unlock()

	for _, e := range mmFileComplete.FileCompleteMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.pp2, e.results.err
		}
	}

	if mmFileComplete.FileCompleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmFileComplete.FileCompleteMock.defaultExpectation.Counter, 1)
		mm_want := mmFileComplete.FileCompleteMock.defaultExpectation.params
		mm_got := FileEventbusServiceHandlerMockFileCompleteParams{ctx, pp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmFileComplete.t.Errorf("FileEventbusServiceHandlerMock.FileComplete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmFileComplete.FileCompleteMock.defaultExpectation.results
		if mm_results == nil {
			mmFileComplete.t.Fatal("No results are set for the FileEventbusServiceHandlerMock.FileComplete")
		}
		return (*mm_results).pp2, (*mm_results).err
	}
	if mmFileComplete.funcFileComplete != nil {
		return mmFileComplete.funcFileComplete(ctx, pp1)
	}
	mmFileComplete.t.Fatalf("Unexpected call to FileEventbusServiceHandlerMock.FileComplete. %v %v", ctx, pp1)
	return
}

// FileCompleteAfterCounter returns a count of finished FileEventbusServiceHandlerMock.FileComplete invocations
func (mmFileComplete *FileEventbusServiceHandlerMock) FileCompleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFileComplete.afterFileCompleteCounter)
}

// FileCompleteBeforeCounter returns a count of FileEventbusServiceHandlerMock.FileComplete invocations
func (mmFileComplete *FileEventbusServiceHandlerMock) FileCompleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFileComplete.beforeFileCompleteCounter)
}

// Calls returns a list of arguments used in each call to FileEventbusServiceHandlerMock.FileComplete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmFileComplete *mFileEventbusServiceHandlerMockFileComplete) Calls() []*FileEventbusServiceHandlerMockFileCompleteParams {
	mmFileComplete.mutex.RLock()

	argCopy := make([]*FileEventbusServiceHandlerMockFileCompleteParams, len(mmFileComplete.callArgs))
	copy(argCopy, mmFileComplete.callArgs)

	mmFileComplete.mutex.RUnlock()

	return argCopy
}

// MinimockFileCompleteDone returns true if the count of the FileComplete invocations corresponds
// the number of defined expectations
func (m *FileEventbusServiceHandlerMock) MinimockFileCompleteDone() bool {
	for _, e := range m.FileCompleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FileCompleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFileCompleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFileComplete != nil && mm_atomic.LoadUint64(&m.afterFileCompleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockFileCompleteInspect logs each unmet expectation
func (m *FileEventbusServiceHandlerMock) MinimockFileCompleteInspect() {
	for _, e := range m.FileCompleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to FileEventbusServiceHandlerMock.FileComplete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FileCompleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFileCompleteCounter) < 1 {
		if m.FileCompleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to FileEventbusServiceHandlerMock.FileComplete")
		} else {
			m.t.Errorf("Expected call to FileEventbusServiceHandlerMock.FileComplete with params: %#v", *m.FileCompleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFileComplete != nil && mm_atomic.LoadUint64(&m.afterFileCompleteCounter) < 1 {
		m.t.Error("Expected call to FileEventbusServiceHandlerMock.FileComplete")
	}
}

type mFileEventbusServiceHandlerMockFileUploaded struct {
	mock               *FileEventbusServiceHandlerMock
	defaultExpectation *FileEventbusServiceHandlerMockFileUploadedExpectation
	expectations       []*FileEventbusServiceHandlerMockFileUploadedExpectation

	callArgs []*FileEventbusServiceHandlerMockFileUploadedParams
	mutex    sync.RWMutex
}

// FileEventbusServiceHandlerMockFileUploadedExpectation specifies expectation struct of the FileEventbusServiceHandler.FileUploaded
type FileEventbusServiceHandlerMockFileUploadedExpectation struct {
	mock    *FileEventbusServiceHandlerMock
	params  *FileEventbusServiceHandlerMockFileUploadedParams
	results *FileEventbusServiceHandlerMockFileUploadedResults
	Counter uint64
}

// FileEventbusServiceHandlerMockFileUploadedParams contains parameters of the FileEventbusServiceHandler.FileUploaded
type FileEventbusServiceHandlerMockFileUploadedParams struct {
	ctx context.Context
	pp1 *connect_go.Request[v14.FileServiceUploadEvent]
}

// FileEventbusServiceHandlerMockFileUploadedResults contains results of the FileEventbusServiceHandler.FileUploaded
type FileEventbusServiceHandlerMockFileUploadedResults struct {
	pp2 *connect_go.Response[v11.FileEventbusFileUploadedResponse]
	err error
}

// Expect sets up expected params for FileEventbusServiceHandler.FileUploaded
func (mmFileUploaded *mFileEventbusServiceHandlerMockFileUploaded) Expect(ctx context.Context, pp1 *connect_go.Request[v14.FileServiceUploadEvent]) *mFileEventbusServiceHandlerMockFileUploaded {
	if mmFileUploaded.mock.funcFileUploaded != nil {
		mmFileUploaded.mock.t.Fatalf("FileEventbusServiceHandlerMock.FileUploaded mock is already set by Set")
	}

	if mmFileUploaded.defaultExpectation == nil {
		mmFileUploaded.defaultExpectation = &FileEventbusServiceHandlerMockFileUploadedExpectation{}
	}

	mmFileUploaded.defaultExpectation.params = &FileEventbusServiceHandlerMockFileUploadedParams{ctx, pp1}
	for _, e := range mmFileUploaded.expectations {
		if minimock.Equal(e.params, mmFileUploaded.defaultExpectation.params) {
			mmFileUploaded.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmFileUploaded.defaultExpectation.params)
		}
	}

	return mmFileUploaded
}

// Inspect accepts an inspector function that has same arguments as the FileEventbusServiceHandler.FileUploaded
func (mmFileUploaded *mFileEventbusServiceHandlerMockFileUploaded) Inspect(f func(ctx context.Context, pp1 *connect_go.Request[v14.FileServiceUploadEvent])) *mFileEventbusServiceHandlerMockFileUploaded {
	if mmFileUploaded.mock.inspectFuncFileUploaded != nil {
		mmFileUploaded.mock.t.Fatalf("Inspect function is already set for FileEventbusServiceHandlerMock.FileUploaded")
	}

	mmFileUploaded.mock.inspectFuncFileUploaded = f

	return mmFileUploaded
}

// Return sets up results that will be returned by FileEventbusServiceHandler.FileUploaded
func (mmFileUploaded *mFileEventbusServiceHandlerMockFileUploaded) Return(pp2 *connect_go.Response[v11.FileEventbusFileUploadedResponse], err error) *FileEventbusServiceHandlerMock {
	if mmFileUploaded.mock.funcFileUploaded != nil {
		mmFileUploaded.mock.t.Fatalf("FileEventbusServiceHandlerMock.FileUploaded mock is already set by Set")
	}

	if mmFileUploaded.defaultExpectation == nil {
		mmFileUploaded.defaultExpectation = &FileEventbusServiceHandlerMockFileUploadedExpectation{mock: mmFileUploaded.mock}
	}
	mmFileUploaded.defaultExpectation.results = &FileEventbusServiceHandlerMockFileUploadedResults{pp2, err}
	return mmFileUploaded.mock
}

// Set uses given function f to mock the FileEventbusServiceHandler.FileUploaded method
func (mmFileUploaded *mFileEventbusServiceHandlerMockFileUploaded) Set(f func(ctx context.Context, pp1 *connect_go.Request[v14.FileServiceUploadEvent]) (pp2 *connect_go.Response[v11.FileEventbusFileUploadedResponse], err error)) *FileEventbusServiceHandlerMock {
	if mmFileUploaded.defaultExpectation != nil {
		mmFileUploaded.mock.t.Fatalf("Default expectation is already set for the FileEventbusServiceHandler.FileUploaded method")
	}

	if len(mmFileUploaded.expectations) > 0 {
		mmFileUploaded.mock.t.Fatalf("Some expectations are already set for the FileEventbusServiceHandler.FileUploaded method")
	}

	mmFileUploaded.mock.funcFileUploaded = f
	return mmFileUploaded.mock
}

// When sets expectation for the FileEventbusServiceHandler.FileUploaded which will trigger the result defined by the following
// Then helper
func (mmFileUploaded *mFileEventbusServiceHandlerMockFileUploaded) When(ctx context.Context, pp1 *connect_go.Request[v14.FileServiceUploadEvent]) *FileEventbusServiceHandlerMockFileUploadedExpectation {
	if mmFileUploaded.mock.funcFileUploaded != nil {
		mmFileUploaded.mock.t.Fatalf("FileEventbusServiceHandlerMock.FileUploaded mock is already set by Set")
	}

	expectation := &FileEventbusServiceHandlerMockFileUploadedExpectation{
		mock:   mmFileUploaded.mock,
		params: &FileEventbusServiceHandlerMockFileUploadedParams{ctx, pp1},
	}
	mmFileUploaded.expectations = append(mmFileUploaded.expectations, expectation)
	return expectation
}

// Then sets up FileEventbusServiceHandler.FileUploaded return parameters for the expectation previously defined by the When method
func (e *FileEventbusServiceHandlerMockFileUploadedExpectation) Then(pp2 *connect_go.Response[v11.FileEventbusFileUploadedResponse], err error) *FileEventbusServiceHandlerMock {
	e.results = &FileEventbusServiceHandlerMockFileUploadedResults{pp2, err}
	return e.mock
}

// FileUploaded implements FileEventbusServiceHandler
func (mmFileUploaded *FileEventbusServiceHandlerMock) FileUploaded(ctx context.Context, pp1 *connect_go.Request[v14.FileServiceUploadEvent]) (pp2 *connect_go.Response[v11.FileEventbusFileUploadedResponse], err error) {
	mm_atomic.AddUint64(&mmFileUploaded.beforeFileUploadedCounter, 1)
	defer mm_atomic.AddUint64(&mmFileUploaded.afterFileUploadedCounter, 1)

	if mmFileUploaded.inspectFuncFileUploaded != nil {
		mmFileUploaded.inspectFuncFileUploaded(ctx, pp1)
	}

	mm_params := &FileEventbusServiceHandlerMockFileUploadedParams{ctx, pp1}

	// Record call args
	mmFileUploaded.FileUploadedMock.mutex.Lock()
	mmFileUploaded.FileUploadedMock.callArgs = append(mmFileUploaded.FileUploadedMock.callArgs, mm_params)
	mmFileUploaded.FileUploadedMock.mutex.Unlock()

	for _, e := range mmFileUploaded.FileUploadedMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.pp2, e.results.err
		}
	}

	if mmFileUploaded.FileUploadedMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmFileUploaded.FileUploadedMock.defaultExpectation.Counter, 1)
		mm_want := mmFileUploaded.FileUploadedMock.defaultExpectation.params
		mm_got := FileEventbusServiceHandlerMockFileUploadedParams{ctx, pp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmFileUploaded.t.Errorf("FileEventbusServiceHandlerMock.FileUploaded got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmFileUploaded.FileUploadedMock.defaultExpectation.results
		if mm_results == nil {
			mmFileUploaded.t.Fatal("No results are set for the FileEventbusServiceHandlerMock.FileUploaded")
		}
		return (*mm_results).pp2, (*mm_results).err
	}
	if mmFileUploaded.funcFileUploaded != nil {
		return mmFileUploaded.funcFileUploaded(ctx, pp1)
	}
	mmFileUploaded.t.Fatalf("Unexpected call to FileEventbusServiceHandlerMock.FileUploaded. %v %v", ctx, pp1)
	return
}

// FileUploadedAfterCounter returns a count of finished FileEventbusServiceHandlerMock.FileUploaded invocations
func (mmFileUploaded *FileEventbusServiceHandlerMock) FileUploadedAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFileUploaded.afterFileUploadedCounter)
}

// FileUploadedBeforeCounter returns a count of FileEventbusServiceHandlerMock.FileUploaded invocations
func (mmFileUploaded *FileEventbusServiceHandlerMock) FileUploadedBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFileUploaded.beforeFileUploadedCounter)
}

// Calls returns a list of arguments used in each call to FileEventbusServiceHandlerMock.FileUploaded.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmFileUploaded *mFileEventbusServiceHandlerMockFileUploaded) Calls() []*FileEventbusServiceHandlerMockFileUploadedParams {
	mmFileUploaded.mutex.RLock()

	argCopy := make([]*FileEventbusServiceHandlerMockFileUploadedParams, len(mmFileUploaded.callArgs))
	copy(argCopy, mmFileUploaded.callArgs)

	mmFileUploaded.mutex.RUnlock()

	return argCopy
}

// MinimockFileUploadedDone returns true if the count of the FileUploaded invocations corresponds
// the number of defined expectations
func (m *FileEventbusServiceHandlerMock) MinimockFileUploadedDone() bool {
	for _, e := range m.FileUploadedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FileUploadedMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFileUploadedCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFileUploaded != nil && mm_atomic.LoadUint64(&m.afterFileUploadedCounter) < 1 {
		return false
	}
	return true
}

// MinimockFileUploadedInspect logs each unmet expectation
func (m *FileEventbusServiceHandlerMock) MinimockFileUploadedInspect() {
	for _, e := range m.FileUploadedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to FileEventbusServiceHandlerMock.FileUploaded with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FileUploadedMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFileUploadedCounter) < 1 {
		if m.FileUploadedMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to FileEventbusServiceHandlerMock.FileUploaded")
		} else {
			m.t.Errorf("Expected call to FileEventbusServiceHandlerMock.FileUploaded with params: %#v", *m.FileUploadedMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFileUploaded != nil && mm_atomic.LoadUint64(&m.afterFileUploadedCounter) < 1 {
		m.t.Error("Expected call to FileEventbusServiceHandlerMock.FileUploaded")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *FileEventbusServiceHandlerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockFileCompleteInspect()

		m.MinimockFileUploadedInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *FileEventbusServiceHandlerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *FileEventbusServiceHandlerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockFileCompleteDone() &&
		m.MinimockFileUploadedDone()
}
