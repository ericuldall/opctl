// Code generated by counterfeiter. DO NOT EDIT.
package outputs

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type FakeInterpreter struct {
	InterpretStub        func(outputArgs map[string]*model.Value, outputParams map[string]*model.Param, opPath string) (map[string]*model.Value, error)
	interpretMutex       sync.RWMutex
	interpretArgsForCall []struct {
		outputArgs   map[string]*model.Value
		outputParams map[string]*model.Param
		opPath       string
	}
	interpretReturns struct {
		result1 map[string]*model.Value
		result2 error
	}
	interpretReturnsOnCall map[int]struct {
		result1 map[string]*model.Value
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInterpreter) Interpret(outputArgs map[string]*model.Value, outputParams map[string]*model.Param, opPath string) (map[string]*model.Value, error) {
	fake.interpretMutex.Lock()
	ret, specificReturn := fake.interpretReturnsOnCall[len(fake.interpretArgsForCall)]
	fake.interpretArgsForCall = append(fake.interpretArgsForCall, struct {
		outputArgs   map[string]*model.Value
		outputParams map[string]*model.Param
		opPath       string
	}{outputArgs, outputParams, opPath})
	fake.recordInvocation("Interpret", []interface{}{outputArgs, outputParams, opPath})
	fake.interpretMutex.Unlock()
	if fake.InterpretStub != nil {
		return fake.InterpretStub(outputArgs, outputParams, opPath)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.interpretReturns.result1, fake.interpretReturns.result2
}

func (fake *FakeInterpreter) InterpretCallCount() int {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return len(fake.interpretArgsForCall)
}

func (fake *FakeInterpreter) InterpretArgsForCall(i int) (map[string]*model.Value, map[string]*model.Param, string) {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return fake.interpretArgsForCall[i].outputArgs, fake.interpretArgsForCall[i].outputParams, fake.interpretArgsForCall[i].opPath
}

func (fake *FakeInterpreter) InterpretReturns(result1 map[string]*model.Value, result2 error) {
	fake.InterpretStub = nil
	fake.interpretReturns = struct {
		result1 map[string]*model.Value
		result2 error
	}{result1, result2}
}

func (fake *FakeInterpreter) InterpretReturnsOnCall(i int, result1 map[string]*model.Value, result2 error) {
	fake.InterpretStub = nil
	if fake.interpretReturnsOnCall == nil {
		fake.interpretReturnsOnCall = make(map[int]struct {
			result1 map[string]*model.Value
			result2 error
		})
	}
	fake.interpretReturnsOnCall[i] = struct {
		result1 map[string]*model.Value
		result2 error
	}{result1, result2}
}

func (fake *FakeInterpreter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInterpreter) recordInvocation(key string, args []interface{}) {
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

var _ Interpreter = new(FakeInterpreter)
