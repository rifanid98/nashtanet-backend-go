package logger

import "go.uber.org/zap"

type zapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() (Logger, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	sugar := log.Sugar()
	defer log.Sync()

	return &zapLogger{logger: sugar}, nil
}

func (z zapLogger) Infof(format string, args ...interface{}) {
	z.logger.Infof(format, args...)
}

func (z zapLogger) Warnf(format string, args ...interface{}) {
	z.logger.Warnf(format, args...)
}

func (z zapLogger) Errorf(format string, args ...interface{}) {
	z.logger.Errorf(format, args...)
}

func (z zapLogger) Fatalln(args ...interface{}) {
	z.logger.Fatal(args...)
}

func (z zapLogger) WithFields(keyValues Fields) Logger {
	var f = make([]interface{}, 0)
	for index, field := range keyValues {
		f = append(f, index)
		f = append(f, field)
	}

	log := z.logger.With(f...)
	return &zapLogger{logger: log}
}

func (z zapLogger) WithError(err error) Logger {
	var log = z.logger.With(err.Error())
	return &zapLogger{logger: log}
}
