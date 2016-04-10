// This file was generated by counterfeiter
package core

import (
  "sync"

  "github.com/dev-op-spec/engine/core/models"
)

type fakeRunOpUseCase struct {
  ExecuteStub        func(req models.RunOpReq, namesOfAlreadyRunOps []*models.Url) (opRun models.OpRunDetailedView, err error)
  executeMutex       sync.RWMutex
  executeArgsForCall []struct {
    req                  models.RunOpReq
    namesOfAlreadyRunOps []*models.Url
  }
  executeReturns     struct {
                       result1 models.OpRunDetailedView
                       result2 error
                     }
}

func (fake *fakeRunOpUseCase) Execute(req models.RunOpReq, namesOfAlreadyRunOps []*models.Url) (opRun models.OpRunDetailedView, err error) {
  fake.executeMutex.Lock()
  fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
    req                  models.RunOpReq
    namesOfAlreadyRunOps []*models.Url
  }{req, namesOfAlreadyRunOps})
  fake.executeMutex.Unlock()
  if fake.ExecuteStub != nil {
    return fake.ExecuteStub(req, namesOfAlreadyRunOps)
  } else {
    return fake.executeReturns.result1, fake.executeReturns.result2
  }
}

func (fake *fakeRunOpUseCase) ExecuteCallCount() int {
  fake.executeMutex.RLock()
  defer fake.executeMutex.RUnlock()
  return len(fake.executeArgsForCall)
}

func (fake *fakeRunOpUseCase) ExecuteArgsForCall(i int) (models.RunOpReq, []*models.Url) {
  fake.executeMutex.RLock()
  defer fake.executeMutex.RUnlock()
  return fake.executeArgsForCall[i].req, fake.executeArgsForCall[i].namesOfAlreadyRunOps
}

func (fake *fakeRunOpUseCase) ExecuteReturns(result1 models.OpRunDetailedView, result2 error) {
  fake.ExecuteStub = nil
  fake.executeReturns = struct {
    result1 models.OpRunDetailedView
    result2 error
  }{result1, result2}
}
