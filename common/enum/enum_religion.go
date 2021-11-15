package enum

import "database/sql/driver"

var religions []interface{}

type Religion string

const (
	Islam        Religion = "islam"
	Protestant   Religion = "protestant"
	Catholic     Religion = "catholic"
	Hindu        Religion = "hindu"
	Buddha       Religion = "buddha"
	Confucianism Religion = "confucianism"
	Other        Religion = "other"
)

func (g *Religion) Scan(value interface{}) error {
	*g = Religion(value.([]byte))
	return nil
}

func (p Religion) Value() (driver.Value, error) {
	return string(p), nil
}

func (Religion) TableName() string {
	return "religion"
}

func setReligionsEnum() {
	religions = make([]interface{}, 0)
	religions = append(religions, "buddha")
	religions = append(religions, "catholic")
	religions = append(religions, "confucianism")
	religions = append(religions, "hindu")
	religions = append(religions, "islam")
	religions = append(religions, "other")
}
