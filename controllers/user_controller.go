package controllers

import (
  "net/http"
  "encoding/json"
  "log"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "github.com/guni973/go-restful-api-sample/models"
  // "github.com/guni973/go-restful-api-sample/database"
)

var DB *gorm.DB

func init(){
  var err error
  DB, err = gorm.Open("postgres", "user=postgres dbname=go-rest-api sslmode=disable")
  if err != nil {
    log.Fatal("DB Connection Error: ", err)
  }
  DB.LogMode(true)
  DB.AutoMigrate(&models.User{})
}


func UserIndex(w http.ResponseWriter, r *http.Request) {
  users := []models.User{}
  DB.Find(&users)

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  enc := json.NewEncoder(w)
  enc.SetIndent("", "    ")

  if err := enc.Encode(users); err != nil {
    log.Fatal("JSON Encode: ", err)
  }
}

func UserDetail(w http.ResponseWriter, r *http.Request) {}
func UserCreate(w http.ResponseWriter, r *http.Request) {}
func UserUpdate(w http.ResponseWriter, r *http.Request) {}
func UserDelete(w http.ResponseWriter, r *http.Request) {}
