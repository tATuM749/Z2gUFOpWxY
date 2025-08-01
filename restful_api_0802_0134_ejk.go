// 代码生成时间: 2025-08-02 01:34:52
package main

import (
    "net/http"
    "encoding/json"
    "log"
    "github.com/gorilla/mux"
)

// 定义一个简单的用户模型
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// 用户数据的存储
var users = []User{
    {ID: 1, Name: "John Doe", Email: "john@example.com"},
    {ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
}

// GetUsers 响应所有用户的请求
func GetUsers(w http.ResponseWriter, r *http.Request) {
    // 将用户数据序列化为JSON格式
    err := json.NewEncoder(w).Encode(users)
    if err != nil {
        // 错误处理
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// GetUserByID 响应单个用户请求
func GetUserByID(w http.ResponseWriter, r *http.Request) {
    // 从URL参数中提取用户ID
    vars := mux.Vars(r)
    id := vars["id"]

    // 将字符串ID转换为整数
    idInt, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // 查找用户
    for _, user := range users {
        if user.ID == idInt {
            err = json.NewEncoder(w).Encode(user)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
            return
        }
    }

    // 如果没有找到用户，则返回404错误
    http.NotFound(w, r)
}

// main 函数初始化并启动服务器
func main() {
    // 创建一个新的路由器
    router := mux.NewRouter().StrictSlash(true)

    // 注册路由
    router.HandleFunc("/users", GetUsers).Methods("GET")
    router.HandleFunc("/users/{id}", GetUserByID).Methods("GET")

    // 启动服务器
    log.Println("Server is running on port 8080...")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}
