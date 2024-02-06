package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	Id        int            `gorm:"primaryKey;column:message_id;autoIncrement;not null"`
	Sender    string         `gorm:"column:sender;"`
	Receiver  string         `gorm:"column:receiver"`
	Message   string         `gorm:"column:message;size=2500"`
	CreatedAt time.Time      `gorm:"column:created_at;->;<-:create"`
	UpdatedAt time.Time      `gorm:"column:updated_at;<-"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at;<-"`
}
