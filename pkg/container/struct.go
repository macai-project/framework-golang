package container

import (
	"database/sql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"go.uber.org/zap"
)

// Container is the container used to inject dependencies into the flow
type Container struct {
	awsConfig            aws.Config
	Logger               *zap.SugaredLogger
	DB                   *sql.DB
	CloudwatchClient     *cloudwatch.Client
	EventBridgeClient    *eventbridge.Client
	S3Client             *s3.Client
	SecretsManagerClient *secretsmanager.Client
	DynamoDbClient       *dynamodb.Client
}
