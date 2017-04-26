package core

//go:generate counterfeiter -o ./fakeContainerCaller.go --fake-name fakeContainerCaller ./ containerCaller

import (
	"github.com/opctl/opctl/util/containerprovider"
	"github.com/opctl/opctl/util/pubsub"
	"github.com/opspec-io/sdk-golang/model"
	"strings"
	"time"
)

type containerCaller interface {
	// Executes a container call
	Call(
		inboundScope map[string]*model.Data,
		containerId string,
		scgContainerCall *model.SCGContainerCall,
		pkgRef string,
		rootOpId string,
	) (
		err error,
	)
}

func newContainerCaller(
	containerProvider containerprovider.ContainerProvider,
	pubSub pubsub.PubSub,
	dcgNodeRepo dcgNodeRepo,
) containerCaller {

	return _containerCaller{
		containerProvider: containerProvider,
		pubSub:            pubSub,
		dcgNodeRepo:       dcgNodeRepo,
	}

}

type _containerCaller struct {
	containerProvider containerprovider.ContainerProvider
	pubSub            pubsub.PubSub
	dcgNodeRepo       dcgNodeRepo
}

func (this _containerCaller) Call(
	inboundScope map[string]*model.Data,
	containerId string,
	scgContainerCall *model.SCGContainerCall,
	pkgRef string,
	rootOpId string,
) (
	err error,
) {
	defer func() {
		// defer must be defined before conditional return statements so it always runs

		this.dcgNodeRepo.DeleteIfExists(containerId)

		this.containerProvider.DeleteContainerIfExists(containerId)

		this.pubSub.Publish(
			&model.Event{
				Timestamp: time.Now().UTC(),
				ContainerExited: &model.ContainerExitedEvent{
					ContainerId: containerId,
					PkgRef:      pkgRef,
					RootOpId:    rootOpId,
				},
			},
		)

	}()

	this.dcgNodeRepo.Add(
		&dcgNodeDescriptor{
			Id:        containerId,
			PkgRef:    pkgRef,
			RootOpId:  rootOpId,
			Container: &dcgContainerDescriptor{},
		},
	)

	dcgContainerCall, err := constructDCGContainerCall(inboundScope, scgContainerCall, containerId, rootOpId, pkgRef)
	if nil != err {
		return
	}

	go this.txOutputs(dcgContainerCall, scgContainerCall)

	this.pubSub.Publish(
		&model.Event{
			Timestamp: time.Now().UTC(),
			ContainerStarted: &model.ContainerStartedEvent{
				ContainerId: containerId,
				PkgRef:      pkgRef,
				RootOpId:    rootOpId,
			},
		},
	)
	err = this.containerProvider.RunContainer(
		dcgContainerCall,
		this.pubSub,
	)

	return
}

func (this _containerCaller) txOutputs(
	dcgContainerCall *model.DCGContainerCall,
	scgContainerCall *model.SCGContainerCall,
) {

	// send socket outputs
	for socketAddr, name := range scgContainerCall.Sockets {
		if "0.0.0.0" == socketAddr {
			this.pubSub.Publish(&model.Event{
				Timestamp: time.Now().UTC(),
				OutputInitialized: &model.OutputInitializedEvent{
					Name:     name,
					Value:    &model.Data{Socket: &dcgContainerCall.ContainerId},
					RootOpId: dcgContainerCall.RootOpId,
					CallId:   dcgContainerCall.ContainerId,
				},
			})
		}
	}

	// send file outputs
	for scgContainerFilePath, name := range scgContainerCall.Files {
		for dcgContainerFilePath, dcgHostFilePath := range dcgContainerCall.Files {
			if scgContainerFilePath == dcgContainerFilePath {
				this.pubSub.Publish(&model.Event{
					Timestamp: time.Now().UTC(),
					OutputInitialized: &model.OutputInitializedEvent{
						Name:     name,
						Value:    &model.Data{File: &dcgHostFilePath},
						RootOpId: dcgContainerCall.RootOpId,
						CallId:   dcgContainerCall.ContainerId,
					},
				})
			}
		}
	}

	// send dir outputs
	for scgContainerDirPath, name := range scgContainerCall.Dirs {
		for dcgContainerDirPath, dcgHostDirPath := range dcgContainerCall.Dirs {
			if scgContainerDirPath == dcgContainerDirPath {
				this.pubSub.Publish(&model.Event{
					Timestamp: time.Now().UTC(),
					OutputInitialized: &model.OutputInitializedEvent{
						Name:     name,
						Value:    &model.Data{Dir: &dcgHostDirPath},
						RootOpId: dcgContainerCall.RootOpId,
						CallId:   dcgContainerCall.ContainerId,
					},
				})
			}
		}
	}

	// subscribe to events
	eventChannel := make(chan *model.Event, 150)
	eventFilterSince := time.Now().UTC()
	this.pubSub.Subscribe(
		&model.EventFilter{
			RootOpIds: []string{dcgContainerCall.RootOpId},
			Since:     &eventFilterSince,
		},
		eventChannel,
	)

	// need to track EOF reads
	var isStdErrEOFRead, isStdOutEOFRead bool

	// send string outputs
eventLoop:
	for event := range eventChannel {
		switch {
		case nil != event.ContainerStdErrEOFRead &&
			event.ContainerStdErrEOFRead.ContainerId == dcgContainerCall.ContainerId:
			isStdErrEOFRead = true
			if isStdOutEOFRead {
				// no more events we care about; break eventLoop
				break eventLoop
			}
		case nil != event.ContainerStdOutEOFRead &&
			event.ContainerStdOutEOFRead.ContainerId == dcgContainerCall.ContainerId:
			isStdOutEOFRead = true
			if isStdErrEOFRead {
				// no more events we care about; break eventLoop
				break eventLoop
			}
		case nil != event.ContainerStdErrWrittenTo &&
			event.ContainerStdErrWrittenTo.ContainerId == dcgContainerCall.ContainerId:
			for boundPrefix, name := range scgContainerCall.StdErr {
				rawOutput := string(event.ContainerStdErrWrittenTo.Data)
				trimmedOutput := strings.TrimPrefix(rawOutput, boundPrefix)
				if trimmedOutput != rawOutput {
					// if output trimming had effect we've got a match
					this.pubSub.Publish(&model.Event{
						Timestamp: time.Now().UTC(),
						OutputInitialized: &model.OutputInitializedEvent{
							Name:     name,
							Value:    &model.Data{String: &trimmedOutput},
							RootOpId: dcgContainerCall.RootOpId,
							CallId:   dcgContainerCall.ContainerId,
						},
					})
				}
			}
		case nil != event.ContainerStdOutWrittenTo &&
			event.ContainerStdOutWrittenTo.ContainerId == dcgContainerCall.ContainerId:
			for boundPrefix, name := range scgContainerCall.StdOut {
				rawOutput := string(event.ContainerStdOutWrittenTo.Data)
				trimmedOutput := strings.TrimPrefix(rawOutput, boundPrefix)
				if trimmedOutput != rawOutput {
					// if output trimming had effect we've got a match
					this.pubSub.Publish(&model.Event{
						Timestamp: time.Now().UTC(),
						OutputInitialized: &model.OutputInitializedEvent{
							Name:     name,
							Value:    &model.Data{String: &trimmedOutput},
							RootOpId: dcgContainerCall.RootOpId,
							CallId:   dcgContainerCall.ContainerId,
						},
					})
				}
			}
		}
	}

	return
}