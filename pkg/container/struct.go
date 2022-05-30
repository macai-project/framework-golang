package container

import (
	"database/sql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/xray"
	awsV1 "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"go.uber.org/zap"
)

// Container is the container used to inject dependencies into the flow
type Container struct {
	AwsConfig            aws.Config
	AwsConfigV1          awsV1.Config
	DB                   *sql.DB
	CloudwatchClient     *cloudwatch.Client
	CloudsearchClient    *cloudsearchdomain.Client
	EventBridgeClient    *eventbridge.Client
	AppSyncClient        *appsync.Client
	DynamoDBClient       *dynamodb.Client
	DynamoDBORM          *dynamo.DB
	Session              *session.Session
	S3Client             *s3.Client
	LambdaClient         *lambda.Client
	Logger               *zap.SugaredLogger
	SecretsManagerClient *secretsmanager.Client
	PinpointClient       *pinpoint.Client
	SESClient            *ses.Client
	SQSClient            *sqs.Client
	XRayClient           *xray.Client
}
