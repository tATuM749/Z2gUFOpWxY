// 代码生成时间: 2025-08-28 23:24:26
package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

// Order represents the structure of an order.
type Order struct {
    ID      int    "json:"id""
    Product string "json:"product""
    Quantity int   "json:"quantity""
}

// OrderService is an interface that defines the methods for order processing.
type OrderService interface {
    CreateOrder(order Order) (int, error)
    ProcessOrder(orderId int) error
}

// InMemoryOrderService implements the OrderService interface using in-memory storage.
type InMemoryOrderService struct {
    orders map[int]Order
}

// NewInMemoryOrderService creates a new instance of InMemoryOrderService.
func NewInMemoryOrderService() *InMemoryOrderService {
    return &InMemoryOrderService{
        orders: make(map[int]Order),
    }
}

// CreateOrder adds a new order to the in-memory storage.
func (s *InMemoryOrderService) CreateOrder(order Order) (int, error) {
    if _, exists := s.orders[order.ID]; exists {
        return 0, fmt.Errorf("order with ID %d already exists", order.ID)
    }
    s.orders[order.ID] = order
    return order.ID, nil
}

// ProcessOrder processes an existing order by updating its status to processed.
func (s *InMemoryOrderService) ProcessOrder(orderId int) error {
    if _, exists := s.orders[orderId]; !exists {
        return fmt.Errorf("order with ID %d not found", orderId)
    }
    s.orders[orderId].Quantity = 0 // Simulating order processing
    return nil
}

// OrderHandler handles HTTP requests related to orders.
func OrderHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        var order Order
        if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        service := NewInMemoryOrderService()
        orderId, err := service.CreateOrder(order)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(map[string]int{"id": orderId})
    case http.MethodGet:
        vars := mux.Vars(r)
        orderId, err := strconv.Atoi(vars["id"])
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        service := NewInMemoryOrderService()
        if err := service.ProcessOrder(orderId); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusOK)
    }
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/orders", OrderHandler).Methods("POST")
    router.HandleFunc("/orders/{id}", OrderHandler).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", router))
}