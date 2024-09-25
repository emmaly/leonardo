package leonardo

import (
	"context"
	"fmt"
	"net/url"
)

// InitImagesService provides methods to interact with the Init Images endpoints.
type InitImagesService struct {
	client *Client
}

// NewInitImagesService creates a new InitImagesService.
func (c *Client) NewInitImagesService() *InitImagesService {
	return &InitImagesService{client: c}
}

// UploadInitImage uploads an init image and retrieves presigned S3 upload details.
// POST /init-image
func (s *InitImagesService) UploadInitImage(ctx context.Context, req UploadInitImageRequest) (*struct {
	UploadInitImageID *string `json:"uploadInitImageId"`
	Message           string  `json:"message"`
}, error) {
	var resp struct {
		UploadInitImageID *string `json:"uploadInitImageId"`
		Message           string  `json:"message"`
	}
	path := "/init-image"

	httpReq, err := s.client.NewRequest(ctx, "POST", path, req)
	if err != nil {
		return nil, fmt.Errorf("creating UploadInitImage request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("uploading init image failed: %w", err)
	}

	return &resp, nil
}

// GetSingleInitImage retrieves a single init image by its ID.
// GET /init-image/{id}
func (s *InitImagesService) GetSingleInitImage(ctx context.Context, id string) (*GetSingleInitImageResponse, error) {
	var resp GetSingleInitImageResponse
	path := fmt.Sprintf("/init-image/%s", url.PathEscape(id))

	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating GetSingleInitImage request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("retrieving init image failed: %w", err)
	}

	return &resp, nil
}

// DeleteInitImage deletes an init image by its ID.
// DELETE /init-image/{id}
func (s *InitImagesService) DeleteInitImage(ctx context.Context, id string) (*DeleteInitImageResponse, error) {
	var resp DeleteInitImageResponse
	path := fmt.Sprintf("/init-image/%s", url.PathEscape(id))

	httpReq, err := s.client.NewRequest(ctx, "DELETE", path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating DeleteInitImage request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("deleting init image failed: %w", err)
	}

	return &resp, nil
}

// UploadCanvasInitAndMaskImage uploads canvas init and mask images and retrieves presigned S3 upload details.
// POST /canvas-init-image
func (s *InitImagesService) UploadCanvasInitAndMaskImage(ctx context.Context, req UploadCanvasInitAndMaskImageRequest) (*UploadCanvasInitAndMaskImageResponse, error) {
	var resp UploadCanvasInitAndMaskImageResponse
	path := "/canvas-init-image"

	httpReq, err := s.client.NewRequest(ctx, "POST", path, req)
	if err != nil {
		return nil, fmt.Errorf("creating UploadCanvasInitAndMaskImage request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("uploading canvas init and mask images failed: %w", err)
	}

	return &resp, nil
}
