package main

import (
	"encoding/json"
	dapr_common "github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/gorilla/mux"
	"github.com/m-to-n/session-management/dapr"
	"log"
	"net/http"
)

type daprConfig struct {
	Entities                   []string `json:"entities,omitempty"`
	ActorIdleTimeout           string   `json:"actorIdleTimeout,omitempty"`
	ActorScanInterval          string   `json:"actorScanInterval,omitempty"`
	DrainOngoingCallTimeout    string   `json:"drainOngoingCallTimeout,omitempty"`
	DrainRebalancedActors      bool     `json:"drainRebalancedActors,omitempty"`
	RemindersStoragePartitions int      `json:"remindersStoragePartitions,omitempty"`
}

var daprConfigResponse = daprConfig{
	ActorIdleTimeout: "60sec",
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("configHandler called")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(daprConfigResponse)
}

func main() {

	log.Println("Starting session-management service...")

	// https://docs.dapr.io/developing-applications/sdks/go/go-service/http-service/
	// https://docs.dapr.io/developing-applications/building-blocks/actors/howto-actors/#actor-runtime-configuration-1
	mux := mux.NewRouter()
	mux.HandleFunc("/dapr/config", configHandler)

	// actor is supported with HTTP API only!
	// TODO - common now does support DaprService only which is using github.com/dapr/go-sdk/service/grpc!
	var s dapr_common.Service
	if 1 == 1 /* without custom dapr config */ {
		s = daprd.NewService(dapr.DAPR_APP_HTTP_ADDR)
	} else /* with custom dapr config - not working for now */ {
		s = daprd.NewServiceWithMux(dapr.DAPR_APP_HTTP_ADDR, mux)
	}

	s.RegisterActorImplFactory(dapr.ActorFactory)

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("dapr server error: %v", err)
	}
}
