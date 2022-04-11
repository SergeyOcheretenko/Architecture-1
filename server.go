package main
 
import (
   "net/http"
   "time"
)
 func main() {
   http.HandleFunc("/time", func(rw http.ResponseWriter, r *http.Request) {
       now := time.Now()
       formattedTime := now.Format("2006-01-02T15:04:05Z07:00")
       rw.Write([]byte(formattedTime))
   })
   http.ListenAndServe(":8795", nil)
}
