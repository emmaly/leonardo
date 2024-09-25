package leonardo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ImagesService provides methods to interact with the Image endpoints of the Leonardo.ai API.
type ImagesService struct {
	client *Client
}

// NewImagesService creates a new ImagesService.
func (c *Client) NewImagesService() *ImagesService {
	return &ImagesService{client: c}
}

// CreateImageGeneration generates images based on a prompt.
// POST /generations
func (s *ImagesService) CreateImageGeneration(ctx context.Context, req CreateGenerationRequest) (*CreateGenerationResponse, error) {
	var resp CreateGenerationResponse
	path := "/generations"

	httpReq, err := s.client.NewRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, fmt.Errorf("creating CreateImageGeneration request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("creating image generation failed: %w", err)
	}

	return &resp, nil
}

// GetImageGeneration retrieves information about a specific image generation.
// GET /generations/{id}
func (s *ImagesService) GetImageGeneration(ctx context.Context, id string) (*GetGenerationResponse, error) {
	var resp GetGenerationResponse
	path := fmt.Sprintf("/generations/%s", urlPathEscape(id))

	httpReq, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating GetImageGeneration request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("getting image generation failed: %w", err)
	}

	return &resp, nil
}

// DeleteGeneration deletes a specific image generation by its ID.
// DELETE /generations/{id}
func (s *ImagesService) DeleteGeneration(ctx context.Context, id string) (*DeleteGenerationResponse, error) {
	var resp DeleteGenerationResponse
	path := fmt.Sprintf("/generations/%s", urlPathEscape(id))

	httpReq, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating DeleteGeneration request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("deleting generation failed: %w", err)
	}

	return &resp, nil
}

// GetGenerationsByUserID retrieves all generations associated with a specific user ID.
// GET /generations/user/{userId}
// Optional parameters: limit, offset for pagination.
func (s *ImagesService) GetGenerationsByUserID(ctx context.Context, userID string, limit, offset int) (*GetGenerationsByUserResponse, error) {
	var resp GetGenerationsByUserResponse

	// Construct query parameters if available
	queryParams := url.Values{}
	if limit > 0 {
		queryParams.Set("limit", fmt.Sprintf("%d", limit))
	}
	if offset > 0 {
		queryParams.Set("offset", fmt.Sprintf("%d", offset))
	}
	queryString := ""
	if len(queryParams) > 0 {
		queryString = "?" + queryParams.Encode()
	}

	path := fmt.Sprintf("/generations/user/%s%s", url.PathEscape(userID), queryString)

	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating GetGenerationsByUserID request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("retrieving generations by user ID failed: %w", err)
	}

	return &resp, nil
}
