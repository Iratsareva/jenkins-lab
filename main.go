package main

import (
    "log"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello from Service 1!\n"))
}

func main() {
    log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(Handler)))
}