package container

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-xray-sdk-go/instrumentation/awsv2"
	"github.com/aws/aws-xray-sdk-go/xray"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func (c *Container) NewSqlClient(ctx context.Context) error {
	var err error

	// Get credentials from Env Vars if in local environment
	auroraHostname, present := os.LookupEnv("AURORA_HOSTNAME")
	if !present {
		c.Logger.Fatal("AURORA_HOSTNAME not present")
		return errors.New("AURORA_HOSTNAME not present")
	}
	auroraUsername, present := os.LookupEnv("AURORA_USERNAME")
	if !present {
		c.Logger.Fatal("AURORA_USERNAME not present")
		return errors.New("AURORA_USERNAME not present")
	}
	auroraPassword, present := os.LookupEnv("AURORA_PASSWORD")
	if !present {
		c.Logger.Fatal("AURORA_PASSWORD not present")
		return errors.New("AURORA_PASSWORD not present")
	}
	auroraDatabase, present := os.LookupEnv("AURORA_DATABASE")
	if !present {
		c.Logger.Fatal("AURORA_DATABASE not present")
		return errors.New("AURORA_DATABASE not present")
	}

	// Instrumenting AWS SDK v2
	awsv2.AWSV2Instrumentor(&c.awsConfig.APIOptions)

	// Open SQL Connection
	c.DB, err = xray.SQLContext("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", auroraUsername, auroraPassword, auroraHostname, auroraDatabase))
	if err != nil {
		c.Logger.Fatal("Error in SQL Connection", err)
		return err
	}
	c.Logger.Debug("SQL Connection Opened")

	// Check connection
	err = c.DB.PingContext(ctx)
	if err != nil {
		c.Logger.Fatal("Error in SQL Ping", err)
		return err
	}
	c.Logger.Debug("SQL Connected!")

	return err
}
