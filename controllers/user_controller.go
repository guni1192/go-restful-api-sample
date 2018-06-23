package controllers

import (
  "net/http"
  "encoding/json"
  "log"
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "github.com/guni973/go-restful-api-sample/models"
  "github.com/gorilla/mux"
)

var DB *gorm.DB

func init(){ var err error
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

  jsonBytes, err:= json.MarshalIndent(users, "", "    ")
  if err != nil {
    log.Fatal("JSON Encode: ", err)
  }
  w.Write(jsonBytes)
}

func UserDetail(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  user := models.User{}
  DB.First(&user, vars["id"])

  if err := DB.First(&user, vars["id"]).Error; err != nil {
    fmt.Println(w, "[]")
    return
  }

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  jsonBytes, err:= json.MarshalIndent(user, "", "    ")
  // enc := json.NewEncoder(w)
  // enc.SetIndent("", "    ")

  // if err := enc.Encode(user); err != nil {
  if err != nil {
    log.Fatal("JSON Encode: ", err)
  }

  w.Write(jsonBytes)
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
  user := models.User{}

  // TODO: ErrorHandle for BadReqest
  // err != nil {
  //   DB.Rollback()
  //   log.Fatal("Faild create User: ", err)
  //   w.WriteHeader(http.StatusBadRequest)
  // }
  _ = json.NewDecoder(r.Body).Decode(&user)
  DB.Create(&user);

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)

  jsonBytes, err:= json.MarshalIndent(user, "", "    ")

  if err != nil {
    log.Fatal("JSON Encode: ", err)
  }

  w.Write(jsonBytes)
}


func UserUpdate(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  user := models.User{}
  DB.First(&user, vars["id"])

  if err := DB.First(&user, vars["id"]).Error; err != nil {
    fmt.Fprintln(w, err)
    return
  }

  _ = json.NewDecoder(r.Body).Decode(&user)

  // TODO: ErrorHandle for BadReqest
  DB.Save(&user)
  //{
  //  DB.Rollback()
  //  w.WriteHeader(http.StatusBadRequest)
  //  fmt.Fprintln(w, err)
  //}

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)

  jsonBytes, err:= json.MarshalIndent(user, "", "    ")

  if err != nil {
    log.Fatal("JSON Encode: ", err)
  }

  w.Write(jsonBytes)
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  user := models.User{}
  if err := DB.First(&user, vars["id"]).Error; err != nil {
    fmt.Fprintln(w, err)
    return
  }

  DB.Delete(&user)
  w.WriteHeader(http.StatusNoContent)
}
