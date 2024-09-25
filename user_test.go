package leonardo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetUserInfo tests the GetUserInfo method.
func TestGetUserInfo(t *testing.T) {
	// Mock server setup
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/me"
		if r.URL.Path != expectedPath || r.Method != "GET" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}

		// Mock successful response
		w.Header().Set("Content-Type", "application/json")
		response := GetUserInfoResponse{
			UserDetails: []struct {
				APIPlanTokenRenewalDate *string `json:"apiPlanTokenRenewalDate"`
				APIConcurrencySlots     *int    `json:"apiConcurrencySlots"`
				APIPaidTokens           *int    `json:"apiPaidTokens"`
				APISubscriptionTokens   *int    `json:"apiSubscriptionTokens"`
				PaidTokens              *int    `json:"paidTokens"`
				SubscriptionGPTTokens   *int    `json:"subscriptionGptTokens"`
				SubscriptionModelTokens *int    `json:"subscriptionModelTokens"`
				SubscriptionTokens      *int    `json:"subscriptionTokens"`
				TokenRenewalDate        *string `json:"tokenRenewalDate"`
				User                    User    `json:"user"`
			}{
				{
					APIPlanTokenRenewalDate: Ptr("2023-11-01T00:00:00Z"),
					APIConcurrencySlots:     Ptr(5),
					APIPaidTokens:           Ptr(200),
					APISubscriptionTokens:   Ptr(300),
					PaidTokens:              Ptr(150),
					SubscriptionGPTTokens:   Ptr(250),
					SubscriptionModelTokens: Ptr(100),
					SubscriptionTokens:      Ptr(400),
					TokenRenewalDate:        Ptr("2023-10-15T00:00:00Z"),
					User: User{
						ID:       Ptr("user-123"),
						Username: Ptr("testuser"),
					},
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
	client.User = client.NewUserService()

	// Execute GetUserInfo
	ctx := context.Background()
	resp, err := client.User.GetUserInfo(ctx)
	if err != nil {
		t.Fatalf("GetUserInfo failed: %v", err)
	}

	// Validate response
	if len(resp.UserDetails) != 1 {
		t.Fatalf("Expected 1 user detail, got %d", len(resp.UserDetails))
	}
	userDetail := resp.UserDetails[0]

	if *userDetail.User.ID != "user-123" {
		t.Errorf("Expected User ID 'user-123', got '%s'", *userDetail.User.ID)
	}
	if *userDetail.User.Username != "testuser" {
		t.Errorf("Expected Username 'testuser', got '%s'", *userDetail.User.Username)
	}

	if *userDetail.APIPaidTokens != 200 {
		t.Errorf("Expected APICredits 200, got %d", *userDetail.APIPaidTokens)
	}
	if *userDetail.APIPlanTokenRenewalDate != "2023-11-01T00:00:00Z" {
		t.Errorf("Expected APIPlanTokenRenewalDate '2023-11-01T00:00:00Z', got '%s'", *userDetail.APIPlanTokenRenewalDate)
	}
	if *userDetail.APIConcurrencySlots != 5 {
		t.Errorf("Expected APIConcurrencySlots 5, got %d", *userDetail.APIConcurrencySlots)
	}
	// Continue validating other fields as needed
}
