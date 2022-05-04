package main

import (
    "fmt"
    "net/http"
)
import _ "net/http/pprof"


func main() {
    fmt.Println("Starting the server at port 8080")

    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}


func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World in repo !!!!!")
}