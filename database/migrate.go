package database

import (
  "log"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "github.com/guni973/go-restful-api-sample/models"
)

func main() {
  db, err := gorm.Open("postgres", "user=postgres dbname=go-rest-api sslmode=disable")

  if err != nil {
    log.Fatal("DB Connection Error: ", err)
  }
  db.LogMode(true)

  var users []models.User
  db.Find(&users)
  log.Println(users)
  defer db.Close()
}
