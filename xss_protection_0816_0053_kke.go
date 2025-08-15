// 代码生成时间: 2025-08-16 00:53:57
package main

import (
    "net/http"
    "html"
    "log"
    "github.com/gorilla/mux"
)

// XSSProtectionHandler 是一个处理XSS攻击防护的中间件
func XSSProtectionHandler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 设置HTTP头部，禁用不安全的头部
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; object-src 'none'")
        // 调用下一个中间件/处理函数
        next.ServeHTTP(w, r)
    })
}

// sanitizeInput 函数用于清理输入，防护XSS攻击
func sanitizeInput(input string) string {
    // 转义HTML标签，防止XSS攻击
    return html.EscapeString(input)
}

func main() {
    // 创建一个新的路由器
    router := mux.NewRouter()
    
    // 使用XSSProtectionHandler中间件
    router.Use(XSSProtectionHandler)

    // 定义一个简单的GET请求处理函数
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // 获取查询参数
        param := r.URL.Query().Get("param")
        if param == "" {
            param = "default value"
        }

        // 清理输入参数，防护XSS攻击
        sanitizedParam := sanitizeInput(param)
        
        // 将清理后的参数写入响应体
        _, err := w.Write([]byte("You entered: " + sanitizedParam))
        if err != nil {
            log.Printf("Error writing response: %v", err)
        }
    })

    // 启动服务器
    log.Println("Server starting on port 8080")
    err := http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
