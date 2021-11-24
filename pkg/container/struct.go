package container

import (
	"database/sql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	awsV1 "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"go.uber.org/zap"
)

// Container is the container used to inject dependencies into the flow
type Container struct {
	awsConfig              aws.Config
	awsConfigV1            awsV1.Config
	DB                   *sql.DB
	CloudwatchClient     *cloudwatch.Client
	EventBridgeClient    *eventbridge.Client
	DynamoDBClient       *dynamodb.Client
	DynamoDBORM          *dynamo.DB
	Session 			 *session.Session
	S3Client             *s3.Client
	Logger               *zap.SugaredLogger
	SecretsManagerClient *secretsmanager.Client
}
