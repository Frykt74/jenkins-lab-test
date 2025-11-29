package main

import (
    "net/http"
)

// Сервис 1: простой ответ
func Handler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello!\n"))
}

// Сервис 2: другой ответ
func Handler2(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello from Service 2!\n"))
}

// Основная функция для запуска
func main() {
    http.ListenAndServe("0.0.0.0:8080", http.HandlerFunc(Handler))
}
