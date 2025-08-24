// 代码生成时间: 2025-08-24 10:59:34
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    "github.com/gorilla/mux"
    "github.com/robfig/cron/v3"
)

// Scheduler defines a struct to hold the cron instance
type Scheduler struct {
    cron *cron.Cron
}

// NewScheduler creates a new scheduler instance
func NewScheduler() *Scheduler {
    return &Scheduler{
        cron: cron.New(cron.WithSeconds()), // WithSeconds() enables seconds in cron schedule
    }
}

// StartCron starts the cron scheduler
func (s *Scheduler) StartCron() {
    s.cron.Start()
}

// AddJob adds a job to the scheduler with a specific schedule
func (s *Scheduler) AddJob(schedule string, cmd func()) error {
    _, err := s.cron.AddFunc(schedule, cmd)
    return err
}

// HTTPServer defines a struct to hold the HTTP server configurations
type HTTPServer struct {
    router *mux.Router
    scheduler *Scheduler
}

// NewHTTPServer creates a new HTTP server instance
func NewHTTPServer(scheduler *Scheduler) *HTTPServer {
    return &HTTPServer{
        router: mux.NewRouter(),
        scheduler: scheduler,
    }
}

// RunServer starts the HTTP server
func (h *HTTPServer) RunServer(port string) error {
    fmt.Println("Starting HTTP server on port", port)
    log.Fatal(http.ListenAndServe(":" + port, h.router))
    return nil // This will never be reached due to ListenAndServe
}

// SetupRoutes sets up the routes for the HTTP server
func (h *HTTPServer) SetupRoutes() {
    h.router.HandleFunc("/add-job", h.addJobHandler).Methods("POST")
    h.router.HandleFunc("/", h.healthCheckHandler).Methods("GET")
}

// addJobHandler handles the adding of a new job to the scheduler
func (h *HTTPServer) addJobHandler(w http.ResponseWriter, r *http.Request) {
    // Implement your logic here to add a new job to the scheduler
    // For example, you could parse the request body and add a job
    // For now, we'll just return a success message
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Job added successfully")
}

// healthCheckHandler handles the health check endpoint
func (h *HTTPServer) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Scheduler is running")
}

func main() {
    scheduler := NewScheduler()
    scheduler.StartCron()
    
    // Add your jobs here, for example:
    // err := scheduler.AddJob("*/1 * * * *", func() { log.Println("Running scheduled job") })
    // if err != nil {
    //     log.Fatal("Failed to add job: ", err)
    // }
    
    httpServer := NewHTTPServer(scheduler)
    httpServer.SetupRoutes()
    
    if err := httpServer.RunServer("8080"); err != nil {
        log.Fatal("Failed to run server: ", err)
    }
}
