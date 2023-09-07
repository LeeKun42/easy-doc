package log

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"time"
)

type loggerConfig struct {
	Dir        string
	Level      string
	FileFormat string `mapstructure:"file_format"`
}

func instance() *logrus.Logger {
	var conf loggerConfig
	viper.UnmarshalKey("logger", &conf)

	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/" + conf.Dir
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := now.Format(conf.FileFormat) + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logClient := logrus.New()

	//设置输出
	logClient.Out = src

	//设置日志级别
	level, _ := logrus.ParseLevel(conf.Level)
	logClient.SetLevel(level)

	//设置日志格式
	logClient.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logClient
}

func Debug(args ...interface{}) {
	instance().Debug(args)
}

func Debugf(format string, args ...interface{}) {
	instance().Debugf(format, args)
}

func Info(args ...interface{}) {
	instance().Info(args)
}

func Infof(format string, args ...interface{}) {
	instance().Infof(format, args)
}

func Warn(args ...interface{}) {
	instance().Warn(args)
}

func Warnf(format string, args ...interface{}) {
	instance().Warnf(format, args)
}

func Error(args ...interface{}) {
	instance().Error(args)
}

func Errorf(format string, args ...interface{}) {
	instance().Errorf(format, args)
}

func HttpLogHandler(ctx iris.Context) {
	instance().Infof("request url：%s	", ctx.FullRequestURI())
	instance().Info("request headers：", ctx.Request().Header)
	body, _ := ctx.GetBody()
	instance().Info("request body：", string(body))
	ctx.Next()
}
