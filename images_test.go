package leonardo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// TestCreateImageGeneration tests the CreateImageGeneration method.
func TestCreateImageGeneration(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/generations"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req CreateGenerationRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}

		// Validate request
		if req.Prompt == "" {
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
		resp := CreateGenerationResponse{
			SDGenerationJob: struct {
				APICreditCost *int    `json:"apiCreditCost"`
				GenerationID  *string `json:"generationId"`
			}{
				APICreditCost: Ptr(2),
				GenerationID:  Ptr("gen-123"),
			},
		}
		json.NewEncoder(w).Encode(resp)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	// Initialize client with mock server URL
	client := &Client{
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
		APIKey:     "test-api-key",
	}
	client.Images = client.NewImagesService()

	// Execute CreateImageGeneration
	ctx := context.Background()
	createReq := CreateGenerationRequest{
		NumImages: Ptr(2),
		Prompt:    "A serene beach at sunset.",
	}

	resp, err := client.Images.CreateImageGeneration(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateImageGeneration failed: %v", err)
	}

	// Validate response
	if *resp.SDGenerationJob.GenerationID != "gen-123" {
		t.Errorf("Expected Deleted Generation ID 'gen-123', got '%v'", *resp.SDGenerationJob.GenerationID)
	}
	if *resp.SDGenerationJob.APICreditCost != 2 {
		t.Errorf("Expected Deleted APICreditCost 2, got '%v'", *resp.SDGenerationJob.APICreditCost)
	}
}

// TestGetImageGeneration tests the GetImageGeneration method.
func TestGetImageGeneration(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedID := "gen-123"
		expectedPath := "/generations/" + expectedID
		if r.URL.Path != expectedPath || r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			resp := APIErrorResponse{
				Code:    "not-found",
				Message: "Generation not found.",
			}
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		resp := GetGenerationResponse{
			GenerationsByPK: struct {
				ID        *string `json:"id"`
				Status    *string `json:"status"`
				CreatedAt *Time   `json:"createdAt"`
				// Add other fields as necessary
			}{
				ID:        Ptr("gen-123"),
				Status:    Ptr("COMPLETE"),
				CreatedAt: &Time{time.Now().Add(-2 * time.Hour)},
			},
		}
		json.NewEncoder(w).Encode(resp)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	// Initialize client with mock server URL
	client := &Client{
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
		APIKey:     "test-api-key",
	}
	client.Images = client.NewImagesService()

	// Execute GetImageGeneration - Success Case
	ctx := context.Background()
	resp, err := client.Images.GetImageGeneration(ctx, "gen-123")
	if err != nil {
		t.Fatalf("GetImageGeneration failed: %v", err)
	}

	// Validate response
	if resp.GenerationsByPK.ID == nil || *resp.GenerationsByPK.ID != "gen-123" {
		t.Errorf("Expected Generation ID 'gen-123', got '%v'", resp.GenerationsByPK.ID)
	}
	if resp.GenerationsByPK.Status == nil || *resp.GenerationsByPK.Status != "COMPLETE" {
		t.Errorf("Expected Status 'COMPLETE', got '%v'", resp.GenerationsByPK.Status)
	}
	if resp.GenerationsByPK.CreatedAt == nil {
		t.Errorf("Expected CreatedAt to be set, got nil")
	}

	// Execute GetImageGeneration - Not Found Case
	_, err = client.Images.GetImageGeneration(ctx, "gen-404")
	if err == nil {
		t.Fatal("Expected error when retrieving nonexistent generation, got nil")
	}

	expectedErrMsg := "API Error 404: Generation not found."
	if !strings.HasSuffix(err.Error(), expectedErrMsg) {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrMsg, err.Error())
	}
}

// TestDeleteGeneration tests the DeleteGeneration method.
func TestDeleteGeneration(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedID := "gen-123"
		expectedPath := "/generations/" + expectedID
		if r.URL.Path != expectedPath || r.Method != "DELETE" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Mock successful deletion response
		w.Header().Set("Content-Type", "application/json")
		response := DeleteGenerationResponse{
			DeleteGenerationsByPK: struct {
				ID *string `json:"id"`
			}{
				ID: Ptr("gen-123"),
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
	client.Images = client.NewImagesService()

	// Execute DeleteGeneration
	ctx := context.Background()
	resp, err := client.Images.DeleteGeneration(ctx, "gen-123")
	if err != nil {
		t.Fatalf("DeleteGeneration failed: %v", err)
	}

	// Assert response
	if resp.DeleteGenerationsByPK.ID == nil || *resp.DeleteGenerationsByPK.ID != "gen-123" {
		t.Errorf("Expected Deleted Generation ID 'gen-123', got '%v'", resp.DeleteGenerationsByPK.ID)
	}
}
