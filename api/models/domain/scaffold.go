package domain

import (
	"time"

	"gorm.io/gorm"
)

type Scaffold struct {
	gorm.Model
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:varchar(255);not null"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	User_id     uint64    `json:"user_id" gorm:"type:uint;not null"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
}
