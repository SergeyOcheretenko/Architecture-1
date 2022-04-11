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
  const RFC3339 = "2006-01-02T15:04:05Z07:00"

  http.HandleFunc("/time", func(rw http.ResponseWriter, r *http.Request) {
    now := time.Now()
    formattedTime := now.Format(RFC3339)
    structTime := Time{NowTime: formattedTime}
    fmt.Println(structTime)
    rw.Write([]byte(formattedTime))
  })
  http.ListenAndServe(":8795", nil)
}
