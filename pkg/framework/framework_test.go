package framework

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/macai-project/events"
	"github.com/macai-project/framework-golang/internal/container"
	"github.com/stretchr/testify/assert"
)

func businessLogic(ctx context.Context, c *container.Container, e events.TestEvent) (string, error) {
	return "ok", nil
}

func TestHandleRequest(t *testing.T) {

	RegisterBusinessLogic(businessLogic)

	inputJson := `{
		"version": "0",
		"id": "12345678-1234-1234-1234-123456789012",
		"detail-type": "EC2 Instance Launch Successful",
		"source": "aws.autoscaling",
		"account": "123456789012",
		"time": "1970-01-01T00:00:00Z",
		"region": "us-west-2",
		"resources": [
			"auto-scaling-group-arn",
			"instance-arn"
		],
		"detail": {
			"StatusCode": "InProgress",
			"Description": "Launching a new EC2 instance: i-12345678",
			"AutoScalingGroupName": "my-auto-scaling-group",
			"ActivityId": "87654321-4321-4321-4321-210987654321",
			"Details": {
				"Availability Zone": "us-west-2b",
				"Subnet ID": "subnet-12345678"
			},
			"RequestId": "12345678-1234-1234-1234-123456789012",
			"StatusMessage": "",
			"EndTime": "1970-01-01T00:00:00Z",
			"EC2InstanceId": "i-1234567890abcdef0",
			"StartTime": "1970-01-01T00:00:00Z",
			"Cause": "description-text"
		}
	}`

	var inputEvent events.TestEvent
	if err := json.Unmarshal([]byte(inputJson), &inputEvent); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	c := container.Container{}
	c.NewLogger()
	ctx := context.Background()
	string, err := HandleRequest(ctx, &c, inputEvent)
	assert.Nil(t, err)
	assert.Equal(t, string, "ok")
}
