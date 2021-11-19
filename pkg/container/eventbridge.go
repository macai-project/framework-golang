package container

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
)

// NewEventbridgeClient create a new Eventbridge client
func (c *Container) NewEventbridgeClient() {
	if c.EventBridgeClient == nil {
		c.EventBridgeClient = eventbridge.NewFromConfig(c.awsConfig)
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
