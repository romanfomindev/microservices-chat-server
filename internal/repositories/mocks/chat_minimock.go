package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/romanfomindev/microservices-chat-server/internal/repositories.Chat -o ./mocks/chat_minimock.go -n ChatMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ChatMock implements repositories.Chat
type ChatMock struct {
	t minimock.Tester

	funcCreate          func(ctx context.Context, name string) (u1 uint64, err error)
	inspectFuncCreate   func(ctx context.Context, name string)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mChatMockCreate

	funcDelete          func(ctx context.Context, id uint64) (err error)
	inspectFuncDelete   func(ctx context.Context, id uint64)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mChatMockDelete
}

// NewChatMock returns a mock for repositories.Chat
func NewChatMock(t minimock.Tester) *ChatMock {
	m := &ChatMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mChatMockCreate{mock: m}
	m.CreateMock.callArgs = []*ChatMockCreateParams{}

	m.DeleteMock = mChatMockDelete{mock: m}
	m.DeleteMock.callArgs = []*ChatMockDeleteParams{}

	return m
}

type mChatMockCreate struct {
	mock               *ChatMock
	defaultExpectation *ChatMockCreateExpectation
	expectations       []*ChatMockCreateExpectation

	callArgs []*ChatMockCreateParams
	mutex    sync.RWMutex
}

// ChatMockCreateExpectation specifies expectation struct of the Chat.Create
type ChatMockCreateExpectation struct {
	mock    *ChatMock
	params  *ChatMockCreateParams
	results *ChatMockCreateResults
	Counter uint64
}

// ChatMockCreateParams contains parameters of the Chat.Create
type ChatMockCreateParams struct {
	ctx  context.Context
	name string
}

// ChatMockCreateResults contains results of the Chat.Create
type ChatMockCreateResults struct {
	u1  uint64
	err error
}

// Expect sets up expected params for Chat.Create
func (mmCreate *mChatMockCreate) Expect(ctx context.Context, name string) *mChatMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &ChatMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &ChatMockCreateParams{ctx, name}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the Chat.Create
func (mmCreate *mChatMockCreate) Inspect(f func(ctx context.Context, name string)) *mChatMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for ChatMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by Chat.Create
func (mmCreate *mChatMockCreate) Return(u1 uint64, err error) *ChatMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &ChatMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &ChatMockCreateResults{u1, err}
	return mmCreate.mock
}

// Set uses given function f to mock the Chat.Create method
func (mmCreate *mChatMockCreate) Set(f func(ctx context.Context, name string) (u1 uint64, err error)) *ChatMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the Chat.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the Chat.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the Chat.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mChatMockCreate) When(ctx context.Context, name string) *ChatMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatMock.Create mock is already set by Set")
	}

	expectation := &ChatMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &ChatMockCreateParams{ctx, name},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up Chat.Create return parameters for the expectation previously defined by the When method
func (e *ChatMockCreateExpectation) Then(u1 uint64, err error) *ChatMock {
	e.results = &ChatMockCreateResults{u1, err}
	return e.mock
}

// Create implements repositories.Chat
func (mmCreate *ChatMock) Create(ctx context.Context, name string) (u1 uint64, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, name)
	}

	mm_params := &ChatMockCreateParams{ctx, name}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.u1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := ChatMockCreateParams{ctx, name}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("ChatMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the ChatMock.Create")
		}
		return (*mm_results).u1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, name)
	}
	mmCreate.t.Fatalf("Unexpected call to ChatMock.Create. %v %v", ctx, name)
	return
}

// CreateAfterCounter returns a count of finished ChatMock.Create invocations
func (mmCreate *ChatMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of ChatMock.Create invocations
func (mmCreate *ChatMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to ChatMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mChatMockCreate) Calls() []*ChatMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*ChatMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *ChatMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *ChatMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatMock.Create")
		} else {
			m.t.Errorf("Expected call to ChatMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to ChatMock.Create")
	}
}

type mChatMockDelete struct {
	mock               *ChatMock
	defaultExpectation *ChatMockDeleteExpectation
	expectations       []*ChatMockDeleteExpectation

	callArgs []*ChatMockDeleteParams
	mutex    sync.RWMutex
}

// ChatMockDeleteExpectation specifies expectation struct of the Chat.Delete
type ChatMockDeleteExpectation struct {
	mock    *ChatMock
	params  *ChatMockDeleteParams
	results *ChatMockDeleteResults
	Counter uint64
}

// ChatMockDeleteParams contains parameters of the Chat.Delete
type ChatMockDeleteParams struct {
	ctx context.Context
	id  uint64
}

// ChatMockDeleteResults contains results of the Chat.Delete
type ChatMockDeleteResults struct {
	err error
}

// Expect sets up expected params for Chat.Delete
func (mmDelete *mChatMockDelete) Expect(ctx context.Context, id uint64) *mChatMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &ChatMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &ChatMockDeleteParams{ctx, id}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the Chat.Delete
func (mmDelete *mChatMockDelete) Inspect(f func(ctx context.Context, id uint64)) *mChatMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for ChatMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by Chat.Delete
func (mmDelete *mChatMockDelete) Return(err error) *ChatMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &ChatMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &ChatMockDeleteResults{err}
	return mmDelete.mock
}

// Set uses given function f to mock the Chat.Delete method
func (mmDelete *mChatMockDelete) Set(f func(ctx context.Context, id uint64) (err error)) *ChatMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the Chat.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the Chat.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the Chat.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mChatMockDelete) When(ctx context.Context, id uint64) *ChatMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatMock.Delete mock is already set by Set")
	}

	expectation := &ChatMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &ChatMockDeleteParams{ctx, id},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up Chat.Delete return parameters for the expectation previously defined by the When method
func (e *ChatMockDeleteExpectation) Then(err error) *ChatMock {
	e.results = &ChatMockDeleteResults{err}
	return e.mock
}

// Delete implements repositories.Chat
func (mmDelete *ChatMock) Delete(ctx context.Context, id uint64) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ctx, id)
	}

	mm_params := &ChatMockDeleteParams{ctx, id}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := ChatMockDeleteParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("ChatMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the ChatMock.Delete")
		}
		return (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ctx, id)
	}
	mmDelete.t.Fatalf("Unexpected call to ChatMock.Delete. %v %v", ctx, id)
	return
}

// DeleteAfterCounter returns a count of finished ChatMock.Delete invocations
func (mmDelete *ChatMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of ChatMock.Delete invocations
func (mmDelete *ChatMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to ChatMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mChatMockDelete) Calls() []*ChatMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*ChatMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *ChatMock) MinimockDeleteDone() bool {
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
func (m *ChatMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatMock.Delete")
		} else {
			m.t.Errorf("Expected call to ChatMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to ChatMock.Delete")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ChatMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()

		m.MinimockDeleteInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ChatMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ChatMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockDeleteDone()
}