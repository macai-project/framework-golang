package framework

import (
	"context"
	"encoding/json"
	"github.com/macai-project/framework-golang/pkg/container"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func businessLogic(ctx context.Context, c *container.Container, e events.CloudWatchEvent) (string, error) {
	c.NewSqlClient(ctx)
	return "ok", nil
}

func TestHandleRequest(t *testing.T) {

	RegisterBusinessLogic(businessLogic)

	inputJson := `{
		"version": "0",
		"id": "12345678-1234-1234-1234-123456789012",
		"detail-type": "itemPlaced",
		"source": "com.macaiapp.warehouse.staging",
		"account": "123456789012",
		"time": "1970-01-01T00:00:00Z",
		"region": "eu-west-1",
		"detail": {
			"Ean":"12345",
			"Name":"Ciao",
			"Quantity":12,
			"ShelfId":"12",
			"ShelfName":"scaffale12",
			"CityId":"allos",
			"WarehouseId":"asd23"
		}
	}`

	var inputEvent events.CloudWatchEvent
	if err := json.Unmarshal([]byte(inputJson), &inputEvent); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	ctx := context.Background()

	s, err := HandleRequest(ctx, inputEvent)

	assert.Nil(t, err)
	assert.Equal(t, s, "ok")
}
