package models

import (
  "log"
  "time"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
)

type User struct {
  ID int `gorm:"primary_key;AUTO_INCREMENT"`
  Name string  `gorm:"size:255"`
  Email string `gorm:"type:varchar(100);unique_index"`
  CreateAt time.Time
  UpdateAt time.Time
}
