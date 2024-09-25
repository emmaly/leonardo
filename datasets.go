// services/datasets.go
package leonardo

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// DatasetsService provides methods to interact with the Dataset endpoints of the Leonardo.ai API.
type DatasetsService struct {
	client *Client
}

// NewDatasetsService creates a new DatasetsService.
func (c *Client) NewDatasetsService() *DatasetsService {
	return &DatasetsService{client: c}
}

// CreateDataset creates a new dataset.
// POST /datasets
func (s *DatasetsService) CreateDataset(ctx context.Context, req CreateDatasetRequest) (*CreateDatasetResponse, error) {
	var resp CreateDatasetResponse
	path := "/datasets"

	httpReq, err := s.client.NewRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, fmt.Errorf("creating CreateDataset request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("creating dataset failed: %w", err)
	}

	return &resp, nil
}

// GetDataset retrieves a dataset by its ID.
// GET /datasets/{id}
func (s *DatasetsService) GetDataset(ctx context.Context, id string) (*GetDatasetResponse, error) {
	var resp GetDatasetResponse
	path := fmt.Sprintf("/datasets/%s", urlPathEscape(id))

	httpReq, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating GetDataset request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("retrieving dataset failed: %w", err)
	}

	return &resp, nil
}

// DeleteDataset deletes a dataset by its ID.
// DELETE /datasets/{id}
func (s *DatasetsService) DeleteDataset(ctx context.Context, id string) (*DeleteDatasetResponse, error) {
	var resp DeleteDatasetResponse
	path := fmt.Sprintf("/datasets/%s", urlPathEscape(id))

	httpReq, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating DeleteDataset request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("deleting dataset failed: %w", err)
	}

	return &resp, nil
}

// UploadDatasetImage retrieves presigned S3 upload details to upload a dataset image.
// POST /datasets/{datasetId}/upload
func (s *DatasetsService) UploadDatasetImage(ctx context.Context, datasetID string, req UploadDatasetImageRequest) (*UploadDatasetImageResponse, error) {
	var resp UploadDatasetImageResponse
	path := fmt.Sprintf("/datasets/%s/upload", urlPathEscape(datasetID))

	httpReq, err := s.client.NewRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, fmt.Errorf("creating UploadDatasetImage request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("retrieving upload details failed: %w", err)
	}

	return &resp, nil
}

// UploadImageToS3 uploads the image data to S3 using the presigned URL.
// This is a helper function and not directly interacting with Leonardo.ai API.
func UploadImageToS3(s3URL string, fields map[string]string, imagePath string) error {
	file, err := os.Open(imagePath)
	if err != nil {
		return fmt.Errorf("opening image file failed: %w", err)
	}
	defer file.Close()

	// Read the file content
	fileData, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("reading image file failed: %w", err)
	}

	// Prepare form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add form fields
	for key, value := range fields {
		if err := writer.WriteField(key, value); err != nil {
			return fmt.Errorf("writing form field %s failed: %w", key, err)
		}
	}

	// Add the file
	part, err := writer.CreateFormFile("file", imagePath)
	if err != nil {
		return fmt.Errorf("creating form file failed: %w", err)
	}
	if _, err := part.Write(fileData); err != nil {
		return fmt.Errorf("writing file data failed: %w", err)
	}

	// Close the writer to finalize the form
	if err := writer.Close(); err != nil {
		return fmt.Errorf("closing multipart writer failed: %w", err)
	}

	// Create the upload request
	req, err := http.NewRequest(http.MethodPost, s3URL, body)
	if err != nil {
		return fmt.Errorf("creating S3 upload request failed: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("S3 upload request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("S3 upload failed with status code: %d", resp.StatusCode)
	}

	return nil
}

// UploadGeneratedImageToDataset uploads a previously generated image to a dataset.
// POST /datasets/{datasetId}/upload/gen
func (s *DatasetsService) UploadGeneratedImageToDataset(ctx context.Context, datasetID string, req UploadGeneratedImageRequest) (*UploadGeneratedImageResponse, error) {
	var resp UploadGeneratedImageResponse
	path := fmt.Sprintf("/datasets/%s/upload/gen", urlPathEscape(datasetID))

	httpReq, err := s.client.NewRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, fmt.Errorf("creating UploadGeneratedImageToDataset request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("uploading generated image to dataset failed: %w", err)
	}

	return &resp, nil
}
