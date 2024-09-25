package leonardo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestCreateUnzoomVariation tests the CreateUnzoomVariation method.
func TestCreateUnzoomVariation(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/variations/unzoom"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req VariationRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		if req.ID == "" {
			w.WriteHeader(http.StatusBadRequest)
			resp := APIErrorResponse{
				Code:    "bad-request",
				Message: "ID is required.",
			}
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := CreateUnzoomVariationResponse{
			SdUnzoomJob: VariationJob{
				ID:            Ptr("unzoom-job-001"),
				APICreditCost: Ptr(5),
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
	client.Variation = client.NewVariationService()

	// Execute CreateUnzoomVariation
	ctx := context.Background()
	req := VariationRequest{
		ID: "image-123",
	}

	resp, err := client.Variation.CreateUnzoomVariation(ctx, req)
	if err != nil {
		t.Fatalf("CreateUnzoomVariation failed: %v", err)
	}

	// Validate response
	if resp.SdUnzoomJob.ID == nil || *resp.SdUnzoomJob.ID != "unzoom-job-001" {
		t.Errorf("Expected Job ID 'unzoom-job-001', got '%v'", resp.SdUnzoomJob.ID)
	}
	if *resp.SdUnzoomJob.APICreditCost != 5 {
		t.Errorf("Expected APICreditCost 5, got %d", *resp.SdUnzoomJob.APICreditCost)
	}
}

// TestCreateUpscaleVariation tests the CreateUpscaleVariation method.
func TestCreateUpscaleVariation(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/variations/upscale"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req UpscaleVariationRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		if req.ID == "" {
			w.WriteHeader(http.StatusBadRequest)
			resp := APIErrorResponse{
				Code:    "bad-request",
				Message: "ID is required.",
			}
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := UpscaleVariationResponse{
			SdUpscaleJob: struct {
				ID            *string `json:"id"`
				APICreditCost *int    `json:"apiCreditCost"`
			}{
				ID:            Ptr("upscale-job-001"),
				APICreditCost: Ptr(8),
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
	client.Variation = client.NewVariationService()

	// Execute CreateUpscaleVariation
	ctx := context.Background()
	id := "image-456"

	resp, err := client.Variation.CreateUpscaleVariation(ctx, id)
	if err != nil {
		t.Fatalf("CreateUpscaleVariation failed: %v", err)
	}

	// Validate response
	if resp.SdUpscaleJob.ID == nil || *resp.SdUpscaleJob.ID != "upscale-job-001" {
		t.Errorf("Expected Job ID 'upscale-job-001', got '%v'", resp.SdUpscaleJob.ID)
	}
	if *resp.SdUpscaleJob.APICreditCost != 8 {
		t.Errorf("Expected APICreditCost 8, got %d", *resp.SdUpscaleJob.APICreditCost)
	}
}

// TestCreateNoBackgroundVariation tests the CreateNoBackgroundVariation method.
func TestCreateNoBackgroundVariation(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/variations/nobg"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req VariationRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		if req.ID == "" {
			w.WriteHeader(http.StatusBadRequest)
			resp := APIErrorResponse{
				Code:    "bad-request",
				Message: "ID is required.",
			}
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := CreateNoBackgroundVariationResponse{
			SdNobgJob: VariationJob{
				ID:            Ptr("nobg-job-001"),
				APICreditCost: Ptr(6),
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
	client.Variation = client.NewVariationService()

	// Execute CreateNoBackgroundVariation
	ctx := context.Background()
	req := VariationRequest{
		ID: "image-789",
	}

	resp, err := client.Variation.CreateNoBackgroundVariation(ctx, req)
	if err != nil {
		t.Fatalf("CreateNoBackgroundVariation failed: %v", err)
	}

	// Validate response
	if resp.SdNobgJob.ID == nil || *resp.SdNobgJob.ID != "nobg-job-001" {
		t.Errorf("Expected Job ID 'nobg-job-001', got '%v'", resp.SdNobgJob.ID)
	}
	if *resp.SdNobgJob.APICreditCost != 6 {
		t.Errorf("Expected APICreditCost 6, got %d", *resp.SdNobgJob.APICreditCost)
	}
}

// TestCreateUniversalUpscalerVariation tests the CreateUniversalUpscalerVariation method.
func TestCreateUniversalUpscalerVariation(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/variations/universal-upscaler"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req UniversalUpscalerRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		if req.ImageURL == "" || req.ScaleFactor < 1 {
			w.WriteHeader(http.StatusBadRequest)
			resp := APIErrorResponse{
				Code:    "bad-request",
				Message: "image_url and scale_factor are required with scale_factor >= 1.",
			}
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := UniversalUpscalerResponse{
			UpscaledImageURL: "https://example.com/upscaled-image.jpg",
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
	client.Variation = client.NewVariationService()

	// Execute CreateUniversalUpscalerVariation
	ctx := context.Background()
	req := UniversalUpscalerRequest{
		ImageURL:    "https://example.com/image-to-upscale.jpg",
		ScaleFactor: 2,
	}

	resp, err := client.Variation.CreateUniversalUpscalerVariation(ctx, req)
	if err != nil {
		t.Fatalf("CreateUniversalUpscalerVariation failed: %v", err)
	}

	// Validate response
	if resp.UpscaledImageURL != "https://example.com/upscaled-image.jpg" {
		t.Errorf("Expected UpscaledImageURL 'https://example.com/upscaled-image.jpg', got '%s'", resp.UpscaledImageURL)
	}
}

// TestGetVariation tests the GetVariation method.
func TestGetVariation(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedID := "var-001"
		expectedPath := "/variations/" + expectedID
		if r.URL.Path != expectedPath || r.Method != "GET" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := GetVariationResponse{
			GeneratedImageVariationGeneric: []struct {
				CreatedAt     *Time   `json:"createdAt"`
				ID            *string `json:"id"`
				Status        *string `json:"status"`
				TransformType *string `json:"transformType"`
				URL           *string `json:"url"`
			}{
				{
					CreatedAt:     &Time{time.Now().Add(-1 * time.Hour)},
					ID:            Ptr("var-001"),
					Status:        Ptr("COMPLETE"),
					TransformType: Ptr("UPSCALE"),
					URL:           Ptr("https://example.com/upscaled-var-001.jpg"),
				},
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
	client.Variation = client.NewVariationService()

	// Execute GetVariation
	ctx := context.Background()
	id := "var-001"

	resp, err := client.Variation.GetVariation(ctx, id)
	if err != nil {
		t.Fatalf("GetVariation failed: %v", err)
	}

	// Validate response
	if len(resp.GeneratedImageVariationGeneric) != 1 {
		t.Fatalf("Expected 1 variation, got %d", len(resp.GeneratedImageVariationGeneric))
	}
	variation := resp.GeneratedImageVariationGeneric[0]
	if *variation.ID != "var-001" {
		t.Errorf("Expected Variation ID 'var-001', got '%s'", *variation.ID)
	}
	if *variation.Status != "COMPLETE" {
		t.Errorf("Expected Status 'COMPLETE', got '%s'", *variation.Status)
	}
	if *variation.TransformType != "UPSCALE" {
		t.Errorf("Expected TransformType 'UPSCALE', got '%s'", *variation.TransformType)
	}
	if *variation.URL != "https://example.com/upscaled-var-001.jpg" {
		t.Errorf("Expected URL 'https://example.com/upscaled-var-001.jpg', got '%s'", *variation.URL)
	}
}
