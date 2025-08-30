// 代码生成时间: 2025-08-31 07:20:24
package main

import (
    "fmt"
    "math/rand"
    "time"
    "net/http"
    "github.com/gorilla/mux"
)

// TestDataGenerator 结构体用于生成测试数据
type TestDataGenerator struct {
    // 可以添加更多的字段以扩展生成器的功能
}

// GenerateName 生成随机名称
func (g *TestDataGenerator) GenerateName() string {
    names := []string{"John", "Jane", "Doe", "Alice", "Bob"}
    return names[rand.Intn(len(names))]
}

// GenerateEmail 生成随机邮箱
func (g *TestDataGenerator) GenerateEmail() string {
    domains := []string{"gmail.com", "yahoo.com", "hotmail.com"}
    return fmt.Sprintf("%s%d@%s", g.GenerateName(), rand.Intn(10000), domains[rand.Intn(len(domains))])
}

// GenerateUserData 生成一组用户数据
func (g *TestDataGenerator) GenerateUserData(count int) []map[string]string {
    userData := make([]map[string]string, count)
    for i := 0; i < count; i++ {
        data := map[string]string{
            "name": g.GenerateName(),
            "email": g.GenerateEmail(),
        }
        userData[i] = data
    }
    return userData
}

// setupRouter 设置路由
func setupRouter() *mux.Router {
    router := mux.NewRouter()
    // 添加更多的路由
    router.HandleFunc("/generate", generateHandler).Methods("GET")
    return router
}

// generateHandler 处理数据生成请求
func generateHandler(w http.ResponseWriter, r *http.Request) {
    // 解析查询参数，例如用户数量
    count := 10 // 默认生成10个用户数据
    if values := r.URL.Query(); values.Has("count") {
        if c, err := strconv.Atoi(values.Get("count")); err == nil {
            count = c
        }
    }
    // 创建测试数据生成器实例
    generator := TestDataGenerator{}
    // 生成用户数据
    userData := generator.GenerateUserData(count)
    // 将用户数据以JSON格式返回
    if err := json.NewEncoder(w).Encode(userData); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
    router := setupRouter()
    fmt.Println("Server is running on port 8080")
    // 启动服务器
    if err := http.ListenAndServe(":8080", router); err != nil {
        fmt.Println("Server startup failed: ", err)
    }
}
