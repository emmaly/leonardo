package leonardo

import (
	"context"
	"fmt"
)

// UserService provides methods to interact with the User endpoints of the Leonardo.ai API.
type UserService struct {
	client *Client
}

// NewUserService creates a new UserService.
func (c *Client) NewUserService() *UserService {
	return &UserService{client: c}
}

// GetUserInfo retrieves information about the authenticated user.
// GET /me
func (s *UserService) GetUserInfo(ctx context.Context) (*GetUserInfoResponse, error) {
	var resp GetUserInfoResponse
	path := "/me"

	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating GetUserInfo request failed: %w", err)
	}

	err = s.client.Do(httpReq, &resp)
	if err != nil {
		return nil, fmt.Errorf("retrieving user info failed: %w", err)
	}

	return &resp, nil
}
