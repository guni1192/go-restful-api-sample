package models

import (
  "time"
)

type Model struct {
  ID          uint       `gorm:"primary_key" json:"id"`
  CreatedAt   time.Time  `json:"created_at"`
  UpdatedAt   time.Time  `json:"updated_at"`
}

type User struct {
  Model
  Name  string  `gorm:"size:255" json:"name" form:"name"`
  Email string  `gorm:"type:varchar(100);unique_index" json:"email" form:"name"`
}
