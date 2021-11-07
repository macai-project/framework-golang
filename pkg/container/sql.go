package container

import (
	"context"
	"fmt"
	"github.com/aws/aws-xray-sdk-go/xray"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func (c *Container) NewSqlClient(ctx context.Context) error {
	var err error

	// Open SQL Connection
	c.DB, err = xray.SQLContext("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", os.Getenv("AURORA_USERNAME"), os.Getenv("AURORA_PASSWORD"), os.Getenv("AURORA_HOSTNAME"), os.Getenv("AURORA_DATABASE")))
	if err != nil {
		return fmt.Errorf("error in SQL Connection: %w", err)
	}
	c.Logger.Debug("SQL Connection Opened")

	c.DB.SetMaxOpenConns(1)
	c.DB.SetMaxIdleConns(1)
	c.DB.SetConnMaxLifetime(-1)
	//c.DB.SetConnMaxIdleTime(-1)

	// Check connection
	err = c.DB.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("error in SQL Ping: %w", err)
	}
	c.Logger.Debug("SQL Connected!")

	return err
}
