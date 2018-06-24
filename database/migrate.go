package main

import (
  "github.com/guni973/go-restful-api-sample/database"
)

func main() {
  DB.CreateTable(&User{})
}
