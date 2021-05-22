package model

import (
	"gorm.io/gorm"
	"time"
)

type PublicSerializer struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"CreatedAt"`
	UpdatedAt time.Time      `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `json:"DeletedAt"`
}

type UserSerializer struct {
	PublicSerializer
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Count int    `json:"count"`
}
