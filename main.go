package main

import (
	common_dapr "github.com/m-to-n/common/dapr"
	"github.com/m-to-n/session-management/dapr"
	"log"
)

func main() {

	log.Println("Starting session-management service...")

	s := common_dapr.DaprService(dapr.DAPR_APP_GRPC_ADDR)

	if err := s.Start(); err != nil {
		log.Fatalf("dapr server error: %v", err)
	}
}
