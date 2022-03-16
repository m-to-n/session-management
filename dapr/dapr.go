package dapr

import (
	"context"
	"fmt"
	"github.com/dapr/go-sdk/actor"
	dapr "github.com/dapr/go-sdk/client"
	"log"
)

// dapr sidecar http port
const DAPR_HTTP_PORT = "3400"

// dapr sidecar grp port
const DAPR_GRPC_PORT = "34000"

// dapr app http address
const DAPR_APP_HTTP_ADDR = ":3401"

// dapr app grpc address
const DAPR_APP_GRPC_ADDR = ":34001"

const SESSION_ACTOR_TYPE = "M2NSessionActor"

type SessionActorClientStub struct {
	ActorId     string
	SendMessage func(context.Context, string) (string, error)
}

func NewSessionActorClientStub(actorId string) *SessionActorClientStub {
	return &SessionActorClientStub{
		ActorId: actorId,
	}
}

func (a *SessionActorClientStub) Type() string {
	return SESSION_ACTOR_TYPE
}

func (a *SessionActorClientStub) ID() string {
	return a.ActorId
}

type SessionActor struct {
	actor.ServerImplBase
	daprClient dapr.Client
}

func (a *SessionActor) Type() string {
	return SESSION_ACTOR_TYPE
}

func (a *SessionActor) SendMessage(ctx context.Context, message string) (string, error) {
	// TODO - implement
	log.Printf("SessionActor(%s).SendMessage called(actorId=%s): %s", a.Type(), a.ID(), message)
	return fmt.Sprintf("[%s] You said: %s!", a.ID(), message), nil
}

func ActorFactory() actor.Server {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	return &SessionActor{
		daprClient: client,
	}
}
