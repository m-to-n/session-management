package main

import (
	"context"
	"fmt"
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
	client.ImplActorClientStub(actor1)

	rsp, err := actor1.SendMessage(ctx, "hey!")
	if err != nil {
		panic(err)
	}
	fmt.Println("SendMessage result: ", rsp)

}
