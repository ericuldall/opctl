// Code generated by counterfeiter. DO NOT EDIT.
package data

import (
	"context"
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type fakePuller struct {
	PullStub        func(ctx context.Context, path string, dataRef string, pullCreds *model.PullCreds) error
	pullMutex       sync.RWMutex
	pullArgsForCall []struct {
		ctx       context.Context
		path      string
		dataRef   string
		pullCreds *model.PullCreds
	}
	pullReturns struct {
		result1 error
	}
	pullReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakePuller) Pull(ctx context.Context, path string, dataRef string, pullCreds *model.PullCreds) error {
	fake.pullMutex.Lock()
	ret, specificReturn := fake.pullReturnsOnCall[len(fake.pullArgsForCall)]
	fake.pullArgsForCall = append(fake.pullArgsForCall, struct {
		ctx       context.Context
		path      string
		dataRef   string
		pullCreds *model.PullCreds
	}{ctx, path, dataRef, pullCreds})
	fake.recordInvocation("Pull", []interface{}{ctx, path, dataRef, pullCreds})
	fake.pullMutex.Unlock()
	if fake.PullStub != nil {
		return fake.PullStub(ctx, path, dataRef, pullCreds)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pullReturns.result1
}

func (fake *fakePuller) PullCallCount() int {
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	return len(fake.pullArgsForCall)
}

func (fake *fakePuller) PullArgsForCall(i int) (context.Context, string, string, *model.PullCreds) {
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	return fake.pullArgsForCall[i].ctx, fake.pullArgsForCall[i].path, fake.pullArgsForCall[i].dataRef, fake.pullArgsForCall[i].pullCreds
}

func (fake *fakePuller) PullReturns(result1 error) {
	fake.PullStub = nil
	fake.pullReturns = struct {
		result1 error
	}{result1}
}

func (fake *fakePuller) PullReturnsOnCall(i int, result1 error) {
	fake.PullStub = nil
	if fake.pullReturnsOnCall == nil {
		fake.pullReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pullReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *fakePuller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *fakePuller) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
