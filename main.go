package main

import (
    "log"
    "net/http"
    "sync"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello from Service 1!\n"))
}

func Handler2(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello from Service 2!\n"))
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    
    // Сервис 1 на порту 8080
    go func() {
        defer wg.Done()
        log.Println("Запуск Service 1 на порту :8080")
        log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(Handler)))
    }()
    
    // Сервис 2 на порту 8081
    go func() {
        defer wg.Done()
        log.Println("Запуск Service 2 на порту :8081")
        log.Fatal(http.ListenAndServe(":8081", http.HandlerFunc(Handler2)))
    }()
    
    wg.Wait()
}