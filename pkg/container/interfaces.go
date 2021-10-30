package container

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
)

type eventClient interface {
	PutEvents(context.Context, *eventbridge.PutEventsInput, ...func(*eventbridge.Options)) (*eventbridge.PutEventsOutput, error)
}
