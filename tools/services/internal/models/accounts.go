package models

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	gorm.Model
	Email string
	Password string
	Nickname string
	LastLoginTime time.Time

	BiliLiveRoomID string
}
