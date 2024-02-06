package models

import (
	"database/sql"
)

type Session struct {
	Id        string       `gorm:"primaryKey;column:id;->;<-:create"`
	Username  string       `gorm:"column:username;unique;size:256;default:'';"`
	ExpiredAt sql.NullTime `gorm:"column:expired_at;"`
}
