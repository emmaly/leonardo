package leonardo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestListElements tests the ListElements method.
func TestListElements(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/elements"
		if r.URL.Path != expectedPath || r.Method != "GET" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Mock successful response with sample Loras
		w.Header().Set("Content-Type", "application/json")
		response := ListElementsResponse{
			Loras: []Lora{
				{
					AKUUID:        Ptr("lora-001"),
					BaseModel:     Ptr("v1_5"),
					CreatorName:   Ptr("Artist1"),
					Description:   Ptr("A Lora for creating vibrant landscapes."),
					Name:          Ptr("Vibrant Landscapes"),
					URLImage:      Ptr("https://example.com/lora1.jpg"),
					WeightDefault: Ptr(5),
					WeightMax:     Ptr(10),
					WeightMin:     Ptr(1),
				},
				{
					AKUUID:        Ptr("lora-002"),
					BaseModel:     Ptr("v2.1"),
					CreatorName:   Ptr("Artist2"),
					Description:   Ptr("A Lora for creating detailed architectural designs."),
					Name:          Ptr("Architectural Designs"),
					URLImage:      Ptr("https://example.com/lora2.jpg"),
					WeightDefault: Ptr(3),
					WeightMax:     Ptr(8),
					WeightMin:     Ptr(2),
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
	client.Elements = client.NewElementsService()

	// Execute ListElements
	ctx := context.Background()
	resp, err := client.Elements.ListElements(ctx)
	if err != nil {
		t.Fatalf("ListElements failed: %v", err)
	}

	// Validate response
	if len(resp.Loras) != 2 {
		t.Fatalf("Expected 2 Loras, got %d", len(resp.Loras))
	}

	lora1 := resp.Loras[0]
	if *lora1.Name != "Vibrant Landscapes" {
		t.Errorf("Expected first Lora Name 'Vibrant Landscapes', got '%s'", *lora1.Name)
	}
	if *lora1.CreatorName != "Artist1" {
		t.Errorf("Expected first Lora CreatorName 'Artist1', got '%s'", *lora1.CreatorName)
	}
	if *lora1.WeightDefault != 5 {
		t.Errorf("Expected first Lora WeightDefault 5, got %d", *lora1.WeightDefault)
	}

	lora2 := resp.Loras[1]
	if *lora2.Name != "Architectural Designs" {
		t.Errorf("Expected second Lora Name 'Architectural Designs', got '%s'", *lora2.Name)
	}
	if *lora2.CreatorName != "Artist2" {
		t.Errorf("Expected second Lora CreatorName 'Artist2', got '%s'", *lora2.CreatorName)
	}
	if *lora2.WeightDefault != 3 {
		t.Errorf("Expected second Lora WeightDefault 3, got %d", *lora2.WeightDefault)
	}
}
