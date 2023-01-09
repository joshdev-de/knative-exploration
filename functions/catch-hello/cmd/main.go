package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/joshdev-de/knative-exploration/functions/catch-hello/pkg/receiver"
)

func main() {
	log.Print("Catch Hello started")
	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	log.Fatal(c.StartReceiver(context.Background(), receiver.Receive))
}
