package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joshdev-de/knative-exploration/functions/send-hello/pkg/handler"
)

func main() {
	log.Print("Send Hello started")

	h := handler.New()

	http.HandleFunc("/", h.Handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
