package models

import "time"

// CommonModel ...
type CommonModel struct {
	ID        uint64     `json:"-" gorm:"type:bigserial; primary_key;"`
	CreatedAt time.Time  `json:"-" gorm:"type:timestamp"`
	UpdatedAt time.Time  `json:"-" gorm:"type:timestamp"`
	DeletedAt *time.Time `json:"-" gorm:"type:timestamp" sql:"index"`
}
