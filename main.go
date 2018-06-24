package main

import (
  "encoding/json"
  "flag"
  "github.com/gorilla/mux"
  "github.com/guni973/go-restful-api-sample/controllers"
  "github.com/guni973/go-restful-api-sample/database"
  "github.com/guni973/go-restful-api-sample/models"
  "log"
  "net/http"
)

func main() {
  var addr = flag.String("addr", ":8080", "localhost")

  flag.Parse()

  if flag.Arg(0) == "migrate" {
    database.DB.CreateTable(&models.User{})
    return
  }

  router := mux.NewRouter().StrictSlash(true)

  router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    enc := json.NewEncoder(w)
    enc.SetIndent("", "    ")
    if err := enc.Encode(map[string]string{"message": "Hello World!!"}); err != nil {
      log.Fatal("JSON Encode: ", err)
    }
  }).Methods("GET")

  router.HandleFunc("/users",      controllers.UserIndex).Methods("GET")
  router.HandleFunc("/users/{id}", controllers.UserDetail).Methods("GET")
  router.HandleFunc("/users",      controllers.UserCreate).Methods("POST")
  router.HandleFunc("/users/{id}", controllers.UserUpdate).Methods("PUT")
  router.HandleFunc("/users/{id}", controllers.UserDelete).Methods("DELETE")

  log.Println("Server is running. Port: ", *addr)
  if err := http.ListenAndServe(*addr, router); err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
