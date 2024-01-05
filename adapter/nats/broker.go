package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/ormushq/ormus/event"
)

// SubscribeToEvent subscribes to the specified NATS subject and processes incoming events.
func (a Adapter) SubscribeToEvent(topic string, action func(event *event.CoreEvent)) error {
	const op = "nats.SubscribeToEvent"

	// TODO - maybe we could only consider get topic from config for now

	return a.Subscribe(topic, func(msg *nats.Msg) {
		coreEvent, err := decodeCoreEvent(msg)
		if err != nil {
			// TODO - use ormus logger
			fmt.Println("Error decoding CoreEvent:", err)
			return
		}

		action(coreEvent)
	})
}

// decodeCoreEvent decodes the Protobuf message into a CoreEvent struct.
func decodeCoreEvent(msg *nats.Msg) (*event.CoreEvent, error) {
	// TODO - Implement me
	coreEvent := &event.CoreEvent{}
	return coreEvent, nil
}
