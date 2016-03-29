// This file was generated by counterfeiter
package core

import (
  "sync"

  "github.com/dev-op-spec/engine/core/models"
)

type FakeApi struct {
  AddDevOpStub                        func(req models.AddDevOpReq) (err error)
  addDevOpMutex                       sync.RWMutex
  addDevOpArgsForCall                 []struct {
    req models.AddDevOpReq
  }
  addDevOpReturns                     struct {
                                        result1 error
                                      }
  AddPipelineStub                     func(req models.AddPipelineReq) (err error)
  addPipelineMutex                    sync.RWMutex
  addPipelineArgsForCall              []struct {
    req models.AddPipelineReq
  }
  addPipelineReturns                  struct {
                                        result1 error
                                      }
  AddStageToPipelineStub              func(req models.AddStageToPipelineReq) (err error)
  addStageToPipelineMutex             sync.RWMutex
  addStageToPipelineArgsForCall       []struct {
    req models.AddStageToPipelineReq
  }
  addStageToPipelineReturns           struct {
                                        result1 error
                                      }
  ListDevOpsStub                      func(pathToProjectRootDir string) (devOps []models.DevOpView, err error)
  listDevOpsMutex                     sync.RWMutex
  listDevOpsArgsForCall               []struct {
    pathToProjectRootDir string
  }
  listDevOpsReturns                   struct {
                                        result1 []models.DevOpView
                                        result2 error
                                      }
  ListPipelinesStub                   func(pathToProjectRootDir string) (pipelines []models.PipelineView, err error)
  listPipelinesMutex                  sync.RWMutex
  listPipelinesArgsForCall            []struct {
    pathToProjectRootDir string
  }
  listPipelinesReturns                struct {
                                        result1 []models.PipelineView
                                        result2 error
                                      }
  RunDevOpStub                        func(req models.RunDevOpReq) (devOpRun models.DevOpRunView, err error)
  runDevOpMutex                       sync.RWMutex
  runDevOpArgsForCall                 []struct {
    req models.RunDevOpReq
  }
  runDevOpReturns                     struct {
                                        result1 models.DevOpRunView
                                        result2 error
                                      }
  RunPipelineStub                     func(req models.RunPipelineReq) (pipelineRun models.PipelineRunView, err error)
  runPipelineMutex                    sync.RWMutex
  runPipelineArgsForCall              []struct {
    req models.RunPipelineReq
  }
  runPipelineReturns                  struct {
                                        result1 models.PipelineRunView
                                        result2 error
                                      }
  SetDescriptionOfDevOpStub           func(req models.SetDescriptionOfDevOpReq) (err error)
  setDescriptionOfDevOpMutex          sync.RWMutex
  setDescriptionOfDevOpArgsForCall    []struct {
    req models.SetDescriptionOfDevOpReq
  }
  setDescriptionOfDevOpReturns        struct {
                                        result1 error
                                      }
  SetDescriptionOfPipelineStub        func(req models.SetDescriptionOfPipelineReq) (err error)
  setDescriptionOfPipelineMutex       sync.RWMutex
  setDescriptionOfPipelineArgsForCall []struct {
    req models.SetDescriptionOfPipelineReq
  }
  setDescriptionOfPipelineReturns     struct {
                                        result1 error
                                      }
}

func (fake *FakeApi) AddDevOp(req models.AddDevOpReq) (err error) {
  fake.addDevOpMutex.Lock()
  fake.addDevOpArgsForCall = append(fake.addDevOpArgsForCall, struct {
    req models.AddDevOpReq
  }{req})
  fake.addDevOpMutex.Unlock()
  if fake.AddDevOpStub != nil {
    return fake.AddDevOpStub(req)
  } else {
    return fake.addDevOpReturns.result1
  }
}

func (fake *FakeApi) AddDevOpCallCount() int {
  fake.addDevOpMutex.RLock()
  defer fake.addDevOpMutex.RUnlock()
  return len(fake.addDevOpArgsForCall)
}

func (fake *FakeApi) AddDevOpArgsForCall(i int) models.AddDevOpReq {
  fake.addDevOpMutex.RLock()
  defer fake.addDevOpMutex.RUnlock()
  return fake.addDevOpArgsForCall[i].req
}

func (fake *FakeApi) AddDevOpReturns(result1 error) {
  fake.AddDevOpStub = nil
  fake.addDevOpReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeApi) AddPipeline(req models.AddPipelineReq) (err error) {
  fake.addPipelineMutex.Lock()
  fake.addPipelineArgsForCall = append(fake.addPipelineArgsForCall, struct {
    req models.AddPipelineReq
  }{req})
  fake.addPipelineMutex.Unlock()
  if fake.AddPipelineStub != nil {
    return fake.AddPipelineStub(req)
  } else {
    return fake.addPipelineReturns.result1
  }
}

func (fake *FakeApi) AddPipelineCallCount() int {
  fake.addPipelineMutex.RLock()
  defer fake.addPipelineMutex.RUnlock()
  return len(fake.addPipelineArgsForCall)
}

func (fake *FakeApi) AddPipelineArgsForCall(i int) models.AddPipelineReq {
  fake.addPipelineMutex.RLock()
  defer fake.addPipelineMutex.RUnlock()
  return fake.addPipelineArgsForCall[i].req
}

func (fake *FakeApi) AddPipelineReturns(result1 error) {
  fake.AddPipelineStub = nil
  fake.addPipelineReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeApi) AddStageToPipeline(req models.AddStageToPipelineReq) (err error) {
  fake.addStageToPipelineMutex.Lock()
  fake.addStageToPipelineArgsForCall = append(fake.addStageToPipelineArgsForCall, struct {
    req models.AddStageToPipelineReq
  }{req})
  fake.addStageToPipelineMutex.Unlock()
  if fake.AddStageToPipelineStub != nil {
    return fake.AddStageToPipelineStub(req)
  } else {
    return fake.addStageToPipelineReturns.result1
  }
}

func (fake *FakeApi) AddStageToPipelineCallCount() int {
  fake.addStageToPipelineMutex.RLock()
  defer fake.addStageToPipelineMutex.RUnlock()
  return len(fake.addStageToPipelineArgsForCall)
}

func (fake *FakeApi) AddStageToPipelineArgsForCall(i int) models.AddStageToPipelineReq {
  fake.addStageToPipelineMutex.RLock()
  defer fake.addStageToPipelineMutex.RUnlock()
  return fake.addStageToPipelineArgsForCall[i].req
}

func (fake *FakeApi) AddStageToPipelineReturns(result1 error) {
  fake.AddStageToPipelineStub = nil
  fake.addStageToPipelineReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeApi) ListDevOps(pathToProjectRootDir string) (devOps []models.DevOpView, err error) {
  fake.listDevOpsMutex.Lock()
  fake.listDevOpsArgsForCall = append(fake.listDevOpsArgsForCall, struct {
    pathToProjectRootDir string
  }{pathToProjectRootDir})
  fake.listDevOpsMutex.Unlock()
  if fake.ListDevOpsStub != nil {
    return fake.ListDevOpsStub(pathToProjectRootDir)
  } else {
    return fake.listDevOpsReturns.result1, fake.listDevOpsReturns.result2
  }
}

func (fake *FakeApi) ListDevOpsCallCount() int {
  fake.listDevOpsMutex.RLock()
  defer fake.listDevOpsMutex.RUnlock()
  return len(fake.listDevOpsArgsForCall)
}

func (fake *FakeApi) ListDevOpsArgsForCall(i int) string {
  fake.listDevOpsMutex.RLock()
  defer fake.listDevOpsMutex.RUnlock()
  return fake.listDevOpsArgsForCall[i].pathToProjectRootDir
}

func (fake *FakeApi) ListDevOpsReturns(result1 []models.DevOpView, result2 error) {
  fake.ListDevOpsStub = nil
  fake.listDevOpsReturns = struct {
    result1 []models.DevOpView
    result2 error
  }{result1, result2}
}

func (fake *FakeApi) ListPipelines(pathToProjectRootDir string) (pipelines []models.PipelineView, err error) {
  fake.listPipelinesMutex.Lock()
  fake.listPipelinesArgsForCall = append(fake.listPipelinesArgsForCall, struct {
    pathToProjectRootDir string
  }{pathToProjectRootDir})
  fake.listPipelinesMutex.Unlock()
  if fake.ListPipelinesStub != nil {
    return fake.ListPipelinesStub(pathToProjectRootDir)
  } else {
    return fake.listPipelinesReturns.result1, fake.listPipelinesReturns.result2
  }
}

func (fake *FakeApi) ListPipelinesCallCount() int {
  fake.listPipelinesMutex.RLock()
  defer fake.listPipelinesMutex.RUnlock()
  return len(fake.listPipelinesArgsForCall)
}

func (fake *FakeApi) ListPipelinesArgsForCall(i int) string {
  fake.listPipelinesMutex.RLock()
  defer fake.listPipelinesMutex.RUnlock()
  return fake.listPipelinesArgsForCall[i].pathToProjectRootDir
}

func (fake *FakeApi) ListPipelinesReturns(result1 []models.PipelineView, result2 error) {
  fake.ListPipelinesStub = nil
  fake.listPipelinesReturns = struct {
    result1 []models.PipelineView
    result2 error
  }{result1, result2}
}

func (fake *FakeApi) RunDevOp(req models.RunDevOpReq) (devOpRun models.DevOpRunView, err error) {
  fake.runDevOpMutex.Lock()
  fake.runDevOpArgsForCall = append(fake.runDevOpArgsForCall, struct {
    req models.RunDevOpReq
  }{req})
  fake.runDevOpMutex.Unlock()
  if fake.RunDevOpStub != nil {
    return fake.RunDevOpStub(req)
  } else {
    return fake.runDevOpReturns.result1, fake.runDevOpReturns.result2
  }
}

func (fake *FakeApi) RunDevOpCallCount() int {
  fake.runDevOpMutex.RLock()
  defer fake.runDevOpMutex.RUnlock()
  return len(fake.runDevOpArgsForCall)
}

func (fake *FakeApi) RunDevOpArgsForCall(i int) models.RunDevOpReq {
  fake.runDevOpMutex.RLock()
  defer fake.runDevOpMutex.RUnlock()
  return fake.runDevOpArgsForCall[i].req
}

func (fake *FakeApi) RunDevOpReturns(result1 models.DevOpRunView, result2 error) {
  fake.RunDevOpStub = nil
  fake.runDevOpReturns = struct {
    result1 models.DevOpRunView
    result2 error
  }{result1, result2}
}

func (fake *FakeApi) RunPipeline(req models.RunPipelineReq) (pipelineRun models.PipelineRunView, err error) {
  fake.runPipelineMutex.Lock()
  fake.runPipelineArgsForCall = append(fake.runPipelineArgsForCall, struct {
    req models.RunPipelineReq
  }{req})
  fake.runPipelineMutex.Unlock()
  if fake.RunPipelineStub != nil {
    return fake.RunPipelineStub(req)
  } else {
    return fake.runPipelineReturns.result1, fake.runPipelineReturns.result2
  }
}

func (fake *FakeApi) RunPipelineCallCount() int {
  fake.runPipelineMutex.RLock()
  defer fake.runPipelineMutex.RUnlock()
  return len(fake.runPipelineArgsForCall)
}

func (fake *FakeApi) RunPipelineArgsForCall(i int) models.RunPipelineReq {
  fake.runPipelineMutex.RLock()
  defer fake.runPipelineMutex.RUnlock()
  return fake.runPipelineArgsForCall[i].req
}

func (fake *FakeApi) RunPipelineReturns(result1 models.PipelineRunView, result2 error) {
  fake.RunPipelineStub = nil
  fake.runPipelineReturns = struct {
    result1 models.PipelineRunView
    result2 error
  }{result1, result2}
}

func (fake *FakeApi) SetDescriptionOfDevOp(req models.SetDescriptionOfDevOpReq) (err error) {
  fake.setDescriptionOfDevOpMutex.Lock()
  fake.setDescriptionOfDevOpArgsForCall = append(fake.setDescriptionOfDevOpArgsForCall, struct {
    req models.SetDescriptionOfDevOpReq
  }{req})
  fake.setDescriptionOfDevOpMutex.Unlock()
  if fake.SetDescriptionOfDevOpStub != nil {
    return fake.SetDescriptionOfDevOpStub(req)
  } else {
    return fake.setDescriptionOfDevOpReturns.result1
  }
}

func (fake *FakeApi) SetDescriptionOfDevOpCallCount() int {
  fake.setDescriptionOfDevOpMutex.RLock()
  defer fake.setDescriptionOfDevOpMutex.RUnlock()
  return len(fake.setDescriptionOfDevOpArgsForCall)
}

func (fake *FakeApi) SetDescriptionOfDevOpArgsForCall(i int) models.SetDescriptionOfDevOpReq {
  fake.setDescriptionOfDevOpMutex.RLock()
  defer fake.setDescriptionOfDevOpMutex.RUnlock()
  return fake.setDescriptionOfDevOpArgsForCall[i].req
}

func (fake *FakeApi) SetDescriptionOfDevOpReturns(result1 error) {
  fake.SetDescriptionOfDevOpStub = nil
  fake.setDescriptionOfDevOpReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeApi) SetDescriptionOfPipeline(req models.SetDescriptionOfPipelineReq) (err error) {
  fake.setDescriptionOfPipelineMutex.Lock()
  fake.setDescriptionOfPipelineArgsForCall = append(fake.setDescriptionOfPipelineArgsForCall, struct {
    req models.SetDescriptionOfPipelineReq
  }{req})
  fake.setDescriptionOfPipelineMutex.Unlock()
  if fake.SetDescriptionOfPipelineStub != nil {
    return fake.SetDescriptionOfPipelineStub(req)
  } else {
    return fake.setDescriptionOfPipelineReturns.result1
  }
}

func (fake *FakeApi) SetDescriptionOfPipelineCallCount() int {
  fake.setDescriptionOfPipelineMutex.RLock()
  defer fake.setDescriptionOfPipelineMutex.RUnlock()
  return len(fake.setDescriptionOfPipelineArgsForCall)
}

func (fake *FakeApi) SetDescriptionOfPipelineArgsForCall(i int) models.SetDescriptionOfPipelineReq {
  fake.setDescriptionOfPipelineMutex.RLock()
  defer fake.setDescriptionOfPipelineMutex.RUnlock()
  return fake.setDescriptionOfPipelineArgsForCall[i].req
}

func (fake *FakeApi) SetDescriptionOfPipelineReturns(result1 error) {
  fake.SetDescriptionOfPipelineStub = nil
  fake.setDescriptionOfPipelineReturns = struct {
    result1 error
  }{result1}
}

var _ Api = new(FakeApi)
