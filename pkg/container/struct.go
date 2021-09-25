package container

import (
	"database/sql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"go.uber.org/zap"
)

// Container is the container used to inject dependencies into the flow
type Container struct {
	Logger            *zap.SugaredLogger
	EventBridgeClient *eventbridge.Client
	CloudwatchClient  *cloudwatch.Client
	DB                *sql.DB
	awsConfig         aws.Config
}
