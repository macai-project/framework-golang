module github.com/macai-project/framework-golang

go 1.16

//replace github.com/macai-project/events => ../events

require (
	github.com/aws/aws-sdk-go-v2 v1.9.1
	github.com/aws/aws-sdk-go-v2/config v1.8.2
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.7.1
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.6.1
	github.com/macai-project/events v0.0.4
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.19.1
)
