package models

import (
	"time"
)

type Items struct {
	Id        string    `xorm:"not null pk CHAR(36)"`
	CreatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
	UpdatedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
	DeletedAt time.Time `xorm:"default 'NULL' TIMESTAMP"`
	Name      string    `xorm:"default 'NULL' VARCHAR(255)"`
}
