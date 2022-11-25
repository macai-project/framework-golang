package container

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
)

// NewEventbridgeClient create a new Eventbridge client
func (c *Container) NewEventbridgeClient() {
	if c.EventBridgeClient == nil {
		c.EventBridgeClient = eventbridge.NewFromConfig(c.AwsConfig)
	}
}

func (c *Container) EventInput(event interface{}, detailType string, eventBus string, source string, resources ...string) (*eventbridge.PutEventsInput, error) {

	eventJson, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling event: %s", err.Error())
	}
	return &eventbridge.PutEventsInput{
		Entries: []types.PutEventsRequestEntry{
			{
				Detail:       aws.String(string(eventJson)),
				DetailType:   aws.String(detailType),
				EventBusName: aws.String(eventBus),
				Resources:    resources,
				Source:       aws.String(source),
			},
		},
	}, nil
}

func (c *Container) SendBatchEvents(ctx context.Context, events []interface{}, detailType string, eventBus string, source string, batchSize int) error {
	var entries []types.PutEventsRequestEntry
	counterSize := 0

	for i, event := range events {
		entry, err := c.EventRequestEntry(event, detailType, eventBus, source)
		if err != nil {
			return err
		}

		entries = append(entries, entry)

		if i-counterSize >= batchSize {
			// Limit of batch size reached, send event
			counterSize = counterSize + i
			_, err = c.EventBridgeClient.PutEvents(ctx, &eventbridge.PutEventsInput{Entries: entries})
			if err != nil {
				return fmt.Errorf("error in Eventbridge SendBatchEvents: %s", err.Error())
			}
			entries = []types.PutEventsRequestEntry{}
		}
	}

	if len(entries) > 0 {
		_, err := c.EventBridgeClient.PutEvents(ctx, &eventbridge.PutEventsInput{Entries: entries})
		if err != nil {
			return fmt.Errorf("error in Eventbridge SendBatchEvents: %s", err.Error())
		}
	}
	return nil
}

func (c *Container) EventRequestEntry(event interface{}, detailType string, eventBus string, source string, resources ...string) (types.PutEventsRequestEntry, error) {
	var requestEntry types.PutEventsRequestEntry

	eventJson, err := json.Marshal(event)
	if err != nil {
		return requestEntry, fmt.Errorf("error unmarshaling event: %s", err.Error())
	}

	requestEntry = types.PutEventsRequestEntry{
		Detail:       aws.String(string(eventJson)),
		DetailType:   aws.String(detailType),
		EventBusName: aws.String(eventBus),
		Resources:    resources,
		Source:       aws.String(source),
	}

	return requestEntry, nil
}
