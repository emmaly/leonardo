package leonardo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestCalculateAPICostSuccess tests the successful execution of CalculateAPICost method.
func TestCalculateAPICostSuccess(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/pricing-calculator"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req CalculateAPICostRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}

		// Validate request
		if req.Service == "" {
			w.WriteHeader(http.StatusBadRequest)
			resp := APIErrorResponse{
				Code:    "bad-request",
				Message: "Service is required.",
			}
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := CalculateAPICostResponse{
			CalculateProductionApiServiceCost: struct {
				Cost *int `json:"cost"`
			}{
				Cost: Ptr(20),
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
	client.PricingCalculator = client.NewPricingCalculatorService()

	// Execute CalculateAPICost
	ctx := context.Background()
	req := CalculateAPICostRequest{
		Service: "IMAGE_GENERATION",
		ServiceParams: map[string]interface{}{
			"param1": "value1",
			"param2": 123,
		},
	}

	resp, err := client.PricingCalculator.CalculateAPICost(ctx, req)
	if err != nil {
		t.Fatalf("CalculateAPICost failed: %v", err)
	}

	// Validate response
	if resp.CalculateProductionApiServiceCost.Cost == nil {
		t.Error("Expected Cost in response, got nil")
	}
	if *resp.CalculateProductionApiServiceCost.Cost != 20 {
		t.Errorf("Expected Cost 20, got %d", *resp.CalculateProductionApiServiceCost.Cost)
	}
}

// TestCalculateAPICostBadRequest tests the error handling of CalculateAPICost method when API responds with 400 Bad Request.
func TestCalculateAPICostBadRequest(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/pricing-calculator"
		if r.URL.Path != expectedPath || r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Decode request body
		var req CalculateAPICostRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}

		// Mock error response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := APIErrorResponse{
			Code:    "bad-request",
			Message: "Invalid service type provided.",
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
	client.PricingCalculator = client.NewPricingCalculatorService()

	// Execute CalculateAPICost with invalid service
	ctx := context.Background()
	req := CalculateAPICostRequest{
		Service: "", // Invalid service to trigger bad request
	}

	_, err := client.PricingCalculator.CalculateAPICost(ctx, req)
	if err == nil {
		t.Fatal("Expected error when service is invalid, got nil")
	}

	expectedErrMsg := "API Error 400: Invalid service type provided."
	if !strings.HasSuffix(err.Error(), expectedErrMsg) {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrMsg, err.Error())
	}
}
