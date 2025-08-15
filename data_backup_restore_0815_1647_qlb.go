// 代码生成时间: 2025-08-15 16:47:26
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "path/filepath"
    "time"

    "github.com/gorilla/mux"
)

// Backup represents the backup data structure
type Backup struct {
    Name    string
    Content string
    Created time.Time
}

// NewBackup creates a new backup
func NewBackup(name string, content string) *Backup {
    return &Backup{
        Name:    name,
        Content: content,
        Created: time.Now(),
    }
}

// BackupService handles backup and restore operations
type BackupService struct {
    backups map[string]*Backup
}

// NewBackupService creates a new BackupService
func NewBackupService() *BackupService {
    return &BackupService{
        backups: make(map[string]*Backup),
    }
}

// Backup performs a backup operation
func (s *BackupService) Backup(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]
    content, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    backup := NewBackup(name, string(content))
    s.backups[backup.Name] = backup
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Backup created successfully"))
}

// Restore performs a restore operation
func (s *BackupService) Restore(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]
    backup, exists := s.backups[name]
    if !exists {
        http.Error(w, "Backup not found", http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(backup.Content))
}

func main() {
    router := mux.NewRouter()
    service := NewBackupService()

    // Register backup and restore routes
    router.HandleFunc("/backup/{name}", service.Backup).Methods("POST")
    router.HandleFunc("/restore/{name}", service.Restore).Methods("GET")

    // Start the server
    http.ListenAndServe(":8080", router)
}
