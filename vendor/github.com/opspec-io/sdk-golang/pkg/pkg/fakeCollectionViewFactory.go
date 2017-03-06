// This file was generated by counterfeiter
package pkg

import (
	"sync"

	"github.com/opspec-io/sdk-golang/pkg/model"
)

type fakeCollectionViewFactory struct {
	ConstructStub        func(collectionPackagePath string) (collectionView model.CollectionView, err error)
	constructMutex       sync.RWMutex
	constructArgsForCall []struct {
		collectionPackagePath string
	}
	constructReturns struct {
		result1 model.CollectionView
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeCollectionViewFactory) Construct(collectionPackagePath string) (collectionView model.CollectionView, err error) {
	fake.constructMutex.Lock()
	fake.constructArgsForCall = append(fake.constructArgsForCall, struct {
		collectionPackagePath string
	}{collectionPackagePath})
	fake.recordInvocation("Construct", []interface{}{collectionPackagePath})
	fake.constructMutex.Unlock()
	if fake.ConstructStub != nil {
		return fake.ConstructStub(collectionPackagePath)
	} else {
		return fake.constructReturns.result1, fake.constructReturns.result2
	}
}

func (fake *fakeCollectionViewFactory) ConstructCallCount() int {
	fake.constructMutex.RLock()
	defer fake.constructMutex.RUnlock()
	return len(fake.constructArgsForCall)
}

func (fake *fakeCollectionViewFactory) ConstructArgsForCall(i int) string {
	fake.constructMutex.RLock()
	defer fake.constructMutex.RUnlock()
	return fake.constructArgsForCall[i].collectionPackagePath
}

func (fake *fakeCollectionViewFactory) ConstructReturns(result1 model.CollectionView, result2 error) {
	fake.ConstructStub = nil
	fake.constructReturns = struct {
		result1 model.CollectionView
		result2 error
	}{result1, result2}
}

func (fake *fakeCollectionViewFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.constructMutex.RLock()
	defer fake.constructMutex.RUnlock()
	return fake.invocations
}

func (fake *fakeCollectionViewFactory) recordInvocation(key string, args []interface{}) {
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