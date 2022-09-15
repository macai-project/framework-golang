package container

import (
	"database/sql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/xray"
	awsV1 "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	appsync "github.com/sony/appsync-client-go"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Container is the container used to inject dependencies into the flow
type Container struct {
	AwsConfig            aws.Config
	AwsConfigV1          awsV1.Config
	AppSyncClient        *appsync.Client
	DB                   *sql.DB
	CloudwatchClient     *cloudwatch.Client
	CloudsearchClient    *cloudsearchdomain.Client
	CognitoClient        *cognitoidentityprovider.Client
	EventBridgeClient    *eventbridge.Client
	FirehoseClient       *firehose.Client
	DynamoDBClient       *dynamodb.Client
	DynamoDBORM          *dynamo.DB
	GormDB               *gorm.DB
	Session              *session.Session
	S3Client             *s3.Client
	LambdaClient         *lambda.Client
	Logger               *zap.SugaredLogger
	SecretsManagerClient *secretsmanager.Client
	PinpointClient       *pinpoint.Client
	SESClient            *ses.Client
	SQSClient            *sqs.Client
	SNSClient            *sns.Client
	SSMClient            *ssm.Client
	XRayClient           *xray.Client
}
