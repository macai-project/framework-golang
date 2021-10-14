package container

import (
	"context"
	"fmt"
	"github.com/aws/aws-xray-sdk-go/instrumentation/awsv2"
	"github.com/aws/aws-xray-sdk-go/xray"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func (c *Container) NewSqlClient(ctx context.Context) error {
	var err error

	// Instrumenting AWS SDK v2
	awsv2.AWSV2Instrumentor(&c.awsConfig.APIOptions)

	// Open SQL Connection
	c.DB, err = xray.SQLContext("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", os.Getenv("AURORA_USERNAME"), os.Getenv("AURORA_PASSWORD"), os.Getenv("AURORA_HOSTNAME"), os.Getenv("AURORA_DATABASE")))
	if err != nil {
		c.Logger.Fatal("Error in SQL Connection", err)
		return err
	}
	c.Logger.Debug("SQL Connection Opened")

	c.DB.SetMaxOpenConns(1)
	c.DB.SetMaxIdleConns(1)
	c.DB.SetConnMaxLifetime(-1)
	//c.DB.SetConnMaxIdleTime(-1)

	// Check connection
	err = c.DB.PingContext(ctx)
	if err != nil {
		c.Logger.Fatal("Error in SQL Ping", err)
		return err
	}
	c.Logger.Debug("SQL Connected!")

	return err
}
