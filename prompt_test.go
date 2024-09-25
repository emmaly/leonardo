package leonardo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGenerateRandomPrompt tests the GenerateRandomPrompt method.
func TestGenerateRandomPrompt(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the endpoint and method
		if r.URL.Path != "/prompt/random" || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Mock successful generation response
		w.Header().Set("Content-Type", "application/json")
		response := GenerateRandomPromptResponse{
			PromptGeneration: &struct {
				APICreditCost *int    `json:"apiCreditCost"`
				Prompt        *string `json:"prompt"`
			}{
				APICreditCost: Ptr(4),
				Prompt:        Ptr("A vibrant sunset over a tranquil lake."),
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
	client.Prompt = client.NewPromptService()

	// Execute GenerateRandomPrompt
	ctx := context.Background()
	resp, err := client.Prompt.GenerateRandomPrompt(ctx)
	if err != nil {
		t.Fatalf("GenerateRandomPrompt failed: %v", err)
	}

	// Validate response
	if resp.PromptGeneration == nil {
		t.Fatal("Expected PromptGeneration in response, got nil")
	}
	expectedPrompt := "A vibrant sunset over a tranquil lake."
	if resp.PromptGeneration.Prompt == nil || *resp.PromptGeneration.Prompt != expectedPrompt {
		t.Errorf("Expected prompt '%s', got '%v'", expectedPrompt, resp.PromptGeneration.Prompt)
	}
	if resp.PromptGeneration.APICreditCost == nil || *resp.PromptGeneration.APICreditCost != 4 {
		t.Errorf("Expected APICreditCost 4, got '%v'", resp.PromptGeneration.APICreditCost)
	}
}

// TestImprovePrompt tests the ImprovePrompt method.
func TestImprovePrompt(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the endpoint and method
		if r.URL.Path != "/prompt/improve" || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode the request body
		var req ImprovePromptRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		expectedOriginalPrompt := "A serene landscape with mountains."
		if req.Prompt == nil || *req.Prompt != expectedOriginalPrompt {
			t.Errorf("Expected prompt '%s', got '%v'", expectedOriginalPrompt, req.Prompt)
		}

		// Mock successful improvement response
		w.Header().Set("Content-Type", "application/json")
		response := ImprovePromptResponse{
			PromptGeneration: &struct {
				APICreditCost *int    `json:"apiCreditCost"`
				Prompt        *string `json:"prompt"`
			}{
				APICreditCost: Ptr(5),
				Prompt:        Ptr("A tranquil mountain landscape at sunrise, with soft light and mist."),
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
	client.Prompt = client.NewPromptService()

	// Prepare ImprovePromptRequest
	originalPrompt := "A serene landscape with mountains."
	improveReq := ImprovePromptRequest{
		Prompt: &originalPrompt,
	}

	// Execute ImprovePrompt
	ctx := context.Background()
	resp, err := client.Prompt.ImprovePrompt(ctx, improveReq)
	if err != nil {
		t.Fatalf("ImprovePrompt failed: %v", err)
	}

	// Validate response
	if resp.PromptGeneration == nil {
		t.Fatal("Expected PromptGeneration in response, got nil")
	}
	expectedImprovedPrompt := "A tranquil mountain landscape at sunrise, with soft light and mist."
	if resp.PromptGeneration.Prompt == nil || *resp.PromptGeneration.Prompt != expectedImprovedPrompt {
		t.Errorf("Expected improved prompt '%s', got '%v'", expectedImprovedPrompt, resp.PromptGeneration.Prompt)
	}
	if resp.PromptGeneration.APICreditCost == nil || *resp.PromptGeneration.APICreditCost != 5 {
		t.Errorf("Expected APICreditCost 5, got '%v'", resp.PromptGeneration.APICreditCost)
	}
}
