package models

import (
  "github.com/jinzhu/gorm"
)

type User struct {
  gorm.Model
  Name string  `gorm:"size:255"`
  Email string `gorm:"type:varchar(100);unique_index"`
}
