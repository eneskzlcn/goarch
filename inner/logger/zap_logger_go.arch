package logger

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
)

type ZapLogger interface {
	Error(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Fatal(args ...interface{})
	Errorf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Debugf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Sync() error
}

type ZapLoggerAdapter struct {
	logger ZapLogger
}

func newZapLoggerForEnv(env string, callerSkip int) (ZapLogger, error) {
	if env == "" || env == "local" || env == "test" || env == "qa" || env == "dev" {
		logger, err := zap.NewDevelopment(zap.AddCallerSkip(1))
		return logger.Sugar(), err
	} else if env == "prod" {
		logger, err := zap.NewProduction(zap.AddCallerSkip(callerSkip), zap.AddStacktrace(zap.ErrorLevel))
		return logger.Sugar(), err
	}
	return nil, errors.New("not valid")
}

func NewZapLoggerAdapter(env string, callerSkip int) *ZapLoggerAdapter {
	logger, err := newZapLoggerForEnv(env, callerSkip)

	if err != nil {
		fmt.Println("error logger is nil")
		return nil
	}
	zapLoggerAdapter := ZapLoggerAdapter{logger: logger}
	return &zapLoggerAdapter
}

func (z *ZapLoggerAdapter) StringModifier(key, value string) interface{} {
	return zap.String(key, value)
}

func (z *ZapLoggerAdapter) ErrorModifier(err error) interface{} {
	return zap.Error(err)
}

func (z *ZapLoggerAdapter) AnyModifier(key string, value any) interface{} {
	return zap.Any(key, value)
}

func (z *ZapLoggerAdapter) Info(args ...interface{}) {
	z.logger.Info(args)
}

func (z *ZapLoggerAdapter) Infof(template string, args ...interface{}) {
	z.logger.Infof(template, args)
}

func (z *ZapLoggerAdapter) Error(args ...interface{}) {
	z.logger.Error(args)
}

func (z *ZapLoggerAdapter) Errorf(template string, args ...interface{}) {
	z.logger.Errorf(template, args)
}

func (z *ZapLoggerAdapter) Debug(args ...interface{}) {
	z.logger.Debug(args)
}

func (z *ZapLoggerAdapter) Debugf(template string, args ...interface{}) {
	z.logger.Debugf(template, args)
}

func (z *ZapLoggerAdapter) Fatal(args ...interface{}) {
	z.logger.Fatal(args)
}

func (z *ZapLoggerAdapter) Fatalf(template string, args ...interface{}) {
	z.logger.Fatalf(template, args)
}

func (z *ZapLoggerAdapter) Sync() error {
	return z.logger.Sync()
}
