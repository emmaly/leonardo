package leonardo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTrainCustomModel(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/models" || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}
		var req TrainCustomModelRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		response := TrainCustomModelResponse{
			SDTrainingJob: struct {
				APICreditCost *int    `json:"apiCreditCost"`
				CustomModelID *string `json:"customModelId"`
			}{
				APICreditCost: Ptr(100),
				CustomModelID: Ptr("model-123"),
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
	client.Models = client.NewModelsService()

	// Prepare request
	req := TrainCustomModelRequest{
		DatasetID:      "dataset-456",
		InstancePrompt: "A description of instances.",
		Name:           "Test Model",
	}

	// Execute method
	resp, err := client.Models.TrainCustomModel(context.Background(), req)
	if err != nil {
		t.Fatalf("TrainCustomModel failed: %v", err)
	}

	// Validate response
	if resp.SDTrainingJob.CustomModelID == nil || *resp.SDTrainingJob.CustomModelID != "model-123" {
		t.Errorf("Expected CustomModelID 'model-123', got '%v'", resp.SDTrainingJob.CustomModelID)
	}

	if resp.SDTrainingJob.APICreditCost == nil || *resp.SDTrainingJob.APICreditCost != 100 {
		t.Errorf("Expected APICreditCost 100, got '%v'", resp.SDTrainingJob.APICreditCost)
	}
}

func TestGetCustomModel(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedID := "model-123"
		expectedPath := "/models/" + expectedID
		if r.URL.Path != expectedPath || r.Method != "GET" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Mock response
		w.Header().Set("Content-Type", "application/json")
		response := GetCustomModelResponse{
			CustomModelsByPK: struct {
				CreatedAt      *Time   `json:"createdAt"`
				Description    *string `json:"description"`
				ID             *string `json:"id"`
				InstancePrompt *string `json:"instancePrompt"`
				ModelHeight    *int    `json:"modelHeight"`
				ModelWidth     *int    `json:"modelWidth"`
				Name           *string `json:"name"`
				Public         *bool   `json:"public"`
				SDVersion      *string `json:"sdVersion"`
				Status         *string `json:"status"`
				Type           *string `json:"type"`
				UpdatedAt      *Time   `json:"updatedAt"`
			}{
				CreatedAt:      &Time{time.Now().Add(-48 * time.Hour)},
				Description:    Ptr("A custom model for landscape generation."),
				ID:             Ptr("model-123"),
				InstancePrompt: Ptr("Detailed description for model instances."),
				ModelHeight:    Ptr(512),
				ModelWidth:     Ptr(512),
				Name:           Ptr("Landscape Model"),
				Public:         Ptr(true),
				SDVersion:      Ptr("v1.5"),
				Status:         Ptr("TRAINING"),
				Type:           Ptr("GENERAL"),
				UpdatedAt:      &Time{time.Now().Add(-24 * time.Hour)},
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
	client.Models = client.NewModelsService()

	// Execute GetCustomModel
	ctx := context.Background()
	resp, err := client.Models.GetCustomModel(ctx, "model-123")
	if err != nil {
		t.Fatalf("GetCustomModel failed: %v", err)
	}

	// Validate response
	if resp.CustomModelsByPK.ID == nil || *resp.CustomModelsByPK.ID != "model-123" {
		t.Errorf("Expected CustomModel ID 'model-123', got '%v'", resp.CustomModelsByPK.ID)
	}
	if resp.CustomModelsByPK.Name == nil || *resp.CustomModelsByPK.Name != "Landscape Model" {
		t.Errorf("Expected CustomModel Name 'Landscape Model', got '%v'", resp.CustomModelsByPK.Name)
	}
}

// TestDeleteCustomModel tests the DeleteCustomModel method.
func TestDeleteCustomModel(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedID := "model-123"
		expectedPath := "/models/" + expectedID
		if r.URL.Path != expectedPath || r.Method != "DELETE" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Mock response
		w.Header().Set("Content-Type", "application/json")
		response := DeleteCustomModelResponse{
			DeleteCustomModelsByPK: struct {
				ID *string `json:"id"`
			}{
				ID: Ptr("model-123"),
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
	client.Models = client.NewModelsService()

	// Execute DeleteCustomModel
	ctx := context.Background()
	resp, err := client.Models.DeleteCustomModel(ctx, "model-123")
	if err != nil {
		t.Fatalf("DeleteCustomModel failed: %v", err)
	}

	// Validate response
	if resp.DeleteCustomModelsByPK.ID == nil || *resp.DeleteCustomModelsByPK.ID != "model-123" {
		t.Errorf("Expected Deleted CustomModel ID 'model-123', got '%v'", resp.DeleteCustomModelsByPK.ID)
	}
}
