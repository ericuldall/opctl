// Code generated by counterfeiter. DO NOT EDIT.
package ijson

import (
	"sync"
)

type Fake struct {
	MarshalStub        func(v interface{}) ([]byte, error)
	marshalMutex       sync.RWMutex
	marshalArgsForCall []struct {
		v interface{}
	}
	marshalReturns struct {
		result1 []byte
		result2 error
	}
	marshalReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) Marshal(v interface{}) ([]byte, error) {
	fake.marshalMutex.Lock()
	ret, specificReturn := fake.marshalReturnsOnCall[len(fake.marshalArgsForCall)]
	fake.marshalArgsForCall = append(fake.marshalArgsForCall, struct {
		v interface{}
	}{v})
	fake.recordInvocation("Marshal", []interface{}{v})
	fake.marshalMutex.Unlock()
	if fake.MarshalStub != nil {
		return fake.MarshalStub(v)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.marshalReturns.result1, fake.marshalReturns.result2
}

func (fake *Fake) MarshalCallCount() int {
	fake.marshalMutex.RLock()
	defer fake.marshalMutex.RUnlock()
	return len(fake.marshalArgsForCall)
}

func (fake *Fake) MarshalArgsForCall(i int) interface{} {
	fake.marshalMutex.RLock()
	defer fake.marshalMutex.RUnlock()
	return fake.marshalArgsForCall[i].v
}

func (fake *Fake) MarshalReturns(result1 []byte, result2 error) {
	fake.MarshalStub = nil
	fake.marshalReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Fake) MarshalReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.MarshalStub = nil
	if fake.marshalReturnsOnCall == nil {
		fake.marshalReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.marshalReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.marshalMutex.RLock()
	defer fake.marshalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Fake) recordInvocation(key string, args []interface{}) {
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

var _ IJSON = new(Fake)
