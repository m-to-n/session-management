package dapr

import (
	"context"
	"fmt"
	"github.com/dapr/go-sdk/actor"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/teris-io/shortid"
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

type SessionActor struct {
	actor.ServerImplBase
	daprClient dapr.Client
	ShortId    string
}

func (a *SessionActor) Type() string {
	return SESSION_ACTOR_TYPE
}

func (a *SessionActor) SendMessage(ctx context.Context, message string) (string, error) {
	// TODO - implement
	log.Printf("SessionActor(%s).SendMessage called(actorId=%s - %s): %s", a.Type(), a.ID(), a.ShortId, message)
	return fmt.Sprintf("[%s] [%s] You said: %s!", a.ID(), a.ShortId, message), nil
}

func ActorFactory() actor.Server {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	sid, _ := shortid.Generate()
	return &SessionActor{
		daprClient: client,
		ShortId:    sid,
	}
}
