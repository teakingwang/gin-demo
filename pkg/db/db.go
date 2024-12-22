package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/teakingwang/gin-demo/config"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	var gdb *gorm.DB

	c := &config.Config.Database

	switch Dialect(c.Dialect) {
	case Postgres:
		gdb = NewPostgres(c)
	default:
		return nil, fmt.Errorf("database not support: %q", c.Dialect)
	}

	return gdb, nil
}
