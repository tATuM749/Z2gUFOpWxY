// 代码生成时间: 2025-08-12 07:22:20
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

// InventoryItem represents an item in the inventory.
type InventoryItem struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    Quantity  int    `json:"quantity"`
}

// inventory is a global variable representing the inventory.
var inventory = []InventoryItem{{ID: "1", Name: "Laptop", Quantity: 10}, {ID: "2", Name: "Mouse", Quantity: 50}}

// GetInventory returns the current inventory as JSON.
func GetInventory(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(inventory)
}

// AddItem adds a new item to the inventory.
func AddItem(w http.ResponseWriter, r *http.Request) {
    var item InventoryItem
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    inventory = append(inventory, item)
    fmt.Fprintf(w, `Added item with ID: "%s"`, item.ID)
}

// UpdateItem updates an existing item in the inventory.
func UpdateItem(w http.ResponseWriter, r *http.Request) {
    var item InventoryItem
    vars := mux.Vars(r)
    id := vars["id"]
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    for i, existingItem := range inventory {
        if existingItem.ID == id {
            inventory[i] = item
            fmt.Fprintf(w, `Updated item with ID: "%s"`, id)
            return
        }
    }
    http.Error(w, "Item not found", http.StatusNotFound)
}

// DeleteItem removes an item from the inventory.
func DeleteItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    for i, existingItem := range inventory {
        if existingItem.ID == id {
            inventory = append(inventory[:i], inventory[i+1:]...)
            fmt.Fprintf(w, `Deleted item with ID: "%s"`, id)
            return
        }
    }
    http.Error(w, "Item not found", http.StatusNotFound)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/inventory", GetInventory).Methods("GET")
    r.HandleFunc("/inventory", AddItem).Methods("POST")
    r.HandleFunc("/inventory/{id}", UpdateItem).Methods("PUT")
    r.HandleFunc("/inventory/{id}", DeleteItem).Methods("DELETE")
    fmt.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", r)
}