package leonardo

import (
	"context"
	"fmt"
	"net/http"
)

// ModelsService provides methods to interact with the Models endpoints of the Leonardo.ai API.
type ModelsService struct {
	client *Client
}

// NewModelsService creates a new ModelsService.
func (c *Client) NewModelsService() *ModelsService {
	return &ModelsService{client: c}
}

// TrainCustomModel trains a new custom model using the specified dataset and parameters.
// POST /models
func (s *ModelsService) TrainCustomModel(ctx context.Context, req TrainCustomModelRequest) (*TrainCustomModelResponse, error) {
	var resp TrainCustomModelResponse
	path := "/models"

	httpReq, err := s.client.NewRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, fmt.Errorf("creating TrainCustomModel request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("training custom model failed: %w", err)
	}

	return &resp, nil
}

// GetCustomModel retrieves details of a specific custom model by its ID.
// GET /models/{id}
func (s *ModelsService) GetCustomModel(ctx context.Context, id string) (*GetCustomModelResponse, error) {
	var resp GetCustomModelResponse
	path := fmt.Sprintf("/models/%s", urlPathEscape(id))

	httpReq, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating GetCustomModel request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("retrieving custom model failed: %w", err)
	}

	return &resp, nil
}

// DeleteCustomModel deletes a specific custom model by its ID.
// DELETE /models/{id}
func (s *ModelsService) DeleteCustomModel(ctx context.Context, id string) (*DeleteCustomModelResponse, error) {
	var resp DeleteCustomModelResponse
	path := fmt.Sprintf("/models/%s", urlPathEscape(id))

	httpReq, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating DeleteCustomModel request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("deleting custom model failed: %w", err)
	}

	return &resp, nil
}

// ListPlatformModels retrieves platform models with pagination support.
// GET /platformModels
func (s *ModelsService) ListPlatformModels(ctx context.Context, req PaginationParams) (*ListPlatformModelsResponse, error) {
	var resp ListPlatformModelsResponse
	path := fmt.Sprintf("/platformModels?limit=%d&offset=%d", req.Limit, req.Offset)

	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating ListPlatformModels request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("listing platform models failed: %w", err)
	}

	return &resp, nil
}

// UpdateCustomModel updates the details of a specific custom model.
// PUT /models/{id}
func (s *ModelsService) UpdateCustomModel(ctx context.Context, id string, req UpdateCustomModelRequest) (*UpdateCustomModelResponse, error) {
	var resp UpdateCustomModelResponse
	path := fmt.Sprintf("/models/%s", urlPathEscape(id))

	httpReq, err := s.client.NewRequest(ctx, "PUT", path, req)
	if err != nil {
		return nil, fmt.Errorf("creating UpdateCustomModel request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("updating custom model failed: %w", err)
	}

	return &resp, nil
}
