package router

import "errors"

const (
	InstanceEchoGorm int = iota
	InstanceEchoXorm int = iota
)

var (
	errInvalidOrmRouterInstance       = errors.New("invalid orm router instance")
	errNotImplementedOrmRouterIntance = errors.New("orm router not yet implemented")
)

type RouterOrmFactoryOptions struct {
	EchoEngineGorm *EchoEngineGormOptions
}

func NewRouterOrmFactory(instance int, options *RouterOrmFactoryOptions) (Server, error) {
	switch instance {
	case InstanceEchoGorm:
		return NewEchoServerGorm(options.EchoEngineGorm), nil
	case InstanceEchoXorm:
		return nil, errNotImplementedOrmRouterIntance
	default:
		return nil, errInvalidOrmRouterInstance
	}
}
