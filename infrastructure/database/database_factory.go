package database

import "errors"

const (
	InstanceOrmDatabase int = iota
	InstanceNativeDatabase int = iota
)

var (
	errInvalidDatabaseInstance = errors.New("invalid database instance")
	errNotImplementedDatabaseInstance = errors.New("native database instance not yet implemented")
)

func NewDatabaseFactory(dbType int, dbInstance int) (interface{}, error)  {
	switch dbType {
	case InstanceOrmDatabase:
		return NewDatabaseOrmFactory(dbInstance)
	case InstanceNativeDatabase:
		return nil, errNotImplementedDatabaseInstance
	default:
		return nil, errInvalidDatabaseInstance
	}
}
