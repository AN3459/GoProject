package http

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "example.com/m/v2/src/application"
)

func NewRouter(userService application.UserService) http.Handler {
    r := chi.NewRouter()
    r.Use(middleware.Logger)

    userController := NewUserController(userService)
    userController.RegisterRoutes(r)

    return r
}
