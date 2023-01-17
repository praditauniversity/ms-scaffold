package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Scaffold struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:text;not null"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	UpdatedBy   string `json:"updated_by"`
	DeletedBy   string `json:"deleted_by"`
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

func (b *Scaffold) TableName() string {
	return "scaffold"
}
