package leonardo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCreateLCMGeneration tests the CreateLCMGeneration method.
func TestCreateLCMGeneration(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/generations-lcm"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req CreateLCMGenerationRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		// Since fields are optional, no need to validate here

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := CreateLCMGenerationResponse{
			LCMGenerationJob: &LCMGenerationJob{
				APICreditCost:    Ptr(10),
				ImageDataURL:     []string{"https://example.com/init1.jpg"},
				RequestTimestamp: Ptr("2023-10-10T10:00:00Z"),
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
	client.RealtimeCanvas = client.NewRealtimeCanvasService()

	// Execute CreateLCMGeneration
	ctx := context.Background()
	req := CreateLCMGenerationRequest{
		Prompt: "A beautiful sunrise over the mountains.",
		// Add other fields as needed
	}

	resp, err := client.RealtimeCanvas.CreateLCMGeneration(ctx, req)
	if err != nil {
		t.Fatalf("CreateLCMGeneration failed: %v", err)
	}

	// Validate response
	if resp.LCMGenerationJob == nil {
		t.Fatal("Expected LCMGenerationJob in response, got nil")
	}
	if *resp.LCMGenerationJob.APICreditCost != 10 {
		t.Errorf("Expected APICreditCost 10, got %d", *resp.LCMGenerationJob.APICreditCost)
	}
	if len(resp.LCMGenerationJob.ImageDataURL) != 1 || resp.LCMGenerationJob.ImageDataURL[0] != "https://example.com/init1.jpg" {
		t.Errorf("Unexpected ImageDataURL: %v", resp.LCMGenerationJob.ImageDataURL)
	}
	if *resp.LCMGenerationJob.RequestTimestamp != "2023-10-10T10:00:00Z" {
		t.Errorf("Unexpected RequestTimestamp: %s", *resp.LCMGenerationJob.RequestTimestamp)
	}
}

// TestPerformInstantRefine tests the PerformInstantRefine method.
func TestPerformInstantRefine(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/lcm-instant-refine"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req PerformInstantRefineRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		if req.ImageDataURL == "" || req.Prompt == "" {
			w.WriteHeader(http.StatusBadRequest)
			resp := APIErrorResponse{
				Code:    "bad-request",
				Message: "imageDataUrl and prompt are required.",
			}
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := PerformInstantRefineResponse{
			LCMGenerationJob: &LCMGenerationJob{
				APICreditCost:    Ptr(15),
				ImageDataURL:     []string{"https://example.com/refined1.jpg"},
				RequestTimestamp: Ptr("2023-10-10T10:30:00Z"),
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
	client.RealtimeCanvas = client.NewRealtimeCanvasService()

	// Execute PerformInstantRefine
	ctx := context.Background()
	req := PerformInstantRefineRequest{
		ImageDataURL:   "https://example.com/init1.jpg",
		Prompt:         "Refine the image with more vibrant colors.",
		RefineStrength: Ptr(0.7),
		Guidance:       Ptr(12.0),
		Height:         Ptr(512),
		Width:          Ptr(512),
		RefineCreative: Ptr(true),
	}

	resp, err := client.RealtimeCanvas.PerformInstantRefine(ctx, req)
	if err != nil {
		t.Fatalf("PerformInstantRefine failed: %v", err)
	}

	// Validate response
	if resp.LCMGenerationJob == nil {
		t.Fatal("Expected LCMGenerationJob in response, got nil")
	}
	if *resp.LCMGenerationJob.APICreditCost != 15 {
		t.Errorf("Expected APICreditCost 15, got %d", *resp.LCMGenerationJob.APICreditCost)
	}
	if len(resp.LCMGenerationJob.ImageDataURL) != 1 || resp.LCMGenerationJob.ImageDataURL[0] != "https://example.com/refined1.jpg" {
		t.Errorf("Unexpected ImageDataURL: %v", resp.LCMGenerationJob.ImageDataURL)
	}
	if *resp.LCMGenerationJob.RequestTimestamp != "2023-10-10T10:30:00Z" {
		t.Errorf("Unexpected RequestTimestamp: %s", *resp.LCMGenerationJob.RequestTimestamp)
	}
}

// TestPerformInpainting tests the PerformInpainting method.
func TestPerformInpainting(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/lcm-inpainting"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req PerformInpaintingRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		if req.ImageDataURL == "" || req.MaskDataURL == "" || req.Prompt == "" {
			w.WriteHeader(http.StatusBadRequest)
			resp := APIErrorResponse{
				Code:    "bad-request",
				Message: "imageDataUrl, maskDataURL, and prompt are required.",
			}
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := PerformInpaintingResponse{
			LCMGenerationJob: &LCMGenerationJob{
				APICreditCost:    Ptr(20),
				ImageDataURL:     []string{"https://example.com/inpainted1.jpg"},
				RequestTimestamp: Ptr("2023-10-10T11:00:00Z"),
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
	client.RealtimeCanvas = client.NewRealtimeCanvasService()

	// Execute PerformInpainting
	ctx := context.Background()
	req := PerformInpaintingRequest{
		ImageDataURL:     "https://example.com/init1.jpg",
		MaskDataURL:      "https://example.com/mask1.png",
		Prompt:           "Remove the object and fill the area seamlessly.",
		Strength:         Ptr(0.8),
		Guidance:         Ptr(14.0),
		Height:           Ptr(512),
		Width:            Ptr(512),
		Style:            Ptr("DIGITAL_ART"),
		RequestTimestamp: Ptr("2023-10-10T11:00:00Z"),
	}

	resp, err := client.RealtimeCanvas.PerformInpainting(ctx, req)
	if err != nil {
		t.Fatalf("PerformInpainting failed: %v", err)
	}

	// Validate response
	if resp.LCMGenerationJob == nil {
		t.Fatal("Expected LCMGenerationJob in response, got nil")
	}
	if *resp.LCMGenerationJob.APICreditCost != 20 {
		t.Errorf("Expected APICreditCost 20, got %d", *resp.LCMGenerationJob.APICreditCost)
	}
	if len(resp.LCMGenerationJob.ImageDataURL) != 1 || resp.LCMGenerationJob.ImageDataURL[0] != "https://example.com/inpainted1.jpg" {
		t.Errorf("Unexpected ImageDataURL: %v", resp.LCMGenerationJob.ImageDataURL)
	}
	if *resp.LCMGenerationJob.RequestTimestamp != "2023-10-10T11:00:00Z" {
		t.Errorf("Unexpected RequestTimestamp: %s", *resp.LCMGenerationJob.RequestTimestamp)
	}
}

// TestPerformAlchemyUpscale tests the PerformAlchemyUpscale method.
func TestPerformAlchemyUpscale(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/lcm-upscale"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req PerformAlchemyUpscaleRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		if req.ImageDataURL == "" || req.Prompt == "" {
			w.WriteHeader(http.StatusBadRequest)
			resp := APIErrorResponse{
				Code:    "bad-request",
				Message: "imageDataUrl and prompt are required.",
			}
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := PerformAlchemyUpscaleResponse{
			LCMGenerationJob: &struct {
				APICreditCost    *int     `json:"apiCreditCost"`
				GeneratedImageID *string  `json:"generatedImageId"`
				GenerationID     []string `json:"generationId"`
				ImageDataURL     []string `json:"imageDataUrl"`
				RequestTimestamp *string  `json:"requestTimestamp"`
				VariationID      []string `json:"variationId"`
			}{
				APICreditCost:    Ptr(25),
				GeneratedImageID: Ptr("gen-upscale-001"),
				GenerationID:     []string{"gen-123"},
				ImageDataURL:     []string{"https://example.com/upscaled1.jpg"},
				RequestTimestamp: Ptr("2023-10-10T11:30:00Z"),
				VariationID:      []string{"var-789"},
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
	client.RealtimeCanvas = client.NewRealtimeCanvasService()

	// Execute PerformAlchemyUpscale
	ctx := context.Background()
	req := PerformAlchemyUpscaleRequest{
		ImageDataURL:     "https://example.com/inpainted1.jpg",
		Prompt:           "Enhance the resolution without losing details.",
		Strength:         Ptr(0.9),
		Guidance:         Ptr(16.0),
		Height:           Ptr(1024),
		Width:            Ptr(1024),
		Style:            Ptr("VIBRANT"),
		RequestTimestamp: Ptr("2023-10-10T11:30:00Z"),
	}

	resp, err := client.RealtimeCanvas.PerformAlchemyUpscale(ctx, req)
	if err != nil {
		t.Fatalf("PerformAlchemyUpscale failed: %v", err)
	}

	// Validate response
	if resp.LCMGenerationJob == nil {
		t.Fatal("Expected LCMGenerationJob in response, got nil")
	}
	if *resp.LCMGenerationJob.APICreditCost != 25 {
		t.Errorf("Expected APICreditCost 25, got %d", *resp.LCMGenerationJob.APICreditCost)
	}
	if *resp.LCMGenerationJob.GeneratedImageID != "gen-upscale-001" {
		t.Errorf("Expected GeneratedImageID 'gen-upscale-001', got '%v'", *resp.LCMGenerationJob.GeneratedImageID)
	}
	if len(resp.LCMGenerationJob.GenerationID) != 1 || resp.LCMGenerationJob.GenerationID[0] != "gen-123" {
		t.Errorf("Unexpected GenerationID: %v", resp.LCMGenerationJob.GenerationID)
	}
	if len(resp.LCMGenerationJob.ImageDataURL) != 1 || resp.LCMGenerationJob.ImageDataURL[0] != "https://example.com/upscaled1.jpg" {
		t.Errorf("Unexpected ImageDataURL: %v", resp.LCMGenerationJob.ImageDataURL)
	}
	if *resp.LCMGenerationJob.RequestTimestamp != "2023-10-10T11:30:00Z" {
		t.Errorf("Unexpected RequestTimestamp: %s", *resp.LCMGenerationJob.RequestTimestamp)
	}
	if len(resp.LCMGenerationJob.VariationID) != 1 || resp.LCMGenerationJob.VariationID[0] != "var-789" {
		t.Errorf("Unexpected VariationID: %v", resp.LCMGenerationJob.VariationID)
	}
}
