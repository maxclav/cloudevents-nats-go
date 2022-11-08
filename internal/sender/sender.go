package sender

import (
	"context"
	"fmt"

	nats "github.com/cloudevents/sdk-go/protocol/nats/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"

	"github.com/maxclav/cloudevents-nats-go/internal/event"
)

func Start(ctx context.Context, url, subject string) error {
	nats, err := nats.NewSender(url, subject, nats.NatsOptions())
	if err != nil {
		return err
	}
	defer nats.Close(ctx)

	cl, err := cloudevents.NewClient(nats)
	if err != nil {
		return err
	}

	fmt.Println("Sender started!")
	for {
		fmt.Print("Enter your message: ")
		var input string
		fmt.Scanln(&input)
		e := event.New(input)
		result := cl.Send(ctx, e)
		var resultStr string
		if cloudevents.IsUndelivered(result) {
			resultStr = "failed"
		} else {
			resultStr = "succeeded"
		}
		fmt.Printf("Sending of %s %s\n", input, resultStr)
	}
}
