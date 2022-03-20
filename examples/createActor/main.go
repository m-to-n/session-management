package main

import (
	"context"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/m-to-n/common/sessions"
	"log"
)

func main() {
	log.Println("Running createActor sample.")

	ctx := context.Background()

	// create the client
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	actor1 := sessions.NewSessionActorClientStub("id123")
	actor2 := sessions.NewSessionActorClientStub("id456")
	// actor3 has same actorId -> new server actor won't be created!
	// instead, actor2 will be reused!
	actor3 := sessions.NewSessionActorClientStub("id456")
	client.ImplActorClientStub(actor1)
	client.ImplActorClientStub(actor2)
	client.ImplActorClientStub(actor3)

	rsp, err := actor1.SendMessage(ctx, "hey!")
	if err != nil {
		panic(err)
	}
	log.Println("SendMessage result: ", rsp)

	rsp, err = actor2.SendMessage(ctx, "you!")
	if err != nil {
		panic(err)
	}
	log.Println("SendMessage result: ", rsp)

	rsp, err = actor3.SendMessage(ctx, "hi!")
	if err != nil {
		panic(err)
	}
	log.Println("SendMessage result: ", rsp)

}
