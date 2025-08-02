// 代码生成时间: 2025-08-02 18:11:03
package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

// 定义一个简单的数据结构
type User struct {
    ID    int    "json:"id""
    Name  string "json:"name""
    Email string "json:"email""
}

// 用户列表
var users = []User{
    {ID: 1, Name: "John", Email: "john@example.com"},
    {ID: 2, Name: "Jane", Email: "jane@example.com"},
}

// GetAllUsers 处理 GET /users 请求
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    // 将用户列表序列化为 JSON
    json.NewEncoder(w).Encode(users)
}

// GetUserByID 处理 GET /users/{id} 请求
func GetUserByID(w http.ResponseWriter, r *http.Request) {
    // 从请求中提取用户 ID
    var userID int
    var err error
    userID, err = strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    
    // 查找指定的用户
    for _, user := range users {
        if user.ID == userID {
            json.NewEncoder(w).Encode(user)
            return
        }
    }
    // 如果没有找到用户
    http.NotFound(w, r)
}

// CreateAUser 处理 POST /users 请求
func CreateAUser(w http.ResponseWriter, r *http.Request) {
    var newUser User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()
    
    // 添加新用户到列表
    newUser.ID = len(users) + 1
    users = append(users, newUser)
    
    // 将新用户序列化为 JSON
    json.NewEncoder(w).Encode(newUser)
}

func main() {
    // 创建一个新的路由
    router := mux.NewRouter()
    
    // 定义路由规则
    router.HandleFunc("/users", GetAllUsers).Methods("GET")
    router.HandleFunc("/users/{id}", GetUserByID).Methods("GET")
    router.HandleFunc("/users", CreateAUser).Methods("POST")
    
    // 启动服务器
    http.Handle("/", router)
    http.ListenAndServe(":8080", nil)
}
