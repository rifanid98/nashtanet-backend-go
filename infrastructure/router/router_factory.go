package router

import "errors"

const (
	InstanceRouterOrm    int = iota
	InstanceRouterNative int = iota
)

type Server interface {
	Listen()
}

type Port int64


var (
	errInvalidRouterInstance        = errors.New("invalid router instance")
	errNotImplementedRouterInstance = errors.New("native router instance not yet implemented")
)

type RouterFactoryOptions struct {
	RouterOrmOptions *RouterOrmFactoryOptions
}

func NewRouterFactory(routerType int, routerInstance int, options *RouterFactoryOptions) (Server, error) {
	switch routerType {
	case InstanceRouterOrm:
		return NewRouterOrmFactory(routerInstance, options.RouterOrmOptions)
	case InstanceRouterNative:
		return nil, errNotImplementedRouterInstance
	default:
		return nil, errInvalidRouterInstance
	}
}
