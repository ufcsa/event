package models

import (
	"time"
)

// Result : ORM struct for table "result"
type Result struct {
	ID         int       `xorm:"not null pk autoincr INT(5)"`
	PrizeID    int       `xorm:"not null default 0 comment('Prize id') index INT(5)"`
	PrizeName  string    `xorm:"not null default '' comment('prize name') VARCHAR(255)"`
	UserID     string    `xorm:"not null default '' comment('user uuid') index VARCHAR(255)"`
	LeftNum    int       `xorm:"not null default 0 comment('number of prize left') INT(5)"`
	Verifycode int       `xorm:"not null default 0 comment('4 digit random number as verification code') INT(4)"`
	ClientIP   string    `xorm:"not null default '' comment('user's ip') VARCHAR(50)"`
	Created    time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('create time') DATETIME"`
	Updated    time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('update time') DATETIME"`
}
