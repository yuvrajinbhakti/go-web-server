package main

import (
    "fmt"
    "log"
    "net/http"
)
func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w,"Hello")
    })
    fmt.Printf("Starting server at port 8084\n")
    if err := http.ListenAndServe(":8084",nil);err!=nil{
        log.Fatal(err)
    }
}
