package dockercompose

import (
  "github.com/dev-op-spec/engine/core/ports"
  "github.com/dev-op-spec/engine/core/models"
)

func New(
) (containerEngine ports.ContainerEngine, err error) {

  var compositionRoot compositionRoot
  compositionRoot, err = newCompositionRoot()
  if (nil != err) {
    return
  }

  containerEngine = _containerEngine{
    compositionRoot:compositionRoot,
  }

  return

}

type _containerEngine struct {
  compositionRoot compositionRoot
}

func (this _containerEngine) InitOp(
pathToOpDir string,
opName string,
) (err error) {

  return this.compositionRoot.
  InitOpUseCase().
  Execute(
    pathToOpDir,
    opName,
  )

}

func (this _containerEngine) RunOp(
pathToOpDir string,
opName string,
logChannel chan *models.LogEntry,
) (exitCode int, err error) {

  return this.compositionRoot.
  RunOpUseCase().
  Execute(
    pathToOpDir,
    opName,
    logChannel,
  )

}

