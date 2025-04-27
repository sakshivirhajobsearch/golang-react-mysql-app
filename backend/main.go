package main

import (
    "backend/routes"
    "log"
    "net/http"
)

func main() {
    routes.SetupRoutes()

    log.Println("Server starting on port 8082...")
    if err := http.ListenAndServe(":8082", nil); err != nil {
        log.Fatal(err)
    }
}
