module github.com/macai-project/framework-golang

go 1.16

//replace github.com/macai-project/events => ../events

require (
	github.com/macai-project/events v0.0.4
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.19.1
)
