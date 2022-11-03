package domain

import (
	"gorm.io/gorm"
)

type Scaffold struct {
	gorm.Model
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	UserID     uint64    `json:"user_id" gorm:"type:uint;not null"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
}
