// 代码生成时间: 2025-08-08 04:35:51
package main

import (
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "gopkg.in/gorilla/schema.v2"
)

// 数据库迁移工具
type MigrationTool struct {
    db *gorm.DB
}

// NewMigrationTool 创建MigrationTool实例
func NewMigrationTool() *MigrationTool {
    // 连接数据库
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database: ", err)
    }

    // 自动迁移模式
    db.AutoMigrate(&User{})

    return &MigrationTool{db: db}
}

// User 定义用户模型
type User struct {
    gorm.Model
    Name string
    Email string `gorm:"type:varchar(100);uniqueIndex"`
}

// Up 执行数据库迁移
func (m *MigrationTool) Up() error {
    // 迁移用户表
    if err := m.db.AutoMigrate(&User{}); err != nil {
        return fmt.Errorf("failed to migrate user table: %w", err)
    }
    fmt.Println("Database migration completed successfully.")
    return nil
}

// Down 回滚数据库迁移
func (m *MigrationTool) Down() error {
    // 回滚用户表
    if err := m.db.Migrator().DropTable("users"); err != nil {
        return fmt.Errorf("failed to drop user table: %w", err)
    }
    fmt.Println("Database migration rolled back successfully.")
    return nil
}

func main() {
    // 创建数据库迁移工具实例
    migrationTool := NewMigrationTool()
    
    // 执行数据库迁移
    if err := migrationTool.Up(); err != nil {
        log.Fatalf("failed to execute database migration: %s", err)
    }
    
    // 回滚数据库迁移
    if err := migrationTool.Down(); err != nil {
        log.Fatalf("failed to rollback database migration: %s", err)
    }
}
