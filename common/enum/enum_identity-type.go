package enum

import "database/sql/driver"

var identityTypes []interface{}

type IdentityType string

const (
	KTP   IdentityType = "ktp"
	KK    IdentityType = "kk"
	SIM   IdentityType = "simi"
	OTHER IdentityType = "other"
)

func (g *IdentityType) Scan(value interface{}) error {
	*g = IdentityType(value.([]byte))
	return nil
}

func (p IdentityType) Value() (driver.Value, error) {
	return string(p), nil
}

func (IdentityType) TableName() string {
	return "identityType"
}

func setIdentityTypesEnum() {
	identityTypes = make([]interface{}, 0)
	identityTypes = append(identityTypes, "ktp")
	identityTypes = append(identityTypes, "kk")
	identityTypes = append(identityTypes, "sim")
	identityTypes = append(identityTypes, "other")
}
