package main
//go:generate gopherjs build main.go -o js/app.js -m
// +build ignore

import (
  "log"
)

func main() {
    log.Println("Hello World")
}
