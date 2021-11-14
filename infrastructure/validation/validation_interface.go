package validation

import "nashtanet-backend-go/common/types"

type Validator interface {
	Validate(payload interface{}) error
	ValidatePartial(payload interface{}, fields ...string) error
	ValidatePointValue(point types.Point) error
	ValidateEnumString(enum []interface{}, value interface{}) bool
	GetFieldValue(value interface{}, field string) string
	Messages() []string
}