// 代码生成时间: 2025-08-05 15:57:34
package main

import (
    "net/http"
    "strings"
    "log"
    "encoding/json"
    "github.com/gorilla/mux"
)

// ShoppingCart represents a shopping cart
type ShoppingCart struct {
    Items map[string]int `json:"items"`
}

// NewShoppingCart initializes a new shopping cart
func NewShoppingCart() *ShoppingCart {
    return &ShoppingCart{Items: make(map[string]int)}
}

// AddItem adds an item to the shopping cart
func (c *ShoppingCart) AddItem(itemID string, quantity int) error {
    if quantity <= 0 {
        return errors.New("item quantity must be greater than zero")
    }
    c.Items[itemID] += quantity
    return nil
}

// RemoveItem removes an item from the shopping cart
func (c *ShoppingCart) RemoveItem(itemID string) error {
    if _, exists := c.Items[itemID]; !exists {
        return errors.New("item does not exist in the cart")
    }
    delete(c.Items, itemID)
    return nil
}

// HandleAddItem handles HTTP requests to add an item to the cart
func HandleAddItem(w http.ResponseWriter, r *http.Request) {
    cart := NewShoppingCart() // In a real application, this should be retrieved from a session or database
    var addItem struct {
        ItemID  string `json:"itemID"`
        Quantity int    `json:"quantity"`
    }
    err := json.NewDecoder(r.Body).Decode(&addItem)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    if err := cart.AddItem(addItem.ItemID, addItem.Quantity); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(cart)
}

// HandleRemoveItem handles HTTP requests to remove an item from the cart
func HandleRemoveItem(w http.ResponseWriter, r *http.Request) {
    cart := NewShoppingCart() // In a real application, this should be retrieved from a session or database
    vars := mux.Vars(r)
    itemID := vars["itemID"]
    if err := cart.RemoveItem(itemID); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(cart)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/cart/add", HandleAddItem).Methods("POST")
    r.HandleFunc("/cart/remove/{itemID}", HandleRemoveItem).Methods("POST")
    log.Println("Starting shopping cart service on port 8080")
    http.ListenAndServe(":8080", r)
}