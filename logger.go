package zapwrap

import (
	"fmt"
	"go.uber.org/zap"
	"github.com/aflogger-go/logger"
)

func New(config *logger.Config) (*zapWrapper, error) {
	conf := zap.NewProductionConfig()
	if config.Debug {
		conf.Development = true
		conf.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	log, err := conf.Build()
	if err != nil {
		return nil, err
	}

	var fields []zap.Field
	for k, v := range config.ConstFields {
		fields = append(fields, zap.String(k, v))
	}

	return NewWrap(log.With(fields...)), nil
}

func NewWrap(instance *zap.Logger) *zapWrapper {
	return &zapWrapper{
		instance: instance,
	}
}

type zapWrapper struct {
	instance *zap.Logger
}

func (log *zapWrapper) getMessage(v ...interface{}) (sentence string) {
	for i := range v {
		var msg string
		switch t := v[i].(type) {
		case string:
			msg = t
		case error:
			msg = t.Error()
		default:
			msg = fmt.Sprint(v[i])
		}

		if sentence != "" {
			sentence += " " + msg
		} else {
			sentence = msg
		}
	}

	return sentence
}

func (log *zapWrapper) withLvl(lvl logger.Level) *zap.Logger {
	return log.instance.With(zap.String("log-lvl", lvl.String()))
}

func (log *zapWrapper) Debug(v ...interface{}) {
	log.withLvl(logger.Debug).Debug(log.getMessage(v))
	_ = log.instance.Sync()
}

func (log *zapWrapper) Debugf(format string, v ...interface{}) {
	log.withLvl(logger.Debug).Debug(fmt.Sprintf(format, v))
	_ = log.instance.Sync()
}

func (log *zapWrapper) Info(v ...interface{}) {
	log.withLvl(logger.Info).Info(log.getMessage(v))
	_ = log.instance.Sync()
}

func (log *zapWrapper) Infof(format string, v ...interface{}) {
	log.withLvl(logger.Info).Info(fmt.Sprintf(format, v))
	_ = log.instance.Sync()
}

func (log *zapWrapper) Warn(v ...interface{}) {
	log.withLvl(logger.Warn).Info(log.getMessage(v))
	_ = log.instance.Sync()
}

func (log *zapWrapper) Warnf(format string, v ...interface{}) {
	log.withLvl(logger.Warn).Info(fmt.Sprintf(format, v))
	_ = log.instance.Sync()
}

func (log *zapWrapper) Err(v ...interface{}) {
	log.withLvl(logger.Err).Error(log.getMessage(v))
	_ = log.instance.Sync()
}

func (log *zapWrapper) Errf(format string, v ...interface{}) {
	log.withLvl(logger.Err).Error(fmt.Sprintf(format, v))
	_ = log.instance.Sync()
}

func (log *zapWrapper) Crit(v ...interface{}) {
	log.withLvl(logger.Crit).Error(log.getMessage(v))
	_ = log.instance.Sync()
}

func (log *zapWrapper) Critf(format string, v ...interface{}) {
	log.withLvl(logger.Crit).Error(fmt.Sprintf(format, v))
	_ = log.instance.Sync()
}

func (log *zapWrapper) Level(lvl logger.Level, v ...interface{}) {
	log.withLvl(lvl).Info(log.getMessage(v))
	_ = log.instance.Sync()
}

func (log *zapWrapper) Levelf(lvl logger.Level, format string, v ...interface{}) {
	log.withLvl(lvl).Info(fmt.Sprintf(format, v))
	_ = log.instance.Sync()
}

var _ logger.Logger = (*zapWrapper)(nil)