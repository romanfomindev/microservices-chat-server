package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/romanfomindev/microservices-chat-server/internal/services.Implementation -o ./mocks/implementation_minimock.go -n ImplementationMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	desc "github.com/romanfomindev/microservices-chat-server/pkg/chat_api_v1"
)

// ImplementationMock implements services.Implementation
type ImplementationMock struct {
	t minimock.Tester

	funcConnectChat          func(chatId uint64, email string, stream desc.ChatApi_ConnectChatServer) (err error)
	inspectFuncConnectChat   func(chatId uint64, email string, stream desc.ChatApi_ConnectChatServer)
	afterConnectChatCounter  uint64
	beforeConnectChatCounter uint64
	ConnectChatMock          mImplementationMockConnectChat

	funcSendMessage          func(chatId uint64, message *desc.Message) (err error)
	inspectFuncSendMessage   func(chatId uint64, message *desc.Message)
	afterSendMessageCounter  uint64
	beforeSendMessageCounter uint64
	SendMessageMock          mImplementationMockSendMessage
}

// NewImplementationMock returns a mock for services.Implementation
func NewImplementationMock(t minimock.Tester) *ImplementationMock {
	m := &ImplementationMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ConnectChatMock = mImplementationMockConnectChat{mock: m}
	m.ConnectChatMock.callArgs = []*ImplementationMockConnectChatParams{}

	m.SendMessageMock = mImplementationMockSendMessage{mock: m}
	m.SendMessageMock.callArgs = []*ImplementationMockSendMessageParams{}

	return m
}

type mImplementationMockConnectChat struct {
	mock               *ImplementationMock
	defaultExpectation *ImplementationMockConnectChatExpectation
	expectations       []*ImplementationMockConnectChatExpectation

	callArgs []*ImplementationMockConnectChatParams
	mutex    sync.RWMutex
}

// ImplementationMockConnectChatExpectation specifies expectation struct of the Implementation.ConnectChat
type ImplementationMockConnectChatExpectation struct {
	mock    *ImplementationMock
	params  *ImplementationMockConnectChatParams
	results *ImplementationMockConnectChatResults
	Counter uint64
}

// ImplementationMockConnectChatParams contains parameters of the Implementation.ConnectChat
type ImplementationMockConnectChatParams struct {
	chatId uint64
	email  string
	stream desc.ChatApi_ConnectChatServer
}

// ImplementationMockConnectChatResults contains results of the Implementation.ConnectChat
type ImplementationMockConnectChatResults struct {
	err error
}

// Expect sets up expected params for Implementation.ConnectChat
func (mmConnectChat *mImplementationMockConnectChat) Expect(chatId uint64, email string, stream desc.ChatApi_ConnectChatServer) *mImplementationMockConnectChat {
	if mmConnectChat.mock.funcConnectChat != nil {
		mmConnectChat.mock.t.Fatalf("ImplementationMock.ConnectChat mock is already set by Set")
	}

	if mmConnectChat.defaultExpectation == nil {
		mmConnectChat.defaultExpectation = &ImplementationMockConnectChatExpectation{}
	}

	mmConnectChat.defaultExpectation.params = &ImplementationMockConnectChatParams{chatId, email, stream}
	for _, e := range mmConnectChat.expectations {
		if minimock.Equal(e.params, mmConnectChat.defaultExpectation.params) {
			mmConnectChat.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmConnectChat.defaultExpectation.params)
		}
	}

	return mmConnectChat
}

// Inspect accepts an inspector function that has same arguments as the Implementation.ConnectChat
func (mmConnectChat *mImplementationMockConnectChat) Inspect(f func(chatId uint64, email string, stream desc.ChatApi_ConnectChatServer)) *mImplementationMockConnectChat {
	if mmConnectChat.mock.inspectFuncConnectChat != nil {
		mmConnectChat.mock.t.Fatalf("Inspect function is already set for ImplementationMock.ConnectChat")
	}

	mmConnectChat.mock.inspectFuncConnectChat = f

	return mmConnectChat
}

// Return sets up results that will be returned by Implementation.ConnectChat
func (mmConnectChat *mImplementationMockConnectChat) Return(err error) *ImplementationMock {
	if mmConnectChat.mock.funcConnectChat != nil {
		mmConnectChat.mock.t.Fatalf("ImplementationMock.ConnectChat mock is already set by Set")
	}

	if mmConnectChat.defaultExpectation == nil {
		mmConnectChat.defaultExpectation = &ImplementationMockConnectChatExpectation{mock: mmConnectChat.mock}
	}
	mmConnectChat.defaultExpectation.results = &ImplementationMockConnectChatResults{err}
	return mmConnectChat.mock
}

// Set uses given function f to mock the Implementation.ConnectChat method
func (mmConnectChat *mImplementationMockConnectChat) Set(f func(chatId uint64, email string, stream desc.ChatApi_ConnectChatServer) (err error)) *ImplementationMock {
	if mmConnectChat.defaultExpectation != nil {
		mmConnectChat.mock.t.Fatalf("Default expectation is already set for the Implementation.ConnectChat method")
	}

	if len(mmConnectChat.expectations) > 0 {
		mmConnectChat.mock.t.Fatalf("Some expectations are already set for the Implementation.ConnectChat method")
	}

	mmConnectChat.mock.funcConnectChat = f
	return mmConnectChat.mock
}

// When sets expectation for the Implementation.ConnectChat which will trigger the result defined by the following
// Then helper
func (mmConnectChat *mImplementationMockConnectChat) When(chatId uint64, email string, stream desc.ChatApi_ConnectChatServer) *ImplementationMockConnectChatExpectation {
	if mmConnectChat.mock.funcConnectChat != nil {
		mmConnectChat.mock.t.Fatalf("ImplementationMock.ConnectChat mock is already set by Set")
	}

	expectation := &ImplementationMockConnectChatExpectation{
		mock:   mmConnectChat.mock,
		params: &ImplementationMockConnectChatParams{chatId, email, stream},
	}
	mmConnectChat.expectations = append(mmConnectChat.expectations, expectation)
	return expectation
}

// Then sets up Implementation.ConnectChat return parameters for the expectation previously defined by the When method
func (e *ImplementationMockConnectChatExpectation) Then(err error) *ImplementationMock {
	e.results = &ImplementationMockConnectChatResults{err}
	return e.mock
}

// ConnectChat implements services.Implementation
func (mmConnectChat *ImplementationMock) ConnectChat(chatId uint64, email string, stream desc.ChatApi_ConnectChatServer) (err error) {
	mm_atomic.AddUint64(&mmConnectChat.beforeConnectChatCounter, 1)
	defer mm_atomic.AddUint64(&mmConnectChat.afterConnectChatCounter, 1)

	if mmConnectChat.inspectFuncConnectChat != nil {
		mmConnectChat.inspectFuncConnectChat(chatId, email, stream)
	}

	mm_params := &ImplementationMockConnectChatParams{chatId, email, stream}

	// Record call args
	mmConnectChat.ConnectChatMock.mutex.Lock()
	mmConnectChat.ConnectChatMock.callArgs = append(mmConnectChat.ConnectChatMock.callArgs, mm_params)
	mmConnectChat.ConnectChatMock.mutex.Unlock()

	for _, e := range mmConnectChat.ConnectChatMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmConnectChat.ConnectChatMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmConnectChat.ConnectChatMock.defaultExpectation.Counter, 1)
		mm_want := mmConnectChat.ConnectChatMock.defaultExpectation.params
		mm_got := ImplementationMockConnectChatParams{chatId, email, stream}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmConnectChat.t.Errorf("ImplementationMock.ConnectChat got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmConnectChat.ConnectChatMock.defaultExpectation.results
		if mm_results == nil {
			mmConnectChat.t.Fatal("No results are set for the ImplementationMock.ConnectChat")
		}
		return (*mm_results).err
	}
	if mmConnectChat.funcConnectChat != nil {
		return mmConnectChat.funcConnectChat(chatId, email, stream)
	}
	mmConnectChat.t.Fatalf("Unexpected call to ImplementationMock.ConnectChat. %v %v %v", chatId, email, stream)
	return
}

// ConnectChatAfterCounter returns a count of finished ImplementationMock.ConnectChat invocations
func (mmConnectChat *ImplementationMock) ConnectChatAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmConnectChat.afterConnectChatCounter)
}

// ConnectChatBeforeCounter returns a count of ImplementationMock.ConnectChat invocations
func (mmConnectChat *ImplementationMock) ConnectChatBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmConnectChat.beforeConnectChatCounter)
}

// Calls returns a list of arguments used in each call to ImplementationMock.ConnectChat.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmConnectChat *mImplementationMockConnectChat) Calls() []*ImplementationMockConnectChatParams {
	mmConnectChat.mutex.RLock()

	argCopy := make([]*ImplementationMockConnectChatParams, len(mmConnectChat.callArgs))
	copy(argCopy, mmConnectChat.callArgs)

	mmConnectChat.mutex.RUnlock()

	return argCopy
}

// MinimockConnectChatDone returns true if the count of the ConnectChat invocations corresponds
// the number of defined expectations
func (m *ImplementationMock) MinimockConnectChatDone() bool {
	for _, e := range m.ConnectChatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ConnectChatMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterConnectChatCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcConnectChat != nil && mm_atomic.LoadUint64(&m.afterConnectChatCounter) < 1 {
		return false
	}
	return true
}

// MinimockConnectChatInspect logs each unmet expectation
func (m *ImplementationMock) MinimockConnectChatInspect() {
	for _, e := range m.ConnectChatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ImplementationMock.ConnectChat with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ConnectChatMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterConnectChatCounter) < 1 {
		if m.ConnectChatMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ImplementationMock.ConnectChat")
		} else {
			m.t.Errorf("Expected call to ImplementationMock.ConnectChat with params: %#v", *m.ConnectChatMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcConnectChat != nil && mm_atomic.LoadUint64(&m.afterConnectChatCounter) < 1 {
		m.t.Error("Expected call to ImplementationMock.ConnectChat")
	}
}

type mImplementationMockSendMessage struct {
	mock               *ImplementationMock
	defaultExpectation *ImplementationMockSendMessageExpectation
	expectations       []*ImplementationMockSendMessageExpectation

	callArgs []*ImplementationMockSendMessageParams
	mutex    sync.RWMutex
}

// ImplementationMockSendMessageExpectation specifies expectation struct of the Implementation.SendMessage
type ImplementationMockSendMessageExpectation struct {
	mock    *ImplementationMock
	params  *ImplementationMockSendMessageParams
	results *ImplementationMockSendMessageResults
	Counter uint64
}

// ImplementationMockSendMessageParams contains parameters of the Implementation.SendMessage
type ImplementationMockSendMessageParams struct {
	chatId  uint64
	message *desc.Message
}

// ImplementationMockSendMessageResults contains results of the Implementation.SendMessage
type ImplementationMockSendMessageResults struct {
	err error
}

// Expect sets up expected params for Implementation.SendMessage
func (mmSendMessage *mImplementationMockSendMessage) Expect(chatId uint64, message *desc.Message) *mImplementationMockSendMessage {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("ImplementationMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &ImplementationMockSendMessageExpectation{}
	}

	mmSendMessage.defaultExpectation.params = &ImplementationMockSendMessageParams{chatId, message}
	for _, e := range mmSendMessage.expectations {
		if minimock.Equal(e.params, mmSendMessage.defaultExpectation.params) {
			mmSendMessage.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendMessage.defaultExpectation.params)
		}
	}

	return mmSendMessage
}

// Inspect accepts an inspector function that has same arguments as the Implementation.SendMessage
func (mmSendMessage *mImplementationMockSendMessage) Inspect(f func(chatId uint64, message *desc.Message)) *mImplementationMockSendMessage {
	if mmSendMessage.mock.inspectFuncSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("Inspect function is already set for ImplementationMock.SendMessage")
	}

	mmSendMessage.mock.inspectFuncSendMessage = f

	return mmSendMessage
}

// Return sets up results that will be returned by Implementation.SendMessage
func (mmSendMessage *mImplementationMockSendMessage) Return(err error) *ImplementationMock {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("ImplementationMock.SendMessage mock is already set by Set")
	}

	if mmSendMessage.defaultExpectation == nil {
		mmSendMessage.defaultExpectation = &ImplementationMockSendMessageExpectation{mock: mmSendMessage.mock}
	}
	mmSendMessage.defaultExpectation.results = &ImplementationMockSendMessageResults{err}
	return mmSendMessage.mock
}

// Set uses given function f to mock the Implementation.SendMessage method
func (mmSendMessage *mImplementationMockSendMessage) Set(f func(chatId uint64, message *desc.Message) (err error)) *ImplementationMock {
	if mmSendMessage.defaultExpectation != nil {
		mmSendMessage.mock.t.Fatalf("Default expectation is already set for the Implementation.SendMessage method")
	}

	if len(mmSendMessage.expectations) > 0 {
		mmSendMessage.mock.t.Fatalf("Some expectations are already set for the Implementation.SendMessage method")
	}

	mmSendMessage.mock.funcSendMessage = f
	return mmSendMessage.mock
}

// When sets expectation for the Implementation.SendMessage which will trigger the result defined by the following
// Then helper
func (mmSendMessage *mImplementationMockSendMessage) When(chatId uint64, message *desc.Message) *ImplementationMockSendMessageExpectation {
	if mmSendMessage.mock.funcSendMessage != nil {
		mmSendMessage.mock.t.Fatalf("ImplementationMock.SendMessage mock is already set by Set")
	}

	expectation := &ImplementationMockSendMessageExpectation{
		mock:   mmSendMessage.mock,
		params: &ImplementationMockSendMessageParams{chatId, message},
	}
	mmSendMessage.expectations = append(mmSendMessage.expectations, expectation)
	return expectation
}

// Then sets up Implementation.SendMessage return parameters for the expectation previously defined by the When method
func (e *ImplementationMockSendMessageExpectation) Then(err error) *ImplementationMock {
	e.results = &ImplementationMockSendMessageResults{err}
	return e.mock
}

// SendMessage implements services.Implementation
func (mmSendMessage *ImplementationMock) SendMessage(chatId uint64, message *desc.Message) (err error) {
	mm_atomic.AddUint64(&mmSendMessage.beforeSendMessageCounter, 1)
	defer mm_atomic.AddUint64(&mmSendMessage.afterSendMessageCounter, 1)

	if mmSendMessage.inspectFuncSendMessage != nil {
		mmSendMessage.inspectFuncSendMessage(chatId, message)
	}

	mm_params := &ImplementationMockSendMessageParams{chatId, message}

	// Record call args
	mmSendMessage.SendMessageMock.mutex.Lock()
	mmSendMessage.SendMessageMock.callArgs = append(mmSendMessage.SendMessageMock.callArgs, mm_params)
	mmSendMessage.SendMessageMock.mutex.Unlock()

	for _, e := range mmSendMessage.SendMessageMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSendMessage.SendMessageMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendMessage.SendMessageMock.defaultExpectation.Counter, 1)
		mm_want := mmSendMessage.SendMessageMock.defaultExpectation.params
		mm_got := ImplementationMockSendMessageParams{chatId, message}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSendMessage.t.Errorf("ImplementationMock.SendMessage got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSendMessage.SendMessageMock.defaultExpectation.results
		if mm_results == nil {
			mmSendMessage.t.Fatal("No results are set for the ImplementationMock.SendMessage")
		}
		return (*mm_results).err
	}
	if mmSendMessage.funcSendMessage != nil {
		return mmSendMessage.funcSendMessage(chatId, message)
	}
	mmSendMessage.t.Fatalf("Unexpected call to ImplementationMock.SendMessage. %v %v", chatId, message)
	return
}

// SendMessageAfterCounter returns a count of finished ImplementationMock.SendMessage invocations
func (mmSendMessage *ImplementationMock) SendMessageAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendMessage.afterSendMessageCounter)
}

// SendMessageBeforeCounter returns a count of ImplementationMock.SendMessage invocations
func (mmSendMessage *ImplementationMock) SendMessageBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendMessage.beforeSendMessageCounter)
}

// Calls returns a list of arguments used in each call to ImplementationMock.SendMessage.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendMessage *mImplementationMockSendMessage) Calls() []*ImplementationMockSendMessageParams {
	mmSendMessage.mutex.RLock()

	argCopy := make([]*ImplementationMockSendMessageParams, len(mmSendMessage.callArgs))
	copy(argCopy, mmSendMessage.callArgs)

	mmSendMessage.mutex.RUnlock()

	return argCopy
}

// MinimockSendMessageDone returns true if the count of the SendMessage invocations corresponds
// the number of defined expectations
func (m *ImplementationMock) MinimockSendMessageDone() bool {
	for _, e := range m.SendMessageMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendMessageMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendMessageCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendMessage != nil && mm_atomic.LoadUint64(&m.afterSendMessageCounter) < 1 {
		return false
	}
	return true
}

// MinimockSendMessageInspect logs each unmet expectation
func (m *ImplementationMock) MinimockSendMessageInspect() {
	for _, e := range m.SendMessageMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ImplementationMock.SendMessage with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendMessageMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendMessageCounter) < 1 {
		if m.SendMessageMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ImplementationMock.SendMessage")
		} else {
			m.t.Errorf("Expected call to ImplementationMock.SendMessage with params: %#v", *m.SendMessageMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendMessage != nil && mm_atomic.LoadUint64(&m.afterSendMessageCounter) < 1 {
		m.t.Error("Expected call to ImplementationMock.SendMessage")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ImplementationMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockConnectChatInspect()

		m.MinimockSendMessageInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ImplementationMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ImplementationMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockConnectChatDone() &&
		m.MinimockSendMessageDone()
}