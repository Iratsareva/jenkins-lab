package main

import (
    "log"
    "net/http"
)

func Handler2(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello from Service 2!\n"))
}

func main() {
    log.Fatal(http.ListenAndServe(":8081", http.HandlerFunc(Handler2)))
}