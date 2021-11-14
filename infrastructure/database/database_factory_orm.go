package database

import "errors"

const (
	InstanceGorm int = iota
	InstanceXorm int = iota
)

var (
	errInvalidOrmDatabaseInstance = errors.New("invalid orm database instance")
	errNotImplementedOrmDatabaseIntance = errors.New("orm database not yet implemented")
)

func NewDatabaseOrmFactory(instance int) (interface{}, error)  {
	switch instance {
	case InstanceGorm:
		return NewGormHandler(NewConfigPostgres())
	case InstanceXorm:
		return nil, errNotImplementedOrmDatabaseIntance
	default:
		return nil, errInvalidOrmDatabaseInstance
	}
}
