package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	Username   string         `gorm:"primaryKey;column:username;unique;size:256;default:'';"`
	Email      string         `gorm:"column:email;unique;size:256;"`
	Password   string         `gorm:"column:password;size:2000;default:'';"`
	IsVerified bool           `gorm:"column:verified;default:false"`
	CreatedAt  time.Time      `gorm:"column:created_at;->;<-:create"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;<-"`
	DeletedAt  gorm.DeletedAt `gorm:"index;column:deleted_at;<-"`
}
