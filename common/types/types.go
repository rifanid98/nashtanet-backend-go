package types

type Point struct {
	X string `json:"x" validate:"numeric"`
	Y string `json:"T" validate:"numeric"`
}