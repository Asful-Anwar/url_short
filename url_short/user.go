package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// User struct represents a user with ID, Username, Email, and Password
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// userStore is an in-memory storage for users
var userStore = make(map[int]User)

// CreateUser adds a new user to the store
func CreateUser(user User) error {
	if _, exists := userStore[user.ID]; exists {
		return errors.New("user with this ID already exists")
	}
	userStore[user.ID] = user
	return nil
}

// GetUser retrieves a user by ID
func GetUser(id int) (User, error) {
	user, exists := userStore[id]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

// UpdateUser updates an existing user
func UpdateUser(user User) error {
	if _, exists := userStore[user.ID]; !exists {
		return errors.New("user not found")
	}
	userStore[user.ID] = user
	return nil
}

// DeleteUser removes a user by ID
func DeleteUser(id int) error {
	if _, exists := userStore[id]; !exists {
		return errors.New("user not found")
	}
	delete(userStore, id)
	return nil
}

// handler for /users endpoint
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Create user
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		err = CreateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handler for /users/{id} endpoint
func userHandler(w http.ResponseWriter, r *http.Request) {
	// Extract id from URL path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		user, err := GetUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
	case http.MethodPut:
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if user.ID != id {
			http.Error(w, "User ID in URL and body do not match", http.StatusBadRequest)
			return
		}
		err = UpdateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
	case http.MethodDelete:
		err := DeleteUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/", userHandler)

	fmt.Println("Server running on port 8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
