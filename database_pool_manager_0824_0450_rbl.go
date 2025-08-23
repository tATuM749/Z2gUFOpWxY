// 代码生成时间: 2025-08-24 04:50:00
package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"
    \_ "encoding/json"
    "github.com/go-sql-driver/mysql"
)

// DatabaseConfig 用于存储数据库连接的配置信息
type DatabaseConfig struct {
    Host     string
    Port     int
    Username string
    Password string
    DBName   string
}

// NewDatabasePool 创建一个新的数据库连接池
func NewDatabasePool(config DatabaseConfig) (*sql.DB, error) {
    // 构建连接字符串
    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.Username, config.Password, config.Host, config.Port, config.DBName)

    // 创建数据库连接
    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        return nil, err
    }

    // 设置数据库连接池的最大打开连接数
    db.SetMaxOpenConns(25)

    // 设置数据库连接池的最大空闲连接数
    db.SetMaxIdleConns(25)

    // 设置了连接可复用的最大时间
    db.SetConnMaxLifetime(5 * time.Minute)

    // 测试数据库连接是否成功
    err = db.Ping()
    if err != nil {
        db.Close()
        return nil, err
    }

    return db, nil
}

func main() {
    // 定义数据库配置
    config := DatabaseConfig{
        Host:     "localhost",
        Port:     3306,
        Username: "your_username",
        Password: "your_password",
        DBName:   "your_dbname",
    }

    // 创建数据库连接池
    db, err := NewDatabasePool(config)
    if err != nil {
        log.Fatalf("Failed to create database pool: %s", err)
    }
    defer db.Close()

    // 在这里可以执行数据库操作，例如查询、插入等
    // ...

    fmt.Println("Database pool created successfully.")
}
