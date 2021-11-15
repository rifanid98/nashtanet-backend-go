package enum

import "database/sql/driver"

var maritalStatus []interface{}

type MaritalStatus string

const (
	Married MaritalStatus = "married"
	Single  MaritalStatus = "single"
	Widow   MaritalStatus = "widow"
	Widower MaritalStatus = "widower"
)

func (ms *MaritalStatus) Scan(value interface{}) error {
	*ms = MaritalStatus(value.([]byte))
	return nil
}

func (ms MaritalStatus) Value() (driver.Value, error) {
	return string(ms), nil
}

func (MaritalStatus) TableName() string {
	return "gender"
}

func setMaritalStatusEnum() {
	maritalStatus = make([]interface{}, 0)
	maritalStatus = append(maritalStatus, "married")
	maritalStatus = append(maritalStatus, "single")
	maritalStatus = append(maritalStatus, "widow")
	maritalStatus = append(maritalStatus, "widower")
}
