package models

import (
	"time"
)

// Prize : ORM struct for table "prize"
type Prize struct {
	ID      int       `xorm:"not null pk autoincr INT(5)"`
	Title   string    `xorm:"not null default '' comment('name of the prize') VARCHAR(255)"`
	Count   int       `xorm:"not null default -1 comment('number of this prize') INT(5)"`
	LeftNum int       `xorm:"not null default 0 comment('number of prize left') INT(5)"`
	Begin   int       `xorm:"not null default 0 comment('start time') INT(11)"`
	End     int       `xorm:"not null default 0 comment('end time') INT(11)"`
	Status  int       `xorm:"not null default 0 comment('status: 0: available, 1: deleted, 2: banned') TINYINT(1)"`
	Created time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('create time') DATETIME"`
	Updated time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('update time') DATETIME"`
}
