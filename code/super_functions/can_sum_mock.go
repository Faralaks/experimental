package super_functions

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i experemental/code/super_functions.CanSum -o ./can_sum_mock.go -n CanSumMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// CanSumMock implements CanSum
type CanSumMock struct {
	t minimock.Tester

	funcSum          func(x int, y int) (i1 int)
	inspectFuncSum   func(x int, y int)
	afterSumCounter  uint64
	beforeSumCounter uint64
	SumMock          mCanSumMockSum
}

// NewCanSumMock returns a mock for CanSum
func NewCanSumMock(t minimock.Tester) *CanSumMock {
	m := &CanSumMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SumMock = mCanSumMockSum{mock: m}
	m.SumMock.callArgs = []*CanSumMockSumParams{}

	return m
}

type mCanSumMockSum struct {
	mock               *CanSumMock
	defaultExpectation *CanSumMockSumExpectation
	expectations       []*CanSumMockSumExpectation

	callArgs []*CanSumMockSumParams
	mutex    sync.RWMutex
}

// CanSumMockSumExpectation specifies expectation struct of the CanSum.Sum
type CanSumMockSumExpectation struct {
	mock    *CanSumMock
	params  *CanSumMockSumParams
	results *CanSumMockSumResults
	Counter uint64
}

// CanSumMockSumParams contains parameters of the CanSum.Sum
type CanSumMockSumParams struct {
	x int
	y int
}

// CanSumMockSumResults contains results of the CanSum.Sum
type CanSumMockSumResults struct {
	i1 int
}

// Expect sets up expected params for CanSum.Sum
func (mmSum *mCanSumMockSum) Expect(x int, y int) *mCanSumMockSum {
	if mmSum.mock.funcSum != nil {
		mmSum.mock.t.Fatalf("CanSumMock.Sum mock is already set by Set")
	}

	if mmSum.defaultExpectation == nil {
		mmSum.defaultExpectation = &CanSumMockSumExpectation{}
	}

	mmSum.defaultExpectation.params = &CanSumMockSumParams{x, y}
	for _, e := range mmSum.expectations {
		if minimock.Equal(e.params, mmSum.defaultExpectation.params) {
			mmSum.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSum.defaultExpectation.params)
		}
	}

	return mmSum
}

// Inspect accepts an inspector function that has same arguments as the CanSum.Sum
func (mmSum *mCanSumMockSum) Inspect(f func(x int, y int)) *mCanSumMockSum {
	if mmSum.mock.inspectFuncSum != nil {
		mmSum.mock.t.Fatalf("Inspect function is already set for CanSumMock.Sum")
	}

	mmSum.mock.inspectFuncSum = f

	return mmSum
}

// Return sets up results that will be returned by CanSum.Sum
func (mmSum *mCanSumMockSum) Return(i1 int) *CanSumMock {
	if mmSum.mock.funcSum != nil {
		mmSum.mock.t.Fatalf("CanSumMock.Sum mock is already set by Set")
	}

	if mmSum.defaultExpectation == nil {
		mmSum.defaultExpectation = &CanSumMockSumExpectation{mock: mmSum.mock}
	}
	mmSum.defaultExpectation.results = &CanSumMockSumResults{i1}
	return mmSum.mock
}

//Set uses given function f to mock the CanSum.Sum method
func (mmSum *mCanSumMockSum) Set(f func(x int, y int) (i1 int)) *CanSumMock {
	if mmSum.defaultExpectation != nil {
		mmSum.mock.t.Fatalf("Default expectation is already set for the CanSum.Sum method")
	}

	if len(mmSum.expectations) > 0 {
		mmSum.mock.t.Fatalf("Some expectations are already set for the CanSum.Sum method")
	}

	mmSum.mock.funcSum = f
	return mmSum.mock
}

// When sets expectation for the CanSum.Sum which will trigger the result defined by the following
// Then helper
func (mmSum *mCanSumMockSum) When(x int, y int) *CanSumMockSumExpectation {
	if mmSum.mock.funcSum != nil {
		mmSum.mock.t.Fatalf("CanSumMock.Sum mock is already set by Set")
	}

	expectation := &CanSumMockSumExpectation{
		mock:   mmSum.mock,
		params: &CanSumMockSumParams{x, y},
	}
	mmSum.expectations = append(mmSum.expectations, expectation)
	return expectation
}

// Then sets up CanSum.Sum return parameters for the expectation previously defined by the When method
func (e *CanSumMockSumExpectation) Then(i1 int) *CanSumMock {
	e.results = &CanSumMockSumResults{i1}
	return e.mock
}

// Sum implements CanSum
func (mmSum *CanSumMock) Sum(x int, y int) (i1 int) {
	mm_atomic.AddUint64(&mmSum.beforeSumCounter, 1)
	defer mm_atomic.AddUint64(&mmSum.afterSumCounter, 1)

	if mmSum.inspectFuncSum != nil {
		mmSum.inspectFuncSum(x, y)
	}

	mm_params := &CanSumMockSumParams{x, y}

	// Record call args
	mmSum.SumMock.mutex.Lock()
	mmSum.SumMock.callArgs = append(mmSum.SumMock.callArgs, mm_params)
	mmSum.SumMock.mutex.Unlock()

	for _, e := range mmSum.SumMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1
		}
	}

	if mmSum.SumMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSum.SumMock.defaultExpectation.Counter, 1)
		mm_want := mmSum.SumMock.defaultExpectation.params
		mm_got := CanSumMockSumParams{x, y}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSum.t.Errorf("CanSumMock.Sum got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSum.SumMock.defaultExpectation.results
		if mm_results == nil {
			mmSum.t.Fatal("No results are set for the CanSumMock.Sum")
		}
		return (*mm_results).i1
	}
	if mmSum.funcSum != nil {
		return mmSum.funcSum(x, y)
	}
	mmSum.t.Fatalf("Unexpected call to CanSumMock.Sum. %v %v", x, y)
	return
}

// SumAfterCounter returns a count of finished CanSumMock.Sum invocations
func (mmSum *CanSumMock) SumAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSum.afterSumCounter)
}

// SumBeforeCounter returns a count of CanSumMock.Sum invocations
func (mmSum *CanSumMock) SumBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSum.beforeSumCounter)
}

// Calls returns a list of arguments used in each call to CanSumMock.Sum.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSum *mCanSumMockSum) Calls() []*CanSumMockSumParams {
	mmSum.mutex.RLock()

	argCopy := make([]*CanSumMockSumParams, len(mmSum.callArgs))
	copy(argCopy, mmSum.callArgs)

	mmSum.mutex.RUnlock()

	return argCopy
}

// MinimockSumDone returns true if the count of the Sum invocations corresponds
// the number of defined expectations
func (m *CanSumMock) MinimockSumDone() bool {
	for _, e := range m.SumMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SumMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSumCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSum != nil && mm_atomic.LoadUint64(&m.afterSumCounter) < 1 {
		return false
	}
	return true
}

// MinimockSumInspect logs each unmet expectation
func (m *CanSumMock) MinimockSumInspect() {
	for _, e := range m.SumMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CanSumMock.Sum with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SumMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSumCounter) < 1 {
		if m.SumMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CanSumMock.Sum")
		} else {
			m.t.Errorf("Expected call to CanSumMock.Sum with params: %#v", *m.SumMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSum != nil && mm_atomic.LoadUint64(&m.afterSumCounter) < 1 {
		m.t.Error("Expected call to CanSumMock.Sum")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CanSumMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockSumInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CanSumMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *CanSumMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSumDone()
}
