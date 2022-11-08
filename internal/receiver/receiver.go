package receiver

import (
	"context"
	"fmt"

	nats "github.com/cloudevents/sdk-go/protocol/nats/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"

	"github.com/maxclav/cloudevents-nats-go/internal/event"
)

func Start(ctx context.Context, url, subject string) error {
	nats, err := nats.NewConsumer(url, subject, nats.NatsOptions())
	if err != nil {
		return err
	}
	defer nats.Close(ctx)

	cl, err := cloudevents.NewClient(nats)
	if err != nil {
		return err
	}

	fmt.Println("Receiver started!")
	for {
		if err := cl.StartReceiver(ctx, receive); err != nil {
			return err
		}
	}
}

func receive(ctx context.Context, e cloudevents.Event) error {
	fmt.Print("Received Event:")
	fmt.Printf("Context = %+v", e.Context)

	data := &event.Content{}
	if err := e.DataAs(data); err != nil {
		fmt.Printf("Error = %s\n", err.Error())
	} else {
		fmt.Printf("Data = %+v\n", data)
	}
	return nil
}
