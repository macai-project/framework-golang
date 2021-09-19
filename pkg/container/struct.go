package container

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"go.uber.org/zap"
)

// Container is the container used to inject dependencies into the flow
type Container struct {
	Logger               *zap.SugaredLogger
	EventBridgeClient    *eventbridge.Client
	SecretsManagerClient *secretsmanager.Client
	AwsConfig            aws.Config
}
