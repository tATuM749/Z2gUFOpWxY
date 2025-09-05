// 代码生成时间: 2025-09-05 13:28:42
package main

import (
    "encoding/json"
    "net/http"
    "strings"

    "github.com/gorilla/mux"
)

// InventoryItem represents an item in the inventory.
type InventoryItem struct {
    ID          int    "json:"id""
    Name        string "json:"name""
    Quantity    int    "json:"quantity""
}

// InventoryService represents the inventory service.
type InventoryService struct {
    items map[int]InventoryItem
}

// NewInventoryService creates a new InventoryService instance.
func NewInventoryService() *InventoryService {
    return &InventoryService{
        items: make(map[int]InventoryItem),
    }
}

// AddItem adds a new item to the inventory.
func (s *InventoryService) AddItem(item InventoryItem) error {
    if _, exists := s.items[item.ID]; exists {
        return errors.New("item with ID already exists")
    }
    s.items[item.ID] = item
    return nil
}

// GetItem retrieves an item from the inventory by its ID.
func (s *InventoryService) GetItem(id int) (InventoryItem, error) {
    item, exists := s.items[id]
    if !exists {
        return InventoryItem{}, errors.New("item not found")
    }
    return item, nil
}

// UpdateItem updates an existing item in the inventory.
func (s *InventoryService) UpdateItem(id int, newItem InventoryItem) error {
    if _, exists := s.items[id]; !exists {
        return errors.New("item not found\)
    }
    s.items[id] = newItem
    return nil
}

// DeleteItem removes an item from the inventory by its ID.
func (s *InventoryService) DeleteItem(id int) error {
    if _, exists := s.items[id]; !exists {
        return errors.New("item not found\)
    }
    delete(s.items, id)
    return nil
}

// InventoryHandler handles inventory-related HTTP requests.
type InventoryHandler struct {
    service *InventoryService
}

// NewInventoryHandler creates a new InventoryHandler instance.
func NewInventoryHandler(service *InventoryService) *InventoryHandler {
    return &InventoryHandler{
        service: service,
    }
}

// AddItemHandler handles POST requests for adding new items to the inventory.
func (h *InventoryHandler) AddItemHandler(w http.ResponseWriter, r *http.Request) {
    var item InventoryItem
    err := json.NewDecoder(r.Body).Decode(&item)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    err = h.service.AddItem(item)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(item)
}

// GetItemHandler handles GET requests for retrieving an item from the inventory.
func (h *InventoryHandler) GetItemHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars[