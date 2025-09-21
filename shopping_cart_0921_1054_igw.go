// 代码生成时间: 2025-09-21 10:54:22
package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

// ShoppingCart represents a shopping cart with items
type ShoppingCart struct {
    Items map[string]int `json:"items"`
}

// Item represents an item in the cart
type Item struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}

// NewShoppingCart creates a new shopping cart
func NewShoppingCart() *ShoppingCart {
    return &ShoppingCart{Items: make(map[string]int)}
}

// AddItem adds an item to the shopping cart
func (c *ShoppingCart) AddItem(item Item) error {
    if _, exists := c.Items[item.ID]; exists {
        return fmt.Errorf("item with id %s already exists", item.ID)
    }
    c.Items[item.ID] = 1 // assuming quantity is 1 for simplicity
    return nil
}

// RemoveItem removes an item from the shopping cart
func (c *ShoppingCart) RemoveItem(itemID string) error {
    if _, exists := c.Items[itemID]; !exists {
        return fmt.Errorf("item with id %s does not exist", itemID)
    }
    delete(c.Items, itemID)
    return nil
}

// GetCart returns the current state of the shopping cart
func (c *ShoppingCart) GetCart() ([]Item, error) {
    var items []Item
    for id, quantity := range c.Items {
        // In a real-world scenario, you would fetch the item details from a database or service
        // For simplicity, let's assume we have a predefined item with a matching ID
        if item, ok := predefinedItems[id]; ok {
            item.Quantity = quantity
            items = append(items, item)
        }
    }
    return items, nil
}

// predefinedItems is a map of predefined items for demonstration purposes
var predefinedItems = map[string]Item{
    "1": {ID: "1", Name: "Apple", Price: 0.99},
    "2": {ID: "2", Name: "Banana", Price: 0.59},
    "3": {ID: "3", Name: "Cherry", Price: 2.99},
}

// Router handles HTTP routing
func Router(cart *ShoppingCart) http.Handler {
    r := mux.NewRouter()
    r.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
        items, err := cart.GetCart()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(items)
    })
    r.HandleFunc("/cart/add/{id}", func(w http.ResponseWriter, r *http.Request) {
        var item Item
        vars := mux.Vars(r)
        itemID := vars["id"]
        if _, exists := predefinedItems[itemID]; !exists {
            http.Error(w, "item not found", http.StatusNotFound)
            return
        }
        if err := cart.AddItem(predefinedItems[itemID]); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        w.WriteHeader(http.StatusOK)
    })
    r.HandleFunc("/cart/remove/{id}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        itemID := vars["id"]
        if err := cart.RemoveItem(itemID); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        w.WriteHeader(http.StatusOK)
    })
    return r
}

func main() {
    cart := NewShoppingCart()
    handler := Router(cart)
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}