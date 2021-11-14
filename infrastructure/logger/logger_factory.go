package logger

import "errors"

var (
	errInvalidLoggerInstance = errors.New("Invalid logger instance")
)

const (
	InstanceZapLogger int = iota
)

func NewLoggerFactory(instance int) (Logger, error)  {
	switch instance {
	case InstanceZapLogger:
		return NewZapLogger()
	default:
		return nil, errInvalidLoggerInstance
	}
}