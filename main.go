package main

import (
  "encoding/json"
  "log"
  "flag"
  "net/http"
)

func main() {
  var addr = flag.String("addr", ":8080", "localhost")
  flag.Parse()

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    enc := json.NewEncoder(w)
    enc.SetIndent("", "    ")
    if err := enc.Encode(map[string]string{"message": "Hello World!!"}); err != nil {
      log.Fatal("JSON Encode: ", err)
    }
  })

  log.Println("Server is running. Port: ", *addr)
  if err := http.ListenAndServe(*addr, nil); err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
