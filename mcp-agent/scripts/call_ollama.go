package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://localhost:11434/api/generate"

	// JSON bod
	requestBody := map[string]string{
		"model":  "llama2",
		"prompt": "Summarize this alert: CPU spike on node 10",
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	// Send HTTP POST
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response from Ollama:")
	fmt.Println(string(body))
}
