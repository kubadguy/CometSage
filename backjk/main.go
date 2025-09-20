package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "🚀 Welcome to CometSage backend!")
    })

    port := ":8080"
    fmt.Println("Server running on http://localhost" + port)
    log.Fatal(http.ListenAndServe(port, nil))
}
