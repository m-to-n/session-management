package main

import (
	"context"
	dapr "github.com/dapr/go-sdk/client"
	sm_dapr "github.com/m-to-n/session-management/dapr"
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

	actor1 := sm_dapr.NewSessionActorClientStub("id123")
	actor2 := sm_dapr.NewSessionActorClientStub("id456")
	actor3 := sm_dapr.NewSessionActorClientStub("id456")
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
