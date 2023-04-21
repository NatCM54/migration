package lib

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm/logger"

	"migrate/config"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n========================\n", sql)
}

func NewMySql() *sql.DB {
	env, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true",
		env.DbMysqlUsername,
		env.DbMysqlPassword,
		env.DbMysqlHost,
		env.DbMysqlPort,
		env.DbMysqlDatabase)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("cannot connect mysql:", err)
	}

	return db
}
