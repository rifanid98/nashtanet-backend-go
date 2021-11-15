package enum

import "database/sql/driver"

var bloodType []interface{}

type BloodType string

const (
	A           BloodType = "A"
	A_POSITIVE  BloodType = "A+"
	A_NEGATIVE  BloodType = "A-"
	B           BloodType = "B"
	B_POSITIVE  BloodType = "B+"
	B_NEGATIVE  BloodType = "B-"
	O           BloodType = "O"
	O_POSITIVE  BloodType = "O+"
	O_NEGATIVE  BloodType = "O-"
	AB          BloodType = "AB"
	AB_POSITIVE BloodType = "AB+"
	AB_NEGATIVE BloodType = "AB-"
)

func (g *BloodType) Scan(value interface{}) error {
	*g = BloodType(value.([]byte))
	return nil
}

func (p BloodType) Value() (driver.Value, error) {
	return string(p), nil
}

func (BloodType) TableName() string {
	return "bloodType"
}

func setBloodTypeEnum() {
	bloodType = make([]interface{}, 0)
	bloodType = append(bloodType, "A")
	bloodType = append(bloodType, "A+")
	bloodType = append(bloodType, "A-")
	bloodType = append(bloodType, "B")
	bloodType = append(bloodType, "B+")
	bloodType = append(bloodType, "B-")
	bloodType = append(bloodType, "O")
	bloodType = append(bloodType, "O+")
	bloodType = append(bloodType, "O-")
	bloodType = append(bloodType, "AB")
	bloodType = append(bloodType, "AB+")
	bloodType = append(bloodType, "AB-")
}
