// 代码生成时间: 2025-08-30 12:46:27
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "sync"
    "time"

    "github.com/gorilla/mux"
)

// BackupSyncService 定义文件备份和同步服务
type BackupSyncService struct {
    srcPath  string
    destPath string
    lock     sync.Mutex
}

// NewBackupSyncService 创建备份和同步服务实例
func NewBackupSyncService(src, dest string) *BackupSyncService {
    return &BackupSyncService{
        srcPath:  src,
        destPath: dest,
    }
}

// Start 开始备份和同步操作
func (s *BackupSyncService) Start() {
    s.lock.Lock()
    defer s.lock.Unlock()

    // 检查源路径和目标路径是否存在
    if _, err := os.Stat(s.srcPath); os.IsNotExist(err) {
        log.Fatalf("Source path does not exist: %s", s.srcPath)
    }
    if _, err := os.Stat(s.destPath); os.IsNotExist(err) {
        log.Fatalf("Destination path does not exist: %s", s.destPath)
    }

    // 递归备份和同步文件夹
    s.syncFolders(s.srcPath, s.destPath)
}

// syncFolders 同步两个文件夹
func (s *BackupSyncService) syncFolders(src, dest string) {
    // 读取源文件夹中的所有文件和文件夹
    files, err := os.ReadDir(src)
    if err != nil {
        log.Fatalf("Failed to read source directory: %v", err)
    }

    for _, file := range files {
        srcFile := filepath.Join(src, file.Name())
        destFile := filepath.Join(dest, file.Name())

        // 如果是文件，则直接复制
        if file.IsDir() {
            // 如果目标文件夹不存在，则创建
            if _, err := os.Stat(destFile); os.IsNotExist(err) {
                if err := os.MkdirAll(destFile, 0755); err != nil {
                    log.Fatalf("Failed to create directory: %v", err)
                }
            }
            // 递归同步子文件夹
            s.syncFolders(srcFile, destFile)
        } else {
            // 复制文件
            s.copyFile(srcFile, destFile)
        }
    }
}

// copyFile 复制单个文件
func (s *BackupSyncService) copyFile(src, dest string) {
    srcFile, err := os.Open(src)
    if err != nil {
        log.Fatalf("Failed to open source file: %v", err)
    }
    defer srcFile.Close()

    destFile, err := os.Create(dest)
    if err != nil {
        log.Fatalf("Failed to create destination file: %v", err)
    }
    defer destFile.Close()

    if _, err := io.Copy(destFile, srcFile); err != nil {
        log.Fatalf("Failed to copy file: %v", err)
    }
}

// setupRoutes 设置路由
func setupRoutes(r *mux.Router, service *BackupSyncService) {
    r.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Starting backup and sync...")
        service.Start()
        fmt.Fprintln(w, "Backup and sync completed.")
    })[:
}

func main() {
    r := mux.NewRouter()
    service := NewBackupSyncService("/path/to/source", "/path/to/destination")
    setupRoutes(r, service)

    log.Println("Server is running on port 8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}
