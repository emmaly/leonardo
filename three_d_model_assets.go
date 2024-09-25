package leonardo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ThreeDModelAssetsService provides methods to interact with the 3D Model Assets endpoints.
type ThreeDModelAssetsService struct {
	client *Client
}

// NewThreeDModelAssetsService creates a new ThreeDModelAssetsService.
func (c *Client) NewThreeDModelAssetsService() *ThreeDModelAssetsService {
	return &ThreeDModelAssetsService{client: c}
}

// Upload3DModel uploads a 3D model and retrieves presigned S3 upload details.
// POST /models-3d/upload
func (s *ThreeDModelAssetsService) Upload3DModel(ctx context.Context, req Upload3DModelRequest) (*Upload3DModelResponse, error) {
	var resp Upload3DModelResponse
	path := "/models-3d/upload"

	httpReq, err := s.client.NewRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, fmt.Errorf("creating Upload3DModel request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("uploading 3D model failed: %w", err)
	}

	return &resp, nil
}

// Get3DModelsByUser retrieves all 3D models associated with a specific user ID.
// GET /models-3d/user/{userId}?limit={limit}&offset={offset}
func (s *ThreeDModelAssetsService) Get3DModelsByUser(ctx context.Context, userID string, limit, offset int) (*Get3DModelsByUserResponse, error) {
	var resp Get3DModelsByUserResponse

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

	path := fmt.Sprintf("/models-3d/user/%s%s", url.PathEscape(userID), queryString)

	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating Get3DModelsByUser request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("retrieving 3D models by user ID failed: %w", err)
	}

	return &resp, nil
}

// Get3DModelByID retrieves a specific 3D model by its ID.
// GET /models-3d/{id}
func (s *ThreeDModelAssetsService) Get3DModelByID(ctx context.Context, id string) (*Get3DModelByIDResponse, error) {
	var resp Get3DModelByIDResponse
	path := fmt.Sprintf("/models-3d/%s", urlPathEscape(id))

	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating Get3DModelByID request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("retrieving 3D model by ID failed: %w", err)
	}

	return &resp, nil
}

// Delete3DModel deletes a specific 3D model by its ID.
// DELETE /models-3d/{id}
func (s *ThreeDModelAssetsService) Delete3DModel(ctx context.Context, id string) (*Delete3DModelResponse, error) {
	var resp Delete3DModelResponse
	path := fmt.Sprintf("/models-3d/%s", urlPathEscape(id))

	httpReq, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating Delete3DModel request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("deleting 3D model failed: %w", err)
	}

	return &resp, nil
}
