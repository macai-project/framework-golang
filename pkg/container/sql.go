package container

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"

	"github.com/aws/aws-xray-sdk-go/xray"
	_ "github.com/go-sql-driver/mysql"
)

func (c *Container) NewSqlClient(ctx context.Context) error {
	var err error

	// If the connection already exists return nil and do nothing
	if c.DB != nil && c.DB.PingContext(ctx) == nil {
		c.Logger.With("Stats", c.DB.Stats()).Debug("Reusing SQL Connection")
		return nil
	}

	// Open SQL Connection
	c.DB, err = xray.SQLContext("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", os.Getenv("AURORA_USERNAME"), os.Getenv("AURORA_PASSWORD"), os.Getenv("AURORA_HOSTNAME"), os.Getenv("AURORA_DATABASE")))
	if err != nil {
		return fmt.Errorf("error in SQL Connection: %w", err)
	}
	c.Logger.Debug("SQL Connection Opened")

	c.DB.SetMaxOpenConns(2)
	c.DB.SetMaxIdleConns(2)
	c.DB.SetConnMaxLifetime(30 * time.Minute)

	// Check connection
	err = c.DB.PingContext(ctx)
	if err != nil {
		c.DB.Close()
		return fmt.Errorf("error in SQL Ping: %w", err)
	}
	c.Logger.With(
		"Stats", c.DB.Stats(),
	).Info("SQL Connected!")

	return err
}

func (c *Container) NewGormDBClient() error {
	var err error

	// If the connection already exists return nil and do nothing
	if c.GormDB != nil {
		c.Logger.With("Stats", c.DB.Stats()).Debug("Reusing GormDB Connection")
		return nil
	}

	// Open SQL Connection
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", os.Getenv("AURORA_USERNAME"), os.Getenv("AURORA_PASSWORD"), os.Getenv("AURORA_HOSTNAME"), os.Getenv("AURORA_DATABASE")))
	if err != nil {
		return fmt.Errorf("error in SQL Connection: %w", err)
	}
	c.Logger.Debug("SQL Connection Opened")

	c.DB.SetMaxOpenConns(2)
	c.DB.SetMaxIdleConns(2)
	c.DB.SetConnMaxLifetime(30 * time.Minute)

	// Check connection
	err = conn.Ping()
	if err != nil {
		c.DB.Close()
		return fmt.Errorf("error in SQL Ping: %w", err)
	}

	c.GormDB, err = gorm.Open(mysql.New(mysql.Config{Conn: conn}), &gorm.Config{})
	if err != nil {
		c.DB.Close()
		return fmt.Errorf("error opening Gorm connection: %w", err)
	}

	c.Logger.With(
		"Stats", c.DB.Stats(),
	).Info("Gorm SQL Connected!")

	return err
}
