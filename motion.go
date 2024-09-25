package leonardo

import (
	"context"
	"fmt"
)

// MotionService provides methods to interact with the Motion endpoints.
type MotionService struct {
	client *Client
}

// NewMotionService creates a new MotionService.
func (c *Client) NewMotionService() *MotionService {
	return &MotionService{client: c}
}

// CreateSVDMotionGeneration generates an SVD motion generation.
// POST /generations-motion-svd
func (s *MotionService) CreateSVDMotionGeneration(ctx context.Context, req MotionRequest) (*CreateSVDMotionGenerationResponse, error) {
	var resp CreateSVDMotionGenerationResponse
	path := "/generations-motion-svd"

	httpReq, err := s.client.NewRequest(ctx, "POST", path, req)
	if err != nil {
		return nil, fmt.Errorf("creating CreateSVDMotionGeneration request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("creating SVD motion generation failed: %w", err)
	}

	return &resp, nil
}
