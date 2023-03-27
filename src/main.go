package main

import (
    "log"
    "net/http"

    "./application"
    "./infrastructure"
    "./interface/http"
)

func main() {
    db, err := infrastructure.NewDB()
    if err != nil {
        log.Fatalf("failed to connect to db: %v", err)
    }
    defer db.Close()

    userRepo := infrastructure.NewUserRepo(db)
    userService := application.NewUserService(userRepo)

    router := http.NewRouter(userService)

    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}
