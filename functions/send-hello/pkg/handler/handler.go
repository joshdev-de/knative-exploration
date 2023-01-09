package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"github.com/joshdev-de/knative-exploration/functions/send-hello/pkg/model"
)

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	client cloudevents.Client
	sink   string
}

func New() Handler {
	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	sink := os.Getenv("K_SINK")
	if sink == "" {
		log.Fatal("env var K_SINK not set")
	}

	return &handler{
		client: c,
		sink:   sink,
	}
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	log.Print("received a request")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body := &model.SendHelloRequest{}
	if errJson := json.NewDecoder(r.Body).Decode(body); errJson != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event := cloudevents.NewEvent()
	event.SetID(uuid.New().String())
	event.SetSource("dev.knative.samples/send-hello-source")
	event.SetType("dev.knative.samples.hello.name")
	event.SetData(cloudevents.ApplicationJSON, model.HelloName{Name: body.Name})

	ctx := cloudevents.ContextWithTarget(context.Background(), h.sink)

	// Send that Event.
	if result := h.client.Send(ctx, event); cloudevents.IsUndelivered(result) {
		log.Printf("failed to send, %v\n", result)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
