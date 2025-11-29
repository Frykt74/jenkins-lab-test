package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

// Unit-тест для первого сервиса
func TestHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }
    
    rec := httptest.NewRecorder()
    http.HandlerFunc(Handler).ServeHTTP(rec, req)
    
    if rec.Code != http.StatusOK {
        t.Errorf("wrong status code: want %v got %v", http.StatusOK, rec.Code)
    }
    
    if rec.Body.String() != "Hello!\n" {
        t.Errorf("wrong body: got %s", rec.Body.String())
    }
}

// Unit-тест для второго сервиса
func TestHandler2(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }
    
    rec := httptest.NewRecorder()
    http.HandlerFunc(Handler2).ServeHTTP(rec, req)
    
    if rec.Code != http.StatusOK {
        t.Errorf("wrong status code: want %v got %v", http.StatusOK, rec.Code)
    }
    
    if rec.Body.String() != "Hello from Service 2!\n" {
        t.Errorf("wrong body: got %s", rec.Body.String())
    }
}

// Интеграционный тест - проверяет работу обоих сервисов вместе
func TestIntegration(t *testing.T) {
    // Запускаем тестовые серверы для обоих сервисов
    server1 := httptest.NewServer(http.HandlerFunc(Handler))
    defer server1.Close()
    
    server2 := httptest.NewServer(http.HandlerFunc(Handler2))
    defer server2.Close()
    
    // Проверяем второй сервис
    resp, err := http.Get(server2.URL)
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Service 2 returned wrong status: got %v want %v", resp.StatusCode, http.StatusOK)
    }
    
    // Проверяем первый сервис
    resp, err = http.Get(server1.URL)
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Service 1 returned wrong status: got %v want %v", resp.StatusCode, http.StatusOK)
    }
    
    t.Log("Integration test passed: both services are responding correctly")
}
