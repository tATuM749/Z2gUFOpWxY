// 代码生成时间: 2025-09-24 01:12:38
package main

import (
    "fmt"
    "net/http"
    "os"
    "text/template"
    "strings"

    "github.com/gorilla/mux"
)

// TestReportData holds the data for the test report
type TestReportData struct {
    Tests      []string 
    Failures   int
    Successes  int
    Skipped    int
}

// ReportGenerator is a function that generates a test report
func ReportGenerator(w http.ResponseWriter, r *http.Request) {
    var data TestReportData
    tests := r.URL.Query().Get("tests")
    failures := r.URL.Query().Get("failures")
    successes := r.URL.Query().Get("successes")
    skipped := r.URL.Query().Get("skipped")
    
    // Parsing the query parameters to int
    data.Failures, _ = strconv.Atoi(failures)
    data.Successes, _ = strconv.Atoi(successes)
    data.Skipped, _ = strconv.Atoi(skipped)
    
    // Splitting the tests query parameter by comma to get individual test names
    data.Tests = strings.Split(tests, ",")
    
    // Loading the template for the report
    tmpl, err := template.ParseFiles("report.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    // Executing the template with the data
    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/report", ReportGenerator).Methods("GET")

    // Serving static files like HTML templates
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

    // Starting the server
    fmt.Println("Starting the test report generator on port 8080")
    err := http.ListenAndServe(":8080", r)
    if err != nil {
        fmt.Println("Error starting server: ", err)
        os.Exit(1)
    }
}
