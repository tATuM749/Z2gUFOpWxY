// 代码生成时间: 2025-08-15 09:04:47
package main

import (
    "fmt"
    "net/http"
    "text/template"

    "github.com/gorilla/mux"
)

// LayoutData 定义了响应式布局需要的数据结构
type LayoutData struct {
    Title string
    Content string
}

// layoutTemplate 是用于渲染响应式布局的模板
var layoutTemplate = `
<html>
<head>
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <div class="container">
        <h1>{{.Title}}</h1>
        <p>{{.Content}}</p>
    </div>
</body>
</html>
`

// NewRouter 创建并配置 Gorilla Mux 路由器
func NewRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/", homeHandler).Methods("GET")
    return router
}

// homeHandler 处理根路由的请求，渲染响应式布局
func homeHandler(w http.ResponseWriter, r *http.Request) {
    data := LayoutData{
        Title: "Responsive Layout",
        Content: "This is a responsive layout using Gorilla Mux and GoLang.",
    }
    tmpl, err := template.New("layout").Parse(layoutTemplate)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    router := NewRouter()
    fmt.Println("Server is running on port 8080")
    // 监听端口 8080 并处理请求
    err := http.ListenAndServe(":8080", router)
    if err != nil {
        fmt.Printf("Server failed to start: %s
", err)
    }
}
