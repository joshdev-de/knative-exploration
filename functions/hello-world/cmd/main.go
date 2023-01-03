package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/joshdev-de/knative-exploration/functions/hello-world/pkg/receiver"
)

func main() {
	log.Print("Hello world sample started.")
	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	log.Fatal(c.StartReceiver(context.Background(), receiver.Receive))
}
