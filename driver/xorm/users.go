package models

import (
	"time"
)

type Users struct {
	Name      string    `xorm:"default 'NULL' VARCHAR(255)"`
	Age       int       `xorm:"default NULL INT(11)"`
	Id        string    `xorm:"not null pk CHAR(36)"`
	CreatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
	DeletedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
	ItemId    []byte    `xorm:"default NULL VARBINARY(255)"`
}
