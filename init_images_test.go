package leonardo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestUploadInitImage tests the UploadInitImage method.
func TestUploadInitImage(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/init-image"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req UploadInitImageRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		if req.ImageFile == "" {
			t.Errorf("ImageFile is required, got empty string")
		}

		// Mock response
		w.Header().Set("Content-Type", "application/json")
		response := struct {
			UploadInitImageID *string `json:"uploadInitImageId"`
			Message           string  `json:"message"`
		}{
			UploadInitImageID: Ptr("init-001"),
			Message:           "Init image uploaded successfully.",
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
	client.InitImages = client.NewInitImagesService()

	// Execute UploadInitImage
	ctx := context.Background()
	uploadReq := UploadInitImageRequest{
		ImageFile: "base64encodeddata",
	}

	resp, err := client.InitImages.UploadInitImage(ctx, uploadReq)
	if err != nil {
		t.Fatalf("UploadInitImage failed: %v", err)
	}

	// Validate response
	if resp.UploadInitImageID == nil || *resp.UploadInitImageID != "init-001" {
		t.Errorf("Expected UploadInitImageID 'init-001', got '%v'", resp.UploadInitImageID)
	}
	if resp.Message != "Init image uploaded successfully." {
		t.Errorf("Expected message 'Init image uploaded successfully.', got '%s'", resp.Message)
	}
}

// TestGetSingleInitImage tests the GetSingleInitImage method.
func TestGetSingleInitImage(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedID := "init-001"
		expectedPath := "/init-image/" + expectedID
		if r.URL.Path != expectedPath || r.Method != "GET" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := GetSingleInitImageResponse{
			InitImagesByPk: struct {
				CreatedAt *Time   `json:"createdAt"`
				ID        *string `json:"id"`
				URL       *string `json:"url"`
			}{
				CreatedAt: &Time{time.Now().Add(-2 * time.Hour)},
				ID:        Ptr(expectedID),
				URL:       Ptr("https://s3.amazonaws.com/bucket/init-001.jpg"),
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
	client.InitImages = client.NewInitImagesService()

	// Execute GetSingleInitImage
	ctx := context.Background()
	resp, err := client.InitImages.GetSingleInitImage(ctx, "init-001")
	if err != nil {
		t.Fatalf("GetSingleInitImage failed: %v", err)
	}

	// Validate response
	if resp.InitImagesByPk.ID == nil || *resp.InitImagesByPk.ID != "init-001" {
		t.Errorf("Expected ID 'init-001', got '%v'", resp.InitImagesByPk.ID)
	}
	if resp.InitImagesByPk.URL == nil || *resp.InitImagesByPk.URL != "https://s3.amazonaws.com/bucket/init-001.jpg" {
		t.Errorf("Expected URL 'https://s3.amazonaws.com/bucket/init-001.jpg', got '%v'", resp.InitImagesByPk.URL)
	}
	if resp.InitImagesByPk.CreatedAt == nil {
		t.Errorf("Expected CreatedAt to be set, got nil")
	}
}

// TestDeleteInitImage tests the DeleteInitImage method.
func TestDeleteInitImage(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedID := "init-001"
		expectedPath := "/init-image/" + expectedID
		if r.URL.Path != expectedPath || r.Method != "DELETE" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Mock successful deletion response
		w.Header().Set("Content-Type", "application/json")
		response := DeleteInitImageResponse{
			DeleteInitImagesByPk: struct {
				ID *string `json:"id"`
			}{
				ID: Ptr(expectedID),
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
	client.InitImages = client.NewInitImagesService()

	// Execute DeleteInitImage
	ctx := context.Background()
	resp, err := client.InitImages.DeleteInitImage(ctx, "init-001")
	if err != nil {
		t.Fatalf("DeleteInitImage failed: %v", err)
	}

	// Validate response
	if resp.DeleteInitImagesByPk.ID == nil || *resp.DeleteInitImagesByPk.ID != "init-001" {
		t.Errorf("Expected Deleted InitImage ID 'init-001', got '%v'", resp.DeleteInitImagesByPk.ID)
	}
}

// TestUploadCanvasInitAndMaskImage tests the UploadCanvasInitAndMaskImage method.
func TestUploadCanvasInitAndMaskImage(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/canvas-init-image"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req UploadCanvasInitAndMaskImageRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		if req.InitExtension == "" || req.MaskExtension == "" {
			t.Errorf("InitExtension and MaskExtension are required")
		}

		// Mock response
		w.Header().Set("Content-Type", "application/json")
		response := UploadCanvasInitAndMaskImageResponse{
			UploadCanvasInitImage: &struct {
				InitFields  string `json:"initFields"`
				InitImageID string `json:"initImageId"`
				InitKey     string `json:"initKey"`
				InitURL     string `json:"initUrl"`
				MaskFields  string `json:"maskFields"`
				MaskImageID string `json:"maskImageId"`
				MaskKey     string `json:"maskKey"`
				MaskURL     string `json:"maskUrl"`
			}{
				InitFields:  "init-policy",
				InitImageID: "init-img-001",
				InitKey:     "uploads/init-img-001.jpg",
				InitURL:     "https://s3.amazonaws.com/bucket/uploads/init-img-001.jpg",
				MaskFields:  "mask-policy",
				MaskImageID: "mask-img-001",
				MaskKey:     "uploads/mask-img-001.jpg",
				MaskURL:     "https://s3.amazonaws.com/bucket/uploads/mask-img-001.jpg",
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
	client.InitImages = client.NewInitImagesService()

	// Execute UploadCanvasInitAndMaskImage
	ctx := context.Background()
	uploadReq := UploadCanvasInitAndMaskImageRequest{
		InitExtension: "png",
		MaskExtension: "png",
	}

	resp, err := client.InitImages.UploadCanvasInitAndMaskImage(ctx, uploadReq)
	if err != nil {
		t.Fatalf("UploadCanvasInitAndMaskImage failed: %v", err)
	}

	// Validate response
	if resp.UploadCanvasInitImage == nil {
		t.Fatal("Expected UploadCanvasInitImage in response, got nil")
	}
	if resp.UploadCanvasInitImage.InitImageID != "init-img-001" {
		t.Errorf("Expected InitImageID 'init-img-001', got '%s'", resp.UploadCanvasInitImage.InitImageID)
	}
	if resp.UploadCanvasInitImage.MaskImageID != "mask-img-001" {
		t.Errorf("Expected MaskImageID 'mask-img-001', got '%s'", resp.UploadCanvasInitImage.MaskImageID)
	}
}
