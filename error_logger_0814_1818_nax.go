// 代码生成时间: 2025-08-14 18:18:49
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gorilla/mux"
)

// ErrorLogEntry 定义了错误日志条目的结构
type ErrorLogEntry struct {
    Timestamp time.Time `json:"timestamp"`
    Error     string    `json:"error"`
}

// ErrorLogHandler 处理错误日志的收集
func ErrorLogHandler(w http.ResponseWriter, r *http.Request) {
    // 解析请求体中的错误信息
    var logEntry ErrorLogEntry
    err := json.NewDecoder(r.Body).Decode(&logEntry)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // 检查错误信息是否为空
    if logEntry.Error == "" {
        http.Error(w, "Error message is required", http.StatusBadRequest)
        return
    }

    // 记录错误日志到文件
    logFile, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer logFile.Close()

    // 格式化错误日志条目并写入文件
    _, err = logFile.WriteString(fmt.Sprintf("%s - %s
", logEntry.Timestamp.Format(time.RFC3339), logEntry.Error))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // 响应客户端请求
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Error logged successfully"})
}

func main() {
    // 创建一个新的路由器
    router := mux.NewRouter()

    // 注册错误日志收集器的路由
    router.HandleFunc("/log", ErrorLogHandler).Methods("POST")

    // 启动HTTP服务器
    fmt.Println("Starting error logger on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
