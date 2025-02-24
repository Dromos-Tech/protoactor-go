package remote

import (
	"github.com/asynkron/protoactor-go/actor"
	remoteProto "github.com/asynkron/protoactor-go/remote/gen"
)

type EndpointTerminatedEvent struct {
	Address string
}

type EndpointConnectedEvent struct {
	Address string
}

type remoteWatch struct {
	Watcher *actor.PID
	Watchee *actor.PID
}

type remoteUnwatch struct {
	Watcher *actor.PID
	Watchee *actor.PID
}

type remoteDeliver struct {
	header       actor.ReadonlyMessageHeader
	message      interface{}
	target       *actor.PID
	sender       *actor.PID
	serializerID int32
}

type remoteTerminate struct {
	Watcher *actor.PID
	Watchee *actor.PID
}

type JsonMessage struct {
	TypeName string
	Json     string
}

var stopMessage interface{} = &actor.Stop{}

var (
	ActorPidRespErr         interface{} = &remoteProto.ActorPidResponse{StatusCode: ResponseStatusCodeERROR.ToInt32()}
	ActorPidRespTimeout     interface{} = &remoteProto.ActorPidResponse{StatusCode: ResponseStatusCodeTIMEOUT.ToInt32()}
	ActorPidRespUnavailable interface{} = &remoteProto.ActorPidResponse{StatusCode: ResponseStatusCodeUNAVAILABLE.ToInt32()}
)

type (
	// Ping is message sent by the actor system to probe an actor is started.
	Ping struct{}

	// Pong is response for ping.
	Pong struct{}
)
