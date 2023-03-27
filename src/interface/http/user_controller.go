package http

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "../application"
)

type userController struct {
    userService application.UserService
}

func NewUserController(userService application.UserService) *userController {
    return &userController{userService}
}

func (c *userController) RegisterRoutes(r chi.Router) {
    r.Post("/users", c.AddUser)
    r.Get("/users/{id}", c.GetUserByID)
}

func (c *userController) AddUser(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Name string `json:"name"`
        Age  int    `json:"age"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := c.userService.AddUser(req.Name, req.Age); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (c *userController) GetUserByID(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    user, err := c.userService.GetUserByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
