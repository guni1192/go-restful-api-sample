package controllers

import (
  "net/http"
  "encoding/json"
  "log"
  // "fmt"
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

  enc := json.NewEncoder(w)
  enc.SetIndent("", "    ")

  if err := enc.Encode(users); err != nil {
    log.Fatal("JSON Encode: ", err)
  }
}

func UserDetail(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  user := models.User{}
  DB.First(&user, vars["id"])

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  enc := json.NewEncoder(w)
  enc.SetIndent("", "    ")

  if err := enc.Encode(user); err != nil {
    log.Fatal("JSON Encode: ", err)
  }

}

func UserCreate(w http.ResponseWriter, r *http.Request) {
  // vars := mux.Vars(r)

  r.ParseForm()
  user := models.User{Name: r.FormValue("name"), Email: r.FormValue("email")}

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

  enc := json.NewEncoder(w)
  enc.SetIndent("", "    ")

  if err := enc.Encode(user); err != nil {
    log.Fatal("JSON Encode: ", err)
  }

}


func UserUpdate(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  r.ParseForm()
  // TODO Not FOund
  user := models.User{}
  DB.First(&user, vars["id"])
  _ = json.NewDecoder(r.Body).Decode(&user)

  // TODO: ErrorHandle for BadReqest
  DB.Save(&user)
  //{
  //  DB.Rollback()
  //  log.Fatal("Faild create User: ", err)
  //  w.WriteHeader(http.StatusBadRequest)
  //  fmt.Fprintln(w, err)
  //}

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)

  enc := json.NewEncoder(w)
  enc.SetIndent("", "    ")

  if err := enc.Encode(user); err != nil {
    log.Fatal("JSON Encode: ", err)
  }

}

func UserDelete(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  user := models.User{}
  DB.First(&user, vars["id"])
  DB.Delete(&user)

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)

  enc := json.NewEncoder(w)
  enc.SetIndent("", "    ")

  if err := enc.Encode(user); err != nil {
    log.Fatal("JSON Encode: ", err)
  }
}
