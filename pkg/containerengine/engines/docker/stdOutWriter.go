package docker

import (
  "time"
  "io"
  "github.com/opspec-io/sdk-golang/pkg/model"
  "bufio"
  "github.com/opspec-io/engine/util/eventing"
)

func NewStdOutWriter(
eventPublisher eventing.EventPublisher,
opRunId string,
rootOpRunId string,
) io.Writer {

  reader, writer := io.Pipe()
  scanner := bufio.NewScanner(reader)

  go func() {
    for scanner.Scan() {
      eventPublisher.Publish(
        model.Event{
          Timestamp:time.Now().UTC(),
          ContainerStdOutWrittenTo:&model.ContainerStdOutWrittenToEvent{
            Data:scanner.Bytes(),
            OpRunId:opRunId,
            RootOpRunId:rootOpRunId,
          },
        },
      )
    }
  }()

  return writer

}
