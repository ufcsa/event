package models

import (
	"time"
)

// User : ORM struct for table "user"
type User struct {
	ID      int       `xorm:"not null pk autoincr INT(10)"`
	UUID    string    `xorm:"not null default '' comment('UUID') unique(user_key) VARCHAR(10)"`
	Name    string    `xorm:"not null default '' comment('full name') unique(user_key) VARCHAR(30)"`
	Created time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('create time') DATETIME"`
	Updated time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('update time') DATETIME"`
}
