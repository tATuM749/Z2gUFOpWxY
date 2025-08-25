// 代码生成时间: 2025-08-25 16:14:12
It demonstrates how to create a RESTful API that calculates and returns statistical
information about a given dataset.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
"github.com/gorilla/mux"
)

// Dataset represents a collection of numerical data points.
type Dataset []float64

// Statistics contains calculated statistical measures.
type Statistics struct {
	Mean     float64 `json:"mean"`
	Median   float64 `json:"median"`
	Mode     float64 `json:"mode"`
	Variance float64 `json:"variance"`
	StdDev   float64 `json:"std_dev"`
}

// calculateMean calculates the mean of the dataset.
func calculateMean(data Dataset) float64 {
	sum := 0.0
	for _, val := range data {
		sum += val
	}
	return sum / float64(len(data))
}

// calculateMedian calculates the median of the dataset.
func calculateMedian(data Dataset) float64 {
	n := len(data)
	sort.Float64s(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2
	}
	return data[n/2]
}

// calculateMode calculates the mode of the dataset.
func calculateMode(data Dataset) float64 {
	elementCounts := make(map[float64]int)
	for _, num := range data {
		elementCounts[num]++
	}
	maxCount := 0
	mode := 0.0
	for num, count := range elementCounts {
		if count > maxCount {
			maxCount = count
			mode = num
		}
	}
	return mode
}

// calculateVariance calculates the variance of the dataset.
func calculateVariance(data Dataset, mean float64) float64 {
	sum := 0.0
	for _, val := range data {
		sum += math.Pow(val-mean, 2)
	}
	return sum / float64(len(data)-1)
}

// calculateStdDev calculates the standard deviation of the dataset.
func calculateStdDev(data Dataset, variance float64) float64 {
	return math.Sqrt(variance)
}

// analyzeData calculates and returns the statistical measures of the dataset.
func analyzeData(w http.ResponseWriter, r *http.Request) {
	var data Dataset
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	mean := calculateMean(data)
	median := calculateMedian(data)
	mode := calculateMode(data)
	variance := calculateVariance(data, mean)
	stdDev := calculateStdDev(data, variance)

	stats := Statistics{Mean: mean, Median: median, Mode: mode, Variance: variance, StdDev: stdDev}
	json.NewEncoder(w).Encode(stats)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/analyze", analyzeData).Methods("POST")

	fmt.Println("Starting the statistical data analyzer on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}