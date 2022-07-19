// Code generated by go-mockgen 1.3.3; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package policies

import (
	"context"
	"sync"

	store "github.com/sourcegraph/sourcegraph/internal/codeintel/policies/internal/store"
	shared "github.com/sourcegraph/sourcegraph/internal/codeintel/policies/shared"
)

// MockStore is a mock implementation of the Store interface (from the
// package
// github.com/sourcegraph/sourcegraph/internal/codeintel/policies/internal/store)
// used for unit testing.
type MockStore struct {
	// GetConfigurationPoliciesFunc is an instance of a mock function object
	// controlling the behavior of the method GetConfigurationPolicies.
	GetConfigurationPoliciesFunc *StoreGetConfigurationPoliciesFunc
	// ListFunc is an instance of a mock function object controlling the
	// behavior of the method List.
	ListFunc *StoreListFunc
	// UpdateReposMatchingPatternsFunc is an instance of a mock function
	// object controlling the behavior of the method
	// UpdateReposMatchingPatterns.
	UpdateReposMatchingPatternsFunc *StoreUpdateReposMatchingPatternsFunc
}

// NewMockStore creates a new mock of the Store interface. All methods
// return zero values for all results, unless overwritten.
func NewMockStore() *MockStore {
	return &MockStore{
		GetConfigurationPoliciesFunc: &StoreGetConfigurationPoliciesFunc{
			defaultHook: func(context.Context, shared.GetConfigurationPoliciesOptions) (r0 []shared.ConfigurationPolicy, r1 int, r2 error) {
				return
			},
		},
		ListFunc: &StoreListFunc{
			defaultHook: func(context.Context, store.ListOpts) (r0 []shared.Policy, r1 error) {
				return
			},
		},
		UpdateReposMatchingPatternsFunc: &StoreUpdateReposMatchingPatternsFunc{
			defaultHook: func(context.Context, []string, int, *int) (r0 error) {
				return
			},
		},
	}
}

// NewStrictMockStore creates a new mock of the Store interface. All methods
// panic on invocation, unless overwritten.
func NewStrictMockStore() *MockStore {
	return &MockStore{
		GetConfigurationPoliciesFunc: &StoreGetConfigurationPoliciesFunc{
			defaultHook: func(context.Context, shared.GetConfigurationPoliciesOptions) ([]shared.ConfigurationPolicy, int, error) {
				panic("unexpected invocation of MockStore.GetConfigurationPolicies")
			},
		},
		ListFunc: &StoreListFunc{
			defaultHook: func(context.Context, store.ListOpts) ([]shared.Policy, error) {
				panic("unexpected invocation of MockStore.List")
			},
		},
		UpdateReposMatchingPatternsFunc: &StoreUpdateReposMatchingPatternsFunc{
			defaultHook: func(context.Context, []string, int, *int) error {
				panic("unexpected invocation of MockStore.UpdateReposMatchingPatterns")
			},
		},
	}
}

// NewMockStoreFrom creates a new mock of the MockStore interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockStoreFrom(i store.Store) *MockStore {
	return &MockStore{
		GetConfigurationPoliciesFunc: &StoreGetConfigurationPoliciesFunc{
			defaultHook: i.GetConfigurationPolicies,
		},
		ListFunc: &StoreListFunc{
			defaultHook: i.List,
		},
		UpdateReposMatchingPatternsFunc: &StoreUpdateReposMatchingPatternsFunc{
			defaultHook: i.UpdateReposMatchingPatterns,
		},
	}
}

// StoreGetConfigurationPoliciesFunc describes the behavior when the
// GetConfigurationPolicies method of the parent MockStore instance is
// invoked.
type StoreGetConfigurationPoliciesFunc struct {
	defaultHook func(context.Context, shared.GetConfigurationPoliciesOptions) ([]shared.ConfigurationPolicy, int, error)
	hooks       []func(context.Context, shared.GetConfigurationPoliciesOptions) ([]shared.ConfigurationPolicy, int, error)
	history     []StoreGetConfigurationPoliciesFuncCall
	mutex       sync.Mutex
}

// GetConfigurationPolicies delegates to the next hook function in the queue
// and stores the parameter and result values of this invocation.
func (m *MockStore) GetConfigurationPolicies(v0 context.Context, v1 shared.GetConfigurationPoliciesOptions) ([]shared.ConfigurationPolicy, int, error) {
	r0, r1, r2 := m.GetConfigurationPoliciesFunc.nextHook()(v0, v1)
	m.GetConfigurationPoliciesFunc.appendCall(StoreGetConfigurationPoliciesFuncCall{v0, v1, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the
// GetConfigurationPolicies method of the parent MockStore instance is
// invoked and the hook queue is empty.
func (f *StoreGetConfigurationPoliciesFunc) SetDefaultHook(hook func(context.Context, shared.GetConfigurationPoliciesOptions) ([]shared.ConfigurationPolicy, int, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetConfigurationPolicies method of the parent MockStore instance invokes
// the hook at the front of the queue and discards it. After the queue is
// empty, the default hook function is invoked for any future action.
func (f *StoreGetConfigurationPoliciesFunc) PushHook(hook func(context.Context, shared.GetConfigurationPoliciesOptions) ([]shared.ConfigurationPolicy, int, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreGetConfigurationPoliciesFunc) SetDefaultReturn(r0 []shared.ConfigurationPolicy, r1 int, r2 error) {
	f.SetDefaultHook(func(context.Context, shared.GetConfigurationPoliciesOptions) ([]shared.ConfigurationPolicy, int, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreGetConfigurationPoliciesFunc) PushReturn(r0 []shared.ConfigurationPolicy, r1 int, r2 error) {
	f.PushHook(func(context.Context, shared.GetConfigurationPoliciesOptions) ([]shared.ConfigurationPolicy, int, error) {
		return r0, r1, r2
	})
}

func (f *StoreGetConfigurationPoliciesFunc) nextHook() func(context.Context, shared.GetConfigurationPoliciesOptions) ([]shared.ConfigurationPolicy, int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreGetConfigurationPoliciesFunc) appendCall(r0 StoreGetConfigurationPoliciesFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreGetConfigurationPoliciesFuncCall
// objects describing the invocations of this function.
func (f *StoreGetConfigurationPoliciesFunc) History() []StoreGetConfigurationPoliciesFuncCall {
	f.mutex.Lock()
	history := make([]StoreGetConfigurationPoliciesFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreGetConfigurationPoliciesFuncCall is an object that describes an
// invocation of method GetConfigurationPolicies on an instance of
// MockStore.
type StoreGetConfigurationPoliciesFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 shared.GetConfigurationPoliciesOptions
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []shared.ConfigurationPolicy
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 int
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreGetConfigurationPoliciesFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreGetConfigurationPoliciesFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}

// StoreListFunc describes the behavior when the List method of the parent
// MockStore instance is invoked.
type StoreListFunc struct {
	defaultHook func(context.Context, store.ListOpts) ([]shared.Policy, error)
	hooks       []func(context.Context, store.ListOpts) ([]shared.Policy, error)
	history     []StoreListFuncCall
	mutex       sync.Mutex
}

// List delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockStore) List(v0 context.Context, v1 store.ListOpts) ([]shared.Policy, error) {
	r0, r1 := m.ListFunc.nextHook()(v0, v1)
	m.ListFunc.appendCall(StoreListFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the List method of the
// parent MockStore instance is invoked and the hook queue is empty.
func (f *StoreListFunc) SetDefaultHook(hook func(context.Context, store.ListOpts) ([]shared.Policy, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// List method of the parent MockStore instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *StoreListFunc) PushHook(hook func(context.Context, store.ListOpts) ([]shared.Policy, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreListFunc) SetDefaultReturn(r0 []shared.Policy, r1 error) {
	f.SetDefaultHook(func(context.Context, store.ListOpts) ([]shared.Policy, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreListFunc) PushReturn(r0 []shared.Policy, r1 error) {
	f.PushHook(func(context.Context, store.ListOpts) ([]shared.Policy, error) {
		return r0, r1
	})
}

func (f *StoreListFunc) nextHook() func(context.Context, store.ListOpts) ([]shared.Policy, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreListFunc) appendCall(r0 StoreListFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreListFuncCall objects describing the
// invocations of this function.
func (f *StoreListFunc) History() []StoreListFuncCall {
	f.mutex.Lock()
	history := make([]StoreListFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreListFuncCall is an object that describes an invocation of method
// List on an instance of MockStore.
type StoreListFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 store.ListOpts
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []shared.Policy
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreListFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreListFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreUpdateReposMatchingPatternsFunc describes the behavior when the
// UpdateReposMatchingPatterns method of the parent MockStore instance is
// invoked.
type StoreUpdateReposMatchingPatternsFunc struct {
	defaultHook func(context.Context, []string, int, *int) error
	hooks       []func(context.Context, []string, int, *int) error
	history     []StoreUpdateReposMatchingPatternsFuncCall
	mutex       sync.Mutex
}

// UpdateReposMatchingPatterns delegates to the next hook function in the
// queue and stores the parameter and result values of this invocation.
func (m *MockStore) UpdateReposMatchingPatterns(v0 context.Context, v1 []string, v2 int, v3 *int) error {
	r0 := m.UpdateReposMatchingPatternsFunc.nextHook()(v0, v1, v2, v3)
	m.UpdateReposMatchingPatternsFunc.appendCall(StoreUpdateReposMatchingPatternsFuncCall{v0, v1, v2, v3, r0})
	return r0
}

// SetDefaultHook sets function that is called when the
// UpdateReposMatchingPatterns method of the parent MockStore instance is
// invoked and the hook queue is empty.
func (f *StoreUpdateReposMatchingPatternsFunc) SetDefaultHook(hook func(context.Context, []string, int, *int) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// UpdateReposMatchingPatterns method of the parent MockStore instance
// invokes the hook at the front of the queue and discards it. After the
// queue is empty, the default hook function is invoked for any future
// action.
func (f *StoreUpdateReposMatchingPatternsFunc) PushHook(hook func(context.Context, []string, int, *int) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreUpdateReposMatchingPatternsFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, []string, int, *int) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreUpdateReposMatchingPatternsFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, []string, int, *int) error {
		return r0
	})
}

func (f *StoreUpdateReposMatchingPatternsFunc) nextHook() func(context.Context, []string, int, *int) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreUpdateReposMatchingPatternsFunc) appendCall(r0 StoreUpdateReposMatchingPatternsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreUpdateReposMatchingPatternsFuncCall
// objects describing the invocations of this function.
func (f *StoreUpdateReposMatchingPatternsFunc) History() []StoreUpdateReposMatchingPatternsFuncCall {
	f.mutex.Lock()
	history := make([]StoreUpdateReposMatchingPatternsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreUpdateReposMatchingPatternsFuncCall is an object that describes an
// invocation of method UpdateReposMatchingPatterns on an instance of
// MockStore.
type StoreUpdateReposMatchingPatternsFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 []string
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 int
	// Arg3 is the value of the 4th argument passed to this method
	// invocation.
	Arg3 *int
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreUpdateReposMatchingPatternsFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2, c.Arg3}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreUpdateReposMatchingPatternsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
