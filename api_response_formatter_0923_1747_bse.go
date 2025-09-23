// 代码生成时间: 2025-09-23 17:47:21
package main

import (
    "net/http"
    "encoding/json"
    "log"
    "github.com/gorilla/mux"
)

// ApiResponse standardizes API responses.
type ApiResponse struct {
    Success bool        `json:"success"`
    Message string     `json:"message"`
    Data    interface{} `json:"data"`
}

// NewApiResponse creates a new ApiResponse with default values.
func NewApiResponse(success bool, message string, data interface{}) ApiResponse {
    return ApiResponse{
        Success: success,
        Message: message,
        Data:    data,
    }
}

// ErrorResponse returns an error response with a 400 status code.
func ErrorResponse(w http.ResponseWriter, message string) {
    w.WriteHeader(http.StatusBadRequest)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(NewApiResponse(false, message, nil))
}

// Home handler is the landing page for our API.
func Home(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(NewApiResponse(true, "Welcome to the API Response Formatter", nil))
}

// main function sets up and starts the HTTP server.
func main() {
    // Create a new instance of the Gorilla Router.
    r := mux.NewRouter()
    // Define routes with associated handlers.
    r.HandleFunc("/", Home).Methods("GET")

    // Start the HTTP server.
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}