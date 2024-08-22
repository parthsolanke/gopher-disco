package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var UsersDB = make(map[int]User)
var DBMutex sync.RWMutex

type ErrorHandlerFunc func(http.ResponseWriter, *http.Request) error

func (fn ErrorHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	router := http.NewServeMux()

	router.Handle("/", ErrorHandlerFunc(handleRoot))
	router.Handle("POST /users", ErrorHandlerFunc(addUser))
	router.Handle("GET /users/{id}", ErrorHandlerFunc(getUser))
	router.Handle("DELETE /users/{id}", ErrorHandlerFunc(deleteUser))

	http.ListenAndServe("127.0.0.1:8080", router)
}

func handleRoot(w http.ResponseWriter, r *http.Request) error {
	_, err := fmt.Fprintf(w, "Hello World")
	return err
}

func addUser(w http.ResponseWriter, r *http.Request) error {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return fmt.Errorf("decoding user data: %w", err)
	}
	if user.Name == "" || user.Email == "" {
		http.Error(w, "name and email are required", http.StatusBadRequest)
		return nil
	}

	DBMutex.Lock()
	UsersDB[len(UsersDB)+1] = user
	DBMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func getUser(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return nil
	}

	DBMutex.RLock()
	user, ok := UsersDB[id]
	DBMutex.RUnlock()
	if !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	msg, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("marshalling user data: %w", err)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(msg)
	return err
}

func deleteUser(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return nil
	}

	DBMutex.Lock()
	defer DBMutex.Unlock()

	if _, ok := UsersDB[id]; !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return nil
	}

	delete(UsersDB, id)
	w.WriteHeader(http.StatusNoContent)
	return nil
}
