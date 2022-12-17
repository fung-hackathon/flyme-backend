package logger

import (
	"flyme-backend/app/config"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

type Log struct {
	Message string
	Cause   error
}

func init() {
	var logger *zap.Logger

	if config.MODE == config.Production {
		l, _ := zap.NewProduction(zap.AddCallerSkip(1))
		logger = l
	} else {
		l, _ := zap.NewDevelopment(zap.AddCallerSkip(1))
		logger = l
	}

	defer logger.Sync()
	sugar = logger.Sugar()
}

func EchoLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:    true,
		LogStatus:    true,
		LogURI:       true,
		LogRemoteIP:  true,
		LogUserAgent: true,
		LogLatency:   true,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			timestamp := time.Now()
			sugar.Infow("request",
				"timestamp", timestamp,
				"method", v.Method,
				"status", v.Status,
				"URI", v.URI,
				"remote_ip", v.RemoteIP,
				"UA", v.UserAgent,
				"latency", v.Latency)
			return nil
		},
	})
}

func (l Log) Info() {
	timestamp := time.Now()
	logMessage := string(l.Message) + ":" + l.Cause.Error()
	sugar.Infow(logMessage, "timestamp", timestamp, "message", l.Message, "cause", l.Cause)
}

func (l Log) Warn() {
	timestamp := time.Now()
	logMessage := string(l.Message) + ":" + l.Cause.Error()
	sugar.Warnw(logMessage, "timestamp", timestamp, "message", l.Message, "cause", l.Cause)
}

func (l Log) Err() {
	timestamp := time.Now()
	logMessage := string(l.Message) + ":" + l.Cause.Error()
	sugar.Errorw(logMessage, "timestamp", timestamp, "message", l.Message, "cause", l.Cause)
}
