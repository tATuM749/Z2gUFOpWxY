// 代码生成时间: 2025-08-26 16:25:41
package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "github.com/gorilla/context"
)

// User represents a user with access control
type User struct {
    Username string
    Role     string
}

// isLoggedIn checks if the user is logged in
func isLoggedIn(user User) bool {
    // Implement actual login check logic here
    return user.Username != ""
}

// isAdmin checks if the user is an admin
func isAdmin(user User) bool {
    // Implement actual role check logic here
    return user.Role == "admin"
}

// Middleware to check for admin access
func AdminMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user, ok := context.GetOk(r, "user").(User)
        if !ok || !isLoggedIn(user) || !isAdmin(user) {
            http.Error(w, "Access denied", http.StatusForbidden)
            return
        }
        next.ServeHTTP(w, r)
    })
}

// Middleware to check for user access
func UserMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user, ok := context.GetOk(r, "user").(User)
        if !ok || !isLoggedIn(user) {
            http.Error(w, "Access denied", http.StatusForbidden)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
        // Admin-only route handler
        w.Write([]byte("Welcome, Admin!"))
    }).Methods("GET").Name("AdminRoute")
    r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
        // User-only route handler
        w.Write([]byte("Welcome, User!"))
    }).Methods("GET\).Name("UserRoute")

    // Apply middleware to routes
    r.NewRoute().Name("AdminRoute\).HandlerFunc(AdminMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Admin route handler logic
        w.Write([]byte("Admin section"))
    }))
    r.NewRoute().Name("UserRoute\).HandlerFunc(UserMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // User route handler logic
        w.Write([]byte("User section"))
    }))

    log.Fatal(http.ListenAndServe(":8080", context.ClearHandler(r)))
}