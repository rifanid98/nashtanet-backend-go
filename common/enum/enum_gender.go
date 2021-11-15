package enum

import "database/sql/driver"

var gender []interface{}

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

func (g *Gender) Scan(value interface{}) error {
	*g = Gender(value.([]byte))
	return nil
}

func (p Gender) Value() (driver.Value, error) {
	return string(p), nil
}

func (Gender) TableName() string {
	return "gender"
}

func setGenderEnum() {
	gender = make([]interface{}, 0)
	gender = append(gender, "female")
	gender = append(gender, "male")
	gender = append(gender, "other")
}
