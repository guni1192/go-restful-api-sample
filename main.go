package main

import (
  "encoding/json"
  "flag"
  "github.com/gorilla/mux"
  "github.com/guni973/go-restful-api-sample/controllers"
  "log"
  "net/http"
)

func main() {
  var addr = flag.String("addr", ":8080", "localhost")
  flag.Parse()

  router := mux.NewRouter().StrictSlash(true)

  router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    enc := json.NewEncoder(w)
    enc.SetIndent("", "    ")
    if err := enc.Encode(map[string]string{"message": "Hello World!!"}); err != nil {
      log.Fatal("JSON Encode: ", err)
    }
  })

  router.HandleFunc("/users", controllers.UserIndex)
  router.HandleFunc("/users/{id}", controllers.UserDetail)

  log.Println("Server is running. Port: ", *addr)
  if err := http.ListenAndServe(*addr, router); err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
