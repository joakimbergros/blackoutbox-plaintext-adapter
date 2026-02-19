package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// This is the shape YOUR API expects
type DocumentPayload struct {
	FileId   string   `json:"file_id"`
	Document string   `json:"document"` // text content
	PrintAt  *int64   `json:"print_at,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// Configuration via environment variables
	apiBase := os.Getenv("API_BASE_URL") // e.g. http://localhost:3000
	systemId := os.Getenv("SYSTEM_ID")   // e.g. system-123

	if apiBase == "" || systemId == "" {
		log.Fatal("API_BASE_URL and SYSTEM_ID must be set")
	}

	// ---- Example text document from external system ----
	textDocument := `Hello,
This is a text document coming from an external system.
It will be stored as a .txt file.`

	documents := []DocumentPayload{
		{
			FileId:   "DOC-001",
			Document: textDocument,
			Tags:     []string{"example", "text"},
		},
	}

	// Convert to JSON
	body, err := json.Marshal(documents)
	if err != nil {
		log.Fatal(err)
	}

	// Send to your API
	url := apiBase + "/systems/" + systemId + "/sync"

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		log.Fatalf("sync failed: %s", resp.Status)
	}

	log.Println("Adapter: text document synced successfully")
}
