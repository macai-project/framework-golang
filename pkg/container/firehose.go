package container

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
)

// NewFirehoseClient create a new Firehose client
func (c *Container) NewFirehoseClient() {
	if c.FirehoseClient == nil {
		c.FirehoseClient = firehose.NewFromConfig(c.AwsConfig)
	}
}

func (c *Container) FirehoseInput(event interface{}, deliveryStreamName string) (*firehose.PutRecordInput, error) {

	eventJson, err := json.Marshal(event)
	
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling event: %s", err.Error())
	}
	return &firehose.PutRecordInput{
		DeliveryStreamName: aws.String(deliveryStreamName),
		Record:             &types.Record{Data: eventJson},
	}, nil
}
