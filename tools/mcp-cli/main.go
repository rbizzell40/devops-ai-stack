package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

type Trigger struct {
    Event string `json:"event"`
    Data  string `json:"data"`
}

func triggerAction(event, data string) {
    trigger := Trigger{Event: event, Data: data}
    payload, err := json.Marshal(trigger)
    if err != nil {
        fmt.Println("❌ Failed to encode payload:", err)
        os.Exit(1)
    }

    resp, err := http.Post("http://localhost:5678/webhook/mcp", "application/json", bytes.NewBuffer(payload))
    if err != nil {
        fmt.Println("❌ Failed to send trigger:", err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    fmt.Println("✅ Trigger sent. Status:", resp.Status)
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: mcp-cli <event-name> [data]")
        os.Exit(1)
    }
    event := os.Args[1]
    data := ""
    if len(os.Args) > 2 {
        data = os.Args[2]
    }

    triggerAction(event, data)
}