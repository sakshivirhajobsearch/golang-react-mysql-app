package routes

import (
    "backend/controllers" // Ensure the correct import path for the controllers package
    "net/http"
)

// SetupRoutes initializes all the routes for the application
func SetupRoutes() {
    // Initialize a new serve mux for better routing control
    mux := http.NewServeMux()

    // Define routes for each endpoint
    mux.HandleFunc("/users", controllers.GetUsers)      // GET /users (list all users)
    mux.HandleFunc("/user", controllers.GetUser)        // GET /user?id=1 (fetch a single user by ID)
    mux.HandleFunc("/api/users", controllers.GetUsers)  // GET /api/users (same as /users, but prefixed with API)

    // Routes for creating, updating, and deleting users
    mux.HandleFunc("/users/create", controllers.CreateUser)  // POST /users/create
    mux.HandleFunc("/users/update", controllers.UpdateUser)  // PUT /users/update
    mux.HandleFunc("/users/delete", controllers.DeleteUser)  // DELETE /users/delete

    // Set the handler for the server
    http.Handle("/", mux)
}
