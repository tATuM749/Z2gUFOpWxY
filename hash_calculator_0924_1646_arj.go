// 代码生成时间: 2025-09-24 16:46:18
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "log"
    "net/http"
    "strings"
)

// HashCalculator 定义哈希值计算工具的结构体
type HashCalculator struct{}

// CalculateSHA256 计算SHA-256哈希值
func (h *HashCalculator) CalculateSHA256(input string) (string, error) {
    sha := sha256.New()
    _, err := sha.Write([]byte(input))
    if err != nil {
        return "", err
    }
    hashBytes := sha.Sum(nil)
    return hex.EncodeToString(hashBytes), nil
}

// HashHandler 处理哈希计算请求
func HashHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    decoder := json.NewDecoder(r.Body)
    var data struct{
        Input string `json:"input"`
    }
    if err := decoder.Decode(&data); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    if data.Input == "" {
        http.Error(w, "Input cannot be empty", http.StatusBadRequest)
        return
    }

    calculator := &HashCalculator{}
    hash, err := calculator.CalculateSHA256(data.Input)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, {"json": "{"filename": "%s", "hash": "%s"}"}, "hash_calculator_result", hash)
}

func main() {
    router := mux.NewRouter()
    // 定义路由，当接收到POST请求时计算哈希值
    router.HandleFunc("/hash", HashHandler).Methods("POST")

    // 启动服务器
    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
