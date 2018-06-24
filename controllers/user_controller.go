package controllers

import (
  "net/http"
  "encoding/json"
  "log"
  "fmt"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/guni973/go-restful-api-sample/models"
  "github.com/guni973/go-restful-api-sample/database"
  "github.com/gorilla/mux"
)



func UserIndex(w http.ResponseWriter, r *http.Request) {
  users := []models.User{}
  database.DB.Find(&users)

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
  database.DB.First(&user, vars["id"])

  if err := database.DB.First(&user, vars["id"]).Error; err != nil {
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
  //   database.DB.Rollback()
  //   log.Fatal("Faild create User: ", err)
  //   w.WriteHeader(http.StatusBadRequest)
  // }
  _ = json.NewDecoder(r.Body).Decode(&user)
  database.DB.Create(&user);

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
  database.DB.First(&user, vars["id"])

  if err := database.DB.First(&user, vars["id"]).Error; err != nil {
    fmt.Fprintln(w, err)
    return
  }

  _ = json.NewDecoder(r.Body).Decode(&user)

  // TODO: ErrorHandle for BadReqest
  database.DB.Save(&user)
  //{
  //  database.DB.Rollback()
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
  if err := database.DB.First(&user, vars["id"]).Error; err != nil {
    fmt.Fprintln(w, err)
    return
  }

  database.DB.Delete(&user)
  w.WriteHeader(http.StatusNoContent)
}
