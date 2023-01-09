package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	client cloudevents.Client
}

func New() Handler {
	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	return &handler{
		client: c,
	}
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	log.Print("received a request")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", target)

	event := cloudevents.NewEvent()
	event.SetSource("dev.knative.samples/helloworldsource")
	event.SetType("dev.knative.samples.helloworld")
	event.SetData(cloudevents.ApplicationJSON, map[string]string{"msg": "world"})

	sink := os.Getenv("K_SINK")

	ctx := cloudevents.ContextWithTarget(context.Background(), sink)

	// Send that Event.
	if result := h.client.Send(ctx, event); cloudevents.IsUndelivered(result) {
		log.Fatalf("failed to send, %v", result)
	}
}
