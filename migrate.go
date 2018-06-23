package main

import (
  "log"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "models"
)


func main() {
  db, err := gorm.Open("postgres", "user=postgres dbname=go-rest-api sslmode=disable")

  if err != nil {
    log.Fatal("DB Connection Error: ", err)
  }
  defer db.Close()

  db.CreateTable(&models.User{})
  // var users []User
  // db.Find(&users)
  // log.Println(users)
}
