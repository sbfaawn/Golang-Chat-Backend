package models

import (
	"time"
)

type Session struct {
	Id       string
	Username string
	TTL      time.Duration
}
