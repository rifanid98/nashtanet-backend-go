package entity

import (
	"time"
)

type Role struct {
	Id        int64      `gorm:"primaryKey; autoIncrement" json:"id,omitempty" validate:"min=1"`
	Name      *string    `gorm:"type:varchar" json:"name,omitempty" validate:"min=3,max=300"`
	CreatedAt *time.Time `gorm:"type:timestamptz;not null;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"type:timestamptz;null" json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"type:timestamptz;null" json:"deleted_at,omitempty"`
}
