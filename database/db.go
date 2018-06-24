package database

import (
  "log"
  "github.com/jinzhu/gorm"
  "github.com/guni973/go-restful-api-sample/models"
)

var DB *gorm.DB

func init(){
  var err error
  DB, err = gorm.Open("postgres", "host=postgres port=5432 user=postgres dbname=go-restful password=postgres sslmode=disable")
  if err != nil {
    log.Fatal("DB Connection Error: ", err)
  }
  DB.LogMode(true)
  DB.AutoMigrate(&models.User{})
}
