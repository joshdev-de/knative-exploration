package receiver

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/joshdev-de/knative-exploration/functions/catch-hello/pkg/model"
)

func Receive(ctx context.Context, event cloudevents.Event) cloudevents.Result {
	// Here is where your code to process the event will go.
	// In this example we will log the event msg
	log.Printf("Event received. \n%s\n", event)
	data := &model.HelloMessage{}
	if err := event.DataAs(data); err != nil {
		log.Printf("Error while extracting cloudevent Data: %s\n", err.Error())
		return cloudevents.NewHTTPResult(400, "failed to convert data: %s", err)
	}
	log.Printf("Received message: %s\n", data.Msg)

	return nil
}
