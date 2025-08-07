// 代码生成时间: 2025-08-07 15:15:18
package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

// ApiResponse represents the structure of API responses
type ApiResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

// NewApiResponse creates a new ApiResponse instance
func NewApiResponse(status, message string, data interface{}) ApiResponse {
    return ApiResponse{
        Status:  status,
        Message: message,
        Data:    data,
    }
}

// ErrorResponse returns a formatted error response
func ErrorResponse(w http.ResponseWriter, err error, statusCode int) {
    apiResponse := NewApiResponse("error", err.Error(), nil)
    EncodeJSONResponse(w, apiResponse, statusCode)
}

// SuccessResponse returns a formatted success response
func SuccessResponse(w http.ResponseWriter, data interface{}, statusCode int) {
    apiResponse := NewApiResponse("success", "success", data)
    EncodeJSONResponse(w, apiResponse, statusCode)
}

// EncodeJSONResponse encodes a response to JSON and writes it to the response writer
func EncodeJSONResponse(w http.ResponseWriter, response interface{}, statusCode int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        panic(err)
    }
}

func main() {
    router := mux.NewRouter()
    
    // Define your routes here
    router.HandleFunc("/format-response", FormatResponseHandler).Methods("GET")

    // Start the server
    fmt.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", router)
}

// FormatResponseHandler is the handler function for the '/format-response' endpoint
func FormatResponseHandler(w http.ResponseWriter, r *http.Request) {
    // Example usage of ErrorResponse and SuccessResponse
    // For demonstration purposes, this handler will always return a success response
    SuccessResponse(w, map[string]string{"key": "value"}, http.StatusOK)
    
    // Uncomment the following lines to see an error response
    // err := errors.New("something went wrong")
    // ErrorResponse(w, err, http.StatusInternalServerError)
}
