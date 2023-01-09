package receiver

import (
	"context"
	"fmt"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"github.com/joshdev-de/knative-exploration/functions/process-hello/pkg/model"
)

func Receive(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	// Here is where your code to process the event will go.
	// In this example we will log the event msg
	log.Printf("Event received. \n%s\n", event)
	data := &model.HelloName{}
	if err := event.DataAs(data); err != nil {
		log.Printf("Error while extracting cloudevent Data: %s\n", err.Error())
		return nil, cloudevents.NewHTTPResult(400, "failed to convert data: %s", err)
	}
	log.Printf("Hello Name from received event %q", data.Name)

	msg := fmt.Sprintf("Hello %s", data.Name)

	// Respond with another event (optional)
	// This is optional and is intended to show how to respond back with another event after processing.
	// The response will go back into the knative eventing system just like any other event
	newEvent := cloudevents.NewEvent()
	newEvent.SetID(uuid.New().String())
	newEvent.SetSource("dev.knative.samples/process-hello-source")
	newEvent.SetType("dev.knative.samples.hello.message")
	if err := newEvent.SetData(cloudevents.ApplicationJSON, model.HelloMessage{Msg: msg}); err != nil {
		return nil, cloudevents.NewHTTPResult(500, "failed to set response data: %s", err)
	}

	// hardcoded case to trigger error in processing
	if data.Name == "error" {
		return nil, cloudevents.NewHTTPResult(500, "This is a manually triggered error")
	}

	log.Printf("Responding with event\n%s\n", newEvent)
	return &newEvent, nil
}
