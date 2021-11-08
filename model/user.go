package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `gorm:"column:id;primary_key"  json:"id"`
	FullName  string         `gorm:"column:full_name"  json:"full_name"`
	Email     string         `gorm:"column:email"  json:"email"`
	Password  string         `gorm:"column:password"  json:"password"`
	Age       int            `gorm:"column:age"  json:"age"`
	Avatar    string         `gorm:"column:avatar" json:"avatar"`
	Role      string         `gorm:"column:role"  json:"role"`
	CreatedAt time.Time      `gorm:"column:created_at"  json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"  json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"  json:"deleted_at"`
}
