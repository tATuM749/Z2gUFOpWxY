// 代码生成时间: 2025-09-11 14:35:07
package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "log"
)

// ApiResponse represents a standardized API response format
type ApiResponse struct {
    Data interface{} `json:"data"`
    Message string `json:"message"`
    Success bool `json:"success"`
    Error string `json:"error"`
}

// NewApiResponse creates a new ApiResponse instance with the provided data and success status.
func NewApiResponse(data interface{}, success bool, message string, error string) ApiResponse {
    return ApiResponse{
        Data: data,
        Message: message,
        Success: success,
        Error: error,
    }
}

// NewResponseHandler creates a new HTTP handler that wraps the given handler function to provide
// a consistent response format.
func NewResponseHandler(handlerFunc func(http.ResponseWriter, *http.Request) (interface{}, error)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        result, err := handlerFunc(w, r)
        if err != nil {
            // Handle error case
            http.Error(w, err.Error(), http.StatusInternalServerError)
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(NewApiResponse(nil, false, "", err.Error()))
            return
        }
        // Handle success case
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(NewApiResponse(result, true, "Success", ""))
    }
}

// ExampleHandler demonstrates a simple API endpoint that uses the NewResponseHandler.
func ExampleHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    // Simulate some logic that might return an error
    if r.URL.Query().Get("error") == "true" {
        return nil, fmt.Errorf("simulated error")
    }
    return map[string]string{"status": "ok"}, nil
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/example", NewResponseHandler(ExampleHandler)).Methods("GET")

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}
