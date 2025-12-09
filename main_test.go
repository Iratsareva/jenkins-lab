package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    rec := httptest.NewRecorder()
    Handler(rec, req)

    if rec.Code != 200 {
        t.Errorf("ожидали 200, получили %d", rec.Code)
    }
    if rec.Body.String() != "Hello from Service 1!\n" {
        t.Errorf("неправильный ответ: %s", rec.Body.String())
    }
}

func TestHandler2(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    rec := httptest.NewRecorder()
    Handler2(rec, req)

    if rec.Code != 200 {
        t.Errorf("ожидали 200, получили %d", rec.Code)
    }
    if rec.Body.String() != "Hello from Service 2!\n" {
        t.Errorf("неправильный ответ")
    }
}

// === ЭТО И ЕСТЬ ИНТЕГРАЦИОННЫЙ ТЕСТ ===
func TestIntegration(t *testing.T) {
    // Запускаем два тестовых сервера
    server1 := httptest.NewServer(http.HandlerFunc(Handler))
    defer server1.Close()

    server2 := httptest.NewServer(http.HandlerFunc(Handler2))
    defer server2.Close()

    // Проверяем, что оба живые
    resp1, err := http.Get(server1.URL)
    if err != nil || resp1.StatusCode != 200 {
        t.Fatal("Service 1 не отвечает")
    }

    resp2, err := http.Get(server2.URL)
    if err != nil || resp2.StatusCode != 200 {
        t.Fatal("Service 2 не отвечает")
    }

    t.Log("Интеграционный тест прошёл успешно — оба сервиса отвечают!")
}