// services/prompt.go
package leonardo

import (
	"context"
	"fmt"
)

// PromptService provides methods to interact with the Prompt endpoints of the Leonardo.ai API.
type PromptService struct {
	client *Client
}

// NewPromptService creates a new PromptService.
func (c *Client) NewPromptService() *PromptService {
	return &PromptService{client: c}
}

// GenerateRandomPrompt generates a random prompt.
// POST /prompt/random
func (s *PromptService) GenerateRandomPrompt(ctx context.Context) (*GenerateRandomPromptResponse, error) {
	var resp GenerateRandomPromptResponse
	path := "/prompt/random"

	httpReq, err := s.client.NewRequest(ctx, "POST", path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating GenerateRandomPrompt request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("generating random prompt failed: %w", err)
	}

	return &resp, nil
}

// ImprovePrompt improves an existing prompt.
// POST /prompt/improve
func (s *PromptService) ImprovePrompt(ctx context.Context, req ImprovePromptRequest) (*ImprovePromptResponse, error) {
	var resp ImprovePromptResponse
	path := "/prompt/improve"

	httpReq, err := s.client.NewRequest(ctx, "POST", path, req)
	if err != nil {
		return nil, fmt.Errorf("creating ImprovePrompt request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("improving prompt failed: %w", err)
	}

	return &resp, nil
}
