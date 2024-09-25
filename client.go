package leonardo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is the Leonardo.ai API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	APIKey     string

	// Services
	Datasets          *DatasetsService
	Images            *ImagesService
	Elements          *ElementsService
	Prompt            *PromptService
	InitImages        *InitImagesService
	Models            *ModelsService
	PricingCalculator *PricingCalculatorService
	RealtimeCanvas    *RealtimeCanvasService
	Texture           *TextureService
	ThreeDModelAssets *ThreeDModelAssetsService
	User              *UserService
	Variation         *VariationService
	Motion            *MotionService
}

// NewClient creates a new Leonardo.ai API client.
func NewClient(apiKey string) *Client {
	c := &Client{
		BaseURL: "https://cloud.leonardo.ai/api/rest/v1",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		APIKey: apiKey,
	}

	// Initialize services
	c.Datasets = c.NewDatasetsService()
	c.Images = c.NewImagesService()
	c.Elements = c.NewElementsService()
	c.Prompt = c.NewPromptService()
	c.InitImages = c.NewInitImagesService()
	c.Models = c.NewModelsService()
	c.PricingCalculator = c.NewPricingCalculatorService()
	c.RealtimeCanvas = c.NewRealtimeCanvasService()
	c.Texture = c.NewTextureService()
	c.ThreeDModelAssets = c.NewThreeDModelAssetsService()
	c.User = c.NewUserService()
	c.Variation = c.NewVariationService()
	c.Motion = c.NewMotionService()

	return c
}

// NewRequest creates an HTTP request with the given method, path, and body.
func (c *Client) NewRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	var buf io.Reader
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(jsonBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL+path, buf)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

// Do sends an HTTP request and decodes the response into v.
// It also handles API-specific error responses.
func (c *Client) Do(req *http.Request, v interface{}) error {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Attempt to decode the response
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var apiErr APIErrorResponse
		// io.Copy(os.Stdout, resp.Body)
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return fmt.Errorf("API request failed with status %d", resp.StatusCode)
		}
		return fmt.Errorf("API Error %s: %s", apiErr.Code, apiErr.Message)
	}

	// Decode successful response
	if v != nil {
		return json.NewDecoder(resp.Body).Decode(v)
	}
	return nil
}
