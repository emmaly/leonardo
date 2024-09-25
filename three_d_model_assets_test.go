package leonardo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestUpload3DModel tests the Upload3DModel method.
func TestUpload3DModel(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/models-3d/upload"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req Upload3DModelRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}

		// Validate request
		if req.Name == nil || *req.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			resp := APIErrorResponse{
				Code:    "bad-request",
				Message: "Name is required.",
			}
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := Upload3DModelResponse{
			UploadModelAsset: &struct {
				ModelFields *string `json:"modelFields"`
				ModelID     *string `json:"modelId"`
				ModelKey    *string `json:"modelKey"`
				ModelURL    *string `json:"modelUrl"`
			}{
				ModelFields: Ptr("model-policy-string"),
				ModelID:     Ptr("model3d-001"),
				ModelKey:    Ptr("uploads/models-3d/model3d-001.glb"),
				ModelURL:    Ptr("https://s3.amazonaws.com/bucket/uploads/models-3d/model3d-001.glb"),
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
	client.ThreeDModelAssets = client.NewThreeDModelAssetsService()

	// Execute Upload3DModel
	ctx := context.Background()
	req := Upload3DModelRequest{
		Name: func() *string { s := "Test 3D Model"; return &s }(),
		// ModelExtension is optional
	}

	resp, err := client.ThreeDModelAssets.Upload3DModel(ctx, req)
	if err != nil {
		t.Fatalf("Upload3DModel failed: %v", err)
	}

	// Validate response
	if resp.UploadModelAsset == nil {
		t.Fatal("Expected UploadModelAsset in response, got nil")
	}
	if *resp.UploadModelAsset.ModelID != "model3d-001" {
		t.Errorf("Expected ModelID 'model3d-001', got '%s'", *resp.UploadModelAsset.ModelID)
	}
	if *resp.UploadModelAsset.ModelURL != "https://s3.amazonaws.com/bucket/uploads/models-3d/model3d-001.glb" {
		t.Errorf("Unexpected ModelURL: %s", *resp.UploadModelAsset.ModelURL)
	}
}

// TestGet3DModelsByUser tests the Get3DModelsByUser method.
func TestGet3DModelsByUser(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract userId from path
		expectedUserID := "user-123"
		expectedPath := "/models-3d/user/" + expectedUserID
		if r.URL.Path != expectedPath || r.Method != "GET" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Extract query parameters for pagination
		limit := r.URL.Query().Get("limit")
		offset := r.URL.Query().Get("offset")
		if offset == "" {
			offset = "0"
		}
		if limit != "10" || offset != "0" {
			t.Errorf("Unexpected query params: limit=%s, offset=%s", limit, offset)
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := Get3DModelsByUserResponse{
			ModelAssets: []struct {
				CreatedAt *Time   `json:"createdAt"`
				ID        *string `json:"id"`
				MeshURL   *string `json:"meshUrl"`
				Name      *string `json:"name"`
				UpdatedAt *Time   `json:"updatedAt"`
				UserID    *string `json:"userId"`
			}{
				{
					CreatedAt: &Time{time.Now().Add(-48 * time.Hour)},
					ID:        Ptr("model3d-001"),
					MeshURL:   Ptr("https://s3.amazonaws.com/bucket/uploads/models-3d/model3d-001.glb"),
					Name:      Ptr("Sample 3D Model"),
					UpdatedAt: &Time{time.Now().Add(-24 * time.Hour)},
					UserID:    Ptr("user-123"),
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
	client.ThreeDModelAssets = client.NewThreeDModelAssetsService()

	// Execute Get3DModelsByUser
	ctx := context.Background()
	userID := "user-123"
	limit := 10
	offset := 0

	resp, err := client.ThreeDModelAssets.Get3DModelsByUser(ctx, userID, limit, offset)
	if err != nil {
		t.Fatalf("Get3DModelsByUser failed: %v", err)
	}

	// Validate response
	if len(resp.ModelAssets) != 1 {
		t.Fatalf("Expected 1 model asset, got %d", len(resp.ModelAssets))
	}
	model := resp.ModelAssets[0]
	if *model.ID != "model3d-001" {
		t.Errorf("Expected Model ID 'model3d-001', got '%s'", *model.ID)
	}
	if *model.Name != "Sample 3D Model" {
		t.Errorf("Expected Model Name 'Sample 3D Model', got '%s'", *model.Name)
	}
	if *model.MeshURL != "https://s3.amazonaws.com/bucket/uploads/models-3d/model3d-001.glb" {
		t.Errorf("Unexpected MeshURL: %s", *model.MeshURL)
	}
	if *model.UserID != "user-123" {
		t.Errorf("Expected UserID 'user-123', got '%s'", *model.UserID)
	}
}

// TestGet3DModelByID tests the Get3DModelByID method.
func TestGet3DModelByID(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedID := "model3d-001"
		expectedPath := "/models-3d/" + expectedID
		if r.URL.Path != expectedPath || r.Method != "GET" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := Get3DModelByIDResponse{
			ModelAssetsByPK: struct {
				CreatedAt   *Time   `json:"createdAt"`
				Description *string `json:"description"`
				ID          *string `json:"id"`
				MeshURL     *string `json:"meshUrl"`
				Name        *string `json:"name"`
				UpdatedAt   *Time   `json:"updatedAt"`
				UserID      *string `json:"userId"`
			}{
				CreatedAt:   &Time{time.Now().Add(-48 * time.Hour)},
				Description: Ptr("A detailed 3D model of a futuristic vehicle."),
				ID:          Ptr("model3d-001"),
				MeshURL:     Ptr("https://s3.amazonaws.com/bucket/uploads/models-3d/model3d-001.glb"),
				Name:        Ptr("Futuristic Vehicle"),
				UpdatedAt:   &Time{time.Now().Add(-24 * time.Hour)},
				UserID:      Ptr("user-123"),
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
	client.ThreeDModelAssets = client.NewThreeDModelAssetsService()

	// Execute Get3DModelByID
	ctx := context.Background()
	id := "model3d-001"

	resp, err := client.ThreeDModelAssets.Get3DModelByID(ctx, id)
	if err != nil {
		t.Fatalf("Get3DModelByID failed: %v", err)
	}

	// Validate response
	if *resp.ModelAssetsByPK.ID != "model3d-001" {
		t.Errorf("Expected Model ID 'model3d-001', got '%s'", *resp.ModelAssetsByPK.ID)
	}
	if *resp.ModelAssetsByPK.Name != "Futuristic Vehicle" {
		t.Errorf("Expected Model Name 'Futuristic Vehicle', got '%s'", *resp.ModelAssetsByPK.Name)
	}
	if *resp.ModelAssetsByPK.Description != "A detailed 3D model of a futuristic vehicle." {
		t.Errorf("Unexpected Description: %s", *resp.ModelAssetsByPK.Description)
	}
	if *resp.ModelAssetsByPK.MeshURL != "https://s3.amazonaws.com/bucket/uploads/models-3d/model3d-001.glb" {
		t.Errorf("Unexpected MeshURL: %s", *resp.ModelAssetsByPK.MeshURL)
	}
	if *resp.ModelAssetsByPK.UserID != "user-123" {
		t.Errorf("Expected UserID 'user-123', got '%s'", *resp.ModelAssetsByPK.UserID)
	}
}

// TestDelete3DModel tests the Delete3DModel method.
func TestDelete3DModel(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedID := "model3d-001"
		expectedPath := "/models-3d/" + expectedID
		if r.URL.Path != expectedPath || r.Method != "DELETE" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Mock successful deletion response
		w.Header().Set("Content-Type", "application/json")
		response := Delete3DModelResponse{
			DeleteModelAssetsByPK: struct {
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
	client.ThreeDModelAssets = client.NewThreeDModelAssetsService()

	// Execute Delete3DModel
	ctx := context.Background()
	id := "model3d-001"

	resp, err := client.ThreeDModelAssets.Delete3DModel(ctx, id)
	if err != nil {
		t.Fatalf("Delete3DModel failed: %v", err)
	}

	// Validate response
	if resp.DeleteModelAssetsByPK.ID == nil || *resp.DeleteModelAssetsByPK.ID != "model3d-001" {
		t.Errorf("Expected Deleted Model ID 'model3d-001', got '%v'", resp.DeleteModelAssetsByPK.ID)
	}
}
