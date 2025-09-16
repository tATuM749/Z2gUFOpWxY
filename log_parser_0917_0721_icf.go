// 代码生成时间: 2025-09-17 07:21:15
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/gorilla/mux" // 引入gorilla mux框架
)

// LogEntry 定义日志文件中的单个条目
type LogEntry struct {
    Timestamp time.Time
    Level     string
    Message   string
}

// parseLogEntry 解析日志条目
func parseLogEntry(line string) (*LogEntry, error) {
    parts := strings.Fields(line)
    if len(parts) < 3 {
        return nil, fmt.Errorf("invalid log entry: %s", line)
    }

    // 解析时间戳
    timestamp, err := time.Parse("2006-01-02 15:04:05", parts[0] + " " + parts[1])
    if err != nil {
        return nil, fmt.Errorf("failed to parse timestamp: %s", err)
    }

    // 解析日志级别和消息
    level := parts[2]
    message := strings.Join(parts[3:], " ")

    return &LogEntry{Timestamp: timestamp, Level: level, Message: message}, nil
}

// LogParser 结构体包含解析器的配置
type LogParser struct {
    LogFilePath string
}

// Parse 解析日志文件
func (p *LogParser) Parse() ([]LogEntry, error) {
    file, err := os.Open(p.LogFilePath)
    if err != nil {
        return nil, fmt.Errorf("failed to open log file: %s", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var entries []LogEntry
    for scanner.Scan() {
        line := scanner.Text()
        entry, err := parseLogEntry(line)
        if err != nil {
            log.Printf("skipping invalid log entry: %s", err)
            continue
        }
        entries = append(entries, *entry)
    }
    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("failed to scan log file: %s", err)
    }
    return entries, nil
}

// setupRouter 设置路由
func setupRouter(r *mux.Router) {
    r.HandleFunc("/parse", func(w http.ResponseWriter, req *http.Request) {
        logFilePath := req.FormValue("log_file_path")
        if logFilePath == "" {
            http.Error(w, "log file path is required", http.StatusBadRequest)
            return
        }

        parser := LogParser{LogFilePath: logFilePath}
        entries, err := parser.Parse()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // 将解析结果转换为JSON并返回
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(entries)
    })
}

func main() {
    r := mux.NewRouter()
    setupRouter(r)

    log.Printf("Starting log parser server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("failed to start server: %s", err)
    }
}