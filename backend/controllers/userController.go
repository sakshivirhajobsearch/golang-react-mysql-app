package controllers

import (
    "encoding/json"
    "net/http"
    "fmt"
)

// GetUsers handles GET requests to retrieve all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        // Logic to fetch users from the database
        fmt.Fprintln(w, "Returning all users") // Replace with actual user fetching logic
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// GetUser handles GET requests to fetch a single user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        // Logic to fetch a single user by ID
        fmt.Fprintln(w, "Returning a single user") // Replace with actual user fetching logic by ID
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// CreateUser handles POST requests to create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        // Logic to create a new user
        var newUser map[string]interface{}
        if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        fmt.Fprintln(w, "User created successfully") // Replace with actual user creation logic
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// UpdateUser handles PUT requests to update an existing user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPut {
        // Logic to update an existing user
        var updatedUser map[string]interface{}
        if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        fmt.Fprintln(w, "User updated successfully") // Replace with actual user update logic
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// DeleteUser handles DELETE requests to remove a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodDelete {
        // Logic to delete a user
        fmt.Fprintln(w, "User deleted successfully") // Replace with actual user deletion logic
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
