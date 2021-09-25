module github.com/macai-project/framework-golang

go 1.16

//replace github.com/macai-project/events => ../events

require (
	github.com/aws/aws-lambda-go v1.26.0
	github.com/aws/aws-sdk-go-v2 v1.9.0
	github.com/aws/aws-sdk-go-v2/config v1.8.1
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.7.0
	github.com/aws/aws-xray-sdk-go v1.6.0
	github.com/getsentry/sentry-go v0.11.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1
)
