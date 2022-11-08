package event

import (
	"time"

	"github.com/google/uuid"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

const (
	eventType   string = "com.maxclav.test.sent"
	eventSource string = "https://github.com/maxclav/cloudevents-nats-go"
)

func New(str string) cloudevents.Event {
	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetType(eventType)
	e.SetTime(time.Now())
	e.SetSource(eventSource)
	_ = e.SetData(ContentType, NewContent(str))

	return e
}

type Content struct {
	Message string
}

const ContentType string = "application/json"

func NewContent(msg string) *Content {
	return &Content{
		Message: msg,
	}
}
