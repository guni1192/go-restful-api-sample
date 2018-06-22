package main

import (
  "encoding/json"
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    enc := json.NewEncoder(w)
    enc.SetIndent("", "    ")
    if err := enc.Encode(map[string]string{"message": "Hello World!!"}); err != nil {
      log.Fatal("JSON Encode: ", err)
    }
  })

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
