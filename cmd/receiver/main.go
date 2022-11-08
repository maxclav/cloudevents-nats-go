package main

import (
	"context"
	"fmt"

	"github.com/maxclav/cloudevents-nats-go/internal/nats"
	"github.com/maxclav/cloudevents-nats-go/internal/receiver"
)

func main() {
	fmt.Println("Starting receiver...")
	if err := receiver.Start(context.Background(), nats.URL, nats.Subject); err != nil {
		panic(err)
	}
}
