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
    if err := http.ListenAndServe("0.0.0.0:8081", http.HandlerFunc(Handler2)); err != nil {
        log.Fatal(err)
    }
}
