// 代码生成时间: 2025-08-09 17:43:05
// sql_injection_prevention.go

package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"log"
	"net/http"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/mux"
)

// Database configuration
const (
	host     = "localhost"
	pw       = "password"
	dbName   = "mydb"
	user     = "root"
	port     = 3306
	dialect  = "mysql"
)

// App holds application configuration
type App struct {
	db *sql.DB
}
type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewApp(db *sql.DB) *App {
	return &App{db: db}
}

// SignInHandler handles user sign in
func (a *App) SignInHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Prevent SQL injection by using parameterized queries
	stmt, err := a.db.Prepare("SELECT * FROM users WHERE username = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var hashedPassword string
	err = stmt.QueryRow(username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Handle successful login
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte("{"message": "User authenticated successfully"}"))
	if err != nil {
		log.Println("Error writing to response writer: ", err)
		return
	}
}

func main() {
	// Connect to the database
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local", user, pw, host, port, dbName)
	db, err := sql.Open(dialect, dataSourceName)
	if err != nil {
		log.Fatal("You encountered an error while opening the database: ", err.Error())
	}
	defer db.Close()

	// Create a new application instance
	app := NewApp(db)

	// Setup the router
	router := mux.NewRouter()
	router.HandleFunc("/signin", app.SignInHandler).Methods("POST")

	// Start the server
	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}