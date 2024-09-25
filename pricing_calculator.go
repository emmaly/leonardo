package leonardo

import (
	"context"
	"fmt"
)

// PricingCalculatorService provides methods to interact with the Pricing Calculator endpoints.
type PricingCalculatorService struct {
	client *Client
}

// NewPricingCalculatorService creates a new PricingCalculatorService.
func (c *Client) NewPricingCalculatorService() *PricingCalculatorService {
	return &PricingCalculatorService{client: c}
}

// CalculateAPICost calculates the API credit cost for a given service.
// POST /pricing-calculator
func (s *PricingCalculatorService) CalculateAPICost(ctx context.Context, req CalculateAPICostRequest) (*CalculateAPICostResponse, error) {
	var resp CalculateAPICostResponse
	path := "/pricing-calculator"

	httpReq, err := s.client.NewRequest(ctx, "POST", path, req)
	if err != nil {
		return nil, fmt.Errorf("creating CalculateAPICost request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("calculating API cost failed: %w", err)
	}

	return &resp, nil
}
