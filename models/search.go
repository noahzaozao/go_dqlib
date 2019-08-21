package models

import (
	"github.com/jinzhu/gorm"
)

type SearchKeyValue struct {
	gorm.Model
	Uuid     string `gorm:"TYPE:VARCHAR(36);UNIQUE_INDEX;NOT NULL"`
	Top      int
	Key      string `gorm:"TYPE:VARCHAR(255);"`
	Value    string `gorm:"TYPE:VARCHAR(255);"`
}
