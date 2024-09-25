package leonardo

import (
	"context"
	"fmt"
	"net/http"
)

// RealtimeCanvasService provides methods to interact with the Realtime Canvas endpoints.
type RealtimeCanvasService struct {
	client *Client
}

// NewRealtimeCanvasService creates a new RealtimeCanvasService.
func (c *Client) NewRealtimeCanvasService() *RealtimeCanvasService {
	return &RealtimeCanvasService{client: c}
}

// CreateLCMGeneration generates a LCM image.
// POST /generations-lcm
func (s *RealtimeCanvasService) CreateLCMGeneration(ctx context.Context, req CreateLCMGenerationRequest) (*CreateLCMGenerationResponse, error) {
	var resp CreateLCMGenerationResponse
	path := "/generations-lcm"

	httpReq, err := s.client.NewRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, fmt.Errorf("creating CreateLCMGeneration request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("creating LCM generation failed: %w", err)
	}

	return &resp, nil
}

// PerformInstantRefine performs instant refine on a LCM image.
// POST /lcm-instant-refine
func (s *RealtimeCanvasService) PerformInstantRefine(ctx context.Context, req PerformInstantRefineRequest) (*PerformInstantRefineResponse, error) {
	var resp PerformInstantRefineResponse
	path := "/lcm-instant-refine"

	httpReq, err := s.client.NewRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, fmt.Errorf("creating PerformInstantRefine request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("performing instant refine failed: %w", err)
	}

	return &resp, nil
}

// PerformInpainting performs inpainting on a LCM image.
// POST /lcm-inpainting
func (s *RealtimeCanvasService) PerformInpainting(ctx context.Context, req PerformInpaintingRequest) (*PerformInpaintingResponse, error) {
	var resp PerformInpaintingResponse
	path := "/lcm-inpainting"

	httpReq, err := s.client.NewRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, fmt.Errorf("creating PerformInpainting request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("performing inpainting failed: %w", err)
	}

	return &resp, nil
}

// PerformAlchemyUpscale performs Alchemy Upscale on a LCM image.
// POST /lcm-upscale
func (s *RealtimeCanvasService) PerformAlchemyUpscale(ctx context.Context, req PerformAlchemyUpscaleRequest) (*PerformAlchemyUpscaleResponse, error) {
	var resp PerformAlchemyUpscaleResponse
	path := "/lcm-upscale"

	httpReq, err := s.client.NewRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, fmt.Errorf("creating PerformAlchemyUpscale request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("performing Alchemy Upscale failed: %w", err)
	}

	return &resp, nil
}
