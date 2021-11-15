package entity

import (
	"time"
)

type User struct {
	Id       int64   `gorm:"primaryKey; autoIncrement" json:"id,omitempty" validate:"min=1"`
	Name     *string `gorm:"type:varchar" json:"name,omitempty" validate:"min=3,max=300"`
	Email    *string `gorm:"type:varchar" json:"email,omitempty" validate:"email"`
	Password *string `gorm:"type:varchar" json:"password,omitempty" validate:"min=6,max=32"`
	//Gender                 *enum.Gender        `gorm:"type:gender" json:"gender,omitempty" sql:"type:gender"`
	BirthPlace *string    `gorm:"type:varchar" json:"birth_place,omitempty"`
	BirthDate  *time.Time `gorm:"type:date" json:"birth_date,omitempty"`
	Phone      *string    `gorm:"type:varchar" json:"phone,omitempty" validate:"min=11"`
	//MaritalStatus          *enum.MaritalStatus `gorm:"type:marital_status" json:"marital_status,omitempty"`
	//Religion               *enum.Religion      `gorm:"" json:"religion,omitempty" sql:"type:religion"`
	IdentityNumber *string `gorm:"type:varchar" json:"identity_number,omitempty"`
	//IdentityType           *enum.IdentityType  `gorm:"" json:"identity_type,omitempty" sql:"type:identity_type"`
	IdentityExpirationDate *time.Time `gorm:"type:varchar" json:"identity_expiration_date,omitempty"`
	IdentityPostalCode     *string    `gorm:"type:varchar" json:"identity_postal_code,omitempty"`
	Address                *string    `gorm:"type:varchar" json:"address,omitempty"`
	PostalCode             *string    `gorm:"type:varchar" json:"postal_code,omitempty"`
	//BloodType              *enum.BloodType     `gorm:"" json:"blood_type,omitempty" sql:"type:blood_type"`
	Authenticator       bool       `gorm:"type:boolean" json:"authenticator,omitempty"`
	AuthenticatorSecret *string    `gorm:"type:varchar" json:"authenticator_secret,omitempty"`
	CreatedAt           *time.Time `gorm:"type:timestamptz;not null;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt           *time.Time `gorm:"type:timestamptz;null" json:"updated_at,omitempty"`
	DeletedAt           *time.Time `gorm:"type:timestamptz;null" json:"deleted_at,omitempty"`
}

func NewUser() *User {
	return &User{}
}
