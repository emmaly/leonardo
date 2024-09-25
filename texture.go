package leonardo

import (
	"context"
	"fmt"
	"net/http"
)

// TextureService provides methods to interact with the Texture endpoints.
type TextureService struct {
	client *Client
}

// NewTextureService creates a new TextureService.
func (c *Client) NewTextureService() *TextureService {
	return &TextureService{client: c}
}

// CreateTextureGeneration generates a texture.
// POST /generations-texture
func (s *TextureService) CreateTextureGeneration(ctx context.Context, req CreateTextureGenerationRequest) (*CreateTextureGenerationResponse, error) {
	var resp CreateTextureGenerationResponse
	path := "/generations-texture"

	httpReq, err := s.client.NewRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, fmt.Errorf("creating CreateTextureGeneration request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("creating texture generation failed: %w", err)
	}

	return &resp, nil
}
