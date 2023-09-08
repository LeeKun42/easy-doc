package model

import (
	"context"
	"easy-doc/app/lib/log"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type DatabaseConfig struct {
	Type     string
	Host     string
	Port     int32
	User     string
	Password string
	Name     string
	Charset  string
}

func Instance() *gorm.DB {
	var dbConfig DatabaseConfig
	viper.UnmarshalKey("database.default", &dbConfig)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name, dbConfig.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: New().LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Printf("contenct to db error, %v", err)
	}
	sqlDB, err := db.DB()
	//设置连接最大生存时间
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)
	return db
}

type dbLog struct {
	LogLevel logger.LogLevel
}

func New() *dbLog {
	return new(dbLog)
}

func (l *dbLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *dbLog) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}
	log.Infof(msg, data)
}

func (l *dbLog) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	log.Warnf(msg, data)
}

func (l *dbLog) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}
	log.Errorf(msg, data)
}

func (l *dbLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//这块的逻辑可以自己根据业务情况修改
	elapsed := time.Since(begin)
	sql, rows := fc()
	log.Infof("sql: %v  row： %v  elapsed: %d ms  err: %v", sql, rows, elapsed.Milliseconds(), err)
}
