package leonardo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestCreateSVDMotionGeneration tests the CreateSVDMotionGeneration method.
func TestCreateSVDMotionGeneration(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/generations-motion-svd"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := CreateSVDMotionGenerationResponse{
			Details:      map[string]interface{}{"info": "Additional details here"},
			GenerationID: "motion-gen-001",
			Status:       "success",
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
	client.Models = client.NewModelsService()
	client.Variation = client.NewVariationService()
	client.InitImages = client.NewInitImagesService()
	client.Prompt = client.NewPromptService()
	client.Elements = client.NewElementsService()
	client.Datasets = client.NewDatasetsService()
	client.Images = client.NewImagesService()
	client.PricingCalculator = client.NewPricingCalculatorService()
	client.RealtimeCanvas = client.NewRealtimeCanvasService()
	client.Texture = client.NewTextureService()
	client.ThreeDModelAssets = client.NewThreeDModelAssetsService()
	client.User = client.NewUserService()
	client.Motion = client.NewMotionService()

	// Execute CreateSVDMotionGeneration
	ctx := context.Background()
	req := MotionRequest{}

	resp, err := client.Motion.CreateSVDMotionGeneration(ctx, req)
	if err != nil {
		t.Fatalf("CreateSVDMotionGeneration failed: %v", err)
	}

	// Validate response
	if resp.GenerationID != "motion-gen-001" {
		t.Errorf("Expected GenerationID 'motion-gen-001', got '%s'", resp.GenerationID)
	}
	if resp.Status != "success" {
		t.Errorf("Expected Status 'success', got '%s'", resp.Status)
	}
	if resp.Details["info"] != "Additional details here" {
		t.Errorf("Expected Details info 'Additional details here', got '%v'", resp.Details["info"])
	}
}

// TestCreateSVDMotionGenerationError tests the error response of CreateSVDMotionGeneration.
func TestCreateSVDMotionGenerationError(t *testing.T) {
	// Mock server setup for error response
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/generations-motion-svd"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Mock error response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := APIErrorResponse{
			Code:    "bad-request",
			Message: "Bad Request: Invalid parameters.",
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
	client.Motion = client.NewMotionService()

	// Execute CreateSVDMotionGeneration
	ctx := context.Background()
	req := MotionRequest{}

	_, err := client.Motion.CreateSVDMotionGeneration(ctx, req)
	if err == nil {
		t.Fatal("Expected error when creating SVD motion generation, got nil")
	}

	expectedErrMsg := "API Error 400: Bad Request: Invalid parameters."
	if !strings.HasSuffix(err.Error(), expectedErrMsg) {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrMsg, err.Error())
	}
}
