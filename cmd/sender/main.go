package main

import (
	"context"
	"fmt"

	"github.com/maxclav/cloudevents-nats-go/internal/nats"
	"github.com/maxclav/cloudevents-nats-go/internal/sender"
)

func main() {
	fmt.Println("Starting sender...")
	if err := sender.Start(context.Background(), nats.URL, nats.Subject); err != nil {
		panic(err)
	}
}
