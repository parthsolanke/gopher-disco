package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func setup() (router *http.ServeMux, cleanup func()) {
	UsersDB = make(map[int]User)
	DBMutex = sync.RWMutex{}

	router = http.NewServeMux()

	router.Handle("/", ErrorHandlerFunc(handleRoot))
	router.Handle("POST /users", ErrorHandlerFunc(addUser))
	router.Handle("GET /users/{id}", ErrorHandlerFunc(getUser))
	router.Handle("DELETE /users/{id}", ErrorHandlerFunc(deleteUser))

	return router, func() {}
}

func TestHandleRoot(t *testing.T) {
	router, cleanup := setup()
	defer cleanup()

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	if w.Body.String() != "Hello World" {
		t.Fatalf("expected body 'Hello World', got '%s'", w.Body.String())
	}
}

func TestAddUser(t *testing.T) {
	router, cleanup := setup()
	defer cleanup()

	user := User{Name: "John Doe", Email: "john@example.com"}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", w.Code)
	}

	DBMutex.RLock()
	if len(UsersDB) != 1 {
		t.Fatalf("expected 1 user in the database, got %d", len(UsersDB))
	}
	DBMutex.RUnlock()
}

func TestGetUser(t *testing.T) {
	router, cleanup := setup()
	defer cleanup()

	DBMutex.Lock()
	UsersDB[1] = User{Name: "John Doe", Email: "john@example.com"}
	DBMutex.Unlock()

	req := httptest.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var user User
	err := json.NewDecoder(w.Body).Decode(&user)
	if err != nil {
		t.Fatalf("error decoding response: %v", err)
	}

	if user.Name != "John Doe" || user.Email != "john@example.com" {
		t.Fatalf("expected user {Name: John Doe, Email: john@example.com}, got %v", user)
	}
}

func TestDeleteUser(t *testing.T) {
	router, cleanup := setup()
	defer cleanup()

	DBMutex.Lock()
	UsersDB[1] = User{Name: "John Doe", Email: "john@example.com"}
	DBMutex.Unlock()

	req := httptest.NewRequest("DELETE", "/users/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", w.Code)
	}

	DBMutex.RLock()
	if _, exists := UsersDB[1]; exists {
		t.Fatalf("expected user to be deleted")
	}
	DBMutex.RUnlock()
}

func TestGetUserNotFound(t *testing.T) {
	router, cleanup := setup()
	defer cleanup()

	req := httptest.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", w.Code)
	}
}

func TestAddUserInvalidData(t *testing.T) {
	router, cleanup := setup()
	defer cleanup()

	req := httptest.NewRequest("POST", "/users", bytes.NewReader([]byte("")))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)
	}
}

func TestDeleteUserNotFound(t *testing.T) {
	router, cleanup := setup()
	defer cleanup()

	req := httptest.NewRequest("DELETE", "/users/999", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", w.Code)
	}
}
