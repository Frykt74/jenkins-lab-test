package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

//  Handler unit test
func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    rec := httptest.NewRecorder()
    
    Handler(rec, req)
    
    if rec.Code != http.StatusOK {
        t.Errorf("wrong status: got %v want %v", rec.Code, http.StatusOK)
    }
    if rec.Body.String() != "Hello!\n" {
        t.Errorf("wrong body: got %s", rec.Body.String())
    }
}

// Handler2 unit test
func TestHandler2(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    rec := httptest.NewRecorder()
    
    Handler2(rec, req)
    
    if rec.Code != http.StatusOK {
        t.Errorf("wrong status: got %v want %v", rec.Code, http.StatusOK)
    }
    if rec.Body.String() != "Hello from Service 2!\n" {
        t.Errorf("wrong body: got %s", rec.Body.String())
    }
}

// integration test
func TestIntegration(t *testing.T) {
    // Запускаем первый сервис на тестовом сервере
    server1 := httptest.NewServer(http.HandlerFunc(Handler))
    defer server1.Close()
    
    // Запускаем второй сервис на тестовом сервере
    server2 := httptest.NewServer(http.HandlerFunc(Handler2))
    defer server2.Close()
    
    // Проверяем, что первый сервис отвечает
    resp1, err := http.Get(server1.URL)
    if err != nil {
        t.Fatal(err)
    }
    defer resp1.Body.Close()
    
    if resp1.StatusCode != http.StatusOK {
        t.Errorf("Service 1 failed: got %v", resp1.StatusCode)
    }
    
    // Проверяем, что второй сервис отвечает
    resp2, err := http.Get(server2.URL)
    if err != nil {
        t.Fatal(err)
    }
    defer resp2.Body.Close()
    
    if resp2.StatusCode != http.StatusOK {
        t.Errorf("Service 2 failed: got %v", resp2.StatusCode)
    }
    
    t.Log("[SUCCESS!] Integration test passed: both services responding")
}
