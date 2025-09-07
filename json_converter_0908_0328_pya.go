// 代码生成时间: 2025-09-08 03:28:29
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
    "log"

    "github.com/gorilla/mux"
)

// JSONData represents the structure of JSON data to be converted.
type JSONData struct {
    Data string `json:"data"`
}

// ErrorResponse represents the error response structure.
type ErrorResponse struct {
    Error string `json:"error"`
}

// convertJSONData is a middleware function that converts JSON data.
func convertJSONData(w http.ResponseWriter, r *http.Request) {
    // Check if the Content-Type is application/json.
    if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
        respondWithError(w, http.StatusUnsupportedMediaType, "Invalid Content-Type. Must be application/json.")
        return
    }

    // Decode the JSON data from the request body.
    var jsonData JSONData
    if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid JSON data.")
        return
    }
    defer r.Body.Close()

    // Perform the conversion logic here.
    // For demonstration, we're simply echoing back the received data.
    // This is where you would add your actual conversion logic.
    
    // Respond with the converted JSON data.
    respondWithJSON(w, http.StatusOK, jsonData)
}

// respondWithError writes the error response to the client.
func respondWithError(w http.ResponseWriter, code int, message string) {
    w.WriteHeader(code)
    respondWithJSON(w, code, ErrorResponse{Error: message})
}

// respondWithJSON writes the JSON response to the client.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    // Set the Content-Type header to application/json.
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    if err := json.NewEncoder(w).Encode(payload); err != nil {
        log.Printf("Error while encoding JSON response: %s", err)
    }
}

func main() {
    // Create a new Gorilla router.
    router := mux.NewRouter()

    // Define the route for the JSON data conversion.
    // The route accepts POST requests with JSON data.
    router.HandleFunc("/convert", convertJSONData).Methods("POST")

    // Start the server on port 8080.
    fmt.Println("Starting JSON data converter on port 8080...")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}