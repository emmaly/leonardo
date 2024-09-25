package leonardo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// VariationService provides methods to interact with the Variation endpoints.
type VariationService struct {
	client *Client
}

// NewVariationService creates a new VariationService.
func (c *Client) NewVariationService() *VariationService {
	return &VariationService{client: c}
}

// CreateUnzoomVariation creates an unzoom variation for the provided image ID.
// POST /variations/unzoom
func (s *VariationService) CreateUnzoomVariation(ctx context.Context, req VariationRequest) (*CreateUnzoomVariationResponse, error) {
	var resp CreateUnzoomVariationResponse
	path := "/variations/unzoom"

	httpReq, err := s.client.NewRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, fmt.Errorf("creating CreateUnzoomVariation request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("creating unzoom variation failed: %w", err)
	}

	return &resp, nil
}

// CreateUpscaleVariation creates an upscale variation for a given image ID.
// POST /variations/upscale
func (s *VariationService) CreateUpscaleVariation(ctx context.Context, id string) (*UpscaleVariationResponse, error) {
	reqBody := UpscaleVariationRequest{
		ID: id,
	}
	var resp UpscaleVariationResponse
	path := "/variations/upscale"

	httpReq, err := s.client.NewRequest(ctx, "POST", path, reqBody)
	if err != nil {
		return nil, fmt.Errorf("creating CreateUpscaleVariation request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("creating upscale variation failed: %w", err)
	}

	return &resp, nil
}

// CreateNoBackgroundVariation creates a no background variation for the provided image ID.
// POST /variations/nobg
func (s *VariationService) CreateNoBackgroundVariation(ctx context.Context, req VariationRequest) (*CreateNoBackgroundVariationResponse, error) {
	var resp CreateNoBackgroundVariationResponse
	path := "/variations/nobg"

	httpReq, err := s.client.NewRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, fmt.Errorf("creating CreateNoBackgroundVariation request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("creating no background variation failed: %w", err)
	}

	return &resp, nil
}

// CreateUniversalUpscalerVariation creates a high-resolution image using Universal Upscaler.
// POST /variations/universal-upscaler
func (s *VariationService) CreateUniversalUpscalerVariation(ctx context.Context, req UniversalUpscalerRequest) (*UniversalUpscalerResponse, error) {
	var resp UniversalUpscalerResponse
	path := "/variations/universal-upscaler"

	httpReq, err := s.client.NewRequest(ctx, "POST", path, req)
	if err != nil {
		return nil, fmt.Errorf("creating CreateUniversalUpscalerVariation request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("creating universal upscaler variation failed: %w", err)
	}

	return &resp, nil
}

// GetVariation retrieves variation details by ID.
// GET /variations/{id}
func (s *VariationService) GetVariation(ctx context.Context, id string) (*GetVariationResponse, error) {
	var resp GetVariationResponse
	path := fmt.Sprintf("/variations/%s", url.PathEscape(id))

	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating GetVariation request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("retrieving variation failed: %w", err)
	}

	return &resp, nil
}
