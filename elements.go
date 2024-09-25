package leonardo

import (
	"context"
	"fmt"
)

// ElementsService provides methods to interact with the Elements endpoints.
type ElementsService struct {
	client *Client
}

// NewElementsService creates a new ElementsService.
func (c *Client) NewElementsService() *ElementsService {
	return &ElementsService{client: c}
}

// ListElements retrieves a list of public Elements available for use with generations.
// GET /elements
func (s *ElementsService) ListElements(ctx context.Context) (*ListElementsResponse, error) {
	var resp ListElementsResponse
	path := "/elements"

	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating ListElements request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("listing elements failed: %w", err)
	}

	return &resp, nil
}
