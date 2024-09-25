package leonardo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestCreateDataset tests the CreateDataset method.
func TestCreateDataset(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify endpoint and method
		if r.URL.Path != "/datasets" || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req CreateDatasetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		if req.Name != "Test Dataset" {
			t.Errorf("Expected dataset name 'Test Dataset', got '%s'", req.Name)
		}
		if req.Description == nil || *req.Description != "A test dataset description." {
			t.Errorf("Expected description 'A test dataset description.', got '%v'", req.Description)
		}

		// Mock response
		w.Header().Set("Content-Type", "application/json")
		response := CreateDatasetResponse{
			InsertDatasetsOne: struct {
				ID *string `json:"id"`
			}{
				ID: Ptr("dataset-123"),
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
	client.Datasets = client.NewDatasetsService()

	// Execute CreateDataset
	ctx := context.Background()
	createReq := CreateDatasetRequest{
		Name:        "Test Dataset",
		Description: Ptr("A test dataset description."),
	}

	resp, err := client.Datasets.CreateDataset(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateDataset failed: %v", err)
	}

	// Validate response
	if resp.InsertDatasetsOne.ID == nil || *resp.InsertDatasetsOne.ID != "dataset-123" {
		t.Errorf("Expected Dataset ID 'dataset-123', got '%v'", resp.InsertDatasetsOne.ID)
	}
}

// TestGetDataset tests the GetDataset method.
func TestGetDataset(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/datasets/dataset-123"
		if r.URL.Path != expectedPath || r.Method != "GET" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Mock response
		w.Header().Set("Content-Type", "application/json")
		response := GetDatasetResponse{
			DatasetsByPk: struct {
				CreatedAt     *Time `json:"createdAt"`
				DatasetImages []struct {
					CreatedAt *Time   `json:"createdAt"`
					ID        *string `json:"id"`
					URL       *string `json:"url"`
				} `json:"dataset_images"`
				Description *string `json:"description"`
				ID          *string `json:"id"`
				Name        *string `json:"name"`
				UpdatedAt   *Time   `json:"updatedAt"`
			}{
				CreatedAt: &Time{time.Now().Add(-24 * time.Hour)},
				DatasetImages: []struct {
					CreatedAt *Time   `json:"createdAt"`
					ID        *string `json:"id"`
					URL       *string `json:"url"`
				}{
					{
						CreatedAt: &Time{time.Now().Add(-23 * time.Hour)},
						ID:        Ptr("img-001"),
						URL:       Ptr("https://s3.amazonaws.com/bucket/img-001.jpg"),
					},
					{
						CreatedAt: &Time{time.Now().Add(-22 * time.Hour)},
						ID:        Ptr("img-002"),
						URL:       Ptr("https://s3.amazonaws.com/bucket/img-002.jpg"),
					},
				},
				Description: Ptr("A test dataset description."),
				ID:          Ptr("dataset-123"),
				Name:        Ptr("Test Dataset"),
				UpdatedAt:   &Time{time.Now()},
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
	client.Datasets = client.NewDatasetsService()

	// Execute GetDataset
	ctx := context.Background()
	resp, err := client.Datasets.GetDataset(ctx, "dataset-123")
	if err != nil {
		t.Fatalf("GetDataset failed: %v", err)
	}

	// Validate response
	if resp.DatasetsByPk.ID == nil || *resp.DatasetsByPk.ID != "dataset-123" {
		t.Errorf("Expected Dataset ID 'dataset-123', got '%v'", resp.DatasetsByPk.ID)
	}
	if resp.DatasetsByPk.Name == nil || *resp.DatasetsByPk.Name != "Test Dataset" {
		t.Errorf("Expected Dataset Name 'Test Dataset', got '%v'", resp.DatasetsByPk.Name)
	}
	if len(resp.DatasetsByPk.DatasetImages) != 2 {
		t.Errorf("Expected 2 Dataset Images, got %d", len(resp.DatasetsByPk.DatasetImages))
	}
}

// TestDeleteDataset tests the DeleteDataset method.
func TestDeleteDataset(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/datasets/dataset-123"
		if r.URL.Path != expectedPath || r.Method != "DELETE" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Mock response
		w.Header().Set("Content-Type", "application/json")
		response := DeleteDatasetResponse{
			DeleteDatasetsByPK: struct {
				ID *string `json:"id"`
			}{
				ID: Ptr("dataset-123"),
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
	client.Datasets = client.NewDatasetsService()

	// Execute DeleteDataset
	ctx := context.Background()
	resp, err := client.Datasets.DeleteDataset(ctx, "dataset-123")
	if err != nil {
		t.Fatalf("DeleteDataset failed: %v", err)
	}

	// Validate response
	if resp.DeleteDatasetsByPK.ID == nil || *resp.DeleteDatasetsByPK.ID != "dataset-123" {
		t.Errorf("Expected Deleted Dataset ID 'dataset-123', got '%v'", resp.DeleteDatasetsByPK.ID)
	}
}

// TestUploadDatasetImage tests the UploadDatasetImage method.
func TestUploadDatasetImage(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/datasets/dataset-123/upload"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req UploadDatasetImageRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		if req.Extension != "jpg" && req.Extension != "png" && req.Extension != "jpeg" && req.Extension != "webp" {
			t.Errorf("Invalid extension: %s", req.Extension)
		}

		// Mock response
		w.Header().Set("Content-Type", "application/json")
		response := UploadDatasetImageResponse{
			UploadDatasetImage: &struct {
				Fields map[string]string `json:"fields"`
				ID     *string           `json:"id"`
				Key    *string           `json:"key"`
				URL    *string           `json:"url"`
			}{
				Fields: map[string]string{
					"key":       "uploads/dataset-123/image-001.jpg",
					"policy":    "policy-string",
					"signature": "signature-string",
					// Add other necessary fields as per S3 presigned URL requirements
				},
				ID:  Ptr("image-001"),
				Key: Ptr("uploads/dataset-123/image-001.jpg"),
				URL: Ptr("https://s3.amazonaws.com/bucket/uploads/dataset-123/image-001.jpg"),
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
	client.Datasets = client.NewDatasetsService()

	// Execute UploadDatasetImage
	ctx := context.Background()
	uploadReq := UploadDatasetImageRequest{
		Extension: "jpg",
	}

	resp, err := client.Datasets.UploadDatasetImage(ctx, "dataset-123", uploadReq)
	if err != nil {
		t.Fatalf("UploadDatasetImage failed: %v", err)
	}

	// Validate response
	if resp.UploadDatasetImage == nil {
		t.Fatal("Expected UploadDatasetImage in response, got nil")
	}
	if resp.UploadDatasetImage.ID == nil || *resp.UploadDatasetImage.ID != "image-001" {
		t.Errorf("Expected Image ID 'image-001', got '%v'", resp.UploadDatasetImage.ID)
	}
	if resp.UploadDatasetImage.URL == nil || *resp.UploadDatasetImage.URL != "https://s3.amazonaws.com/bucket/uploads/dataset-123/image-001.jpg" {
		t.Errorf("Expected S3 URL 'https://s3.amazonaws.com/bucket/uploads/dataset-123/image-001.jpg', got '%v'", resp.UploadDatasetImage.URL)
	}
	if resp.UploadDatasetImage.Key == nil || *resp.UploadDatasetImage.Key != "uploads/dataset-123/image-001.jpg" {
		t.Errorf("Expected S3 Key 'uploads/dataset-123/image-001.jpg', got '%v'", resp.UploadDatasetImage.Key)
	}
	if len(resp.UploadDatasetImage.Fields) != 3 { // Adjust based on mock response
		t.Errorf("Expected 3 S3 Fields, got %d", len(resp.UploadDatasetImage.Fields))
	}
}

// TestUploadGeneratedImageToDataset tests the UploadGeneratedImageToDataset method.
func TestUploadGeneratedImageToDataset(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/datasets/dataset-123/upload/gen"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req UploadGeneratedImageRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		if req.GeneratedImageID == "" {
			t.Errorf("GeneratedImageID is required, got empty string")
		}

		// Mock response
		w.Header().Set("Content-Type", "application/json")
		response := UploadGeneratedImageResponse{
			UploadDatasetImageFromGen: &struct {
				ID *string `json:"id"`
			}{
				ID: Ptr("uploaded-gen-456"),
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
	client.Datasets = client.NewDatasetsService()

	// Execute UploadGeneratedImageToDataset
	ctx := context.Background()
	uploadGenReq := UploadGeneratedImageRequest{
		GeneratedImageID: "gen-456",
	}

	resp, err := client.Datasets.UploadGeneratedImageToDataset(ctx, "dataset-123", uploadGenReq)
	if err != nil {
		t.Fatalf("UploadGeneratedImageToDataset failed: %v", err)
	}

	// Validate response
	if resp.UploadDatasetImageFromGen == nil || *resp.UploadDatasetImageFromGen.ID != "uploaded-gen-456" {
		t.Errorf("Expected Uploaded Generated Image ID 'uploaded-gen-456', got '%v'", resp.UploadDatasetImageFromGen.ID)
	}
}
