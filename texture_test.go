package leonardo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCreateTextureGeneration tests the CreateTextureGeneration method.
func TestCreateTextureGeneration(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/generations-texture"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req CreateTextureGenerationRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		if req.Prompt == nil || *req.Prompt == "" {
			w.WriteHeader(http.StatusBadRequest)
			resp := APIErrorResponse{
				Code:    "bad-request",
				Message: "Prompt is required.",
			}
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := CreateTextureGenerationResponse{
			TextureGenerationJob: struct {
				APICreditCost *int    `json:"apiCreditCost"`
				ID            *string `json:"id"`
			}{
				APICreditCost: Ptr(12),
				ID:            Ptr("texture-job-001"),
			},
		}
		json.NewEncoder(w).Encode(response)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	// Initialize client with mock server URL
	client := &Client{
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
		APIKey:     "test-api-key",
	}
	client.Texture = client.NewTextureService()

	// Execute CreateTextureGeneration
	ctx := context.Background()
	req := CreateTextureGenerationRequest{
		Prompt: Ptr("A beautiful sunset over the mountains."),
		// Add other fields as needed
	}

	resp, err := client.Texture.CreateTextureGeneration(ctx, req)
	if err != nil {
		t.Fatalf("CreateTextureGeneration failed: %v", err)
	}

	// Validate response
	if resp.TextureGenerationJob.ID == nil || *resp.TextureGenerationJob.ID != "texture-job-001" {
		t.Errorf("Expected Job ID 'texture-job-001', got '%v'", resp.TextureGenerationJob.ID)
	}
	if *resp.TextureGenerationJob.APICreditCost != 12 {
		t.Errorf("Expected APICreditCost 12, got %d", *resp.TextureGenerationJob.APICreditCost)
	}
}
