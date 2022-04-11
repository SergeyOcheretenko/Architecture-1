package main

import (
  "net/http"
  "time"
  "fmt"
)

type Time struct {
  NowTime string
}
func main() {
  http.HandleFunc("/time", func(rw http.ResponseWriter, r *http.Request) {
    now := time.Now()
    formattedTime := now.Format("2006-01-02T15:04:05Z07:00")
    structTime := Time{NowTime: formattedTime}
    fmt.Println(structTime)
    rw.Write([]byte("work"))
  })
  http.ListenAndServe(":8795", nil)
}
