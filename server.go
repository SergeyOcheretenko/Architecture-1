package main

import "net/http"
 
func main() {
   http.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
       rw.Write([]byte("work"))
   })
   http.ListenAndServe(":8795", nil)
}
