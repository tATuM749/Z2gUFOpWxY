// 代码生成时间: 2025-08-18 07:58:54
// responsive_layout.go

package main

import (
    "fmt"
    "log"
    "net/http"
    "text/template"

    "github.com/gorilla/mux"
)

// HTML templates
const (
    indexTemplate = "index.html"
)

// App 定义了应用程序的结构
type App struct {
    templates *template.Template
    router    *mux.Router
}

// NewApp 创建并配置一个新的App实例
func NewApp() *App {
    r := mux.NewRouter()
    return &App{
        router:    r,
        templates: template.Must(template.ParseGlob("templates/*.html")),
    }
}

// Start 启动应用程序
func (app *App) Start(port string) {
    fs := http.FileServer(http.Dir("public"))
    app.router.PathPrefix("/public/").Handler(fs)
    
    app.router.HandleFunc("/", app.homeHandler).Methods("GET")
    
    log.Printf("Server starting on port %s
", port)
    log.Fatal(http.ListenAndServe(":" + port, app.router))
}

// homeHandler 处理首页请求
func (app *App) homeHandler(w http.ResponseWriter, r *http.Request) {
    err := app.templates.ExecuteTemplate(w, indexTemplate, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func main() {
    app := NewApp()
    app.Start("8080")
}
