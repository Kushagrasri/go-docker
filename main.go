package main

import (
    "fmt"
    "net/http"
)
import _ "net/http/pprof"


func main() {
    
    go func() {
		fmt.Println(http.ListenAndServe(":6060", nil))
	}()

    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}


func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, W orlds!!!!!")
}