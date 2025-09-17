// 代码生成时间: 2025-09-18 01:36:15
作者：[您的名字]
日期：[今天的日期]
*/

package main

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "log"
    "encoding/json"
    "github.com/gorilla/mux"
)

// ErrorResponse 结构体用于封装错误响应
type ErrorResponse struct {
    Error string `json:"error"`
}

// RenameRequest 结构体用于封装重命名请求
type RenameRequest struct {
    Files []string `json:"files"`
    Pattern string `json:"pattern"`
}

// RenameResponse 结构体用于封装重命名响应
type RenameResponse struct {
    Files []string `json:"files"`
}

// renameFiles 函数接受文件名列表和命名模式，返回重命名结果
func renameFiles(files []string, pattern string) ([]string, error) {
    var renamedFiles []string
    for i, filename := range files {
        base := filepath.Base(filename)
        newFilename := fmt.Sprintf(pattern, i+1)
        newFilepath := filepath.Dir(filename) + "/" + newFilename
        if err := os.Rename(filename, newFilepath); err != nil {
            return nil, err
        }
        renamedFiles = append(renamedFiles, newFilepath)
    }
    return renamedFiles, nil
}

// handleRename 处理HTTP POST请求，实现文件重命名
func handleRename(w http.ResponseWriter, r *http.Request) {
    var req RenameRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, fmt.Sprintf("Error decoding request: %s", err), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    renamedFiles, err := renameFiles(req.Files, req.Pattern)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error renaming files: %s", err), http.StatusInternalServerError)
        return
    }

    resp := RenameResponse{Files: renamedFiles}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/rename", handleRename).Methods("POST")

    fmt.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}