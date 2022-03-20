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
	daprClient    dapr.Client
	ShortId       string
	JustInitiated bool
}

func (a *SessionActor) Type() string {
	return SESSION_ACTOR_TYPE
}

func (a *SessionActor) SendMessage(ctx context.Context, message string) (string, error) {
	log.Printf("SessionActor(%s).SendMessage called(actorId=%s - %s): %s", a.Type(), a.ID(), a.ShortId, message)

	if a.JustInitiated == true {
		a.JustInitiated = false

		a.daprClient.RegisterActorReminder(ctx, &dapr.RegisterActorReminderRequest{
			ActorType: a.Type(),
			ActorID:   a.ID(),
			Name:      "session_ping_reminder",
			DueTime:   "5s",
			Period:    "5s",
			Data:      []byte("ping"),
		})
	}

	return fmt.Sprintf("[%s] [%s] You said: %s!", a.ID(), a.ShortId, message), nil
}

func (a *SessionActor) ReminderCall(reminderName string, state []byte, dueTime string, period string) {
	fmt.Println(
		"ReminderCall: receive reminder = ", reminderName,
		" state = ", string(state),
		"duetime = ", dueTime,
		"period = ", period,
		"actor = ", a.ID(),
		"actor_shortid = ", a.ShortId)
}

func ActorFactory() actor.Server {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	sid, _ := shortid.Generate()
	return &SessionActor{
		daprClient:    client,
		ShortId:       sid,
		JustInitiated: true,
	}
}
