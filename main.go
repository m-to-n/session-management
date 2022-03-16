package main

import (
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/m-to-n/session-management/dapr"
	"log"
	"net/http"
)

func main() {

	log.Println("Starting session-management service...")

	// actor is supported with HTTP API only!
	// TODO - common now does support DaprService only which is using github.com/dapr/go-sdk/service/grpc!
	s := daprd.NewService(dapr.DAPR_APP_HTTP_ADDR)
	s.RegisterActorImplFactory(dapr.ActorFactory)
	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("dapr server error: %v", err)
	}
}
