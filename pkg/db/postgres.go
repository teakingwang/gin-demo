package db

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/teakingwang/gin-demo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
)

type PGDBConfig struct {
	User     string
	Password string
	DBName   string
	DBSchema string
	Host     string
	Port     int
	Debug    bool
	LogLevel logger.LogLevel
}

func NewPostgres(c *config.DatabaseConfig) *gorm.DB {
	pDB, err := newPGDBWithLevel(
		c.User,
		c.Password,
		c.Database,
		c.Schema,
		c.Host,
		c.Port,
		c.Level,
	)
	if err != nil {
		logrus.Errorf("failed to connect database, %+v", err)
		return nil
	}

	return pDB
}

func newPGDBWithLevel(user, password, db, schema, host, port, level string) (*gorm.DB, error) {
	connURL := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s search_path=%s sslmode=disable",
		host, port, user, db, password, schema)

	var logLevel logger.LogLevel
	switch strings.ToLower(level) {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	default:
		logLevel = logger.Info
	}

	pgdb, err := gorm.Open(postgres.Open(connURL), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, err
	}

	return pgdb, nil
}
