# CloudEvents NATS Go (GoLang)

Simple project to test [CloudEvents](https://cloudevents.io/) with [NATS](https://docs.nats.io/) in [Go](https://go.dev/) (GoLang).

## Steps

1. Run a NATS server: `make run-nats`
2. Run a many receivers as you want: `make run-receiver`
3. Run a sender: `make run-sender`
  1. Enter text in the sender.

to check all commands: `make help`.

## Useful Links

- [NATS Protocol Binding for CloudEventts (documentation)](https://github.com/cloudevents/spec/blob/main/cloudevents/bindings/nats-protocol-binding.md)

### CloudEvents

- [CloudEvents (Python)](https://github.com/cloudevents/spec)
- [CloudEvents Go Client](https://github.com/cloudevents/sdk-go)
  - [`nats` package of CloudEvents Go Client](https://github.com/cloudevents/sdk-go/tree/main/protocol/nats/v2)
  - [samples](https://github.com/cloudevents/sdk-go/tree/main/samples)
    - [NATS sample](https://github.com/cloudevents/sdk-go/tree/main/samples/nats)

### NATS

- [NATS Server](https://github.com/nats-io/nats-server)
  - [NATS Streaming Server](https://github.com/nats-io/nats-streaming-server)
- [NATS Go Client](https://github.com/nats-io/nats.go)
  - [Example of NATS Go Client](https://github.com/nats-io/nats.go/tree/main/examples)
