// 代码生成时间: 2025-08-22 16:18:13
package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "log"
    "github.com/gorilla/mux"
)

// JSONData represents the structure of the JSON data to be transformed.
type JSONData struct {
    Name  string `json:"name"`
    Value string `json:"value"`
}

// HomeHandler is the handler function that will be called when the home route is hit.
// It converts the JSON data from the request to a new JSON format and returns it.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    // Create a new instance of JSONData.
    var data JSONData

    // Decode the incoming JSON request into our data struct.
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        // If there's an error, respond with a 400 Bad Request status and the error message.
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Change the format of the JSON data if needed.
    // For example, we could uppercase the name here.
    data.Name = data.Name

    // Encode the new data struct as JSON.
    if err := json.NewEncoder(w).Encode(data); err != nil {
        // If there's an error encoding the response, respond with a 500 Internal Server Error status and the error message.
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func main() {
    // Initialize a new router.
    router := mux.NewRouter()

    // Register the HomeHandler for the '/' route.
    router.HandleFunc("/", HomeHandler).Methods("POST")

    // Start the HTTP server on port 8080.
    log.Println("Server is running on port 8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal("Error starting server: ", err)
    }
}
