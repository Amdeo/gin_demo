package model

import (
	"gorm.io/gorm"
	"time"
)

type Public struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	Public
	Name  string `gorm:"comment:姓名"`
	Phone string `gorm:"size:32;unique_index;not null;comment:手机号码"`
	Email string `gorm:"type:varchar(100);comment:电子邮箱"`
	Count int    `gorm:"comment:登录次数"`
	CompanyID int
	Company   Company
}

type Company struct {
	ID   int
	Name string
}