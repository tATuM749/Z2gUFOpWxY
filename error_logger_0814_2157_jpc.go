// 代码生成时间: 2025-08-14 21:57:21
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
    "time"
    "github.com/gorilla/mux"
)

// ErrorLoggerMiddleware 是一个中间件，用于记录HTTP请求中的错误。
func ErrorLoggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        rw := responseWriter{ResponseWriter: w}
        next.ServeHTTP(&rw, r)
        if rw.status >= http.StatusBadRequest {
            log.Printf("Error: %d - %s %s - %s", rw.status, r.Method, r.URL.Path, time.Since(start))
        }
    })
}

// responseWriter 包装了http.ResponseWriter，允许我们捕获HTTP状态码。
type responseWriter struct {
    http.ResponseWriter
    status int
}

// WriteHeader 实现了http.ResponseWriter接口，用于捕获写入的状态码。
func (rw *responseWriter) WriteHeader(status int) {
    rw.status = status
    rw.ResponseWriter.WriteHeader(status)
}

// main 函数启动HTTP服务器，并设置路由和中间件。
func main() {
    r := mux.NewRouter()
    r.Use(ErrorLoggerMiddleware)

    // 这里可以添加更多的路由处理函数
    r.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
        // 模拟一个错误
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintln(w, "Internal Server Error")
    })

    // 监听端口并启动服务器
    log.Printf("Server starting on port 8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}
