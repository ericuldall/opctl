// This file was generated by counterfeiter
package dockercompose

import (
  "sync"
)

type fakeCompositionRoot struct {
  InitOpUseCaseStub        func() initOpUseCase
  initOpUseCaseMutex       sync.RWMutex
  initOpUseCaseArgsForCall []struct{}
  initOpUseCaseReturns     struct {
                             result1 initOpUseCase
                           }
  RunOpUseCaseStub         func() runOpUseCase
  runOpUseCaseMutex        sync.RWMutex
  runOpUseCaseArgsForCall  []struct{}
  runOpUseCaseReturns      struct {
                             result1 runOpUseCase
                           }
}

func (fake *fakeCompositionRoot) InitOpUseCase() initOpUseCase {
  fake.initOpUseCaseMutex.Lock()
  fake.initOpUseCaseArgsForCall = append(fake.initOpUseCaseArgsForCall, struct{}{})
  fake.initOpUseCaseMutex.Unlock()
  if fake.InitOpUseCaseStub != nil {
    return fake.InitOpUseCaseStub()
  } else {
    return fake.initOpUseCaseReturns.result1
  }
}

func (fake *fakeCompositionRoot) InitOpUseCaseCallCount() int {
  fake.initOpUseCaseMutex.RLock()
  defer fake.initOpUseCaseMutex.RUnlock()
  return len(fake.initOpUseCaseArgsForCall)
}

func (fake *fakeCompositionRoot) InitOpUseCaseReturns(result1 initOpUseCase) {
  fake.InitOpUseCaseStub = nil
  fake.initOpUseCaseReturns = struct {
    result1 initOpUseCase
  }{result1}
}

func (fake *fakeCompositionRoot) RunOpUseCase() runOpUseCase {
  fake.runOpUseCaseMutex.Lock()
  fake.runOpUseCaseArgsForCall = append(fake.runOpUseCaseArgsForCall, struct{}{})
  fake.runOpUseCaseMutex.Unlock()
  if fake.RunOpUseCaseStub != nil {
    return fake.RunOpUseCaseStub()
  } else {
    return fake.runOpUseCaseReturns.result1
  }
}

func (fake *fakeCompositionRoot) RunOpUseCaseCallCount() int {
  fake.runOpUseCaseMutex.RLock()
  defer fake.runOpUseCaseMutex.RUnlock()
  return len(fake.runOpUseCaseArgsForCall)
}

func (fake *fakeCompositionRoot) RunOpUseCaseReturns(result1 runOpUseCase) {
  fake.RunOpUseCaseStub = nil
  fake.runOpUseCaseReturns = struct {
    result1 runOpUseCase
  }{result1}
}
