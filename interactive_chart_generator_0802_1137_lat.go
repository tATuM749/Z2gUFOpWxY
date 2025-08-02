// 代码生成时间: 2025-08-02 11:37:27
 * interactive_chart_generator.go
 * This program uses the Gorilla WebSocket library to create an interactive chart generator.
 * It allows for real-time data updates and interactions with the chart.
 */

package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

// Define the upgrader for websocket connections
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

// ChartData represents the data structure for the chart
type ChartData struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
}

// ChartUpdate is the structure for sending updates to the chart
type ChartUpdate struct {
    Data []ChartData `json:"data"`
}

// handleWebSocket handles the websocket connections for interactive chart data
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()

    // Register our new connection and pass it to the hub
    // This would be part of a larger system with a hub to manage multiple clients
    // For simplicity, this example does not implement the hub
    // However, in a real-world scenario, you would use a hub here
    // to manage the broadcasting of data to all connected clients.

    // Read data from the connection
    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("read:", err)
            break
        }
        log.Printf("recv: %s", message)

        // Here you would parse the message and update the chart accordingly
        // For now, we just echo back the received message
        if err = conn.WriteMessage(messageType, message); err != nil {
            log.Println("write: