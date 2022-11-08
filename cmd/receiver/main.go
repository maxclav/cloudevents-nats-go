package main

import (
	"context"
	"fmt"

	"github.com/maxclav/cloudevents-nats-go/internal/nats"
	"github.com/maxclav/cloudevents-nats-go/internal/receiver"
)

func main() {
	ctx := context.Background()

	fmt.Println("Starting receiver...")
	if err := receiver.Start(ctx, nats.URL, nats.Subject); err != nil {
		panic(err)
	}
}
