package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
    if r.url.Path != "/hello"{
        http.Error(w,"404 not found",http.StatusNotFound)
        return
    }

    if r.Method !="GET"{
        http.Error(w,"Method is not allowed",http.StatusMethodNotAllowed)
        return
    }

    fmt.Fprintf(w, "Hello")
}
func main() {
    http.HandleFunc("/hello",helloHandler)
    fmt.Printf("Starting server at port 8084\n")
    if err := http.ListenAndServe(":8084",nil);err!=nil{
        log.Fatal(err)
    }
}
